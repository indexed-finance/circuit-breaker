package main

import (
	"encoding/hex"
	"errors"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/bonedaddy/go-indexed/bclient"
	"github.com/bonedaddy/go-indexed/bindings/erc20"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/indexed-finance/circuit-breaker/bindings/controller"
	"github.com/indexed-finance/circuit-breaker/bindings/freetokens"
	"github.com/indexed-finance/circuit-breaker/bindings/multicall"
	"github.com/indexed-finance/circuit-breaker/bindings/sigmacore"
	"github.com/indexed-finance/circuit-breaker/config"
	"github.com/indexed-finance/circuit-breaker/utils"
	"github.com/olekukonko/tablewriter"
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
			Name:        "gas-price-calculations",
			Usage:       "used to calculate overall gas cost of swap disable transactions",
			Description: "Uses https://rinkeby.etherscan.io/tx/0xe16c55dac3f702f56118077dbc441998f7ec8c746cd0f071a47b1535b12c0e1a as a reference",
			Action: func(c *cli.Context) error {
				// columns: gas limit, gas price, gas cost
				data := [][]string{
					{"35000", utils.ToWei("100.0", 9).String(), ""},
					{"35000", utils.ToWei("200.0", 9).String(), ""},
					{"35000", utils.ToWei("300.0", 9).String(), ""},
					{"35000", utils.ToWei("400.0", 9).String(), ""},
					{"35000", utils.ToWei("500.0", 9).String(), ""},
					{"35000", utils.ToWei("600.0", 9).String(), ""},
					{"40000", utils.ToWei("600.0", 9).String(), ""},
					{"50000", utils.ToWei("600.0", 9).String(), ""},
					{"60000", utils.ToWei("600.0", 9).String(), ""},
					{"70000", utils.ToWei("600.0", 9).String(), ""},
					{"80000", utils.ToWei("600.0", 9).String(), ""},
				}
				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader([]string{"Gas Limit", "Gas Price (gwei)", "Gas Cost (eth)"})
				for _, v := range data {
					// get gas limit
					gasLimit, err := strconv.ParseUint(v[0], 10, 64)
					if err != nil {
						return err
					}
					// get gas price
					gasPrice, ok := new(big.Int).SetString(v[1], 10)
					if !ok {
						return errors.New("failed to parse gas cost")
					}
					// calculate gas cost
					gasCost := utils.CalcGasCost(gasLimit, gasPrice)
					v[1] = utils.ToDecimal(gasPrice.String(), 9).String()
					v[2] = utils.ToDecimal(gasCost.String(), 18).String()
					table.Append(v)
				}
				table.Render()
				return nil

				return nil
			},
		},
		&cli.Command{
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
		},
		&cli.Command{
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
		},
		&cli.Command{
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
		},
		&cli.Command{
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
		},
		&cli.Command{
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
