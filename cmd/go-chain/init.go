package main

import (
	"log"

	"github.com/josetom/go-chain/constants"
	"github.com/josetom/go-chain/core"
	"github.com/josetom/go-chain/db/errors"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Genesis " + constants.BlockChainName + " server",
	Run: func(cmd *cobra.Command, args []string) {

		log.Println("Running genesis...")

		// Ensure there is no existing blockchain
		// To avoid genesis creating origin conflict
		state, err := core.NewState()
		if err != nil {
			log.Fatalln(err)
		}
		block, err := state.GetBlockWithNumber(0)
		if err != nil && err.Error() != errors.NotFoundError.Error() {
			log.Fatalln(err)
		}
		if !block.Hash.IsEmpty() {
			log.Fatal("genesis block already present")
		}
		state.Close()

		// Initialize Genesis
		err = core.InitGenesis()
		if err != nil {
			log.Fatalln(err)
		}

		log.Println("Genesis block created")
	},
}
