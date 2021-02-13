package alerts

import (
	"os"
	"testing"

	"github.com/indexed-finance/circuit-breaker/config"
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

func TestAlerter(t *testing.T) {
	t.Run("Twilio", func(t *testing.T) {
		cfg := getConfig(t)
		logger, err := zap.NewDevelopment()
		require.NoError(t, err)
		alerter := New(logger, cfg.Alerts)
		alerter.NotifyCircuitBreak("this is a test message")
	})
}
