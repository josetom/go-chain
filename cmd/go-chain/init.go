package main

import (
	"log"

	"github.com/josetom/go-chain/constants"
	"github.com/josetom/go-chain/core"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Genesis " + constants.BlockChainName + " server",
	Run: func(cmd *cobra.Command, args []string) {

		log.Println("Initialising go-chain genesis...")

		err := core.InitGenesis()
		if err != nil {
			log.Fatalln(err)
		}

		log.Println("Genesis block created")
	},
}
