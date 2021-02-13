package utils

import (
	"context"
	"math/big"
	"testing"

	"github.com/bonedaddy/bdsm/testenv"
	bmat "github.com/indexed-finance/circuit-breaker/bindings/bmath"
	"github.com/stretchr/testify/require"
)

func TestBMath(t *testing.T) {
	tenv, err := testenv.NewBlockchain(context.Background())
	require.NoError(t, err)
	_, tx, bmath, err := bmat.DeployBmat(tenv.Auth, tenv)
	require.NoError(t, err)
	_, err = tenv.DoWaitDeployed(tx)
	require.NoError(t, err)

	t.Run("Constants", func(t *testing.T) {
		// BONE
		cBone, err := bmath.BONE(nil)
		require.NoError(t, err)
		require.Equal(
			t,
			0,
			cBone.Cmp(BONE()),
		)
		// PRECISION
		cPrec, err := bmath.BPOWPRECISION(nil)
		require.NoError(t, err)
		require.Equal(
			t,
			0,
			cPrec.Cmp(BPOWPRECISION),
		)
	})
	t.Run("BTOI", func(t *testing.T) {
		type args struct {
			input *big.Int
		}
		tests := []struct {
			name string
			args args
		}{
			{"2", args{
				new(big.Int).Mul(BONE(), big.NewInt(200000000)),
			}},
			{"4", args{
				new(big.Int).Mul(BONE(), big.NewInt(400000000)),
			}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				res := Btoi(tt.args.input)
				ret, err := bmath.Btoi(nil, tt.args.input)
				require.NoError(t, err)
				require.Equal(t, 0, res.Cmp(ret))
			})
		}
	})
	t.Run("BFLOOR", func(t *testing.T) {
		type args struct {
			input *big.Int
		}
		tests := []struct {
			name string
			args args
		}{
			{"2", args{big.NewInt(200000000)}},
			{"4", args{big.NewInt(400000000)}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				res := Bfloor(tt.args.input)
				ret, err := bmath.Bfloor(nil, tt.args.input)
				require.NoError(t, err)
				require.Equal(t, 0, res.Cmp(ret))
			})
		}
	})
	t.Run("BADD", func(t *testing.T) {
		type args struct {
			a *big.Int
			b *big.Int
		}
		tests := []struct {
			name    string
			args    args
			wantVal *big.Int
		}{
			{"2", args{big.NewInt(20000000000), big.NewInt(20000)}, big.NewInt(4)},
			{"4", args{big.NewInt(40000000000), big.NewInt(40000)}, big.NewInt(8)},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				res := Badd(tt.args.a, tt.args.b)
				ret, err := bmath.Badd(nil, tt.args.a, tt.args.b)
				require.NoError(t, err)
				require.Equal(t, 0, res.Cmp(ret))
			})
		}
	})
	t.Run("BSUB", func(t *testing.T) {
		type args struct {
			a *big.Int
			b *big.Int
		}
		tests := []struct {
			name string
			args args
		}{
			{"2", args{big.NewInt(20000000000), big.NewInt(20)}},
			{"4", args{big.NewInt(40000000000), big.NewInt(40)}},
			{"8", args{big.NewInt(80000000000), big.NewInt(80)}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				res := Bsub(tt.args.a, tt.args.b)
				ret, err := bmath.Bsub(nil, tt.args.a, tt.args.b)
				require.NoError(t, err)
				require.Equal(t, 0, res.Cmp(ret))
			})
		}
	})
	t.Run("BSUBSIGN", func(t *testing.T) {
		type args struct {
			a *big.Int
			b *big.Int
		}
		tests := []struct {
			name string
			args args
		}{
			{"2", args{big.NewInt(20000000000), big.NewInt(20)}},
			{"4", args{big.NewInt(40000000000), big.NewInt(40)}},
			{"8-1", args{big.NewInt(80000000000), big.NewInt(40)}},
			{"8-2", args{big.NewInt(40000000000), big.NewInt(80)}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				res, resb := Bsubsign(tt.args.a, tt.args.b)
				ret, retb, err := bmath.BsubSign(nil, tt.args.a, tt.args.b)
				require.NoError(t, err)
				require.Equal(t, res, ret)
				require.Equal(t, resb, retb)
			})
		}
	})
	t.Run("BMUL", func(t *testing.T) {
		type args struct {
			a *big.Int
			b *big.Int
		}
		tests := []struct {
			name string
			args args
		}{
			{"2", args{new(big.Int).Mul(big.NewInt(20000000000), BONE()), big.NewInt(200)}},
			{"4", args{new(big.Int).Mul(big.NewInt(40000000000), BONE()), big.NewInt(400)}},
			{"8", args{new(big.Int).Mul(big.NewInt(80000000000), BONE()), big.NewInt(800)}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				res := Bmul(tt.args.a, tt.args.b)
				ret, err := bmath.Bmul(nil, tt.args.a, tt.args.b)
				require.NoError(t, err)
				require.Equal(t, res, ret)
			})
		}
	})
	t.Run("BDIV", func(t *testing.T) {
		type args struct {
			a *big.Int
			b *big.Int
		}
		tests := []struct {
			name string
			args args
		}{
			{"2", args{new(big.Int).Mul(big.NewInt(20000000000), big.NewInt(2)), big.NewInt(200)}},
			{"4", args{new(big.Int).Mul(big.NewInt(40000000000), big.NewInt(4)), big.NewInt(400)}},
			{"8", args{new(big.Int).Mul(big.NewInt(80000000000), big.NewInt(8)), big.NewInt(800)}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				res := Bdiv(tt.args.a, tt.args.b)
				ret, err := bmath.Bdiv(nil, tt.args.a, tt.args.b)
				require.NoError(t, err)
				require.Equal(t, res, ret)
			})
		}
	})
	t.Run("BPOWI", func(t *testing.T) {
		// this test doesnt use simulatedbackend due to precision issues
		// https://github.com/ethereum/go-ethereum/issues/22261
		type args struct {
			a *big.Int
			b *big.Int
		}
		tests := []struct {
			name    string
			args    args
			wantVal *big.Int
		}{
			{"1", args{big.NewInt(9999), big.NewInt(1)}, big.NewInt(9999)},
			{"2", args{big.NewInt(9999), big.NewInt(0)}, big.NewInt(1000000000000000000)},
			{"3", args{big.NewInt(823), big.NewInt(1)}, big.NewInt(823)},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				res := Bpowi(tt.args.a, tt.args.b)
				require.Equal(t, 0, res.Cmp(tt.wantVal))
			})
		}
	})
	t.Run("BPOW", func(t *testing.T) {
		type args struct {
			a *big.Int
			b *big.Int
		}
		tests := []struct {
			name string
			args args
		}{
			{"1", args{big.NewInt(9999), big.NewInt(1)}},
			{"2", args{big.NewInt(8313), big.NewInt(2)}},
			{"3", args{big.NewInt(7891), big.NewInt(3)}},
			{"4", args{big.NewInt(1), big.NewInt(9999)}},
			{"5", args{big.NewInt(2), big.NewInt(9999)}},
			{"6", args{big.NewInt(3), big.NewInt(9999)}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				res := Bpow(tt.args.a, tt.args.b)
				ret, err := bmath.Bpow(nil, tt.args.a, tt.args.b)
				require.NoError(t, err)
				require.Equal(t, 0, res.Cmp(ret))
			})
		}
	})
	t.Run("BPOWAPPROX", func(t *testing.T) {
		type args struct {
			base      *big.Int
			exp       *big.Int
			precision *big.Int
		}
		tests := []struct {
			name string
			args args
		}{
			{"1", args{big.NewInt(10), big.NewInt(2), big.NewInt(3)}},
			{"2", args{big.NewInt(9), big.NewInt(3), big.NewInt(4)}},
			{"3", args{big.NewInt(2), big.NewInt(10), big.NewInt(6)}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				res := Bpowapprox(tt.args.base, tt.args.exp, tt.args.precision)
				ret, err := bmath.BpowApprox(nil, tt.args.base, tt.args.exp, tt.args.precision)
				require.NoError(t, err)
				require.Equal(t, 0, res.Cmp(ret))
			})
		}
	})
}
