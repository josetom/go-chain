package core

import (
	"bytes"
	"crypto/sha256"
	"testing"
)

const (
	Address_100                               = "00000000000000000100"
	Address_100_with_extra_2_zeroes_prefixed  = "0000000000000000000100"
	Address_100_with_extra_2_numbers_prefixed = "1200000000000000000100"

	Address_100_Hex_with_0x = "0x3030303030303030303030303030303030313030"
)

var Address_100_as_Bytes = []byte(Address_100)
var Address_100_with_extra_2_zeroes_prefixed_as_bytes = []byte(Address_100_with_extra_2_zeroes_prefixed)
var Address_100_with_extra_2_numbers_prefixed_as_bytes = []byte(Address_100_with_extra_2_numbers_prefixed)

var Address_100_Hex_with_0x_as_Bytes = []byte(Address_100_Hex_with_0x)

var Hash_Address_100_as_Bytes = sha256.Sum256(Address_100_as_Bytes)

func TestNewAddress(t *testing.T) {
	a := NewAddress(Address_100_Hex_with_0x)
	if !bytes.Equal(a.Bytes(), Address_100_as_Bytes) {
		t.Fail()
	}
}

func TestBytesToAddress(t *testing.T) {
	a1 := NewAddress(Address_100_Hex_with_0x)
	a2 := BytesToAddress(Address_100_as_Bytes)
	if !a1.Equal(a2) {
		t.Fail()
	}
}

func TestSetBytes(t *testing.T) {
	var a1, a2 Address
	a1.SetBytes(Address_100_with_extra_2_zeroes_prefixed_as_bytes)
	a2.SetBytes(Address_100_with_extra_2_numbers_prefixed_as_bytes)

	a3 := NewAddress(Address_100_Hex_with_0x)

	if !a1.Equal(a3) {
		t.Fail()
	}

	if !a2.Equal(a3) {
		t.Fail()
	}
}

func TestHash(t *testing.T) {
	a := NewAddress(Address_100_Hex_with_0x)
	if Hash_Address_100_as_Bytes != a.Hash() {
		t.Fail()
	}
}

func TestString(t *testing.T) {
	a := NewAddress(Address_100_Hex_with_0x)
	if a.String() != Address_100_Hex_with_0x {
		t.Fail()
	}
}

func TestMarshalText(t *testing.T) {
	a := NewAddress(Address_100_Hex_with_0x)
	m, err := a.MarshalText()
	if err != nil || !bytes.Equal(m, Address_100_Hex_with_0x_as_Bytes) {
		t.Fail()
	}
}

func TestUnmarshalText(t *testing.T) {
	var a Address
	a.UnmarshalText(Address_100_Hex_with_0x_as_Bytes)
	if !a.Equal(NewAddress(Address_100_Hex_with_0x)) {
		t.Fail()
	}

}
