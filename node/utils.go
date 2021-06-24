package node

import (
	"sort"

	"github.com/josetom/go-chain/core"
)

func txnMapToArray(txnMap map[string]core.Transaction) []core.Transaction {
	s := make([]core.Transaction, len(txnMap))
	i := 0
	for _, item := range txnMap {
		s[i] = item
		i++
	}
	// Sort slice to ensure that the txns are ordered and don't lead to inconsistencies
	sort.Slice(s, func(i, j int) bool {
		return s[i].TxnContent.Timestamp < s[j].TxnContent.Timestamp
	})
	return s
}
