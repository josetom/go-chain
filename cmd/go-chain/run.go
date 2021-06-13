package main

import (
	"log"

	"github.com/josetom/go-chain/constants"
	"github.com/josetom/go-chain/node"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: constants.BlockChainName + " server",
	Run: func(cmd *cobra.Command, args []string) {

		log.Println("Launching go-chain node...")

		n := node.NewNode()
		err := n.Run()

		if err != nil {
			log.Fatalln(err)
		}
	},
}
