package service

import (
	"context"
	"errors"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"sync"

	"github.com/BurntSushi/locker"
	"github.com/bonedaddy/go-indexed/bclient"
	"github.com/bonedaddy/go-indexed/bindings/erc20"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/indexed-finance/circuit-breaker/alerts"
	"github.com/indexed-finance/circuit-breaker/bindings/controller"
	"github.com/indexed-finance/circuit-breaker/bindings/sigmacore"
	"github.com/indexed-finance/circuit-breaker/config"
	"github.com/indexed-finance/circuit-breaker/database"
	"github.com/indexed-finance/circuit-breaker/eventwatcher"
	"github.com/indexed-finance/circuit-breaker/multicall"
	"github.com/indexed-finance/circuit-breaker/utils"
	"go.uber.org/multierr"
	"go.uber.org/zap"
)

// Service wraps multiple modules to provide the circuit breaking functionality
type Service struct {
	db     *database.Database
	ew     *eventwatcher.EventWatcher
	logger *zap.Logger
	mc     *multicall.Multicall
	at     *alerts.Alerter
	auther *utils.Authorizer
	ctx    context.Context
	cancel context.CancelFunc
	start  sync.Once
	// pools represents the index pools to process
	pools         []config.Pool
	mx            sync.RWMutex
	wg            sync.WaitGroup
	namedLock     *locker.Locker
	minimumGwei   *big.Int
	gasMultiplier *big.Int
}

// New intializes all needed modules and returns a ready to consume service
func New(
	ctx context.Context,
	alert *alerts.Alerter,
	mc *multicall.Multicall,
	db *database.Database,
	bc *bclient.Client,
	auther *utils.Authorizer,
	logger *zap.Logger,
	pools []config.Pool,
	gasConfig config.GasPrice) (*Service, error) {
	ctx, cancel := context.WithCancel(ctx)
	minimumGwei, ok := new(big.Int).SetString(gasConfig.MinimumGwei, 10)
	if !ok {
		cancel()
		return nil, errors.New("failed to parse minimum gwei config")
	}
	gasMultiplier, ok := new(big.Int).SetString(gasConfig.Multiplier, 10)
	if !ok {
		cancel()
		return nil, errors.New("failed to parse gas multiplier")
	}
	return &Service{
		db:            db,
		ew:            eventwatcher.New(bc),
		logger:        logger.Named("service"),
		mc:            mc,
		at:            alert,
		auther:        auther,
		ctx:           ctx,
		cancel:        cancel,
		pools:         pools,
		namedLock:     locker.NewLocker(),
		minimumGwei:   minimumGwei,
		gasMultiplier: gasMultiplier,
	}, nil
}

// Prepare is used to prepare the service the primary responsibility of this function
// is to check the database and determine if we need to update the records. This will
// usually only happen during one of two conditions
//     1) the very first time the service is started on a fresh database
//     2) service starts up after being offline for 1 or more blocks since last run
func (s *Service) Prepare() error {
	var startErr error
	s.start.Do(func() {
		s.mx.RLock()
		defer s.mx.RUnlock()
		block, err := s.ew.BC().CurrentBlock()
		if err != nil {
			s.logger.Error("failed to get current block", zap.Error(err))
			startErr = err
			return
		}
		for _, pool := range s.pools {
			// construct the pool contract bindings
			contract, err := sigmacore.NewSigmacore(common.HexToAddress(pool.ContractAddress), s.ew.BC().EthClient())
			if err != nil {
				s.logger.Error("failed to get pool contract bindings")
				startErr = err
				return
			}
			// make sure the pool exists
			dPool, err := s.db.GetPool(pool.Name)
			if err != nil {
				s.logger.Info("creating fresh pool database entry", zap.String("pool", pool.Name))
				tokens, err := utils.PoolTokensFor(contract, s.ew.BC().EthClient())
				if err != nil {
					s.logger.Error("failed to get pool tokens", zap.Error(err))
					return
				}
				balances, denormWeights, totalSupplies, err := s.GetBalancesWeightsAndSupplies(contract, block, pool.ContractAddress, tokens)
				if err != nil {
					s.logger.Error("GetBalancesWeightsAndSupplies call failed", zap.Error(err), zap.String("pool", pool.Name))
					startErr = err
					return
				}
				// gather decimal information
				decimals := make(map[string]interface{})
				for tok, addr := range tokens {
					erc20Contract, err := erc20.NewErc20(addr, s.ew.BC().EthClient())
					if err != nil {
						s.logger.Error("failed to get erc20 contract binding", zap.Error(err), zap.String("pool", pool.Name), zap.String("token", tok))
						startErr = err
						return
					}
					udec, err := erc20Contract.Decimals(nil)
					if err != nil {
						s.logger.Error("failed to get token decimals", zap.Error(err), zap.String("pool", pool.Name), zap.String("token", tok))
						startErr = err
						return
					}
					decimals[strings.ToLower(tok)] = strconv.FormatUint(uint64(udec), 10)
				}
				// create the record
				if err := s.db.NewPool(
					pool.Name,
					pool.ContractAddress,
					block,
					balances,
					denormWeights,
					totalSupplies,
					decimals,
				); err != nil {
					s.logger.Error("failed to create pool db entry", zap.Error(err), zap.String("pool", pool.Name))
					startErr = err
					return
				}
				// no more processing needed so skip this pool
				continue
			}
			if block <= dPool.LastUpdateAt {
				// skip
				continue
			}
			s.logger.Info("updating pool database entry", zap.String("pool", pool.Name))
			if _, _, _, err := s.UpdateBalancesWeightsAndSupplies(block, pool, contract); err != nil {
				s.logger.Error("failed to update pool balance and weights", zap.Error(err), zap.String("pool", pool.Name))
				startErr = err
				return
			}
		}
	})
	if startErr != nil {
		s.logger.Error("failed to prepare service", zap.Error(startErr))
	} else {
		s.logger.Info("service preparation succeeded")
	}
	return startErr
}

