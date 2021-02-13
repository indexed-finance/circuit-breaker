package eventwatcher

import (
	"context"
	"math/big"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/bonedaddy/bdsm/testenv"
	"github.com/bonedaddy/circuit-breaker/alerts"
	"github.com/bonedaddy/circuit-breaker/bindings/logswap"
	"github.com/bonedaddy/circuit-breaker/config"
	"github.com/bonedaddy/circuit-breaker/database"
	"github.com/bonedaddy/circuit-breaker/utils"
	poolbindings "github.com/bonedaddy/go-indexed/bindings/pool"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

var (
	twilioSID           = os.Getenv("TWILIO_SID")
	twilioAuthToken     = os.Getenv("TWILIO_AUTH_TOKEN")
	twilioTestRecipient = os.Getenv("TWILIO_TEST_RECIPIENT")
	twilioNumber        = os.Getenv("TWILIO_NUMBER")
)

func getConfig(t *testing.T) *config.Config {
	exCfg := *config.ExampleConfig
	exCfg.InfuraAPIKey = os.Getenv("INFURA_API_KEY")
	exCfg.Alerts.Twilio.Recipients = []string{twilioTestRecipient}
	exCfg.Alerts.Twilio.From = twilioNumber
	exCfg.Alerts.Twilio.SID = twilioSID
	exCfg.Alerts.Twilio.AuthToken = twilioAuthToken
	return &exCfg
}

func TestWatchedContract(t *testing.T) {
	t.Cleanup(func() {
		os.Remove("circuit-breaker.db")
		os.RemoveAll("testDir")
	})
	os.MkdirAll("testDir", os.ModePerm)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	tenv, err := testenv.NewBlockchain(ctx)
	require.NoError(t, err)
	authorizer := utils.NewAuthorizerFromPK(tenv.PK())
	cfg := getConfig(t)
	db, err := database.New(database.OptsFromConfig(cfg.Database))
	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate())
	addr, tx, logswapper, err := logswap.DeployLogswap(tenv.Auth, tenv)
	require.NoError(t, err)
	_, err = tenv.DoWaitDeployed(tx)
	require.NoError(t, err)
	ew := &EventWatcher{}
	pool, err := poolbindings.NewPoolbindings(addr, tenv)
	require.NoError(t, err)
	/*logger, err := zap.NewDevelopment()
	require.NoError(t, err)*/
	watched, err := ew.NewWatchedContracts(zap.NewNop(), tenv, map[string]*poolbindings.Poolbindings{"cc10": pool})
	require.NoError(t, err)
	// start listene

	watchedContract := watched[0]
	// triggert getPreviousSwapFee error case while simulation is set to false
	_, err = watchedContract.getPreviousSwapFee(tenv.Blockchain().CurrentBlock().NumberU64() - 1)
	require.Error(t, err)
	// set simulation mode for additional testing
	watchedContract.isSimulation = true
	// create fake pool
	require.NoError(t, db.NewPool(
		watchedContract.name,
		addr.String(),
		0,
		map[string]interface{}{"cc10": utils.ToWei("900.0", 18).String()},
		map[string]interface{}{"cc10": utils.ToWei("900.0", 18).String()},
		map[string]interface{}{"cc10": utils.ToWei("900.0", 18).String()},
		map[string]interface{}{"cc10": 18},
	))

	require.NoError(t, db.RecordInfo(
		watchedContract.name,
		uint64(1),
		map[string]interface{}{"cc10": utils.ToWei("900.0", 18).String()},
		map[string]interface{}{"cc10": utils.ToWei("900.0", 18).String()},
		map[string]interface{}{"cc10": utils.ToWei("900.0", 18).String()},
	))

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		watchedContract.Listen(ctx, db, alerts.New(zap.NewNop(), cfg.Alerts), authorizer, 0.1)
	}()

	// we need to spoof some information for easier testing
	watchedContract.tokenNames = map[common.Address]string{tenv.Auth.From: "cc10"}

	t.Run("TestSpotPriceCalculations", func(t *testing.T) {
		// TODO: verify results iunstead of checking correctness
		go func() {
			time.Sleep(time.Second) // need a better way of waiting for this to mine until we are listening
			tx, err = logswapper.EmitLogSwap(tenv.Auth, utils.ToWei("100.0", 18), utils.ToWei("200.0", 18))
			require.NoError(t, err)
			require.NoError(t, tenv.DoWaitMined(tx))
		}()
		evLog1 := <-watchedContract.NotifChan()
		currSpot1, err := watchedContract.getCurrentSpotPrice(evLog1)
		require.NoError(t, err)
		require.NotEqual(t, 0, currSpot1.Cmp(big.NewInt(0)))

		// record fake price to trigger circuit break
		// record fake infos
		require.NoError(t, db.RecordInfo(
			watchedContract.name,
			2,
			map[string]interface{}{"cc10": utils.ToWei("900000000000.0", 18).String()},
			map[string]interface{}{"cc10": utils.ToWei("800000000000.0", 18).String()},
			map[string]interface{}{"cc10": utils.ToWei("900000000000.0", 18).String()},
		))
		require.NoError(t, db.RecordInfo(
			watchedContract.name,
			3,
			map[string]interface{}{"cc10": utils.ToWei("900000000000.0", 18).String()},
			map[string]interface{}{"cc10": utils.ToWei("500000000000.0", 18).String()},
			map[string]interface{}{"cc10": utils.ToWei("10000000.0", 18).String()},
		))
		go func() {
			time.Sleep(time.Second) // need a better way of waiting for this to mine until we are listening
			tx, err = logswapper.EmitLogSwap(tenv.Auth, utils.ToWei("900.0", 18), utils.ToWei("1000.0", 18))
			require.NoError(t, err)
			require.NoError(t, tenv.DoWaitMined(tx))
		}()

		evLog2 := <-watchedContract.NotifChan()
		currSpot2, err := watchedContract.getCurrentSpotPrice(evLog2)
		require.NoError(t, err)
		require.NotEqual(t, 0, currSpot2.Cmp(big.NewInt(0)))

		_, err = watchedContract.getPreviousSpotPrice(db, evLog2, utils.ToWei("0.025", 18))
		require.NoError(t, err)
	})

	// now lets send an event to trigger the removed log code
	watchedContract.evCh <- &poolbindings.PoolbindingsLOGSWAP{Raw: types.Log{Removed: true}}
	time.Sleep(time.Second * 3)
	cancel()
	wg.Wait()
}
