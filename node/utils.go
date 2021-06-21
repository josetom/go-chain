package node

import "github.com/josetom/go-chain/core"

func txnMapToArray(txnMap map[string]core.Transaction) []core.Transaction {
	s := make([]core.Transaction, len(txnMap))
	i := 0
	for _, item := range txnMap {
		s[i] = item
		i++
	}
	return s
}
