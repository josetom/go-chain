package core

type Transaction struct {
	From  Account `json: "from"`
	To    Account `json: "to"`
	Value uint    `json: "value"`
	Data  string  `json: "data"`
}

func (tx *Transaction) IsReward() bool {
	return tx.Data == "reward"
}

func NewTransaction(from Account, to Account, value uint, data string) Transaction {
	return Transaction{from, to, value, data}
}
