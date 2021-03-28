package eventwatcher

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"sync"

	dutils "github.com/bonedaddy/go-defi/utils"

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
	name    string
	binding *sigmacore.Sigmacore

	notifSwapCh   chan *sigmacore.SigmacoreLOGSWAP              // duplicate of swapCh to send events to watchers
	notifToggleCh chan *sigmacore.SigmacoreLOGPUBLICSWAPTOGGLED // duplicate of toggleCh to send events to watchers

	logger         *zap.Logger
	tokenAddresses map[string]common.Address
	tokenNames     map[common.Address]string
	isSimulation   bool // indicates if this is a simulation for more refined control over parameters
	backend        bind.ContractBackend
	controller     *controller.Controller
	// the breaker has no public contract right now
	// but it has the same public swap function so we
	// can reuse the contract binding for it
	breaker       *controller.Controller
	minimumGwei   *big.Int
	gasMultiplier *big.Int

	breakLock sync.Mutex
	breaking  bool

	brokenLock sync.Mutex
	broken     bool

	toggleCh  chan *sigmacore.SigmacoreLOGPUBLICSWAPTOGGLED
	toggleSub event.Subscription

	swapCh  chan *sigmacore.SigmacoreLOGSWAP
	swapSub event.Subscription

	bc dutils.Blockchain
}

// NewWatchedContracts initializes watched contracts and prepares them for event listening
func (ew *EventWatcher) NewWatchedContracts(
	logger *zap.Logger,
	bindings map[string]*sigmacore.Sigmacore,
	minimumGwei, gasMultiplier *big.Int,
) ([]*WatchedContract, error) {
	out := make([]*WatchedContract, 0)
	for name, contract := range bindings {
		swapCh := make(chan *sigmacore.SigmacoreLOGSWAP, 100) // todo: we may want to change this in the future
		swapSub, err := contract.WatchLOGSWAP(nil, swapCh, nil, nil, nil)
		if err != nil {
			return nil, err
		}
		toggleCh := make(chan *sigmacore.SigmacoreLOGPUBLICSWAPTOGGLED, 100)
		toggleSub, err := contract.WatchLOGPUBLICSWAPTOGGLED(nil, toggleCh)
		if err != nil {
			return nil, err
		}
		// get the current pool tokens
		// this logic is copied from go-indexed/blcient
		// we could use the ew.BC() function but that makes it harder to do simulated testing
		addrs, err := utils.PoolTokensFor(contract, ew.bc)
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
		control, err := controller.NewController(controllerAddress, ew.bc)
		if err != nil {
			return nil, err
		}
		circuitBreakerAddr, err := control.CircuitBreaker(nil)
		if err != nil {
			return nil, err
		}
		breaker, err := controller.NewController(circuitBreakerAddr, ew.bc)
		if err != nil {
			return nil, err
		}
		// lowercase the name
		name = strings.ToLower(name)
		out = append(out, &WatchedContract{
			name:           name,
			binding:        contract,
			swapCh:         swapCh,
			notifSwapCh:    make(chan *sigmacore.SigmacoreLOGSWAP),
			notifToggleCh:  make(chan *sigmacore.SigmacoreLOGPUBLICSWAPTOGGLED),
			swapSub:        swapSub,
			logger:         logger.Named("contract.watcher").With(zap.String("pool", name)),
			tokenAddresses: addrs,
			tokenNames:     names,
			backend:        ew.bc,
			controller:     control,
			breaker:        breaker,
			minimumGwei:    minimumGwei,
			gasMultiplier:  gasMultiplier,
			toggleCh:       toggleCh,
			toggleSub:      toggleSub,
		})
	}
	return out, nil
}

// Listen stars the watche contract listening process waiting for incoming events
func (wc *WatchedContract) Listen(ctx context.Context, db *database.Database, alerter *alerts.Alerter, authorizer *utils.Authorizer, breakPercentage float64) error {
	// register channel closures to prevent possible issues

	wc.logger.Info("contract watcher started listening for events")
	// get the pool contract address
	dbInfo, err := db.GetPool(wc.name)
	if err != nil {
		return err
	}
	poolAddress := common.HexToAddress(dbInfo.ContractAddress)

	// monitor the toggle events in a specific goroutine
	go wc.handleSwapToggles(ctx)

	// this is a blocking function call
	return wc.handleLogSwaps(ctx, db, alerter, authorizer, breakPercentage, poolAddress)
}

// NotifChanSwap returns events when listened on, this is used to notify any
// listening goroutines of received events without having to worry about multi channel listeners
// as if we have multiple listeners on a channel golang will randomly choose 1
func (wc *WatchedContract) NotifChanSwap() <-chan *sigmacore.SigmacoreLOGSWAP {
	return wc.notifSwapCh
}

// NotifChanToggle returns events when listened on, this is used to notify any
// listening goroutines of received events without having to worry about multi channel listeners
// as if we have multiple listeners on a channel golang will randomly choose 1
func (wc *WatchedContract) NotifChanToggle() <-chan *sigmacore.SigmacoreLOGPUBLICSWAPTOGGLED {
	return wc.notifToggleCh
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
	wc.breakLock.Lock()
	if wc.breaking {
		wc.logger.Warn("attempting to trigger circuit break while one is ongoing")
		wc.breakLock.Unlock()
		return
	}
	wc.breaking = true
	wc.breakLock.Unlock()
	auth.GasPrice = gasPrice
	if tx, err := wc.breaker.SetPublicSwap(auth.TransactOpts, poolAddress, false); err != nil {
		wc.logger.Error(
			"failed to broadcast public swap disable transaction",
			zap.Error(err),
			zap.String("pool", poolName),
		)
	} else {
		wc.logger.Warn(
			"public swap disable transaction sent, waiting for transaction to be mined",
			zap.String("pool", poolName),
			zap.String("tx.hash", tx.Hash().String()),
		)
		if rcpt, err := bind.WaitMined(ctx, wc.bc, tx); err != nil {
			wc.logger.Error("failed to wait for transaction to be mined", zap.Error(err), zap.String("pool", poolName), zap.String("tx.hash", tx.Hash().String()))
		} else {
			if rcpt.Status != 1 {
				wc.logger.Warn("public swap transaction failed execution as return status is not 1")
			}
		}
		// unset the breaking boolean
		wc.breakLock.Lock()
		wc.breaking = false
		wc.breakLock.Unlock()

	}
}
