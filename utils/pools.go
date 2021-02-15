package utils

import (
	"fmt"
	"sync"
	"time"

	"github.com/bonedaddy/go-indexed/bindings/erc20"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/indexed-finance/circuit-breaker/bindings/sigmacore"
)

var (
	// symbols maps token address -> token symbol
	symbols  = make(map[common.Address]string)
	poolLock sync.RWMutex
)

// PoolTokensFor is similar go-indexed/bclient's version but for easier testing
func PoolTokensFor(bindings *sigmacore.Sigmacore, backend bind.ContractBackend) (map[string]common.Address, error) {
	tokens, err := bindings.GetCurrentTokens(nil)
	if err != nil {
		return nil, err
	}
	var out = make(map[string]common.Address)
	for _, token := range tokens {
		// check to see if we have this symbol cached already
		poolLock.RLock()
		symbol := symbols[token]
		// if not empty it means we have cached the results so use cached result
		if symbol != "" {
			out[symbols[token]] = token
			poolLock.RUnlock()
			continue
		}
		poolLock.RUnlock()
		ec, err := erc20.NewErc20(token, backend)
		if err != nil {
			return nil, err
		}
		// get the token name
		// ignore error as some tokens such as maker cause this problem
		// https://github.com/ethereum/go-ethereum/issues/21754#issuecomment-716231021
		symbol, _ = ec.Symbol(nil)
		if symbol == "" {
			symbol = guessTokenSymbol(token.String())
		}
		out[symbol] = token
		// set the symbol in our cache
		poolLock.Lock()
		symbols[token] = symbol
		poolLock.Unlock()
	}
	return out, nil
}

func guessTokenSymbol(address string) string {
	if address == "0x9f8F72aA9304c8B593d555F12eF6589cC3A579A2" {
		return "MKR"
	}
	return fmt.Sprintf("unknown-%v-%v", time.Now().UnixNano(), len(address))
}
