package common

import (
	"bytes"
	"testing"

	"github.com/josetom/go-chain/test_helper"
)

func TestMarshalText(t *testing.T) {
	h := BytesToHash(test_helper.Address_100_as_Bytes)
	m, err := h.MarshalText()
	if err != nil || !bytes.Equal(m, test_helper.Hash_Address_100_with_0x_as_Bytes) {
		t.Fail()
	}
}

func TestUnmarshalText(t *testing.T) {
	var h Hash
	h.UnmarshalText(test_helper.Hash_Address_100_with_0x_as_Bytes)
	if !h.Equal((test_helper.Hash_Address_100_as_Bytes)) {
		t.Fail()
	}

}

func TestIsEmpty(t *testing.T) {
	h1 := Hash{}
	if !h1.IsEmpty() {
		t.Fail()
	}
}

func TestString(t *testing.T) {
	h := Hash{}
	if h.String() != test_helper.Hash_0x {
		t.Fail()
	}
}

func TestSetBytes(t *testing.T) {
	var h1 Hash
	h3 := Hash{}

	// setting a dummy prefix
	ba := make([]byte, HashLength+2)
	copy(ba, []byte("as"))
	h1.setBytes(ba)

	if !h1.Equal(h3) {
		t.Error(h1, h3)
	}
}
