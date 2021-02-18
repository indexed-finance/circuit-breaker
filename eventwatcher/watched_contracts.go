package eventwatcher

import (
	"context"
	"errors"
	"fmt"
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
	"github.com/indexed-finance/circuit-breaker/alerts"
	"github.com/indexed-finance/circuit-breaker/bindings/controller"
	"github.com/indexed-finance/circuit-breaker/bindings/sigmacore"
	"github.com/indexed-finance/circuit-breaker/database"
	"github.com/indexed-finance/circuit-breaker/utils"
	"go.uber.org/zap"
)

// WatchedContract is a singular watched contract
type WatchedContract struct {
	name           string
	binding        *sigmacore.Sigmacore
	evCh           chan *sigmacore.SigmacoreLOGSWAP
	notifCh        chan *sigmacore.SigmacoreLOGSWAP // duplicate of evCh to send events to watchers
	sub            event.Subscription
	logger         *zap.Logger
	tokenAddresses map[string]common.Address
	tokenNames     map[common.Address]string
	isSimulation   bool // indicates if this is a simulation for more refined control over parameters
	backend        bind.ContractBackend
	controller     *controller.Controller
}

// NewWatchedContracts initializes watched contracts and prepares them for event listening
func (ew *EventWatcher) NewWatchedContracts(logger *zap.Logger, backend bind.ContractBackend, bindings map[string]*sigmacore.Sigmacore) ([]*WatchedContract, error) {
	out := make([]*WatchedContract, 0)
	for name, contract := range bindings {
		ch := make(chan *sigmacore.SigmacoreLOGSWAP, 100) // todo: we may want to change this in the future
		sub, err := contract.WatchLOGSWAP(nil, ch, nil, nil, nil)
		if err != nil {
			return nil, err
		}
		// get the current pool tokens
		// this logic is copied from go-indexed/blcient
		// we could use the ew.BC() function but that makes it harder to do simulated testing
		addrs, err := utils.PoolTokensFor(contract, backend)
		if err != nil {
			return nil, err
		}
		names := make(map[common.Address]string)
		for name, addr := range addrs {
			names[addr] = strings.ToLower(name)
		}
		controllerAddress, err := contract.GetController(nil)
		if err != nil {
			return nil, err
		}
		control, err := controller.NewController(controllerAddress, backend)
		if err != nil {
			return nil, err
		}
		// lowercase the name
		name = strings.ToLower(name)
		out = append(out, &WatchedContract{
			name:           name,
			binding:        contract,
			evCh:           ch,
			notifCh:        make(chan *sigmacore.SigmacoreLOGSWAP),
			sub:            sub,
			logger:         logger.Named("contract.watcher").With(zap.String("pool", name)),
			tokenAddresses: addrs,
			tokenNames:     names,
			backend:        backend,
			controller:     control,
		})
	}
	return out, nil
}

