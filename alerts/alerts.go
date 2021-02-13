package alerts

import (
	"github.com/bonedaddy/circuit-breaker/config"
	"github.com/kevinburke/twilio-go"
	"go.uber.org/multierr"
	"go.uber.org/zap"
)

// Alerter provides various alerting capabilities
type Alerter struct {
	logger *zap.Logger
	cfg    config.Alerts
	twilio *twilio.Client
}

// New returns a new alerter service
func New(logger *zap.Logger, cfg config.Alerts) *Alerter {
	return &Alerter{
		logger: logger.Named("alerter"),
		cfg:    cfg,
		twilio: twilio.NewClient(cfg.SID, cfg.AuthToken, nil),
	}
}

// NotifyCircuitBreak notifies recipients abouta broken circuit
func (a *Alerter) NotifyCircuitBreak(msg string) error {
	var errors error
	for _, number := range a.cfg.Twilio.Recipients {
		resp, err := a.twilio.Messages.SendMessage(
			a.cfg.Twilio.From,
			number,
			msg,
			nil,
		)
		if err != nil {
			// dont permanently exit so we continue to retry others
			a.logger.Error(
				"failed to send sms message",
				zap.Error(err),
				zap.String("recipient", number),
			)
			errors = multierr.Combine(errors, err)
		} else {
			a.logger.Debug(
				"message sent",
				zap.Any("response", resp),
			)
		}
	}
	return errors
}