// StartWatchers starts all of our contract watchers
// which listen for LOG_SWAP events to potentially trigger
// circuit breaking activities
func (s *Service) StartWatchers() error {

	var (
		// maps pool name -> contract address
		addresses = make(map[string]string)
		// maps pool name -> break percentage
		spotPriceBreakPercentages = make(map[string]float64)
	)

	for _, pool := range s.pools {
		addresses[strings.ToLower(pool.Name)] = pool.ContractAddress
		spotPriceBreakPercentages[strings.ToLower(pool.Name)] = pool.SpotPriceBreakPercentage
	}

	// get the bindings
	// name -> contract
	bindings, err := s.ew.NewBindings(addresses)
	if err != nil {
		return err
	}

	watchers, err := s.ew.NewWatchedContracts(s.logger, s.ew.BC().EthClient(), bindings, s.minimumGwei, s.gasMultiplier)
	if err != nil {
		return err
	}

	var (
		// create a dedicated context for the contract watchers
		// this will allow us to easily cancel this context upon receiving an error
		// from one of the spawned watchers, allowing us to gracefully exit and restart
		ctx, cancel = context.WithCancel(s.ctx)
		// this error channel will hold any errors returned from a single contract watcher
		// it will allow us to block until either a context is cancelled or an error is returned
		// we create a channel with a capacity equal to the number of watchers we will spawn
		errCh = make(chan error, len(watchers))
	)

	s.wg.Add(len(watchers))

	for _, watcher := range watchers {
		go func(wtchr *eventwatcher.WatchedContract, bkPercent float64) {
			defer s.wg.Done()
			if err := wtchr.Listen(ctx, s.db, s.at, s.auther, bkPercent, s.ew.BC().EthClient()); err != nil {
				errCh <- err
			}
		}(watcher, spotPriceBreakPercentages[strings.ToLower(watcher.Name())])
	}

	select {
	case <-ctx.Done():
		cancel()
		return nil
	case err := <-errCh:
		s.logger.Error("encountered error from contract watcher, aborting...", zap.Error(err))
		cancel()
		return err
	}

}

