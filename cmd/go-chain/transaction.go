package main

import (
	"log"

	"github.com/josetom/go-chain/core"
	"github.com/spf13/cobra"
)

const flagFrom = "from"
const flagTo = "to"
const flagValue = "value"
const flagData = "data"

func txCmd() *cobra.Command {
	var txsCmd = &cobra.Command{
		Use:   "tx",
		Short: "Interact with txs (add...).",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return incorrectUsageErr()
		},
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	txsCmd.AddCommand(txAddCmd())

	return txsCmd
}

func txAddCmd() *cobra.Command {

	var cmd = &cobra.Command{
		Use:   "add",
		Short: "Adds new TX to database.",
		Run: func(cmd *cobra.Command, args []string) {
			from, _ := cmd.Flags().GetString(flagFrom)
			to, _ := cmd.Flags().GetString(flagTo)
			value, _ := cmd.Flags().GetUint(flagValue)
			data, _ := cmd.Flags().GetString(flagData)

			fromAcc := core.NewAddress(from)
			toAcc := core.NewAddress(to)

			tx := core.NewTransaction(fromAcc, toAcc, value, data)

			state, err := core.LoadState()
			if err != nil {
				log.Fatalln(err)
			}

			// defer means, at the end of this function execution,
			// execute the following statement (close DB file with all TXs)
			defer state.Close()

			// Add the TX to an in-memory array (pool)
			err = state.AddTransaction(tx)
			if err != nil {
				log.Fatalln(err)
			}

			// Flush the mempool TXs to disk
			// TODO : This is like automining txn. Needs to be updated
			_, err = state.Persist()
			if err != nil {
				log.Fatalln(err)
			}

			log.Println("TX successfully added to the ledger.")
		},
	}

	cmd.Flags().String(flagFrom, "", "From what address to send tokens")
	cmd.MarkFlagRequired(flagFrom)

	cmd.Flags().String(flagTo, "", "To what address to send tokens")
	cmd.MarkFlagRequired(flagTo)

	cmd.Flags().Uint(flagValue, 0, "How many tokens to send")
	cmd.MarkFlagRequired(flagValue)

	cmd.Flags().String(flagData, "", "Transaction data")

	return cmd
}
