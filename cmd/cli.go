/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/emanuelvss13/go-hexagonal/adapters/cli"
	"github.com/spf13/cobra"
)

var action string
var id string
var name string
var price float64

// cliCmd represents the cli command
var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := cli.Run(productService, action, id, name, price)

		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(res)

	},
}

func init() {
	rootCmd.AddCommand(cliCmd)

	cliCmd.Flags().StringVarP(&action, "action", "a", "enable", "Enable / Disable")

	cliCmd.Flags().StringVarP(&id, "id", "i", "", "product id")
	cliCmd.Flags().StringVarP(&name, "name", "n", "", "product name")
	cliCmd.Flags().Float64VarP(&price, "price", "p", 0, "product price")
}
