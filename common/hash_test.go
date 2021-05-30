package common

import (
	"bytes"
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
