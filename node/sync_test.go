package node

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/josetom/go-chain/common"
	"github.com/josetom/go-chain/core"
	"github.com/josetom/go-chain/db"
	"github.com/josetom/go-chain/fs"
	"github.com/josetom/go-chain/test_helper"
	"github.com/josetom/go-chain/test_helper/test_helper_core"
)

var node Node = createDummmyNodeAndLoadState()
var handlers *http.ServeMux

func TestSync(t *testing.T) {
	// node = createDummmyNodeAndLoadState()
	handlers = node.registerHandlers()
	httpClient = getDummyClient()

	t.Run("testFetchBlocksFromPeer", testFetchBlocksFromPeer)
	t.Run("testJoinKnownPeer", testJoinKnownPeer)
	t.Run("testQueryPeerStatus", testQueryPeerStatus)

	defer node.state.Close()
}

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

func createDummmyNodeAndLoadState() Node {
	db.Config.Type = db.LEVEL_DB
	fs.Config.DataDir = test_helper.GetTestDataDir()
	core.Config.State.DbFile = core.Defaults().State.DbFile

	node := NewNode()
	state, err := test_helper_core.GetTestState()
	if err != nil {
		log.Fatalln("sync_test loadstate failed \n", err)
	}

	node.state = state

	return node

}

func getDummyClient() *http.Client {
	return NewTestClient(func(req *http.Request) *http.Response {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest(req.Method, req.URL.String(), req.Body)
		if err != nil {
			writeErrRes(rr, err)
			return rr.Result()
		}

		handlers.ServeHTTP(rr, req)

		return rr.Result()
	})
}

func testFetchBlocksFromPeer(t *testing.T) {
	db.Config.Type = db.LEVEL_DB
	fs.Config.DataDir = test_helper.GetTestDataDir()
	tempDbPath := test_helper.CreateAndGetTestDbFile()
	core.Config.State.DbFile = tempDbPath

	node := NewNode()
	peer := NewPeerNode(getDefaultBootstrapNodes()[0].Host, true, false)

	state, err := core.LoadState()
	if err != nil {
		t.Error(err)
	}
	node.state = state

	blocks, err := fetchBlocksFromPeer(peer, common.Hash{})
	if err != nil {
		t.Error(err)
	}
	if len(blocks) != 2 || err != nil {
		hash, err := blocks[0].Hash()
		if err != nil {
			t.Error(err)
		}
		if hash.String() != test_helper.Hash_Block_0 {
			t.Fail()
		}
	}

	cleanup := func() {
		state.Close()
		test_helper.DeleteTestDbFile(tempDbPath)
	}
	t.Cleanup(cleanup)
}

func testQueryPeerStatus(t *testing.T) {
	db.Config.Type = db.LEVEL_DB
	fs.Config.DataDir = test_helper.GetTestDataDir()
	tempDbPath := test_helper.CreateAndGetTestDbFile()
	core.Config.State.DbFile = tempDbPath

	node := NewNode()
	peer := NewPeerNode(getDefaultBootstrapNodes()[0].Host, true, false)

	state, err := core.LoadState()
	if err != nil {
		t.Error(err)
	}
	node.state = state

	nodeStatusRes, err := queryPeerStatus(peer)
	if err != nil {
		t.Error(err)
	}
	if nodeStatusRes.Hash.String() != test_helper.Hash_Block_1 {
		t.Error(nodeStatusRes.Hash.String())
	}

	cleanup := func() {
		state.Close()
		test_helper.DeleteTestDbFile(tempDbPath)
	}
	t.Cleanup(cleanup)
}

func testJoinKnownPeer(t *testing.T) {
	db.Config.Type = db.LEVEL_DB
	fs.Config.DataDir = test_helper.GetTestDataDir()
	tempDbPath := test_helper.CreateAndGetTestDbFile()
	core.Config.State.DbFile = tempDbPath

	node := NewNode()
	peer := NewPeerNode(getDefaultBootstrapNodes()[0].Host, true, false)

	state, err := core.LoadState()
	if err != nil {
		t.Error(err)
	}
	node.state = state

	err = (&node).joinKnownPeer(peer)
	if err != nil {
		t.Error(err)
	}

	if !node.knownPeers[peer.Host].connected {
		t.Fail()
	}

	cleanup := func() {
		state.Close()
		test_helper.DeleteTestDbFile(tempDbPath)
	}
	t.Cleanup(cleanup)
}
