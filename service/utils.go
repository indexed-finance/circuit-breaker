package service

import (
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/indexed-finance/circuit-breaker/bindings/sigmacore"
	"github.com/indexed-finance/circuit-breaker/config"
	"github.com/indexed-finance/circuit-breaker/utils"
	"go.uber.org/zap"
)

// GetBalancesWeightsAndSupplies returns the balance, weights
// and total supplies of all tokens held by an IndexPool
// all keys of the maps are the symbols of a tokein lower case
func (s *Service) GetBalancesWeightsAndSupplies(
	contract *sigmacore.Sigmacore,
	block uint64,
	poolAddress string,
	tokens map[string]common.Address,
) (
	balances map[string]interface{},
	denormWeights map[string]interface{},
	totalSupplies map[string]interface{},
	err error,
) {
	var (
		toks []string
		// addr -> symbol
		tokSymbols = make(map[common.Address]string)
	)
	balances = make(map[string]interface{})
	denormWeights = make(map[string]interface{})
	totalSupplies = make(map[string]interface{})

	for symbol, tokAddr := range tokens {
		toks = append(toks, tokAddr.String())
		tokSymbols[tokAddr] = strings.ToLower(symbol)
	}
	bundle, err := s.mc.GetBundle(block, poolAddress, toks)
	if err != nil {
		// try one last time after a sleep, this is sometimes a temporary error
		if strings.Contains(err.Error(), "header not found") {
			s.logger.Warn("header not found error encountered, retrying after sleep")
			time.Sleep(time.Millisecond * 750)
			bundle, err = s.mc.GetBundle(block, poolAddress, toks)
		}
		if err != nil {
			s.logger.Error("failed to get bundle after a retry", zap.Error(err))
			return
		}
	}
	for i, addr := range bundle.Tokens {
		symbol := tokSymbols[addr]
		// if 0 skip as this means it is an uninitialized token
		if bundle.Balances[i].Cmp(big.NewInt(0)) == 0 {
			s.logger.Warn("encountered token with balance of 0 removing", zap.String("symbol", symbol))
			continue
		}
		// if 0 skip as this means it is an uninitialized token
		if bundle.DenormalizedWeights[i].Cmp(big.NewInt(0)) == 0 {
			s.logger.Warn("encountered token with weight less than or equal to 0", zap.String("symbol", symbol))
			continue
		}
		balances[symbol] = bundle.Balances[i].String()
		denormWeights[symbol] = bundle.DenormalizedWeights[i].String()
		totalSupplies[symbol] = bundle.TotalSupplies[i].String()
	}
	return
}

// UpdateBalancesWeightsAndSupplies is used to update all balances, weights, and supplies for a given pool
// Note: If using from many different threads, caller will want to make sure to call the service named locks
func (s *Service) UpdateBalancesWeightsAndSupplies(
	block uint64,
	pool config.Pool,
	contract *sigmacore.Sigmacore,
) (
	balances map[string]interface{},
	denormWeights map[string]interface{},
	totalSupplies map[string]interface{},
	err error,
) {
	tokens, err := utils.PoolTokensFor(contract, s.ew.BC())
	if err != nil {
		s.logger.Error("failed to get pool tokens", zap.Error(err))
		return
	}
	balances, denormWeights, totalSupplies, err = s.GetBalancesWeightsAndSupplies(contract, block, pool.ContractAddress, tokens)
	if err != nil {
		s.logger.Error("GetBalancesAndWeight call failed", zap.Error(err), zap.String("pool", pool.Name))
		return
	}
	err = s.db.RecordInfo(pool.Name, block, balances, denormWeights, totalSupplies)
	return
}
