package database

import (
	"os"
	"testing"

	"github.com/indexed-finance/circuit-breaker/config"
	"github.com/stretchr/testify/require"
)

func TestDB(t *testing.T) {
	t.Run("DBOpts", func(t *testing.T) {
		tests := []struct {
			name    string
			args    *Opts
			wantErr bool
		}{
			{
				"pg-ssl", &Opts{
					Type:           "postgres",
					Host:           "localhost",
					Port:           "5432",
					User:           "user",
					Password:       "pass",
					DBName:         "circuit-breaker",
					SSLModeDisable: true,
				}, false,
			},
			{
				"pg-nossl", &Opts{
					Type:           "postgres",
					Host:           "localhost",
					Port:           "5432",
					User:           "user",
					Password:       "pass",
					DBName:         "circuit-breaker",
					SSLModeDisable: false,
				}, false,
			},
			{
				"sqlite", &Opts{
					Type:   "sqlite",
					DBName: "circuit-breaker",
				}, false,
			},
			{
				"invalid-type", &Opts{
					Type:   "eos",
					DBName: "getrekt",
				}, true,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				_, err := tt.args.Open()
				if (err != nil) != tt.wantErr {
					t.Fatalf("Open() err %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
		t.Run("FromConfig", func(t *testing.T) {
			cfg := config.ExampleConfig.Database
			opts := OptsFromConfig(cfg)
			require.Equal(t, cfg.Type, opts.Type)
			require.Equal(t, cfg.Host, opts.Host)
			require.Equal(t, cfg.Port, opts.Port)
			require.Equal(t, cfg.User, opts.User)
			require.Equal(t, cfg.Pass, opts.Password)
			require.Equal(t, cfg.DBName, opts.DBName)
			require.Equal(t, cfg.DBPath, opts.DBPath)
			require.Equal(t, cfg.SSLModeDisable, opts.SSLModeDisable)
		})
	})
	t.Run("AutoMigrate", func(t *testing.T) {
		t.Cleanup(func() {
			os.Remove("circuit-breaker.db")
		})
		db := newTestDB(t)
		require.NoError(t, db.Close())
	})
}

func newTestDB(t *testing.T) *Database {
	db, err := New(&Opts{Type: "sqlite", DBName: "circuit-breaker"})
	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate())
	return db
}
