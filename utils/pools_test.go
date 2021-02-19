package utils

import (
	"os"
	"strings"
	"testing"

	"github.com/bonedaddy/go-indexed/bclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/indexed-finance/circuit-breaker/bindings/sigmacore"
	"github.com/indexed-finance/circuit-breaker/config"
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

func TestPoolTokensFor(t *testing.T) {
	cfg := getConfig(t)
	bc, err := bclient.NewInfuraClient(cfg.InfuraAPIKey, true)
	require.NoError(t, err)
	var defi5, cc10 *sigmacore.Sigmacore
	for _, pool := range cfg.Pools {
		switch strings.ToLower(pool.Name) {
		case "cc10":
			cc10, err = sigmacore.NewSigmacore(common.HexToAddress(pool.ContractAddress), bc.EthClient())
			require.NoError(t, err)
		case "defi5":
			defi5, err = sigmacore.NewSigmacore(common.HexToAddress(pool.ContractAddress), bc.EthClient())
			require.NoError(t, err)
		}
	}
	// ensure symbols cache is empty
	require.Len(t, symbols, 0)
	// check defi5
	_, err = PoolTokensFor(defi5, bc.EthClient())
	require.NoError(t, err)
	// now ensure we populated the cache
	require.Greater(t, len(symbols), 0)
	// get current lenght to compare after cc10 run
	currLen := len(symbols)
	// check cc10
	_, err = PoolTokensFor(cc10, bc.EthClient())
	require.NoError(t, err)
	// ensure we populate cache
	require.Greater(t, len(symbols), 5)
	newLen := len(symbols)
	require.Greater(t, newLen, currLen)

}
