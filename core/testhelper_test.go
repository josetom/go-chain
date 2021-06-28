package core

import (
	"time"

	"github.com/josetom/go-chain/common"
	"github.com/josetom/go-chain/test_helper"
)

func mineBlockHelper(pendingBlock Block) (Block, error) {
	hash, err := pendingBlock.Hash()
	if err != nil {
		return Block{}, err
	}
	if IsBlockHashValid(hash) {
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
