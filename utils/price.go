package utils

import "math/big"

var (
	// SwapFee 2.5% in wei
	SwapFee = big.NewInt(25000000000000000)
)

// CalcSpotPrice is used to calculate the spot price of a given asset
func CalcSpotPrice(
	tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, swapFee *big.Int,
) *big.Int {
	numer := Bdiv(tokenBalanceIn, tokenWeightIn)
	denom := Bdiv(tokenBalanceOut, tokenWeightOut)
	ratio := Bdiv(numer, denom)
	scale := Bdiv(BONE(), Bsub(BONE(), swapFee))
	return Bmul(ratio, scale)
}
