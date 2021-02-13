package service

import (
	"context"
	"math/big"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/bonedaddy/go-indexed/bclient"
	poolbindings "github.com/bonedaddy/go-indexed/bindings/pool"
	"github.com/ethereum/go-ethereum/common"
	"github.com/indexed-finance/circuit-breaker/alerts"
	"github.com/indexed-finance/circuit-breaker/config"
	"github.com/indexed-finance/circuit-breaker/database"
	"github.com/indexed-finance/circuit-breaker/multicall"
	"github.com/indexed-finance/circuit-breaker/utils"
	"github.com/stretchr/testify/require"
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

func TestService(t *testing.T) {
	t.Cleanup(func() {
		os.Remove("circuit-breaker.db")
		os.Remove("circuit-breaker.log")
		os.RemoveAll("testDir")
	})
	os.MkdirAll("testDir", os.ModePerm)
	acct, err := utils.NewAccount("testDir", "password123")
	require.NoError(t, err)
	auther, err := utils.NewAuthorizer(acct.URL.Path, "password123")
	require.NoError(t, err)
	// create test account
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cfg := getConfig(t)
	db, err := database.New(database.OptsFromConfig(cfg.Database))
	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate())
	// ensure no previous pools exist
	_, err = db.GetPool("cc10")
	require.Error(t, err)
	_, err = db.GetPool("defi5")
	require.Error(t, err)
	bc, err := bclient.NewInfuraClient(cfg.InfuraAPIKey, true)
	require.NoError(t, err)
	logger, err := config.LoggerFromConfig(cfg)
	require.NoError(t, err)
	mc, err := multicall.New(ctx, cfg.MultiCallAddress, bc.EthClient())
	require.NoError(t, err)
	srv, err := New(ctx, alerts.New(logger, cfg.Alerts), mc, db, bc, auther, logger, cfg.Pools)
	require.NoError(t, err)
	var (
		preCreateBlock                                      uint64
		preCreatePool                                       string
		oldBalances, oldDenomWeights, oldSupplies, decimals map[string]interface{}
		nonStalePool                                        string // this denotes the pool that wont be stale and is freshly fetched
	)
	t.Run("PreCreateDatabaseEntry", func(t *testing.T) {
		// get current block
		block, err := srv.ew.BC().CurrentBlock()
		require.NoError(t, err)
		pool := cfg.Pools[0]
		switch strings.ToLower(pool.Name) {
		case "cc10":
			nonStalePool = "defi5"
		case "defi5":
			nonStalePool = "cc10"
		}
		contract, err := poolbindings.NewPoolbindings(common.HexToAddress(pool.ContractAddress), srv.ew.BC().EthClient())
		require.NoError(t, err)
		tokens, err := utils.PoolTokensFor(contract, srv.ew.BC().EthClient())
		require.NoError(t, err)
		// balances, denormWeights, err := srv.GetBalancesAndWeights(contract)
		// require.NoError(t, err)
		preCreateBlock = block - uint64(10)
		preCreatePool = pool.Name
		oldBalances, oldDenomWeights, oldSupplies, decimals = getFakeBalancesWeightsAndSupplies(t, tokens)
		require.NoError(t, srv.db.NewPool(pool.Name, pool.ContractAddress, preCreateBlock, oldBalances, oldDenomWeights, oldSupplies, decimals))
	})

	require.NoError(t, srv.Prepare())

	// validates stale preparation was ok
	t.Run("PrepareStaleValidate", func(t *testing.T) {
		pool, err := srv.db.GetPool(preCreatePool)
		require.NoError(t, err)
		require.Greater(t, pool.LastUpdateAt, preCreateBlock)
		info, err := db.GetInfo(pool.Name, pool.LastUpdateAt)
		require.NoError(t, err)
		t.Run("BalanceValidation", func(t *testing.T) {
			for tok, bal := range info.Balances {
				spBal, ok := bal.(string)
				require.True(t, ok)
				sBal, ok := oldBalances[tok].(string)
				require.True(t, ok)
				spbBal, ok := new(big.Int).SetString(spBal, 10)
				require.True(t, ok)
				sbBal, ok := new(big.Int).SetString(sBal, 10)
				require.True(t, ok)
				require.Equal(t, 1, spbBal.Cmp(sbBal), "spbBal: ", spbBal, " sbBal: ", sbBal)
				require.Equal(t, 1, sbBal.Cmp(big.NewInt(0)))
				require.Equal(t, 1, spbBal.Cmp(big.NewInt(0)))
			}
		})
		t.Run("WeighValidation", func(t *testing.T) {
			for tok, weight := range info.DenormalizedWeights {
				spBal, ok := weight.(string)
				require.True(t, ok)
				sBal, ok := oldDenomWeights[tok].(string)
				require.True(t, ok)
				spbBal, ok := new(big.Int).SetString(spBal, 10)
				require.True(t, ok)
				sbBal, ok := new(big.Int).SetString(sBal, 10)
				require.True(t, ok)
				require.Equal(t, 1, spbBal.Cmp(sbBal))
				require.Equal(t, 1, sbBal.Cmp(big.NewInt(0)))
				require.Equal(t, 1, spbBal.Cmp(big.NewInt(0)))
			}
		})
	})
	// validates that we created the missing record
	t.Run("FreshPoolValidate", func(t *testing.T) {
		pool, err := srv.db.GetPool(nonStalePool)
		require.NoError(t, err)
		contract, err := poolbindings.NewPoolbindings(common.HexToAddress(pool.ContractAddress), srv.ew.BC().EthClient())
		require.NoError(t, err)
		poolTokens, err := utils.PoolTokensFor(contract, srv.ew.BC().EthClient())
		require.NoError(t, err)
		info, err := db.GetInfo(pool.Name, pool.LastUpdateAt)
		require.NoError(t, err)
		require.Len(t, info.Balances, len(poolTokens))
		require.Len(t, info.DenormalizedWeights, len(poolTokens))
		// validate non zero weight
		t.Run("WeightValidation", func(t *testing.T) {
			for _, weight := range info.DenormalizedWeights {
				spWeight, ok := weight.(string)
				require.True(t, ok)
				spbWeight, ok := new(big.Int).SetString(spWeight, 10)
				require.True(t, ok)
				require.Equal(t, 1, spbWeight.Cmp(big.NewInt(0)))
			}
		})
		// validate non zero balance
		t.Run("BalanceValidation", func(t *testing.T) {
			for _, bal := range info.Balances {
				spBal, ok := bal.(string)
				require.True(t, ok)
				spbBal, ok := new(big.Int).SetString(spBal, 10)
				require.True(t, ok)
				require.Equal(t, 1, spbBal.Cmp(big.NewInt(0)))
			}
		})
	})
	go srv.StartWatchers()
	// this needs to be start as a goroutien
	go srv.StartBlockListener()
	t.Log("sleeping for 65 seconds to let processes run")
	time.Sleep(time.Second * 65)
	require.NoError(t, srv.Close())
}

func getFakeBalancesWeightsAndSupplies(
	t *testing.T,
	tokens map[string]common.Address,
) (balances map[string]interface{}, denormWeights map[string]interface{}, totalSupplies map[string]interface{}, decimals map[string]interface{}) {
	balances = make(map[string]interface{})
	denormWeights = make(map[string]interface{})
	totalSupplies = make(map[string]interface{})
	decimals = make(map[string]interface{})
	i := int64(1)
	for tok := range tokens {
		balances[strings.ToLower(tok)] = big.NewInt(i).String()
		denormWeights[strings.ToLower(tok)] = big.NewInt(i).String()
		totalSupplies[strings.ToLower(tok)] = big.NewInt(i).String()
		decimals[strings.ToLower(tok)] = big.NewInt(i).String()
		i++
	}
	return
}
