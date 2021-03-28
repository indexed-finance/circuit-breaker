package circuit

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/bonedaddy/go-defi/testenv"
	"github.com/indexed-finance/circuit-breaker/bindings/controller"
	"github.com/indexed-finance/circuit-breaker/bindings/logswap"
	"github.com/indexed-finance/circuit-breaker/utils"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestBreakCircuit(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	tenv, err := testenv.NewBlockchain(ctx)
	require.NoError(t, err)
	authorizer := utils.NewAuthorizerFromPK(tenv.PK())
	//cfg := testutils.GetConfig(t)
	addr, tx, _, err := logswap.DeployLogswap(tenv.Auth, tenv)
	require.NoError(t, err)
	_, err = tenv.DoWaitDeployed(tx)
	require.NoError(t, err)
	breaker, err := controller.NewController(addr, tenv)
	require.NoError(t, err)
	go func() {
		time.Sleep(time.Second * 5)
		tenv.Commit()
	}()
	require.NoError(t, BreakCircuit(
		BreakConfig{
			Ctx:           ctx,
			BC:            tenv,
			MinimumGwei:   big.NewInt(1000),
			GasMultiplier: big.NewInt(3),
			Auth:          authorizer,
			PoolAddress:   addr,
			PoolName:      "yodawg",
			Logger:        zap.NewNop(),
			Breaker:       breaker,
		},
	))
}
