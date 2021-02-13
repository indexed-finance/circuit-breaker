package multicall

import (
	"context"
	"os"
	"testing"

	"github.com/bonedaddy/bdsm/testenv"
	"github.com/bonedaddy/go-indexed/bclient"
	"github.com/indexed-finance/circuit-breaker/bindings/logswap"
	"github.com/indexed-finance/circuit-breaker/bindings/multicall"
	"github.com/stretchr/testify/require"
)

var (
	infuraAPIKey             = os.Getenv("INFURA_API_KEY")
	defi5ContractAddress     = "0xfa6de2697d59e88ed7fc4dfe5a33dac43565ea41"
	multicallContractAddress = "0x3067b1b7bf344027c7439509fbdf344eb25f5991"
	defi5Tokens              = []string{
		"0xc00e94Cb662C3520282E6f5717214004A7f26888",
		"0xC011a73ee8576Fb46F5E1c5751cA3B9Fe0af2a6F",
		"0xD533a949740bb3306d119CC777fa900bA034cd52",
		"0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984",
		"0x7Fc66500c84A76Ad7e9c93437bFc5Ac33E2DDaE9",
	}
)

func TestMulticallReal(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	bc, err := bclient.NewInfuraClient(infuraAPIKey, true)
	require.NoError(t, err)
	caller, err := New(ctx, multicallContractAddress, bc.EthClient())
	require.NoError(t, err)
	var defi5TokensReversed []string
	for i := len(defi5Tokens) - 1; i >= 0; i-- {
		defi5TokensReversed = append(defi5TokensReversed, defi5Tokens[i])
	}
	// make sure reversal worked
	require.NotEqual(t, defi5Tokens, defi5TokensReversed)

	t.Run("GetDenormalizedWeights", func(t *testing.T) {
		type args struct {
			pool  string
			addrs []string
		}
		tests := []struct {
			name      string
			args      args
			wantCount int
		}{
			{"DEFI5", args{
				defi5ContractAddress,
				defi5Tokens,
			}, 5},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// get the current block number
				block, err := bc.CurrentBlock()
				require.NoError(t, err)
				addrs, values, err := caller.GetDenormalizedWeights(
					block,
					tt.args.pool,
					tt.args.addrs,
				)
				require.NoError(t, err)
				require.Len(t, values, tt.wantCount)
				require.Len(t, addrs, tt.wantCount)
			})
		}
	})

	t.Run("GetBalances", func(t *testing.T) {
		type args struct {
			pool  string
			addrs []string
		}
		tests := []struct {
			name      string
			args      args
			wantCount int
		}{
			{"DEFI5", args{
				defi5ContractAddress,
				defi5Tokens,
			}, 5},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// get the current block number
				block, err := bc.CurrentBlock()
				require.NoError(t, err)
				addrs, values, err := caller.GetBalances(
					block,
					tt.args.pool,
					tt.args.addrs,
				)
				require.NoError(t, err)
				require.Len(t, values, tt.wantCount)
				require.Len(t, addrs, tt.wantCount)
			})
		}
	})

	t.Run("GetUsedBalances", func(t *testing.T) {
		type args struct {
			pool  string
			addrs []string
		}
		tests := []struct {
			name      string
			args      args
			wantCount int
		}{
			{"DEFI5", args{
				defi5ContractAddress,
				defi5Tokens,
			}, 5},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// get the current block number
				block, err := bc.CurrentBlock()
				require.NoError(t, err)
				addrs, values, err := caller.GetUsedBalances(
					block,
					tt.args.pool,
					tt.args.addrs,
				)
				require.NoError(t, err)
				require.Len(t, values, tt.wantCount)
				require.Len(t, addrs, tt.wantCount)
			})
		}
	})

	t.Run("GetSpotPrices", func(t *testing.T) {
		type args struct {
			pool  string
			addrs []string
		}
		tests := []struct {
			name      string
			args      args
			wantCount int
		}{
			{"DEFI5", args{
				defi5ContractAddress,
				defi5Tokens,
			}, 5},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// get the current block number
				block, err := bc.CurrentBlock()
				require.NoError(t, err)
				addrsIn, addrsOut, values, err := caller.GetSpotPrices(
					block,
					tt.args.pool,
					tt.args.addrs,       // in
					defi5TokensReversed, // out
				)
				require.NoError(t, err)
				require.Len(t, values, tt.wantCount)
				require.Len(t, addrsIn, tt.wantCount)
				require.Len(t, addrsOut, tt.wantCount)
			})
		}
	})

	t.Run("GetTotalSupplies", func(t *testing.T) {
		type args struct {
			addrs []string
		}
		tests := []struct {
			name      string
			args      args
			wantCount int
		}{
			{"DEFI5", args{
				defi5Tokens,
			}, 5},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// get the current block number
				block, err := bc.CurrentBlock()
				require.NoError(t, err)
				addrs, values, err := caller.GetTotalSupplies(
					block,
					tt.args.addrs,
				)
				require.NoError(t, err)
				require.Len(t, values, tt.wantCount)
				require.Len(t, addrs, tt.wantCount)
			})
		}
	})

	caller.Close()
}

