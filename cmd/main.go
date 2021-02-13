package main

import (
	"encoding/hex"
	"errors"
	"log"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/indexed-finance/circuit-breaker/config"
	"github.com/indexed-finance/circuit-breaker/utils"
	"github.com/urfave/cli/v2"
)

var (
	// Version ...
	Version string
)

func main() {
	app := cli.NewApp()
	app.Name = "circuit-breaker"
	app.Usage = "indexed finance circuit breaker"
	app.Version = Version
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "config.path",
			Usage: "path to store the config file at",
			Value: "circuit-breaker.yaml",
		},
		&cli.BoolFlag{
			Name:  "db.migrate",
			Usage: "run database migrations",
			Value: false,
		},
		&cli.BoolFlag{
			Name:  "startup.sleep",
			Usage: "whether or not to sleep on startup, useful for giving containers time to initialize",
			Value: false,
		},
		&cli.DurationFlag{
			Name:  "startup.sleep_time",
			Usage: "time.Duration type specifying sleep duration",
			Value: time.Second * 5,
		},
	}
	app.Before = func(c *cli.Context) (err error) {
		if c.Bool("startup.sleep") {
			time.Sleep(c.Duration("startup.sleep_time"))
		}
		return nil
	}
	app.Commands = cli.Commands{
		&cli.Command{
			Name:  "account-new",
			Usage: "generates an ethereum private key or keyfile",
			Action: func(c *cli.Context) error {
				switch c.String("mode") {
				case "keyfile":
					if c.String("key.file_pass") == "" {
						return errors.New("key.file_pass flag is empty")
					}
					var (
						keyFileDir = c.String("key.file_dir")
						err        error
					)
					if keyFileDir == "" {
						log.Println("WARN: key.file_dir is empty defaulting to current directory")
						keyFileDir, err = os.Getwd()
						if err != nil {
							return err
						}
					}
					acct, err := utils.NewAccount(keyFileDir, c.String("key.file_pass"))
					if err != nil {
						return err
					}
					log.Println("account generated")
					log.Println("public address: ", acct.Address.Hex())
					log.Println("file path: ", acct.URL.Path)
				case "privatekey":
					pk, err := crypto.GenerateKey()
					if err != nil {
						return err
					}
					log.Println("hex encoded private key:")
					log.Println(hex.EncodeToString(crypto.FromECDSA(pk)))
				default:
					return errors.New("unsupported mode, must be one of: keyfile, privatekey")
				}
				return nil
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "key.file_dir",
					Usage: "directory to store keyfile in, only applicable to keyfile mode",
				},
				&cli.StringFlag{
					Name:  "key.file_pass",
					Usage: "password to encrypt key file with, only applicable to keyfile mode",
				},
				&cli.StringFlag{
					Name:  "mode",
					Usage: "specifies the account generation mode to use: privatekey, keyfile",
					Value: "privatekey",
				},
			},
		},
		&cli.Command{
			Name:  "contract-watcher",
			Usage: "starts the contract-watcher microservice which listens for monitored events",
			Action: func(c *cli.Context) error {
				cfg, err := config.LoadConfig(c.String("config.path"))
				if err != nil {
					return err
				}
				if cfg.Database.Type == "sqlite" {
					return errors.New("sqlite database is only supported in combined mode")
				}
				return run(c, "contract-watcher", cfg)
			},
		},
		&cli.Command{
			Name:        "block-listener",
			Usage:       "starts the block-listener microservice which listens for new blocks. it must be started first",
			Description: "the block-listener updates database information for all monitored pools, as well as performing total supply inflation/deflation checks",
			Action: func(c *cli.Context) error {
				cfg, err := config.LoadConfig(c.String("config.path"))
				if err != nil {
					return err
				}
				if cfg.Database.Type == "sqlite" {
					return errors.New("sqlite database is only supported in combined mode")
				}
				return run(c, "block-listener", cfg)
			},
		},
		&cli.Command{
			Name:  "combined",
			Usage: "starts a combined block-listener and contract-watcher service",
			Action: func(c *cli.Context) error {
				cfg, err := config.LoadConfig(c.String("config.path"))
				if err != nil {
					return err
				}
				return run(c, "combined", cfg)
			},
		},
		&cli.Command{
			Name:  "config",
			Usage: "generate a config file",
			Action: func(c *cli.Context) error {
				return config.NewConfig(c.String("config.path"))
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
