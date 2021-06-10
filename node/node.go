package node

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/josetom/go-chain/core"
)

type Node struct {
	dataDir     string
	isBootstrap bool
	host        string

	state      *core.State
	knownPeers map[string]PeerNode
}

func NewNode() Node {

	knownPeers := make(map[string]PeerNode)

	n := Node{
		dataDir:     core.GetDataDir(),
		knownPeers:  knownPeers,
		isBootstrap: Config.IsBootstrap,
		host:        Config.Http.Host,
	}

	for _, bn := range Config.BootstrapNodes {
		b := NewPeerNode(bn.Host, true, false)
		n.AddPeer(b)
	}

	return n
}

func (n *Node) Run() error {
	log.Println("Initializing node")

	err := core.InitFS()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Initialized database")

	state, err := core.LoadState()
	if err != nil {
		log.Fatalln(err)
	}
	n.state = state
	defer state.Close()
	log.Println("Loaded state from disk. Latest hash : ", state.LatestBlockHash())

	ctx := context.Background()
	go n.sync(ctx)

	http.HandleFunc("/balances", func(rw http.ResponseWriter, r *http.Request) {
		balancesHandler(rw, r, n.state)
	})

	http.HandleFunc("/transactions", func(rw http.ResponseWriter, r *http.Request) {
		transactionsHandler(rw, r, n.state)
	})

	http.HandleFunc(RequestNodeStatus, func(rw http.ResponseWriter, r *http.Request) {
		nodeStatusHandler(rw, r, n)
	})

	http.HandleFunc(RequestNodeSync, func(rw http.ResponseWriter, r *http.Request) {
		nodeSyncHandler(rw, r, n.state)
	})

	http.HandleFunc(RequestAddPeers, func(rw http.ResponseWriter, r *http.Request) {
		nodePeersHandler(rw, r, n)
	})

	return http.ListenAndServe(fmt.Sprintf(":%v", Config.Http.Port), nil)
}

func (n *Node) AddPeer(peer PeerNode) {
	n.knownPeers[peer.Host] = peer
}

func (n *Node) RemovePeer(peer PeerNode) {
	delete(n.knownPeers, peer.Host)
}

func (n *Node) IsKnownPeer(peer PeerNode) bool {
	_, isKnownPeer := n.knownPeers[peer.Host]
	return isKnownPeer
}
