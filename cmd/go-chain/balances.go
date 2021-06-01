package main

import (
	"log"

	"github.com/josetom/go-chain/core"
	"github.com/spf13/cobra"
)

func balancesCmd() *cobra.Command {
	var balancesCmd = &cobra.Command{
		Use:   "balances",
		Short: "Interact with balances (list...).",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return incorrectUsageErr()
		},
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	balancesCmd.AddCommand(balancesListCmd)

	return balancesCmd
}

var balancesListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all balances.",
	Run: func(cmd *cobra.Command, args []string) {
		state, err := core.LoadState()
		if err != nil {
			log.Fatalln(err)
		}
		defer state.Close()

		log.Println("Address balances at ", state.LatestBlockHash().String())
		log.Println("__________________")
		log.Println("")
		for address, balance := range state.Balances {
			log.Printf("%s: %d", address, balance)
		}
	},
}
