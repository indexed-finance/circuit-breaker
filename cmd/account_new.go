package main

import (
	"encoding/hex"
	"errors"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/indexed-finance/circuit-breaker/utils"
	"github.com/urfave/cli/v2"
)

func accountNew() *cli.Command {
	return &cli.Command{
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
	}
}
