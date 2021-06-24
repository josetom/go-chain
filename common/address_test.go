package common

import (
	"bytes"
	"testing"

	"github.com/josetom/go-chain/test_helper"
)

func TestNewAddress(t *testing.T) {
	a := NewAddress(test_helper.Address_100_Hex_with_0x)
	if !bytes.Equal(a.Bytes(), test_helper.Address_100_as_Bytes) {
		t.Fail()
	}
}

func TestBytesToAddress(t *testing.T) {
	a1 := NewAddress(test_helper.Address_100_Hex_with_0x)
	a2 := BytesToAddress(test_helper.Address_100_as_Bytes)
	if !a1.Equal(a2) {
		t.Fail()
	}
}

func TestAddressSetBytes(t *testing.T) {
	var a1, a2 Address
	a1.SetBytes(test_helper.Address_100_with_extra_2_zeroes_prefixed_as_bytes)
	a2.SetBytes(test_helper.Address_100_with_extra_2_numbers_prefixed_as_bytes)

	a3 := NewAddress(test_helper.Address_100_Hex_with_0x)

	if !a1.Equal(a3) {
		t.Fail()
	}

	if !a2.Equal(a3) {
		t.Fail()
	}
}

func TestHash(t *testing.T) {
	a := NewAddress(test_helper.Address_100_Hex_with_0x)
	if test_helper.Hash_Address_100_as_Bytes != a.Hash() {
		t.Fail()
	}
}

func TestAddressString(t *testing.T) {
	a := NewAddress(test_helper.Address_100_Hex_with_0x)
	if a.String() != test_helper.Address_100_Hex_with_0x {
		t.Fail()
	}
}

func TestAddressMarshalText(t *testing.T) {
	a := NewAddress(test_helper.Address_100_Hex_with_0x)
	m, err := a.MarshalText()
	if err != nil || !bytes.Equal(m, test_helper.Address_100_Hex_with_0x_as_Bytes) {
		t.Fail()
	}
}

func TestAddressUnmarshalText(t *testing.T) {
	var a Address
	a.UnmarshalText(test_helper.Address_100_Hex_with_0x_as_Bytes)
	if !a.Equal(NewAddress(test_helper.Address_100_Hex_with_0x)) {
		t.Fail()
	}

}
