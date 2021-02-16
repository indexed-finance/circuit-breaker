package utils

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Breaker is a minimal interface needed for circuit breaker
// designed to allow for easier testing
type Breaker interface {
	SetPublicSwap(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error)
}

// GetGasPrice returns a gas price as reported by the oracle
// but multiplies by 3 to ensure speedy transactions
func GetGasPrice(ctx context.Context, ec *ethclient.Client) (*big.Int, error) {
	gasPrice, err := ec.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}
	return new(big.Int).Mul(gasPrice, big.NewInt(3)), nil
}
