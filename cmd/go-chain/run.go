package main

import (
	"context"
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

		ctx := context.Background()

		n := node.NewNode()
		err := n.Run(ctx)

		if err != nil {
			log.Fatalln(err)
		}
	},
}
