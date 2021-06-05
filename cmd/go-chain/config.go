package main

import (
	"log"

	"github.com/josetom/go-chain/constants"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: constants.BlockChainName + " config",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO : Needs to be implemented
		log.Println("Coming soon !")
	},
}
