package main

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/bonedaddy/go-indexed/bclient"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/indexed-finance/circuit-breaker/alerts"
	"github.com/indexed-finance/circuit-breaker/config"
	"github.com/indexed-finance/circuit-breaker/database"
	"github.com/indexed-finance/circuit-breaker/multicall"
	"github.com/indexed-finance/circuit-breaker/profiler"
	"github.com/indexed-finance/circuit-breaker/service"
	"github.com/indexed-finance/circuit-breaker/utils"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func run(c *cli.Context, mode string, cfg *config.Config) error {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, os.Interrupt)
	ctx, cancel := context.WithCancel(c.Context)
	defer cancel()
	logger, err := config.LoggerFromConfig(cfg)
	if err != nil {
		return err
	}
	logger.Info("generating ethereum tranasction signer")
	auther, err := getAuthorizer(cfg)
	if err != nil {
		return err
	}
	// dont allow double registration of the cpu profiler
	if cfg.Profiler.Enable {
		logger.Info("enabling http profiling server")
		profServer := profiler.New()
		profServer.Start(cfg.Profiler.Address, cfg.Profiler.EnableStatsViz)
		defer profServer.Close()
	}
	db, err := database.New(database.OptsFromConfig(cfg.Database))
	if err != nil {
		return err
	}
	if c.Bool("db.migrate") {
		logger.Info("running database migrations")
		if err := db.AutoMigrate(); err != nil {
			db.Close()
			return err
		}
	}
	logger.Info("initializing infura client")
	endpointURL := "wss://" + cfg.Network + ".infura.io/ws/v3/" + cfg.InfuraAPIKey
	bc, err := bclient.NewClient(endpointURL)
	if err != nil {
		db.Close()
		return err
	}
	logger.Info("initializing multicall contract binding")
	mc, err := multicall.New(ctx, cfg.MultiCallAddress, bc.EthClient())
	if err != nil {
		db.Close()
		return err
	}
	logger.Info("initializing service")
	srv, err := service.New(ctx, alerts.New(logger, cfg.Alerts), mc, db, bc, auther, logger, cfg.Pools, cfg.EthereumAccount.GasPrice)
	if err != nil {
		db.Close()
		return err
	}
	switch mode {
	case "contract-watcher":
		go func() {
			<-ch
			logger.Info("caught exit signal, shutting down")
			cancel()
		}()
		logger.Info("launching contract watchers")
		if err := srv.StartWatchers(); err != nil {
			srv.Close()
			return err
		}
	case "block-listener":
		logger.Info("starting service preparation")
		if err := srv.Prepare(); err != nil {
			srv.Close()
			return err
		}
		logger.Info("finished service preparation")
		go func() {
			<-ch
			logger.Info("caught exit signal, shutting down")
			cancel()
		}()
		logger.Info("launching service (block listener loop)")
		if err := srv.StartBlockListener(); err != nil {
			logger.Error("encountered error during block listener loop", zap.Error(err))
		}
	case "combined":
		logger.Info("starting service preparation")
		if err := srv.Prepare(); err != nil {
			srv.Close()
			return err
		}
		logger.Info("finished service preparation")
		go func() {
			<-ch
			logger.Info("caught exit signal, shutting down")
			cancel()
		}()
		logger.Info("launching contract watchers")
		go srv.StartWatchers()
		logger.Info("launching service (block listener loop)")
		if err := srv.StartBlockListener(); err != nil {
			logger.Error("encountered error during block listener loop", zap.Error(err))
		}
	default:
		return errors.New("invalid mode provided must be one of: contract-watcher, block-listener, combined")
	}
	if err := srv.Close(); err != nil {
		logger.Error("service shutdown failed", zap.Error(err))
		return err
	}
	return nil

}

func getAuthorizer(cfg *config.Config) (*utils.Authorizer, error) {
	switch cfg.EthereumAccount.Mode {
	case "keyfile":
		return utils.NewAuthorizer(cfg.KeyFilePath, cfg.KeyFilePassword)
	case "privatekey":
		pk, err := crypto.HexToECDSA(cfg.EthereumAccount.PrivateKey)
		if err != nil {
			return nil, err
		}
		return utils.NewAuthorizerFromPK(pk), nil
	default:
		return nil, errors.New("unsupported account mode, must be one of: keyfile, privatekey")
	}
}
