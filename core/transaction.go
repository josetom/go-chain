package core

import (
	"encoding/json"
	"time"

	"github.com/josetom/go-chain/common"
)

type TransactionData struct {
	From  Address `json:"from"`
	To    Address `json:"to"`
	Value uint    `json:"value"`
	Data  string  `json:"data"`
}

type Transaction struct {
	TxnData   TransactionData `json:"txndata"`
	TxnHash   common.Hash     `json:"txnhash"`
	Timestamp uint64          `json:"timestamp"`
}

func (tx *Transaction) IsReward() bool {
	return tx.TxnData.Data == "reward"
}

func NewTransaction(from Address, to Address, value uint, data string) Transaction {
	txnData := TransactionData{from, to, value, data}
	txn := Transaction{
		TxnData:   txnData,
		Timestamp: uint64(time.Now().UnixNano()),
	}
	txn.Hash()
	return txn
}

func (t *Transaction) Hash() error {
	bytes, err := json.Marshal(t.TxnData)
	if err == nil {
		t.TxnHash = common.BytesToHash(bytes)
		return nil
	}
	return err
}

func (t *Transaction) From() Address {
	return t.TxnData.From
}

func (t *Transaction) To() Address {
	return t.TxnData.To
}

func (t *Transaction) Value() uint {
	return t.TxnData.Value
}

func (t *Transaction) Data() string {
	return t.TxnData.Data
}
