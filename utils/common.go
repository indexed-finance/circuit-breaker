package utils

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Breaker is a minimal interface needed for circuit breaker
// designed to allow for easier testing
type Breaker interface {
	SetPublicSwap(opts *bind.TransactOpts, pool common.Address, enabled bool) (*types.Transaction, error)
	CircuitBreaker(opts *bind.CallOpts) (common.Address, error)
}

// GetGasPrice returns a gas price as reported by the oracle, overriding it if it is lower
// than the minimum, while also applying a multiplier
func GetGasPrice(ctx context.Context, backend bind.ContractBackend, minimum, multiplier *big.Int) (*big.Int, error) {
	reported, err := backend.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}
	return BoostGasPrice(reported, minimum, multiplier), nil
}

// BoostGasPrice is used to override the reported gas price if it is lower than the minimum
// while also multiplying the overall gas price to ensure faster block inclusion
func BoostGasPrice(reportedPrice *big.Int, minimum *big.Int, multiplier *big.Int) *big.Int {
	// override if the reported price from the oracle is too low
	if reportedPrice.Cmp(minimum) == -1 {
		reportedPrice = new(big.Int).Set(minimum)
	}
	// mutiply the price to ensure faster block inclusion
	return new(big.Int).Mul(reportedPrice, multiplier)

}
