package config

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	var cfgPath = "kek.yaml"
	t.Cleanup(func() {
		os.Remove(cfgPath)
		os.Remove("circuit-breaker.log")
	})
	require.NoError(t, NewConfig(cfgPath))
	config, err := LoadConfig(cfgPath)
	require.NoError(t, err)
	require.Len(t, config.Pools, 2)
	require.Equal(t, config.InfuraAPIKey, "INFURA-KEY")
	require.Equal(t, config.Database, ExampleConfig.Database)
	for _, pool := range config.Pools {
		switch strings.ToLower(pool.Name) {
		case "defi5":
			require.Equal(t, pool.ContractAddress, "0xfa6de2697d59e88ed7fc4dfe5a33dac43565ea41")
			require.Equal(t, float64(10), pool.SpotPriceBreakPercentage)
			require.Equal(t, float64(5), pool.SupplyBreakPercentage)
		case "cc10":
			require.Equal(t, pool.ContractAddress, "0x17ac188e09a7890a1844e5e65471fe8b0ccfadf3")
			require.Equal(t, float64(11), pool.SpotPriceBreakPercentage)
			require.Equal(t, float64(6), pool.SupplyBreakPercentage)
		}
	}
	_, err = LoggerFromConfig(config)
	require.NoError(t, err)
}

func TestLoggerFullOpts(t *testing.T) {
	t.Cleanup(func() { os.Remove("test.log") })
	logger, err := LoggerFromConfig(&Config{
		Logger: Logger{
			Path:     "test.log",
			Dev:      true,
			Debug:    true,
			FileOnly: true,
			Fields:   map[string]interface{}{"test": "true"},
		},
	})
	require.NoError(t, err)
	require.NoError(t, logger.Sync())
}
