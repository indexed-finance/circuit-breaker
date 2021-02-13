package utils

import "math/big"

// GetChangePercentageBig returns the change percentage for two big.Int ypes
func GetChangePercentageBig(final, start *big.Int) float64 {
	// (final - start) / start
	// OR should we do (final - start) / |start| (|..| = absolute, this is what we currently do in the indexed discord app)
	fFinal, _ := ToDecimal(final, 18).Float64() // TODO: handle variable decimals
	fStart, _ := ToDecimal(start, 18).Float64() // TODO: handle variable decimals
	return (fFinal - fStart) / fStart
}
