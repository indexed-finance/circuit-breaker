package alerts

import (
	"testing"

	testutils "github.com/indexed-finance/circuit-breaker/utils/tests"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestAlerter(t *testing.T) {
	t.Run("Twilio", func(t *testing.T) {
		cfg := testutils.GetConfig(t)
		logger, err := zap.NewDevelopment()
		require.NoError(t, err)
		alerter := New(logger, cfg.Alerts)
		alerter.NotifyCircuitBreak("this is a test message")
	})
}