// Listen stars the watche contract listening process waiting for incoming events
func (wc *WatchedContract) Listen(ctx context.Context, db *database.Database, alerter *alerts.Alerter, authorizer *utils.Authorizer, breakPercentage float64) error {
	wc.logger.Info("contract watcher started listening for events")
	// get the pool contract address
	dbInfo, err := db.GetPool(wc.name)
	if err != nil {
		return err
	}
	poolAddress := common.HexToAddress(dbInfo.ContractAddress)
	for {
		select {
		case <-ctx.Done():
			wc.sub.Unsubscribe()
			close(wc.evCh)
			close(wc.notifCh)
			return nil
		case err := <-wc.sub.Err():
			wc.logger.Error("event subscription error received", zap.Error(err))
			wc.sub.Unsubscribe()
			close(wc.evCh)
			close(wc.notifCh)
			return err
		case evLog := <-wc.evCh:
			wc.logger.Debug("received new event", zap.Any("event", evLog))
			// skip reverted logs
			if evLog.Raw.Removed {
				wc.logger.Warn("event was marked as removed by blockchain due to reorg")
				continue
			}

			// send event to watchers
			select {
			case wc.notifCh <- evLog:
			default:
			}

			prevSwapFee, err := wc.getPreviousSwapFee(evLog.Raw.BlockNumber - 1)
			if err != nil {
				wc.logger.Error("failed to get previous swap fee", zap.Error(err))
				continue
			}

			currSpotPrice, err := wc.getCurrentSpotPrice(evLog)
			if err != nil {
				wc.logger.Error("failed to get current spot price", zap.Error(err))
				continue

			}
			wc.logger.Debug(
				"LOG_SWAP event spot price calculated",
				zap.String("tx.hash", evLog.Raw.TxHash.String()),
				zap.String("price", currSpotPrice.String()),
				zap.String("token.out", evLog.TokenOut.String()),
				zap.String("token_amount.out", evLog.TokenAmountOut.String()),
				zap.String("token.in", evLog.TokenIn.String()),
				zap.String("token_amount.in", evLog.TokenAmountIn.String()),
				zap.Uint64("block.number", evLog.Raw.BlockNumber),
			)

			prevSpotPrice, err := wc.getPreviousSpotPrice(db, evLog, prevSwapFee)
			if err != nil {
				wc.logger.Error("failed to calculate previous block spot price", zap.Error(err))
				continue
			}

			// multipliy by 100 to get correct change percentage
			change := (utils.GetChangePercentageBig(currSpotPrice, prevSpotPrice)) * 100
			if change != 0 {
				wc.logger.Warn(
					"spot price fluctuation detected",
					zap.String("price.previous", prevSpotPrice.String()),
					zap.String("price.current", currSpotPrice.String()),
					zap.String("token.out", evLog.TokenOut.String()),
					zap.String("token.in", evLog.TokenIn.String()),
					zap.String("tx.hash", evLog.Raw.TxHash.String()),
					zap.Uint64("block_number.current", evLog.Raw.BlockNumber),
					zap.Float64("price.fluctuation", change),
				)
				// only handle breaking if the price decreases (change < 0)
				// and if the change is outside the break precentage window
				// to do this we calculate the absolute value of the change percentage
				if change < 0 && math.Abs(change) >= breakPercentage {
					wc.logger.Warn(
						"price fluctuation outside of acceptable bounds, breaking!",
					)
					// lock the authorizer since bind.TransactOpts is not threadsafe
					authorizer.Lock()
					gasPrice, err := utils.GetGasPrice(ctx, wc.backend)
					if err != nil {
						wc.logger.Error("failed to suggest gasprice", zap.Error(err))
					} else {
						wc.setPublicSwap(ctx, authorizer, poolAddress, gasPrice, wc.name)
					}
					// we need to unset the gas price that we overrode the transactor with
					// so that future uses of this transactor have the gas price set to nil
					authorizer.TransactOpts.GasPrice = nil
					authorizer.Unlock()

					if err := alerter.NotifyCircuitBreak(
						fmt.Sprintf(
							"a circuit break is in progress. pool: %s, tokenIn: %s, tokenOut: %s, changePercent: %0.2f, reason: spot price fluctuation",
							wc.name, evLog.TokenIn.String(), evLog.TokenOut.String(), change,
						),
					); err != nil {
						wc.logger.Error("circuit break notification failed")
					}
				}
			}

		}
	}
}

// NotifChan returns events when listened on, this is used to notify any
// listening goroutines of received events without having to worry about multi channel listeners
// as if we have multiple listeners on a channel golang will randomly choose 1
func (wc *WatchedContract) NotifChan() <-chan *sigmacore.SigmacoreLOGSWAP {
	return wc.notifCh
}

// Name returns the name of the indexed pool this watcher monitors
func (wc *WatchedContract) Name() string {
	return wc.name
}

