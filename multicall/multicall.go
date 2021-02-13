package multicall

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	bindings "github.com/indexed-finance/circuit-breaker/bindings/multicall"
)

// Multicall exposes utility functions around the multicall contract
type Multicall struct {
	backend bind.ContractBackend
	binding *bindings.Multicall
	ctx     context.Context
	cancel  context.CancelFunc
}

// New initializes the multicall contract wrapper package
func New(ctx context.Context, addr string, backend bind.ContractBackend) (*Multicall, error) {
	ctx, cancel := context.WithCancel(ctx)
	binding, err := bindings.NewMulticall(common.HexToAddress(addr), backend)
	if err != nil {
		cancel()
		return nil, err
	}
	return &Multicall{
		backend: backend,
		binding: binding,
		ctx:     ctx,
		cancel:  cancel,
	}, nil
}

// GetDenormalizedWeights returns the denormalized weights of all given tokens
func (m *Multicall) GetDenormalizedWeights(
	block uint64,
	pool string,
	tokens []string,
) ([]common.Address, []*big.Int, error) {
	poolAddr := common.HexToAddress(pool)
	var tokenAddrs = make([]common.Address, len(tokens))
	for i, tok := range tokens {
		tokenAddrs[i] = common.HexToAddress(tok)
	}
	return m.binding.GetDenormalizedWeights(
		&bind.CallOpts{
			Context:     m.ctx,
			BlockNumber: new(big.Int).SetUint64(block),
		},
		poolAddr,
		tokenAddrs,
	)
}

// GetBalances returns the balances of various tokens
func (m *Multicall) GetBalances(
	block uint64,
	pool string,
	tokens []string,
) ([]common.Address, []*big.Int, error) {
	poolAddr := common.HexToAddress(pool)
	var tokenAddrs = make([]common.Address, len(tokens))
	for i, tok := range tokens {
		tokenAddrs[i] = common.HexToAddress(tok)
	}
	return m.binding.GetBalances(
		&bind.CallOpts{
			Context:     m.ctx,
			BlockNumber: new(big.Int).SetUint64(block),
		},
		poolAddr,
		tokenAddrs,
	)
}

// GetUsedBalances returns the used balances of various tokens
func (m *Multicall) GetUsedBalances(
	block uint64,
	pool string,
	tokens []string,
) ([]common.Address, []*big.Int, error) {
	poolAddr := common.HexToAddress(pool)
	var tokenAddrs = make([]common.Address, len(tokens))
	for i, tok := range tokens {
		tokenAddrs[i] = common.HexToAddress(tok)
	}
	return m.binding.GetUsedBalances(
		&bind.CallOpts{
			Context:     m.ctx,
			BlockNumber: new(big.Int).SetUint64(block),
		},
		poolAddr,
		tokenAddrs,
	)
}

// GetSpotPrices returns the spot prices of various in/out pairs
func (m *Multicall) GetSpotPrices(
	block uint64,
	pool string,
	tokensIn []string,
	tokensOut []string,
) ([]common.Address, []common.Address, []*big.Int, error) {
	poolAddr := common.HexToAddress(pool)
	var (
		tokenInAddrs  = make([]common.Address, len(tokensIn))
		tokenOutAddrs = make([]common.Address, len(tokensOut))
	)
	for i, tok := range tokensIn {
		tokenInAddrs[i] = common.HexToAddress(tok)
	}
	for i, tok := range tokensOut {
		tokenOutAddrs[i] = common.HexToAddress(tok)
	}
	return m.binding.GetSpotPrices(
		&bind.CallOpts{
			Context:     m.ctx,
			BlockNumber: new(big.Int).SetUint64(block),
		},
		poolAddr,
		tokenInAddrs,
		tokenOutAddrs,
	)
}

// GetTotalSupplies returns the total supplies of various tokens
func (m *Multicall) GetTotalSupplies(
	block uint64,
	tokens []string,
) ([]common.Address, []*big.Int, error) {
	var tokenAddrs = make([]common.Address, len(tokens))
	for i, tok := range tokens {
		tokenAddrs[i] = common.HexToAddress(tok)
	}
	return m.binding.GetTotalSupplies(
		&bind.CallOpts{
			Context:     m.ctx,
			BlockNumber: new(big.Int).SetUint64(block),
		},
		tokenAddrs,
	)
}

// Close terminates the context
func (m *Multicall) Close() {
	m.cancel()
}