// StartBlockListener waits for new blocks to be mined to trigger blockchain db updates
func (s *Service) StartBlockListener() error {
	ch := make(chan *types.Header, 100)
	sub, err := s.ew.BC().EthClient().SubscribeNewHead(s.ctx, ch)
	if err != nil {
		s.logger.Error("failed to create new head subscription", zap.Error(err))
		close(ch)
		return err
	}
	for {
		select {
		case <-s.ctx.Done():
			sub.Unsubscribe()
			close(ch)
			return nil
		case err := <-sub.Err():
			s.logger.Error("new head subscription error received", zap.Error(err))
			sub.Unsubscribe()
			close(ch)
			return err
		case header := <-ch:
			s.logger.Debug("new block header received", zap.Any("header", header))

			// do a single update for wait group incrementation instead of multiple times
			s.wg.Add(len(s.pools))

			for _, pool := range s.pools {

				// claim the named lock
				s.namedLock.Lock(pool.Name)

				// start the total supply update information in the background
				go func(pool config.Pool) {

					// defer named lock unlock
					defer s.namedLock.Unlock(pool.Name)
					defer s.wg.Done()

					// grab the previous pool information so we can determine if we should break the circuit
					dPool, err := s.db.GetPool(pool.Name)
					if err != nil {
						s.logger.Error("failed to get pool database entry", zap.Error(err), zap.String("pool", pool.Name))
						return
					}
					// get the information for the time this record was last updated at
					infos, err := s.db.GetInfo(dPool.Name, dPool.LastUpdateAt)
					if err != nil {
						s.logger.Error("failed to get pool infos database entry", zap.Error(err), zap.String("pool", pool.Name), zap.Uint64("updated.at", dPool.LastUpdateAt))
						return
					}

					// check the block number from the header
					blockNum := header.Number.Uint64()

					s.logger.Info("updating pool information (supply, balances, weights)", zap.String("pool", pool.Name), zap.Uint64("block.num", blockNum))

					// create bindings for the given contract
					contract, err := sigmacore.NewSigmacore(
						common.HexToAddress(pool.ContractAddress),
						s.ew.BC().EthClient(),
					)
					if err != nil {
						s.logger.Error("failed to get contract binding", zap.Error(err), zap.String("pool", pool.Name))
						return
					}
					// update the information stored in the database
					_, _, totalSupplies, err := s.UpdateBalancesWeightsAndSupplies(blockNum, pool, contract)
					if err != nil {
						s.logger.Error("UpdateBalancesWeightsAndSupplies failed", zap.Error(err), zap.String("pool", pool.Name))
					}

					// now start checking total supply shifts
					for tok, supply := range infos.TokenTotalSupplies {
						// check the total supply for the given token
						// however if it fails dont abort processing, continue processing
						// as other tokens may need to be checked
						if controllerAddress, err := contract.GetController(nil); err != nil {
							s.logger.Error("failed to get controller address", zap.Error(err), zap.String("pool", pool.Name), zap.String("token", tok))
						} else {
							conContract, err := controller.NewController(controllerAddress, s.ew.BC().EthClient())
							if err != nil {
								s.logger.Error("failed to get controller contract", zap.Error(err), zap.String("pool", pool.Name), zap.String("token", tok))
							} else {
								// get the circuit breaker address
								circuitBreakerAddr, err := conContract.CircuitBreaker(nil)
								if err != nil {
									s.logger.Error("failed to get circuit breaker contract address", zap.Error(err), zap.String("pool", pool.Name), zap.String("token", tok))
								} else {
									breaker, err := controller.NewController(circuitBreakerAddr, s.ew.BC().EthClient())
									if err != nil {
										s.logger.Error("failed to get circuit breaker contract", zap.Error(err), zap.String("pool", pool.Name), zap.String("token", tok))
									} else {
										if err := s.circuitBreakCheck(tok, supply, totalSupplies, pool, breaker); err != nil {
											s.logger.Error("circuitBreakCheck failed", zap.Error(err), zap.String("pool", pool.Name), zap.String("token", tok))
										}
									}
								}
							}
						}
					}

					// purge old records if need be
					s.purgeInfoCheck(pool.Name)

				}(pool)
			}
		}
	}
}

// Close shuts down all internal modules returning any errors encountered from shutdown
func (s *Service) Close() error {
	var errors error
	s.logger.Info("cancelling context")
	s.cancel()
	// wait for pending goroutines to exit
	s.logger.Info("waiting for existing goroutines to exit")
	s.wg.Wait()
	s.logger.Info("closing multicaller")
	s.mc.Close()
	s.logger.Info("closing event watcher")
	// close eventwatcher/bclient
	s.ew.Close()
	// close db
	s.logger.Info("closing database")
	if err := s.db.Close(); err != nil {
		errors = multierr.Combine(errors, err)
	}
	s.logger.Info("shutdown process finished, goodbye...")
	return errors
}

