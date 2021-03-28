package circuit

import (
	"context"
	"math/big"

	dutils "github.com/bonedaddy/go-defi/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/indexed-finance/circuit-breaker/bindings/controller"
	"github.com/indexed-finance/circuit-breaker/utils"
	"go.uber.org/zap"
)

// BreakConfig provides a configuration to use for breaking circuits
type BreakConfig struct {
	Ctx           context.Context
	BC            dutils.Blockchain
	MinimumGwei   *big.Int
	GasMultiplier *big.Int
	Auth          *utils.Authorizer
	PoolAddress   common.Address
	PoolName      string
	Logger        *zap.Logger
	Breaker       *controller.Controller
}

func BreakCircuit(cfg BreakConfig) error {
	gasPrice, err := utils.GetGasPrice(cfg.Ctx, cfg.BC, cfg.MinimumGwei, cfg.GasMultiplier)
	if err != nil {
		cfg.Logger.Error("failed to suggest gas price", zap.Error(err))
		return err
	}
	cfg.Logger.Info("gas price calculated (includes boost)", zap.String("gas.price", gasPrice.String()))
	cfg.Auth.Lock()
	cfg.Auth.GasPrice = gasPrice
	tx, err := cfg.Breaker.SetPublicSwap(cfg.Auth.TransactOpts, cfg.PoolAddress, false)
	cfg.Auth.GasPrice = nil
	cfg.Auth.Unlock()
	if err != nil {
		cfg.Logger.Error("failed to broadcast public swap disable tx", zap.Error(err))
		return err
	}
	cfg.Logger.Warn(
		"public swap disable transaction sent",
		zap.String("tx.hash", tx.Hash().String()),
		zap.String("pool", cfg.PoolName),
	)
	rcpt, err := bind.WaitMined(cfg.Ctx, cfg.BC, tx)
	if err != nil {
		cfg.Logger.Error("failed to wait for transaction to be mined", zap.Error(err), zap.String("pool", cfg.PoolName), zap.String("tx.hash", tx.Hash().String()))
		return err
	}
	if rcpt.Status != 1 {
		cfg.Logger.Warn("public swap tx failed to execute correctly as status code is not 1", zap.String("pool", cfg.PoolName), zap.String("tx.hash", tx.Hash().String()))
	}
	return nil
}
