package main

import (
	"errors"

	"github.com/indexed-finance/circuit-breaker/config"
	"github.com/urfave/cli/v2"
)

func getServices() cli.Commands {
	return cli.Commands{
		&cli.Command{
			Name:  "services",
			Usage: "this command group exposes all microservice commands required for running circuit-breaker",
			Subcommands: cli.Commands{
				contractWatcher(),
				blockListener(),
				combined(),
			},
		},
	}
}

func contractWatcher() *cli.Command {
	return &cli.Command{
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
	}
}

func blockListener() *cli.Command {
	return &cli.Command{
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
	}
}

func combined() *cli.Command {
	return &cli.Command{
		Name:  "combined",
		Usage: "starts a combined block-listener and contract-watcher service",
		Action: func(c *cli.Context) error {
			cfg, err := config.LoadConfig(c.String("config.path"))
			if err != nil {
				return err
			}
			return run(c, "combined", cfg)
		},
	}
}