// circuitBreakCheck performs a check to see if a circuit needs to be broken
// and if so handles breaking of the circuit
func (s *Service) circuitBreakCheck(
	tok string, supply interface{},
	totalSupplies map[string]interface{},
	pool config.Pool,
	contract utils.Breaker,
) error {

	oldSupplyStr, ok := supply.(string)
	if !ok {
		return errors.New("old supply is not of type string")
	}

	oldSupplyBig, ok := new(big.Int).SetString(oldSupplyStr, 10)
	if !ok {
		return errors.New("failed to convert old supply from string to big.Int")
	}

	newSupplyStr, ok := totalSupplies[tok].(string)
	if !ok {
		return errors.New("new supply is not of type string")
	}

	newSupplyBig, ok := new(big.Int).SetString(newSupplyStr, 10)
	if !ok {
		return errors.New("failed to convert new supply from string to big.Int")
	}

	var (
		direction string
		change    = (utils.GetChangePercentageBig(newSupplyBig, oldSupplyBig)) * 100
	)

	if change < 0 {
		direction = "decreased"
	} else if change > 0 {
		direction = "increased"
	}

	s.logger.Debug(
		"token total supply changed",
		zap.String("direction", direction),
		zap.Float64("change.percent", change),
		zap.String("pool", pool.Name),
		zap.String("token", tok),
		zap.String("supply.old", oldSupplyStr),
		zap.String("supply.new", newSupplyStr),
	)

	// we take the absolute to see if the change is greater than the break percentage
	// as we want to handle circuit breaks whether the total supply increased or decreases
	if math.Abs(change) > pool.SupplyBreakPercentage {

		s.logger.Warn(
			"token supply fluctuation is greater than minimum break percentage, breaking circuits!",
			zap.String("pool", pool.Name), zap.String("token", tok),
			zap.Float64("break_at", pool.SupplyBreakPercentage), zap.Float64("change", change), zap.Float64("change_abs", math.Abs(change)),
			zap.String("supply.old", oldSupplyBig.String()), zap.String("supply.new", newSupplyBig.String()),
		)

		// lock the authorizer since bind.TransactOpts is not threadsafe
		s.auther.Lock()
		// get gas price for break transactoin
		gasPrice, err := utils.GetGasPrice(s.ctx, s.ew.BC().EthClient(), s.minimumGwei, s.gasMultiplier)
		if err != nil {
			s.logger.Error("failed to suggest gasprice", zap.Error(err))
		} else {
			s.logger.Info("gas price calculated (includes boost)", zap.String("gas.price", gasPrice.String()))
			s.setPublicSwap(contract, gasPrice, pool.Name, tok, pool.ContractAddress)
		}
		// we need to unset the gas price that we overrode the transactor with
		// so that future uses of this transactor have the gas price set to nil
		// this is set within the setPublicSwap call
		s.auther.TransactOpts.GasPrice = nil
		s.auther.Unlock()

		if err := s.at.NotifyCircuitBreak(
			fmt.Sprintf(
				"a circuit break is in progress. pool: %s, token: %s, reason: token supply %s",
				pool.Name, tok, direction,
			),
		); err != nil {
			s.logger.Error("circuit break notification failed")
		}

	}
	return nil
}

// performs a check to see if we need to purge records from the database
func (s *Service) purgeInfoCheck(name string) {
	// check to see if we need to purge old information
	if infos, err := s.db.GetNumInfos(name); err == nil {
		if infos > 512 {
			s.logger.Debug("purging old infos")
			if err := s.db.PurgeOldInfos(name, 512); err != nil {
				s.logger.Error("failed to purge old infos", zap.Error(err))
			}
		}
	}
}

func (s *Service) setPublicSwap(
	contract utils.Breaker,
	gasPrice *big.Int, poolName, tokenName, poolAddress string,
) {
	s.auther.GasPrice = gasPrice
	if tx, err := contract.SetPublicSwap(s.auther.TransactOpts, common.HexToAddress(poolAddress), false); err != nil {
		s.logger.Error(
			"failed to broadcast public swap disable transaction",
			zap.Error(err),
			zap.String("pool", poolName),
			zap.String("token", tokenName),
		)
	} else {
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			s.logger.Warn(
				"waiting for public swap disable transaction to be mined",
				zap.String("pool", poolName),
				zap.String("token", tokenName),
				zap.String("tx.hash", tx.Hash().String()),
			)
			if rcpt, err := bind.WaitMined(s.ctx, s.ew.BC().EthClient(), tx); err != nil {
				s.logger.Error(
					"transaction failed to be mined",
					zap.String("pool", poolName),
					zap.String("token", tokenName),
					zap.String("tx.hash", tx.Hash().String()),
				)
			} else {
				s.logger.Warn(
					"transaction mined",
					zap.String("pool", poolName),
					zap.String("token", tokenName),
					zap.Any("receipt", rcpt),
				)
			}
		}()
	}
}
