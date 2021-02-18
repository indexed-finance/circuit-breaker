package main

import (
	"github.com/indexed-finance/circuit-breaker/config"
	"github.com/urfave/cli/v2"
)

func genConfig() *cli.Command {
	return &cli.Command{
		Name:  "config",
		Usage: "generate a config file",
		Action: func(c *cli.Context) error {
			return config.NewConfig(c.String("config.path"))
		},
	}
}
