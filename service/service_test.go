package service

import (
	"context"
	"math/big"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/bonedaddy/go-indexed/bclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/indexed-finance/circuit-breaker/alerts"
	"github.com/indexed-finance/circuit-breaker/bindings/sigmacore"
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
		contract, err := sigmacore.NewSigmacore(common.HexToAddress(pool.ContractAddress), srv.ew.BC().EthClient())
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
		contract, err := sigmacore.NewSigmacore(common.HexToAddress(pool.ContractAddress), srv.ew.BC().EthClient())
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
	t.Run("PurgeInfoCheck", func(t *testing.T) {
		dPool, err := srv.db.GetPool("cc10")
		require.NoError(t, err)
		// record fake infos to go past the max info count
		for i := 1; i < 516; i++ {
			require.NoError(t, srv.db.RecordInfo(
				"cc10",
				dPool.LastUpdateAt+uint64(i),
				oldBalances,
				oldDenomWeights,
				oldSupplies,
			))
		}

		srv.purgeInfoCheck("cc10")
		infos, err := srv.db.GetNumInfos("cc10")
		require.NoError(t, err)
		require.LessOrEqual(t, infos, 512)

	})
	t.Run("CircuitBreakCheck", func(t *testing.T) {
		tx, _, err := srv.ew.BC().EthClient().TransactionByHash(
			ctx,
			common.HexToHash("0x5b6d669d3a27795f6f162737115a5e605157fb26465de23292c55e9739a198e8"),
		)
		require.NoError(t, err)
		fakeSwap := newFakePublicSwap(tx)
		type args struct {
			token         string
			supply        interface{}
			totalSupplies map[string]interface{}
			pool          config.Pool
		}
		tests := []struct {
			name       string
			args       args
			wantErr    bool
			wantErrStr string
		}{
			{"1-InvalidToken", args{
				"notarealtoken",
				"100",
				map[string]interface{}{
					"snx": "100",
				},
				config.Pool{
					Name:                  "cc10",
					SupplyBreakPercentage: 0,
				},
			}, true, "new supply is not of type string"},
			{"2-InvalidOldSupply-NotString", args{
				"snx",
				99.99,
				map[string]interface{}{
					"snx": "100",
				},
				config.Pool{
					Name:                  "cc10",
					SupplyBreakPercentage: 0,
				},
			}, true, "old supply is not of type string"},
			{"3-InvalidOldSupply-NotBigInt", args{
				"snx",
				"abc",
				map[string]interface{}{
					"snx": "100",
				},
				config.Pool{
					Name:                  "cc10",
					SupplyBreakPercentage: 0,
				},
			}, true, "failed to convert old supply from string to big.Int"},
			{"4-InvalidNewSupply-NotString", args{
				"snx",
				"100",
				map[string]interface{}{
					"snx": 99,
				},
				config.Pool{
					Name:                  "cc10",
					SupplyBreakPercentage: 0,
				},
			}, true, "new supply is not of type string"},
			{"5-InvalidNewSupply-NotBigInt", args{
				"snx",
				"100",
				map[string]interface{}{
					"snx": "abc",
				},
				config.Pool{
					Name:                  "cc10",
					SupplyBreakPercentage: 0,
				},
			}, true, "failed to convert new supply from string to big.Int"},
			{"6-SupplyDecreased-Break", args{
				"snx",
				utils.ToWei("1000", 18).String(),
				map[string]interface{}{
					"snx": utils.ToWei("1", 18).String(),
				},
				config.Pool{
					Name:                  "cc10",
					SupplyBreakPercentage: 1,
				},
			}, false, ""},
			{"7-SupplyIncreased-Break", args{
				"snx",
				utils.ToWei("1", 18).String(),
				map[string]interface{}{
					"snx": utils.ToWei("1000", 18).String(),
				},
				config.Pool{
					Name:                  "cc10",
					SupplyBreakPercentage: 1,
				},
			}, false, ""},
			{"8-SupplyDecreased-NoBreak", args{
				"snx",
				utils.ToWei("3", 18).String(),
				map[string]interface{}{
					"snx": utils.ToWei("2", 18).String(),
				},
				config.Pool{
					Name:                  "cc10",
					SupplyBreakPercentage: 25,
				},
			}, false, ""},
			{"9-SupplyIncreased-NoBreak", args{
				"snx",
				utils.ToWei("2", 18).String(),
				map[string]interface{}{
					"snx": utils.ToWei("3", 18).String(),
				},
				config.Pool{
					Name:                  "cc10",
					SupplyBreakPercentage: 25,
				},
			}, false, ""},
			{"10-NoSupplyChange", args{
				"snx",
				utils.ToWei("2", 18).String(),
				map[string]interface{}{
					"snx": utils.ToWei("2", 18).String(),
				},
				config.Pool{
					Name:                  "cc10",
					SupplyBreakPercentage: 25,
				},
			}, false, ""},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				err := srv.circuitBreakCheck(
					tt.args.token,
					tt.args.supply,
					tt.args.totalSupplies,
					tt.args.pool,
					fakeSwap,
				)
				if (err != nil) != tt.wantErr {
					t.Errorf("circuitBreakCheck err %v, wantErr %v", err, tt.wantErr)
				}
				// validate error string
				if tt.wantErr {
					require.Equal(t, tt.wantErrStr, err.Error())
				}
			})
		}
	})

	go srv.StartWatchers()
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

type fakePublicSwap struct {
	tx *types.Transaction
}

// returns a fake public swap struct for use with the circuitBreakCheck
// function, you must provide a transaction hash that has already been mined
func newFakePublicSwap(tx *types.Transaction) *fakePublicSwap {
	return &fakePublicSwap{tx: tx}
}

func (fpb *fakePublicSwap) SetPublicSwap(opts *bind.TransactOpts, pool common.Address, enabled bool) (*types.Transaction, error) {
	return fpb.tx, nil
}
