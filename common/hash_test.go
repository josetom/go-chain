package common

import (
	"bytes"
	"log"
	"testing"
)

func TestMarshalText(t *testing.T) {
	h := BytesToHash(Address_100_as_Bytes)
	m, err := h.MarshalText()
	if err != nil || !bytes.Equal(m, Hash_Address_100_with_0x_as_Bytes) {
		t.Fail()
	}
}

func TestUnmarshalText(t *testing.T) {
	var h Hash
	h.UnmarshalText(Hash_Address_100_with_0x_as_Bytes)
	if !h.Equal((Hash_Address_100_as_Bytes)) {
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
	if h.String() != "0x0000000000000000000000000000000000000000000000000000000000000000" {
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
		log.Println(h1, h3)
		t.Fail()
	}
}
