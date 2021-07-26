package node

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/josetom/go-chain/common"
	"github.com/josetom/go-chain/core"
)

func (n *Node) sync(ctx context.Context) error {
	n.doSync()

	tickerDuration := time.Duration(Config.Sync.TickerDuration)
	ticker := time.NewTicker(tickerDuration * time.Second)

	for {
		select {
		case <-ticker.C:
			n.doSync()

		case <-ctx.Done():
			ticker.Stop()
		}
	}
}

func (n *Node) doSync() {
	for _, peer := range n.knownPeers {
		if n.info.Host == peer.Host {
			continue
		}

		if peer.Host == "" {
			continue
		}

		// log.Printf("Searching for new Peers and their Blocks and Peers: '%s'\n", peer.Host)

		status, err := queryPeerStatus(peer)
		if err != nil {
			log.Printf("ERROR: %s\n", err)
			log.Printf("Peer '%s' was removed from KnownPeers\n", peer.Host)

			n.RemovePeer(peer)

			continue
		}

		err = n.joinKnownPeer(peer)
		if err != nil {
			log.Printf("ERROR: %s\n", err)
			continue
		}

		err = n.syncBlocks(peer, status)
		if err != nil {
			log.Printf("ERROR: %s\n", err)
			continue
		}

		err = n.syncKnownPeers(status)
		if err != nil {
			log.Printf("ERROR: %s\n", err)
			continue
		}

		n.syncPendingTxs(peer, status.PendingTxns)

	}
}

func (n *Node) syncBlocks(peer PeerNode, status NodeStatusRes) error {
	localBlockNumber := n.state.LatestBlock().Header.Number

	// If the peer has no blocks, ignore it
	if status.Hash.IsEmpty() {
		return nil
	}

	// If the peer has less blocks than us, ignore it
	if status.Number < localBlockNumber {
		return nil
	}

	// If it's the genesis block and we already synced it, ignore it
	if status.Number == 0 && !n.state.LatestBlockHash().IsEmpty() {
		return nil
	}

	newBlocksCount := status.Number - localBlockNumber
	if localBlockNumber == 0 && status.Number == 0 {
		// log.Printf("Found genesis new blocks from Peer %s\n", peer.Host)
	} else if newBlocksCount > 0 {
		log.Printf("Found %d new blocks from Peer %s\n", newBlocksCount, peer.Host)
	}

	blocks, err := fetchBlocksFromPeer(peer, n.state.LatestBlockHash())
	if err != nil {
		return err
	}

	for _, block := range blocks {
		_, err := n.state.AddBlock(block)
		if err != nil {
			return err
		}
		n.miner.syncBlockCh <- block
	}

	return nil
}

func (n *Node) syncPendingTxs(peer PeerNode, pendingTxns []core.Transaction) {
	for _, txn := range pendingTxns {
		n.miner.txnsCh <- txn
	}
}

func (n *Node) syncKnownPeers(status NodeStatusRes) error {
	for _, statusPeer := range status.KnownPeers {
		if !n.IsKnownPeer(statusPeer) {
			log.Printf("Found new Peer %s\n", statusPeer.Host)
			n.AddPeer(statusPeer)
		}
	}

	return nil
}

func (n *Node) joinKnownPeer(peer PeerNode) error {
	if peer.connected {
		return nil
	}

	url := fmt.Sprintf(
		"%s%s",
		peer.Host,
		RequestAddPeers,
	)

	body := &NodeAddPeerReq{
		Host:        n.info.Host,
		IsBootstrap: n.info.IsBootstrap,
	}
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(body)

	res, err := httpClient.Post(url, "application/json", payloadBuf)
	if err != nil {
		return err
	}

	addPeerRes := NodeAddPeerRes{}
	err = ReadRes(res, &addPeerRes)
	if err != nil {
		return err
	}
	if addPeerRes.Error != "" {
		return fmt.Errorf(addPeerRes.Error)
	}

	if !addPeerRes.Success {
		return fmt.Errorf("unable to join peer '%s'", peer.Host)
	}

	kp := n.knownPeers[peer.Host]
	kp.connected = addPeerRes.Success

	n.AddPeer(kp)

	return nil

}

func queryPeerStatus(peer PeerNode) (NodeStatusRes, error) {
	url := fmt.Sprintf("%s%s", peer.Host, RequestNodeStatus)
	res, err := httpClient.Get(url)
	if err != nil {
		return NodeStatusRes{}, fmt.Errorf("unable to connect to %s", peer.Host)
	}

	statusRes := NodeStatusRes{}
	err = ReadRes(res, &statusRes)
	if err != nil {
		return NodeStatusRes{}, err
	}

	return statusRes, nil
}

func fetchBlocksFromPeer(peer PeerNode, hash common.Hash) ([]core.Block, error) {

	url := fmt.Sprintf(
		"%s%s?%s=%s",
		peer.Host,
		RequestNodeSync,
		QueryParamFromBlock,
		hash,
	)

	res, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	syncRes := NodeSyncRes{}
	err = ReadRes(res, &syncRes)
	if err != nil {
		return nil, err
	}

	return syncRes.Blocks, nil
}
