package main

import (
	"log"
	"math/big"

	"github.com/bonedaddy/go-indexed/bclient"
	"github.com/bonedaddy/go-indexed/bindings/erc20"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/indexed-finance/circuit-breaker/bindings/controller"
	"github.com/indexed-finance/circuit-breaker/bindings/freetokens"
	"github.com/indexed-finance/circuit-breaker/bindings/multicall"
	"github.com/indexed-finance/circuit-breaker/bindings/sigmacore"
	"github.com/indexed-finance/circuit-breaker/config"
	"github.com/indexed-finance/circuit-breaker/utils"
	"github.com/urfave/cli/v2"
)

// commands for interacting with contracts

func contractCommands() cli.Commands {
	return cli.Commands{
		&cli.Command{
			Name:        "contracts",
			Usage:       "exposes helper functions for interacting with contracts",
			Description: "warning this is largely intended for testing purposes on testnets, do not use on mainnets or you burn a lot of money very fast",
			Subcommands: cli.Commands{
				setPublicSwap(),
				getFreeTokens(),
				deployMulticall(),
				approve(),
				poolSwap(),
			},
		},
	}
}

func setPublicSwap() *cli.Command {
	return &cli.Command{
		Name:  "set-public-swap",
		Usage: "sets public swap",
		Action: func(c *cli.Context) error {
			cfg, err := config.LoadConfig(c.String("config.path"))
			if err != nil {
				return err
			}
			// overrride network to rinkeby as this is only fors testing
			cfg.Network = "rinkeby"
			auther, err := getAuthorizer(cfg)
			if err != nil {
				return err
			}
			endpointURL := "wss://" + cfg.Network + ".infura.io/ws/v3/" + cfg.InfuraAPIKey
			bc, err := bclient.NewClient(endpointURL)
			if err != nil {
				return err
			}
			con, err := controller.NewController(common.HexToAddress(c.String("controller.address")), bc.EthClient())
			if err != nil {
				return err
			}
			tx, err := con.SetPublicSwap(auther.TransactOpts, common.HexToAddress(c.String("pool.address")), c.Bool("enable"))
			if err != nil {
				return err
			}
			log.Println("waiting for tx to be mined: ", tx.Hash().String())
			if _, err := bind.WaitMined(c.Context, bc.EthClient(), tx); err != nil {
				return err
			}
			log.Println("transaction mined")
			return nil
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "enable",
				Value: false,
			},
			&cli.StringFlag{
				Name:  "pool.address",
				Value: "0x59de54EFD24aB1c9234CdCA7f59FFE0CB93A7355",
			},
			&cli.StringFlag{
				Name:  "controller.address",
				Value: "0xc11e494B62565D6Ed2027ae0ebcf85E47B2CECbB",
			},
		},
	}
}

func getFreeTokens() *cli.Command {
	return &cli.Command{
		Name:  "get-free-tokens",
		Usage: "calls the free tokens contract on rinkeby",
		Action: func(c *cli.Context) error {
			cfg, err := config.LoadConfig(c.String("config.path"))
			if err != nil {
				return err
			}
			// overrride network to rinkeby as this is only fors testing
			cfg.Network = "rinkeby"
			auther, err := getAuthorizer(cfg)
			if err != nil {
				return err
			}
			endpointURL := "wss://" + cfg.Network + ".infura.io/ws/v3/" + cfg.InfuraAPIKey
			bc, err := bclient.NewClient(endpointURL)
			if err != nil {
				return err
			}
			contract, err := freetokens.NewFreetokens(common.HexToAddress(c.String("tokens.contract")), bc.EthClient())
			if err != nil {
				return err
			}
			auther.Lock()
			tx, err := contract.GetFreeTokens(auther.TransactOpts, auther.TransactOpts.From, utils.ToWei("100000.0", 18))
			if err != nil {
				return err
			}
			auther.Unlock()
			log.Println("waiting for transaction to be mined: ", tx.Hash().String())
			if _, err := bind.WaitMined(c.Context, bc.EthClient(), tx); err != nil {
				return err
			}
			log.Println("transactio mined")
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "tokens.contract",
				Usage: "the address of the tokens contract. tk1: 0xA9B218E0b6D22b7d7072D0c529adB6624db4Bf66, tk2: 0x1EBed07D5ff3d37A0E03C6EAaC57e562e29f0eee, tk3: 0xE5f9F19c415b5aa6Ca4A5d826440F2F327113DEC",
				Value: "0xE5f9F19c415b5aa6Ca4A5d826440F2F327113DEC",
			},
		},
	}
}

