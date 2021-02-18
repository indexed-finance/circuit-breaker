package main

import (
	"os"
	"time"

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
	app.Commands = getCommands()
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}

func getCommands() (cmds cli.Commands) {
	cmds = append(
		cmds,
		accountNew(),
		genConfig(),
		gasPriceCalculations(),
	)
	cmds = append(cmds, contractCommands()...)
	cmds = append(cmds, getServices()...)
	return
}
