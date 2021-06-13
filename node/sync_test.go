package node

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/josetom/go-chain/common"
	"github.com/josetom/go-chain/core"
	"github.com/josetom/go-chain/fs"
	"github.com/josetom/go-chain/test_helper"
)

type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

//NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func getDummyClient(t *testing.T, handler HandlerFunc) *http.Client {
	return NewTestClient(func(req *http.Request) *http.Response {

		node := NewNode()
		state, err := core.LoadState()
		if err != nil {
			t.Fail()
		}
		node.state = state

		req, err = http.NewRequest(req.Method, req.URL.String(), req.Body)
		if err != nil {
			t.Fail()
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			handler(rw, r, &node)
		})
		handler.ServeHTTP(rr, req)
		return rr.Result()
	})
}

func TestFetchBlocksFromPeer(t *testing.T) {
	fs.Config.DataDir = test_helper.GetTestDataDir()

	node := NewNode()
	peer := NewPeerNode(getDefaultBootstrapNodes()[0].Host, true, false)
	httpClient = getDummyClient(t, nodeSyncHandler)

	state, err := core.LoadState()
	if err != nil {
		t.Fail()
	}
	node.state = state

	blocks, err := fetchBlocksFromPeer(peer, common.Hash{})
	if err != nil {
		t.Fail()
	}
	if len(blocks) != 2 || err != nil {
		hash, err := blocks[0].Hash()
		if err != nil || hash.String() != "0xbfa63a77a70876ac1b5ebaba6d9113b181259aae5afa11207aeb5143a6ed9990" {
			t.Fail()
		}
	}
}

func TestQueryPeerStatus(t *testing.T) {
	fs.Config.DataDir = test_helper.GetTestDataDir()

	node := NewNode()
	peer := NewPeerNode(getDefaultBootstrapNodes()[0].Host, true, false)
	httpClient = getDummyClient(t, nodeStatusHandler)

	state, err := core.LoadState()
	if err != nil {
		t.Fail()
	}
	node.state = state

	nodeStatusRes, err := queryPeerStatus(peer)
	if err != nil {
		t.Fail()
	}
	if nodeStatusRes.Hash.String() != "0x39714f635bda97ef70bf48ecae1a8ea27a42cc5e35dd40895db35d44107bf1bd" {
		t.Fail()
	}
}

func TestJoinKnownPeer(t *testing.T) {
	fs.Config.DataDir = test_helper.GetTestDataDir()

	node := NewNode()
	peer := NewPeerNode(getDefaultBootstrapNodes()[0].Host, true, false)
	httpClient = getDummyClient(t, nodePeersHandler)

	state, err := core.LoadState()
	if err != nil {
		t.Fail()
	}
	node.state = state

	err = (&node).joinKnownPeer(peer)
	if err != nil {
		t.Fail()
	}

	if !node.knownPeers[peer.Host].connected {
		t.Fail()
	}
}