func (wc *WatchedContract) getPreviousSpotPrice(
	db *database.Database,
	evLog *sigmacore.SigmacoreLOGSWAP,
	swapFee *big.Int,
) (*big.Int, error) {
	var (
		info database.PoolInfo
		// specifies the maximum number of times we should attempt to retrieve infos from database
		maxAttempts = 10
		err         error
	)
	// TODO(bonedaddy): try a better method this is meant to resolve possible race conditions
	for i := 1; i <= maxAttempts; i++ {
		info, err = db.GetInfo(wc.name, evLog.Raw.BlockNumber-uint64(i))
		if err != nil && i == maxAttempts {
			return nil, err
		} else if err != nil {
			wc.logger.Error(
				"failed to find record, trying a different block",
				zap.Uint64("failed.lock", evLog.Raw.BlockNumber-uint64(i)),
				zap.Uint64("next.block", evLog.Raw.BlockNumber-uint64(i+1)),
			)
		} else if err == nil {
			wc.logger.Debug(
				"found database info",
				zap.Uint64("want.block_number", evLog.Raw.BlockNumber-1),
				zap.Uint64("found.block_number", evLog.Raw.BlockNumber-uint64(i)),
			)
			break
		}
	}

	// get balance information for previous block
	prevTokenInBalanceStr, ok := info.Balances[wc.tokenNames[evLog.TokenIn]].(string)
	if !ok {
		return nil, errors.New("token in balance is not of type string")
	}
	prevTokenInBalance, ok := new(big.Int).SetString(prevTokenInBalanceStr, 10)
	if !ok {
		return nil, errors.New("failed to convert token in balance from string to big.Int")
	}
	prevTokenOutBalanceStr, ok := info.Balances[wc.tokenNames[evLog.TokenOut]].(string)
	if !ok {
		return nil, errors.New("token out balance is not of type string")
	}
	prevTokenOutBalance, ok := new(big.Int).SetString(prevTokenOutBalanceStr, 10)
	if !ok {
		return nil, errors.New("failed to convert token out balance from string to big.Int")
	}

	// get weight information for previous block
	prevTokenInWeightStr, ok := info.DenormalizedWeights[wc.tokenNames[evLog.TokenIn]].(string)
	if !ok {
		return nil, errors.New("token in weight is not of type string")
	}
	prevTokenInWeight, ok := new(big.Int).SetString(prevTokenInWeightStr, 10)
	if !ok {
		return nil, errors.New("failed to convert token in weight from string to big.Int")
	}
	prevTokenOutWeightStr, ok := info.DenormalizedWeights[wc.tokenNames[evLog.TokenOut]].(string)
	if !ok {
		return nil, errors.New("token out weight is not of type string")
	}
	prevTokenOutWeight, ok := new(big.Int).SetString(prevTokenOutWeightStr, 10)
	if !ok {
		return nil, errors.New("failed to convert token  out weight from string to big.Int")
	}
	return utils.CalcSpotPrice(
		prevTokenInBalance,
		prevTokenInWeight,
		prevTokenOutBalance,
		prevTokenOutWeight,
		swapFee), nil
}

func (wc *WatchedContract) getCurrentSpotPrice(
	evLog *sigmacore.SigmacoreLOGSWAP,
) (*big.Int, error) {
	return wc.binding.GetSpotPrice(&bind.CallOpts{
		BlockNumber: new(big.Int).SetUint64(evLog.Raw.BlockNumber),
	}, evLog.TokenIn, evLog.TokenOut)
}

// getPreviousSwapFee returns the swap fee at the specific block, however if the error
// contains the message `simulatedBackend cannot access blocks other than the latest block`
// we return a spoofed swap fee as this is indicative of simulated backend problems
func (wc *WatchedContract) getPreviousSwapFee(blockNum uint64) (*big.Int, error) {
	prevSwapFee, err := wc.binding.GetSwapFee(&bind.CallOpts{
		BlockNumber: new(big.Int).SetUint64(blockNum),
	})
	if err != nil {
		if wc.isSimulation {
			return utils.ToWei("0.025", 18), nil
		}
		return nil, err
	}
	return prevSwapFee, nil

}

func (wc *WatchedContract) setPublicSwap(
	ctx context.Context,
	auth *utils.Authorizer,
	poolAddress common.Address,
	gasPrice *big.Int, poolName string,
) {
	auth.GasPrice = gasPrice
	if tx, err := wc.controller.SetPublicSwap(auth.TransactOpts, poolAddress, true); err != nil {
		wc.logger.Error(
			"failed to broadcast public swap disable transaction",
			zap.Error(err),
			zap.String("pool", poolName),
		)
	} else {
		// TODO: at the moment due to simulation backend usage
		// this is very hard to test, however once we have a deployment
		// on a testnet i will be able to modify this package to follow
		// the same pattern as the service package does
		wc.logger.Warn(
			"public swap disable transaction sent. TODO: enable waiting for mining",
			zap.String("pool", poolName),
			zap.String("tx.hash", tx.Hash().String()),
		)
	}
}
