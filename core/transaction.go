package core

import (
	"encoding/json"
	"time"

	"github.com/josetom/go-chain/common"
)

type TransactionData struct {
	From  common.Address `json:"from"`
	To    common.Address `json:"to"`
	Value uint           `json:"value"`
	Data  string         `json:"data"`
}

type TransactionContent struct {
	TxnData   TransactionData `json:"txnData"`
	Timestamp uint64          `json:"timestamp"`
	IsReward  bool            `json:"isReward"`
}

type Transaction struct {
	TxnContent TransactionContent `json:"txnContent"`
	TxnHash    common.Hash        `json:"txnhHsh"`
}

func NewTransaction(from common.Address, to common.Address, value uint, data string) Transaction {
	txnData := TransactionData{from, to, value, data}
	txn := Transaction{
		TxnContent: TransactionContent{
			TxnData:   txnData,
			Timestamp: uint64(time.Now().UnixNano()),
		},
	}
	txn.hash()
	return txn
}

func (t *Transaction) hash() error {
	bytes, err := json.Marshal(t.TxnContent)
	if err == nil {
		t.TxnHash = common.BytesToHash(bytes)
		return nil
	}
	return err
}

func (t *Transaction) Hash() common.Hash {
	return t.TxnHash
}

func (t *Transaction) From() common.Address {
	return t.TxnContent.TxnData.From
}

func (t *Transaction) To() common.Address {
	return t.TxnContent.TxnData.To
}

func (t *Transaction) Value() uint {
	return t.TxnContent.TxnData.Value
}

func (t *Transaction) Data() string {
	return t.TxnContent.TxnData.Data
}
