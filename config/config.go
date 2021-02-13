package config

import (
	"io/ioutil"
	"os"

	"go.bobheadxi.dev/zapx/zapx"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

// Config wraps all configurable variables
type Config struct {
	InfuraAPIKey     string `yaml:"infura_api_key"`
	Network          string `yaml:"network"`
	MultiCallAddress string `yaml:"multi_call_address"`
	Pools            []Pool `yaml:"pools"`
	Database         `yaml:"database"`
	Logger           `yaml:"logger"`
	Alerts           `yaml:"alerts"`
	Profiler         `yaml:"profiler"`
	// provides control over ethereum accounts used for tranaction signing
	EthereumAccount `yaml:"ethereum_account"`
}

// Pool provides configuration information for a given index pool
type Pool struct {
	Name            string `yaml:"name"`
	ContractAddress string `yaml:"contract_address"`
	// the spot price fluctuation percent at which we should break a circuit
	SpotPriceBreakPercentage float64 `yaml:"spot_price_break_percentage"`
	// the total supply fluctation percent at which we should break a circuit
	SupplyBreakPercentage float64 `yaml:"supply_break_percentage"`
}

// Database provides configuration over our database connection
type Database struct {
	Type           string `yaml:"type"` // sqlite or postgres, if sqlite all other options except DBName are ignored
	Host           string `yaml:"host"`
	Port           string `yaml:"port"`
	User           string `yaml:"user"`
	Pass           string `yaml:"pass"`
	DBName         string `yaml:"db_name"`
	DBPath         string `yaml:"db_path"`
	SSLModeDisable bool   `yaml:"ssl_mode_disable"`
}

// Logger provides configuration over zap logger
type Logger struct {
	Path     string                 `yaml:"path"`
	Debug    bool                   `yaml:"debug"`
	Dev      bool                   `yaml:"dev"`
	FileOnly bool                   `yaml:"file_only"`
	Fields   map[string]interface{} `yaml:"fields"`
}

// Alerts allow configuration of various cirucit breaker alert capabilities
// currently only supporting twilio sms
type Alerts struct {
	Twilio `yaml:"twilio"`
}

// Twilio provides configuration of twilio sms capabilities
type Twilio struct {
	SID        string   `yaml:"sid"`
	AuthToken  string   `yaml:"auth_token"`
	From       string   `yaml:"from"`
	Recipients []string `yaml:"recipients"`
}

// Profiler configures golang profiler tooling
type Profiler struct {
	// enables the basic net/http/pprof server
	Enable bool `yaml:"enable"`
	// enables optional stats visualization (github.com/arl/statsviz)
	EnableStatsViz bool `yaml:"enable_statsviz"`
	// address to expose pprof handlers on
	Address string `yaml:"address"`
}

// EthereumAccount provides configuration over the account used to sign transactions
type EthereumAccount struct {
	// Mode specifies the ethereum account mode to use
	// currently supported values: keyfile, privatekey
	Mode            string `yaml:"mode"`
	KeyFilePath     string `yaml:"key_file_path"`
	KeyFilePassword string `yaml:"key_file_password"`
	PrivateKey      string `yaml:"private_key"`
}

var (
	// ExampleConfig is primarily used to provide a template for generating the config file
	ExampleConfig = &Config{
		InfuraAPIKey:     "INFURA-KEY",
		Network:          "mainnet",
		MultiCallAddress: "0x3067b1b7bf344027c7439509fbdf344eb25f5991",
		Logger: Logger{
			Path:  "circuit-breaker.log",
			Debug: true,
			Dev:   true,
		},
		Pools: []Pool{
			{
				Name:                     "DEFI5",
				ContractAddress:          "0xfa6de2697d59e88ed7fc4dfe5a33dac43565ea41",
				SpotPriceBreakPercentage: 10,
				SupplyBreakPercentage:    5,
			},
			{
				Name:                     "CC10",
				ContractAddress:          "0x17ac188e09a7890a1844e5e65471fe8b0ccfadf3",
				SpotPriceBreakPercentage: 11,
				SupplyBreakPercentage:    6,
			},
		},
		Database: Database{
			Type:           "sqlite",
			Host:           "localhost",
			Port:           "5432",
			User:           "user",
			Pass:           "pass",
			DBName:         "circuit-breaker",
			DBPath:         "", // current directory
			SSLModeDisable: false,
		},
		Alerts: Alerts{
			Twilio: Twilio{
				SID:        "CHANGEME-SID",
				AuthToken:  "CHANGEME-AT",
				From:       "CHANEME-FROM",
				Recipients: []string{"RECIPIENT-1"},
			},
		},
		Profiler: Profiler{
			Enable:         false,
			EnableStatsViz: false,
			Address:        "localhost:6060",
		},
		EthereumAccount: EthereumAccount{
			Mode:            "privatekey",
			KeyFilePath:     "CHANGEME-PATH",
			KeyFilePassword: "CHANGEME-PASS",
			PrivateKey:      "CHANGEME-PK",
		},
	}
)

// NewConfig generates a new config and stores at path
func NewConfig(path string) error {
	data, err := yaml.Marshal(ExampleConfig)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, data, os.ModePerm)
}

// LoadConfig loads the configuration
func LoadConfig(path string) (*Config, error) {
	r, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(r, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// LoggerFromConfig returns a logger from our config
func LoggerFromConfig(cfg *Config) (*zap.Logger, error) {
	var opts []zapx.Option
	if cfg.Logger.Debug {
		opts = append(opts, zapx.WithDebug(true))
	}
	if cfg.Logger.FileOnly {
		opts = append(opts, zapx.OnlyToFile())
	}
	if cfg.Logger.Fields != nil {
		opts = append(opts, zapx.WithFields(cfg.Logger.Fields))
	}
	return zapx.New(
		cfg.Logger.Path,
		cfg.Logger.Dev,
		opts...,
	)
}
