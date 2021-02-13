package database

import (
	"math/big"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPool(t *testing.T) {
	db := newTestDB(t)
	t.Cleanup(func() {
		db.Close()
		os.Remove("circuit-breaker.db")
	})
	t.Run("NewPool", func(t *testing.T) {
		type args struct {
			name               string
			address            string
			lastUpdateAt       uint64
			balances           map[string]interface{}
			denormWeights      map[string]interface{}
			tokenTotalSupplies map[string]interface{}
			decimals           map[string]interface{}
		}
		tests := []struct {
			name    string
			args    args
			wantErr bool
		}{
			{"DEFI5-NoError", args{
				"dEfi5",
				"addr1",
				100,
				map[string]interface{}{
					"crv": big.NewInt(80).String(),
					"uni": big.NewInt(90).String(),
					"snx": big.NewInt(100).String(),
				},
				map[string]interface{}{
					"crv": big.NewInt(60).String(),
					"uni": big.NewInt(20).String(),
					"snx": big.NewInt(20).String(),
				},
				map[string]interface{}{
					"crv": big.NewInt(10000).String(),
					"uni": big.NewInt(20000).String(),
					"snx": big.NewInt(30000).String(),
				},
				map[string]interface{}{
					"crv": big.NewInt(8).String(),
					"uni": big.NewInt(9).String(),
					"snx": big.NewInt(10).String(),
				},
			}, false},
			{"DEFI5-Error", args{
				"dEfi5",
				"addr2",
				100,
				map[string]interface{}{
					"crv": big.NewInt(100).String(),
					"uni": big.NewInt(100).String(),
					"snx": big.NewInt(100).String(),
				},
				map[string]interface{}{
					"crv": big.NewInt(60).String(),
					"uni": big.NewInt(20).String(),
					"snx": big.NewInt(20).String(),
				},
				map[string]interface{}{
					"crv": big.NewInt(10000).String(),
					"uni": big.NewInt(20000).String(),
					"snx": big.NewInt(30000).String(),
				},
				map[string]interface{}{
					"crv": big.NewInt(8),
					"uni": big.NewInt(9),
					"snx": big.NewInt(10),
				},
			}, true},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				err := db.NewPool(tt.args.name, tt.args.address, tt.args.lastUpdateAt, tt.args.balances, tt.args.denormWeights, tt.args.tokenTotalSupplies, tt.args.decimals)
				if (err != nil) != tt.wantErr {
					t.Fatalf("NewPool err %v, wantErr %v", err, tt.wantErr)
				}
				t.Run("GetInfo", func(t *testing.T) {
					if tt.wantErr {
						return
					}
					info, err := db.GetInfo(tt.args.name, tt.args.lastUpdateAt)
					require.NoError(t, err)
					t.Logf("data: %+v\n", info)
				})
				t.Run("GetTokenDecimals", func(t *testing.T) {
					if tt.wantErr {
						return
					}
					for tok, dec := range tt.args.decimals {
						retDec, err := db.GetTokenDecimals(tt.args.name, tok)
						require.NoError(t, err)
						if tt.wantErr {
							continue
						}
						strDec, ok := dec.(string)
						require.True(t, ok)
						uDec, err := strconv.ParseUint(strDec, 10, 64)
						require.NoError(t, err)
						require.Equal(t, uint8(uDec), retDec)
					}
				})
			})
		}
	})
	t.Run("GetPool", func(t *testing.T) {
		type args struct {
			name string
		}
		tests := []struct {
			name    string
			args    args
			wantErr bool
		}{
			{"DEFI5", args{"dEfi5"}, false},
			{"EOS", args{"eOs"}, true},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				_, err := db.GetPool(tt.args.name)
				if (err != nil) != tt.wantErr {
					t.Fatalf("GetPool err %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	})

	t.Run("PurgeOldInfos", func(t *testing.T) {
		for i := 0; i < 1024; i++ {
			require.NoError(t, db.RecordInfo(
				"defi5",
				1000+uint64(i),
				getFakeBalances(t, int64(i)),
				getFakeWeights(t, int64(i)),
				getFakeSupplies(t, int64(i)),
			))
		}
		// get the current info count
		count, err := db.GetNumInfos("defi5")
		require.NoError(t, err)
		require.NoError(t, db.PurgeOldInfos("defi5", 512))
		wantCount := count - 512
		retCount, err := db.GetNumInfos(("deFi5"))
		require.NoError(t, err)
		require.Equal(t, wantCount, retCount)
	})

}

func getFakeBalances(t *testing.T, iter int64) (balances map[string]interface{}) {
	balances = make(map[string]interface{})
	balances["uni"] = big.NewInt(iter + 1).String()
	balances["crv"] = big.NewInt(iter + 2).String()
	balances["snx"] = big.NewInt(iter + 3).String()
	return
}

func getFakeWeights(t *testing.T, iter int64) (denormWeights map[string]interface{}) {
	denormWeights = make(map[string]interface{})
	denormWeights["uni"] = big.NewInt(iter + 11).String()
	denormWeights["crv"] = big.NewInt(iter + 22).String()
	denormWeights["snx"] = big.NewInt(iter + 33).String()

	return
}

func getFakeSupplies(t *testing.T, iter int64) (tokenTotalSupplies map[string]interface{}) {
	tokenTotalSupplies = make(map[string]interface{})
	tokenTotalSupplies["uni"] = big.NewInt(iter + 111).String()
	tokenTotalSupplies["crv"] = big.NewInt(iter + 222).String()
	tokenTotalSupplies["snx"] = big.NewInt(iter + 333).String()
	return
}
