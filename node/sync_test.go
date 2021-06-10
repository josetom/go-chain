package node

import (
	"bytes"
	"encoding/json"
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

func TestFetchBlocksFromPeer(t *testing.T) {
	fs.Config.DataDir = test_helper.GetTestDataDir()

	node := Node{}
	state, err := core.LoadState()
	if err != nil {
		t.Fail()
	}
	node.state = state

	peer := PeerNode{}

	httpClient = NewTestClient(func(req *http.Request) *http.Response {
		payloadBuf := new(bytes.Buffer)
		json.NewEncoder(payloadBuf).Encode(req.URL.Query().Encode())

		req, err := http.NewRequest(req.Method, req.URL.String(), payloadBuf)
		if err != nil {
			t.Fail()
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			nodeSyncHandler(rw, r, node.state)
		})
		handler.ServeHTTP(rr, req)
		return rr.Result()
	})

	blocks, err := fetchBlocksFromPeer(peer, common.Hash{})
	if len(blocks) != 2 || err != nil {
		hash, err := blocks[0].Hash()
		if err != nil || hash.String() != "0xbfa63a77a70876ac1b5ebaba6d9113b181259aae5afa11207aeb5143a6ed9990" {
			t.Fail()
		}
	}
	if err != nil {
		t.Fail()
	}
}
