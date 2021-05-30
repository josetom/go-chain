package main

import (
	"log"

	"github.com/josetom/go-chain/config"
	"github.com/josetom/go-chain/constants"
	"github.com/spf13/cobra"
)

var consoleCmd = &cobra.Command{
	Use:   "console",
	Short: constants.BlockChainName + " CLI client",
	Run: func(cmd *cobra.Command, args []string) {
		config.Load("config.yaml")
		log.Println("Coming soon !")
	},
}
