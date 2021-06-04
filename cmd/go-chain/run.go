package main

import (
	"log"

	"github.com/josetom/go-chain/constants"
	"github.com/josetom/go-chain/core"
	"github.com/josetom/go-chain/server"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: constants.BlockChainName + " server",
	Run: func(cmd *cobra.Command, args []string) {

		log.Println("Launching go-chain node...")

		state, err := core.LoadState()
		log.Println("Loaded the state")

		err = server.Run(state)
		if err != nil {
			log.Fatalln(err)
		}
	},
}