func TestMulticallSimulated(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	tenv, err := testenv.NewBlockchain(ctx)
	require.NoError(t, err)
	addr, tx, _, err := multicall.DeployMulticall(tenv.Auth, tenv)
	require.NoError(t, err)
	_, err = tenv.DoWaitDeployed(tx, "multicall deployment")
	require.NoError(t, err)
	lAddr, tx, _, err := logswap.DeployLogswap(tenv.Auth, tenv)
	require.NoError(t, err)
	_, err = tenv.DoWaitDeployed(tx, "logswap deployment")
	require.NoError(t, err)
	_ = lAddr

	require.NoError(t, err)
	caller, err := New(ctx, addr.String(), tenv)
	require.NoError(t, err)
	_ = caller

	t.Run("GetDenormalizedWeights", func(t *testing.T) {
		type args struct {
			pool  string
			addrs []string
		}
		tests := []struct {
			name      string
			args      args
			wantCount int
		}{
			{"1-Addrs", args{
				lAddr.String(),
				[]string{lAddr.String()},
			}, 1},
			{"2-Addrs", args{
				lAddr.String(),
				[]string{lAddr.String(), lAddr.String()},
			}, 2},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				addrs, values, err := caller.GetDenormalizedWeights(
					tenv.Blockchain().CurrentBlock().NumberU64(),
					tt.args.pool,
					tt.args.addrs,
				)
				require.NoError(t, err)
				require.Len(t, values, tt.wantCount)
				require.Len(t, addrs, tt.wantCount)

			})
		}
	})

	t.Run("GetBalances", func(t *testing.T) {
		type args struct {
			pool  string
			addrs []string
		}
		tests := []struct {
			name      string
			args      args
			wantCount int
		}{
			{"1-Addrs", args{
				lAddr.String(),
				[]string{lAddr.String()},
			}, 1},
			{"2-Addrs", args{
				lAddr.String(),
				[]string{lAddr.String(), lAddr.String()},
			}, 2},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				addrs, values, err := caller.GetBalances(
					tenv.Blockchain().CurrentBlock().NumberU64(),
					tt.args.pool,
					tt.args.addrs,
				)
				require.NoError(t, err)
				require.Len(t, values, tt.wantCount)
				require.Len(t, addrs, tt.wantCount)

			})
		}
	})

	t.Run("GetUsedBalances", func(t *testing.T) {
		type args struct {
			pool  string
			addrs []string
		}
		tests := []struct {
			name      string
			args      args
			wantCount int
		}{
			{"1-Addrs", args{
				lAddr.String(),
				[]string{lAddr.String()},
			}, 1},
			{"2-Addrs", args{
				lAddr.String(),
				[]string{lAddr.String(), lAddr.String()},
			}, 2},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				addrs, values, err := caller.GetUsedBalances(
					tenv.Blockchain().CurrentBlock().NumberU64(),
					tt.args.pool,
					tt.args.addrs,
				)
				require.NoError(t, err)
				require.Len(t, values, tt.wantCount)
				require.Len(t, addrs, tt.wantCount)

			})
		}
	})

	t.Run("GetSpotPrices", func(t *testing.T) {
		type args struct {
			pool     string
			addrsIn  []string
			addrsOut []string
		}
		tests := []struct {
			name      string
			args      args
			wantCount int
		}{
			{"1-Addrs", args{
				lAddr.String(),
				[]string{lAddr.String()},
				[]string{lAddr.String()},
			}, 1},
			{"2-Addrs", args{
				lAddr.String(),
				[]string{lAddr.String(), lAddr.String()},
				[]string{lAddr.String(), lAddr.String()},
			}, 2},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				addrsIn, addrsOut, values, err := caller.GetSpotPrices(
					tenv.Blockchain().CurrentBlock().NumberU64(),
					tt.args.pool,
					tt.args.addrsIn,
					tt.args.addrsOut,
				)
				require.NoError(t, err)
				require.Len(t, values, tt.wantCount)
				require.Len(t, addrsIn, tt.wantCount)
				require.Len(t, addrsOut, tt.wantCount)
			})
		}
	})

	caller.Close()
}
