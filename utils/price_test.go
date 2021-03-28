package utils

import (
	"context"
	"math/big"
	"testing"

	"github.com/bonedaddy/go-defi/testenv"
	bmat "github.com/indexed-finance/circuit-breaker/bindings/bmath"
	"github.com/stretchr/testify/require"
)

func TestCalcSpotPrice(t *testing.T) {
	tenv, err := testenv.NewBlockchain(context.Background())
	require.NoError(t, err)
	_, tx, bmath, err := bmat.DeployBmat(tenv.Auth, tenv)
	require.NoError(t, err)
	_, err = tenv.DoWaitDeployed(tx)
	require.NoError(t, err)
	type args struct {
		tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut *big.Int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"0", args{
				big.NewInt(1250000000000000000), // balanceIn
				big.NewInt(15),
				big.NewInt(980000000000000000), // balanceOut
				big.NewInt(90),
			},
		},
		{
			"1", args{
				big.NewInt(1250000000000000000), // balanceIn
				big.NewInt(90),
				big.NewInt(980000000000000000), // balanceOut
				big.NewInt(15),
			},
		},
		{
			"0", args{
				big.NewInt(123456789), // balanceIn
				big.NewInt(15),
				big.NewInt(987654321), // balanceOut
				big.NewInt(90),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := CalcSpotPrice(tt.args.tokenBalanceIn, tt.args.tokenWeightIn, tt.args.tokenBalanceOut, tt.args.tokenWeightIn, SwapFee)
			ret, err := bmath.CalcSpotPrice(nil, tt.args.tokenBalanceIn, tt.args.tokenWeightIn, tt.args.tokenBalanceOut, tt.args.tokenWeightIn, SwapFee)
			require.NoError(t, err)
			require.Equal(t, 0, res.Cmp(ret))
		})
	}
}
