package main

import (
	"fmt"
	"log"

	"github.com/josetom/go-chain/config"
	"github.com/josetom/go-chain/constants"
	"github.com/spf13/cobra"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC | log.Lshortfile)
	config.Load("config.yaml")

	var cmd = &cobra.Command{
		Use:   constants.CliName,
		Short: constants.BlockChainName + " CLI",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return incorrectUsageErr()
		},
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	cmd.AddCommand(versionCmd)
	cmd.AddCommand(runCmd)
	cmd.AddCommand(initCmd)
	cmd.AddCommand(configCmd)
	cmd.AddCommand(balancesCmd())
	cmd.AddCommand(txCmd())
	cmd.AddCommand(walletCmd())

	err := cmd.Execute()
	if err != nil {
		log.Fatalln(err)
	}
}

func incorrectUsageErr() error {
	return fmt.Errorf("incorrect_usage")
}
