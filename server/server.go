package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/josetom/go-chain/core"
	"github.com/josetom/go-chain/node"
)

func Run(state *core.State) error {

	log.Println("Initializing http server")

	// http.HandleFunc("/transactions", HttpWrapper(&core.TransactionData{}, handler))

	http.HandleFunc("/transactions", func(w http.ResponseWriter, r *http.Request) {
		reqObject := core.TransactionData{}
		err := readReq(r, &reqObject)
		if err != nil {
			writeErrRes(w, err)
			return
		}

		txn := core.NewTransaction(
			reqObject.From,
			reqObject.To,
			reqObject.Value,
			reqObject.Data,
		)

		err = state.AddTransaction(txn)

		if err != nil {
			writeErrRes(w, err)
			return
		}

		_, err = state.Persist()

		if err != nil {
			writeErrRes(w, err)
			return
		}

		writeRes(w, txn)
	})

	return http.ListenAndServe(fmt.Sprintf(":%v", node.Config.HttpPort), nil)
}