func deployMulticall() *cli.Command {
	return &cli.Command{
		Name:  "deploy-multicall",
		Usage: "used to deploy the multicall contract",
		Action: func(c *cli.Context) error {
			cfg, err := config.LoadConfig(c.String("config.path"))
			if err != nil {
				return err
			}
			// overrride network to rinkeby as this is only fors testing
			auther, err := getAuthorizer(cfg)
			if err != nil {
				return err
			}
			endpointURL := "wss://" + cfg.Network + ".infura.io/ws/v3/" + cfg.InfuraAPIKey
			bc, err := bclient.NewClient(endpointURL)
			if err != nil {
				return err
			}
			addr, tx, _, err := multicall.DeployMulticall(auther.TransactOpts, bc.EthClient())
			if err != nil {
				return err
			}
			if _, err := bind.WaitDeployed(c.Context, bc.EthClient(), tx); err != nil {
				return err
			}
			log.Println("multicall contract: ", addr)
			return nil
		},
	}
}

func approve() *cli.Command {
	return &cli.Command{
		Name:  "approve",
		Usage: "used to do erc20 approvals",
		Action: func(c *cli.Context) error {
			cfg, err := config.LoadConfig(c.String("config.path"))
			if err != nil {
				return err
			}
			// overrride network to rinkeby as this is only fors testing
			auther, err := getAuthorizer(cfg)
			if err != nil {
				return err
			}
			endpointURL := "wss://" + cfg.Network + ".infura.io/ws/v3/" + cfg.InfuraAPIKey
			bc, err := bclient.NewClient(endpointURL)
			if err != nil {
				return err
			}
			ercContract, err := erc20.NewErc20(common.HexToAddress(c.String("tokens.contract")), bc.EthClient())
			if err != nil {
				return err
			}
			tx, err := ercContract.Approve(auther.TransactOpts, common.HexToAddress(c.String("approval.address")), utils.ToWei("10000000000000.0", 18))
			if err != nil {
				return err
			}
			log.Println("waiting for transaction to be mined: ", tx.Hash().String())
			if _, err := bind.WaitMined(c.Context, bc.EthClient(), tx); err != nil {
				return err
			}
			log.Println("transaction mined")
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "tokens.contract",
				Usage: "the address of the tokens contract. tk1: 0xA9B218E0b6D22b7d7072D0c529adB6624db4Bf66, tk2: 0x1EBed07D5ff3d37A0E03C6EAaC57e562e29f0eee, tk3: 0xE5f9F19c415b5aa6Ca4A5d826440F2F327113DEC",
				Value: "0xE5f9F19c415b5aa6Ca4A5d826440F2F327113DEC",
			},
			&cli.StringFlag{
				Name:  "approval.address",
				Usage: "address to approve spending from",
				Value: "0x59de54EFD24aB1c9234CdCA7f59FFE0CB93A7355",
			},
		},
	}
}

func poolSwap() *cli.Command {
	return &cli.Command{
		Name:  "pool-swap",
		Usage: "used to mint",
		Action: func(c *cli.Context) error {
			cfg, err := config.LoadConfig(c.String("config.path"))
			if err != nil {
				return err
			}
			// overrride network to rinkeby as this is only fors testing
			cfg.Network = "rinkeby"
			auther, err := getAuthorizer(cfg)
			if err != nil {
				return err
			}
			endpointURL := "wss://" + cfg.Network + ".infura.io/ws/v3/" + cfg.InfuraAPIKey
			bc, err := bclient.NewClient(endpointURL)
			if err != nil {
				return err
			}
			poolContract, err := sigmacore.NewSigmacore(common.HexToAddress(c.String("pool.contract")), bc.EthClient())
			if err != nil {
				return err
			}

			spotPrice, err := poolContract.GetSpotPrice(
				nil,
				common.HexToAddress(c.String("token.in")),
				common.HexToAddress(c.String("token.out")),
			)
			if err != nil {
				return err
			}

			tx, err := poolContract.SwapExactAmountIn(
				auther.TransactOpts,
				common.HexToAddress(c.String("token.in")),
				utils.ToWei(c.String("token.in_amount"), 18),
				common.HexToAddress(c.String("token.out")),
				utils.ToWei(c.String("token.out_amount"), 18),
				new(big.Int).Mul(spotPrice, big.NewInt(2)),
			)
			if err != nil {
				return err
			}
			log.Println("waiting for tx to be mined: ", tx.Hash().String())
			if _, err := bind.WaitMined(c.Context, bc.EthClient(), tx); err != nil {
				return err
			}
			log.Println("transaction mined")
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "token.in_amount",
				Value: "200",
			},
			&cli.StringFlag{
				Name:  "token.out_amount",
				Value: "200",
			},
			&cli.StringFlag{
				Name:  "token.in",
				Value: "0xA9B218E0b6D22b7d7072D0c529adB6624db4Bf66",
			},
			&cli.StringFlag{
				Name:  "token.out",
				Value: "0xE5f9F19c415b5aa6Ca4A5d826440F2F327113DEC",
			},
			&cli.StringFlag{
				Name:  "pool.contract",
				Value: "0x59de54EFD24aB1c9234CdCA7f59FFE0CB93A7355",
			},
		},
	}
}
