package main

import (
	"errors"
	"math/big"
	"os"
	"strconv"

	"github.com/indexed-finance/circuit-breaker/utils"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v2"
)

func gasPriceCalculations() *cli.Command {
	return &cli.Command{
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
		},
	}
}
