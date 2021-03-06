package service

import (
	"errors"
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
	balances, err = s.GetBalances(contract, block, poolAddress, tokens)
	if err != nil {
		// try one last time after a sleep, this is sometimes a temporary error
		if strings.Contains(err.Error(), "header not found") {
			s.logger.Warn("header not found error encountered, retrying after sleep")
			time.Sleep(time.Millisecond * 750)
			balances, err = s.GetBalances(contract, block, poolAddress, tokens)
		}
		if err != nil {
			s.logger.Error("failed to get balances after after a retry", zap.Error(err))
			return
		}
	}
	denormWeights, err = s.GetWeights(contract, block, poolAddress, tokens)
	if err != nil {
		// try one last time after a sleep, this is sometimes a temporary error
		if strings.Contains(err.Error(), "header not found") {
			s.logger.Warn("header not found error encountered, retrying after sleep")
			time.Sleep(time.Millisecond * 750)
			denormWeights, err = s.GetWeights(contract, block, poolAddress, tokens)
		}
		if err != nil {
			s.logger.Error("failed to get weights after a retry", zap.Error(err))
			return
		}
	}
	totalSupplies, err = s.GetTotalSupplies(contract, block, tokens)
	if err != nil {
		// try one last time after a sleep, this is sometimes a temporary error
		if strings.Contains(err.Error(), "header not found") {
			s.logger.Warn("header not found error encountered, retrying after sleep")
			time.Sleep(time.Millisecond * 750)
			totalSupplies, err = s.GetTotalSupplies(contract, block, tokens)
		}
		if err != nil {
			s.logger.Error("failed to get total supplies after a retry", zap.Error(err))
			return
		}
	}
	for symbol, value := range balances {
		valBig, ok := new(big.Int).SetString(value.(string), 10)
		if !ok {
			s.logger.Error("failed to parse balance", zap.String("symbol", symbol))
			err = errors.New("invalid balance")
			return
		}
		if valBig.Cmp(big.NewInt(0)) == 0 {
			s.logger.Warn("encountered token with balance of 0 removing", zap.String("symbol", symbol))
			delete(balances, symbol)
			delete(denormWeights, symbol)
			delete(totalSupplies, symbol)
		}
	}
	for symbol, value := range denormWeights {
		valBig, ok := new(big.Int).SetString(value.(string), 10)
		if !ok {
			s.logger.Error("failed to parse denormalized weight", zap.String("symbol", symbol))
			err = errors.New("invalid weight")
			return
		}
		if valBig.Cmp(big.NewInt(0)) == 0 {
			s.logger.Warn("encountered token with weight less than or equal to 0", zap.String("symbol", symbol))
			delete(balances, symbol)
			delete(denormWeights, symbol)
			delete(totalSupplies, symbol)
		}
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
	tokens, err := utils.PoolTokensFor(contract, s.ew.BC().EthClient())
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

// GetWeights is used to return a given pool's token weights
func (s *Service) GetWeights(
	contract *sigmacore.Sigmacore,
	block uint64,
	poolAddress string,
	tokens map[string]common.Address,
) (map[string]interface{}, error) {
	var tokenAddress []string
	var names = make(map[common.Address]string)
	for name, addr := range tokens {
		tokenAddress = append(tokenAddress, addr.String())
		names[addr] = strings.ToLower(name)
	}
	addrs, values, err := s.mc.GetDenormalizedWeights(block, poolAddress, tokenAddress)
	if err != nil {
		return nil, err
	}

	weights := make(map[string]interface{})
	for i, addr := range addrs {
		weights[names[addr]] = values[i].String()
	}
	return weights, nil
}

// GetBalances returns the balances of various pool tokens
// This is not to be confused with IndexPool token balances
func (s *Service) GetBalances(
	contract *sigmacore.Sigmacore,
	block uint64,
	poolAddress string,
	tokens map[string]common.Address,
) (map[string]interface{}, error) {
	var tokenAddress []string
	var names = make(map[common.Address]string)
	for name, addr := range tokens {
		tokenAddress = append(tokenAddress, addr.String())
		names[addr] = strings.ToLower(name)
	}
	addrs, values, err := s.mc.GetBalances(block, poolAddress, tokenAddress)
	if err != nil {
		return nil, err
	}
	balances := make(map[string]interface{})
	for i, addr := range addrs {
		balances[names[addr]] = values[i].String()
	}
	return balances, nil
}

// GetTotalSupplies returns the total supplies of all the tokens that are part of the pool
func (s *Service) GetTotalSupplies(
	contract *sigmacore.Sigmacore,
	block uint64,
	tokens map[string]common.Address,
) (map[string]interface{}, error) {
	var tokenAddress []string
	var names = make(map[common.Address]string)
	for name, addr := range tokens {
		tokenAddress = append(tokenAddress, addr.String())
		names[addr] = strings.ToLower(name)
	}
	addrs, values, err := s.mc.GetTotalSupplies(block, tokenAddress)
	if err != nil {
		return nil, err
	}
	supplies := make(map[string]interface{})
	for i, addr := range addrs {
		supplies[names[addr]] = values[i].String()
	}
	return supplies, nil
}
