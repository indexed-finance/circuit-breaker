package testutils

import (
	"os"
	"testing"

	"github.com/indexed-finance/circuit-breaker/config"
)

var (
	twilioSID           = os.Getenv("TWILIO_SID")
	twilioAuthToken     = os.Getenv("TWILIO_AUTH_TOKEN")
	twilioTestRecipient = os.Getenv("TWILIO_TEST_RECIPIENT")
	twilioNumber        = os.Getenv("TWILIO_NUMBER")
)

func GetConfig(t *testing.T) *config.Config {
	exCfg := *config.ExampleConfig
	exCfg.InfuraAPIKey = os.Getenv("INFURA_API_KEY")
	exCfg.Alerts.Twilio.Recipients = []string{twilioTestRecipient}
	exCfg.Alerts.Twilio.From = twilioNumber
	exCfg.Alerts.Twilio.SID = twilioSID
	exCfg.Alerts.Twilio.AuthToken = twilioAuthToken
	return &exCfg
}
