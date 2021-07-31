package core

import (
	"time"

	"github.com/josetom/go-chain/common"
	"github.com/josetom/go-chain/test_helper"
)

func mineBlockHelper(pendingBlock Block) (Block, error) {
	isBlockValid, err := pendingBlock.IsBlockHashValid()
	if err != nil {
		return Block{}, err
	}
	if isBlockValid {
		return pendingBlock, nil
	}
	pendingBlock.Header.Nonce = common.GenNonce()
	return mineBlockHelper(pendingBlock)
}

func getTestTxn() Transaction {

	data := test_helper.DUMMY_DATA

	txn := NewTransaction(
		common.NewAddress(test_helper.Test_Address_1),
		common.NewAddress(test_helper.Test_Address_2),
		100,
		data,
	)
	txn.TxnContent.Timestamp = uint64(time.Time{}.UnixNano())
	txn.Hash()
	txn.Sign()
	return txn
}

func getTestBlock(isValid bool, state *State, txns []Transaction) (Block, error) {
	block := NewBlock(
		state.latestBlockHash,
		1,
		uint64(time.Time{}.UnixNano()),
		0,
		common.NewAddress(""), // miner.Config.Address
		MINING_ALGO_POW,
		uint64(Config.Block.Reward),
		txns,
		nil,
	)

	if isValid {
		return mineBlockHelper(block)
	} else {
		return block, nil
	}
}
