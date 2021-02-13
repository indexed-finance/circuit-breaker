package utils

import (
	"math/big"
)

// this library provides a golang equivalent of the "BMath" contracts
// see https://github.com/indexed-finance/indexed-core/blob/master/contracts/balancer/BNum.sol

var (
	// BONE ...
	bONE = new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	// BPOWPRECISION ...
	BPOWPRECISION = new(big.Int).Div(
		BONE(),
		new(big.Int).Exp(big.NewInt(10), big.NewInt(10), nil),
	)
)

// BONE ...
func BONE() *big.Int {
	return new(big.Int).Set(bONE)
}

// Btoi ...
func Btoi(a *big.Int) *big.Int {
	return new(big.Int).Div(a, BONE())
}

// Bfloor ...
func Bfloor(a *big.Int) *big.Int {
	return new(big.Int).Mul(Btoi(a), BONE())
}

// Badd ...
func Badd(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

// Bsub ...
func Bsub(a, b *big.Int) *big.Int {
	c, _ := Bsubsign(a, b)
	return c
}

// Bsubsign ...
func Bsubsign(a, b *big.Int) (*big.Int, bool) {
	compared := a.Cmp(b)
	if compared == 0 || compared == 1 {
		return new(big.Int).Sub(a, b), false
	}
	return new(big.Int).Sub(b, a), true
}

// Bmul ...
func Bmul(a, b *big.Int) *big.Int {
	c0 := new(big.Int).Mul(a, b)
	c1 := new(big.Int).Add(
		c0,
		new(big.Int).Div(BONE(), big.NewInt(2)),
	)
	return new(big.Int).Div(c1, BONE())
}

// Bdiv ...
// TODO(bonedaddy): right now we check for division by 0 and return 0 however this may not be good
func Bdiv(a, b *big.Int) *big.Int {
	if b.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0)
	}
	c0 := new(big.Int).Mul(a, BONE())
	c1 := new(big.Int).Add(
		c0,
		new(big.Int).Div(b, big.NewInt(2)),
	)
	return new(big.Int).Div(c1, b)
}

// Bpowi ...
func Bpowi(a, n *big.Int) *big.Int {
	// s if n % 2 is not equal to 0, set z to a otherwise set z to BONE
	// uint256 z = n % 2 != 0 ? a : BONE;
	var z *big.Int
	compared := new(big.Int).Mod(n, big.NewInt(2)).Cmp(big.NewInt(0))
	if compared != 0 {
		z = new(big.Int).Set(a)
	} else {
		z = new(big.Int).Set(BONE())
	}
	for n.Div(n, big.NewInt(2)); n.Cmp(big.NewInt(0)) != 0; n.Div(n, big.NewInt(2)) {
		a = Bmul(a, a)
		if new(big.Int).Mod(n, big.NewInt(2)).Cmp(big.NewInt(0)) != 0 {
			z = Bmul(z, a)
		}
	}
	return z
}

// Bpow ...
func Bpow(base, exp *big.Int) *big.Int {
	// TODO: error check according to
	// https://github.com/indexed-finance/indexed-core/blob/master/contracts/balancer/BNum.sol#L87-L88
	whole := Bfloor(exp)
	remain := Bsub(exp, whole)
	wholePow := Bpowi(base, Btoi(whole))
	if remain.Cmp(big.NewInt(0)) == 0 {
		return wholePow
	}
	partialResult := Bpowapprox(base, remain, BPOWPRECISION)
	return Bmul(wholePow, partialResult)
}

// Bpowapprox ...
func Bpowapprox(base, exp, precision *big.Int) *big.Int {
	a := new(big.Int).Set(exp)
	x, xneg := Bsubsign(base, BONE())
	term := new(big.Int).Set(BONE())
	sum := new(big.Int).Set(term)
	negative := false
	for i := big.NewInt(1); term.Cmp(precision) == 0 || term.Cmp(precision) == 1; i.Add(i, big.NewInt(1)) {
		bigK := new(big.Int).Mul(i, BONE())
		c, cneg := Bsubsign(a, Bsub(bigK, BONE()))
		term = Bmul(term, Bmul(c, x))
		term = Bdiv(term, bigK)
		if term.Cmp(big.NewInt(0)) == 0 {
			break
		}
		if xneg {
			negative = !negative
		}
		if cneg {
			negative = !negative
		}
		if negative {
			sum = Bsub(sum, term)
		} else {
			sum = Badd(sum, term)
		}
	}
	return sum
}
