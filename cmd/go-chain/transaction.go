package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/josetom/go-chain/common"
	"github.com/josetom/go-chain/core"
	"github.com/josetom/go-chain/node"
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

			fromAcc := common.NewAddress(from)
			toAcc := common.NewAddress(to)

			url := fmt.Sprintf(
				"%s%s",
				node.Config.Http.Host,
				node.RequestTransactions,
			)

			body := &core.TransactionData{
				From:  fromAcc,
				To:    toAcc,
				Value: value,
				Data:  data,
			}

			payloadBuf := new(bytes.Buffer)
			json.NewEncoder(payloadBuf).Encode(body)

			res, err := http.Post(url, "application/json", payloadBuf)
			if err != nil {
				log.Panicln(err)
			}

			txnRes := core.Transaction{}
			node.ReadRes(res, &txnRes)
			log.Println("TX successfully added to the ledger.", txnRes.TxnHash)
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
