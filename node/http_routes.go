package node

import (
	"fmt"
	"log"
	"net/http"

	"github.com/josetom/go-chain/common"
	"github.com/josetom/go-chain/core"
)

const (
	RequestBalances     = "/balances"
	RequestTransactions = "/transactions"
	RequestAddPeers     = "/node/peers"
	RequestNodeStatus   = "/node/status"
	RequestNodeSync     = "/node/sync"

	QueryParamFromBlock = "fromBlock"
)

type HandlerFunc func(rw http.ResponseWriter, r *http.Request, n *Node)

type ErrRes struct {
	Error string `json:"error"`
}

type BalancesRes struct {
	Balances map[core.Address]uint `json:"balances"`
	Hash     common.Hash           `json:"block_hash"`
}

type NodeStatusRes struct {
	Hash       common.Hash         `json:"block_hash"`
	Number     uint64              `json:"block_number"`
	Timestamp  uint64              `json:"block_timestamp"`
	KnownPeers map[string]PeerNode `json:"peers_known"`
}

type NodeSyncRes struct {
	Blocks []core.Block `json:"blocks"`
}

type NodeAddPeerReq struct {
	Host        string `json:"host"`
	IsBootstrap bool   `json:"isBootstrap"`
}

type NodeAddPeerRes struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func balancesHandler(w http.ResponseWriter, r *http.Request, node *Node) {
	switch r.Method {
	case http.MethodGet:
		res := BalancesRes{
			Balances: node.state.Balances,
			Hash:     node.state.LatestBlockHash(),
		}
		writeRes(w, res)
	default:
		writeErrRes(w, fmt.Errorf("only GET is supported"))
	}
}

func transactionsHandler(w http.ResponseWriter, r *http.Request, node *Node) {
	switch r.Method {
	case http.MethodPost:
		reqObject := core.TransactionData{}
		err := readReqBody(r, &reqObject)
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

		err = node.state.AddTransaction(txn)

		if err != nil {
			writeErrRes(w, err)
			return
		}

		_, err = node.state.Persist()

		if err != nil {
			writeErrRes(w, err)
			return
		}

		writeRes(w, txn)
	default:
		writeErrRes(w, fmt.Errorf("only POST is supported"))
	}
}

func nodeStatusHandler(w http.ResponseWriter, r *http.Request, n *Node) {
	switch r.Method {
	case http.MethodGet:
		res := NodeStatusRes{
			Hash:       n.state.LatestBlockHash(),
			Number:     n.state.LatestBlock().Header.Number,
			Timestamp:  n.state.LatestBlock().Header.Timestamp,
			KnownPeers: n.knownPeers,
		}
		writeRes(w, res)
	default:
		writeErrRes(w, fmt.Errorf("only GET is supported"))
	}
}

func nodeSyncHandler(w http.ResponseWriter, r *http.Request, node *Node) {
	switch r.Method {
	case http.MethodGet:
		reqHash := r.URL.Query().Get(QueryParamFromBlock)

		hash := common.Hash{}
		err := hash.UnmarshalText([]byte(reqHash))
		if err != nil {
			writeErrRes(w, err)
			return
		}

		blocks, err := node.state.GetBlocksAfter(hash)
		if err != nil {
			writeErrRes(w, err)
			return
		}

		res := NodeSyncRes{
			Blocks: blocks,
		}
		writeRes(w, res)
	default:
		writeErrRes(w, fmt.Errorf("only GET is supported"))
	}
}

func nodePeersHandler(w http.ResponseWriter, r *http.Request, n *Node) {
	switch r.Method {
	case http.MethodPost:
		napq := NodeAddPeerReq{}
		err := readReqBody(r, &napq)
		if err != nil {
			writeErrRes(w, err)
			return
		}
		peer := NewPeerNode(napq.Host, napq.IsBootstrap, true)

		n.AddPeer(peer)
		log.Println("Added new peer", peer.Host)

		res := NodeAddPeerRes{
			Success: true,
		}
		writeRes(w, res)
	default:
		writeErrRes(w, fmt.Errorf("only POST is supported"))
	}
}
