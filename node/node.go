package node

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/josetom/go-chain/core"
)

type Node struct {
	dataDir string
	info    PeerNode

	state      *core.State
	knownPeers map[string]PeerNode

	miner Miner
}

func NewNode() Node {

	knownPeers := make(map[string]PeerNode)
	info := NewPeerNode(Config.Http.Host, Config.IsBootstrap, false)

	n := Node{
		dataDir:    core.GetDataDir(),
		knownPeers: knownPeers,
		info:       info,
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
	n.miner = InitMiner(state)

	// TODO : set node ready to accept txns

	defer state.Close()
	log.Println("Loaded state from disk. Latest hash : ", state.LatestBlockHash())

	ctx := context.Background()
	go n.sync(ctx)
	go n.miner.mainLoop(ctx)

	http.HandleFunc(RequestBalances, func(rw http.ResponseWriter, r *http.Request) {
		balancesHandler(rw, r, n)
	})

	http.HandleFunc(RequestTransactions, func(rw http.ResponseWriter, r *http.Request) {
		transactionsHandler(rw, r, n)
	})

	http.HandleFunc(RequestNodeStatus, func(rw http.ResponseWriter, r *http.Request) {
		nodeStatusHandler(rw, r, n)
	})

	http.HandleFunc(RequestNodeSync, func(rw http.ResponseWriter, r *http.Request) {
		nodeSyncHandler(rw, r, n)
	})

	http.HandleFunc(RequestAddPeers, func(rw http.ResponseWriter, r *http.Request) {
		nodePeersHandler(rw, r, n)
	})

	log.Println("starting server in : ", Config.Http.Port)
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
