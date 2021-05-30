package common

import (
	"bytes"
	"crypto/sha256"
	"testing"
)

const (
	Address_0_with_0x = "0x0000000000000000000000000000000000000000"

	Address_100 = "00000000000000000100"
	Address_200 = "00000000000000000200"

	Address_100_Hex_with_0x = "0x3030303030303030303030303030303030313030"
	Address_200_Hex_with_0x = "0x3030303030303030303030303030303030323030"

	Address_100_Hex_without_0x = "3030303030303030303030303030303030313030"
	Address_200_Hex_without_0x = "3030303030303030303030303030303030323030"
)

var Address_100_as_Bytes = []byte(Address_100)
var Address_200_as_Bytes = []byte(Address_200)

var Address_100_Hex_with_0x_as_Bytes = []byte(Address_100_Hex_with_0x)
var Address_100_Hex_without_0x_as_Bytes = []byte(Address_100_Hex_without_0x)

var Hash_Address_100_as_Bytes = sha256.Sum256(Address_100_as_Bytes)

var Hash_Address_100_with_0x = "0xf572455bfe4edc8964b3197d07d1f27c6dc16cfaf250fbdc7eaa36e2f6304864"
var Hash_Address_100_with_0x_as_Bytes = []byte(Hash_Address_100_with_0x)

func TestBytesToHash(t *testing.T) {
	if Hash_Address_100_as_Bytes != BytesToHash(Address_100_as_Bytes) {
		t.Fail()
	}
}

func TestHexToBytes(t *testing.T) {
	// string with prefix 0x
	if !bytes.Equal(Address_100_as_Bytes, Hex2Bytes(Address_100_Hex_without_0x)) {
		t.Fail()
	}
	// string without prefix 0x
	if !bytes.Equal(Address_100_as_Bytes, Hex2Bytes(Address_100_Hex_with_0x)) {
		t.Fail()
	}
	// len of string is odd, removing the first 0 in address_100
	if !bytes.Equal(Hex2Bytes(Address_100), Hex2Bytes(Address_100[1:])) {
		t.Fail()
	}
}

func TestHas0xPrefix(t *testing.T) {
	if !has0xPrefix(Address_100_Hex_with_0x) {
		t.Fail()
	}
	if has0xPrefix(Address_100_Hex_without_0x) {
		t.Fail()
	}
}

func TestBytes2Hex(t *testing.T) {
	if Bytes2Hex(Address_100_as_Bytes, true) != Address_100_Hex_with_0x {
		t.Fail()
	}
	if Bytes2Hex(Address_100_as_Bytes, false) != Address_100_Hex_without_0x {
		t.Fail()
	}
}

func TestMarshalUtil(t *testing.T) {
	if !bytes.Equal(MarshalUtil(Address_100_as_Bytes), Address_100_Hex_with_0x_as_Bytes) {
		t.Fail()
	}
}

func TestUnmarshalUtil(t *testing.T) {
	u, err := UnmarshalUtil(Address_100_Hex_with_0x_as_Bytes)
	if err != nil || !bytes.Equal(u, Address_100_as_Bytes) {
		t.Fail()
	}
}

func TestBytesHave0xPrefix(t *testing.T) {
	if !bytesHave0xPrefix(Hash_Address_100_with_0x_as_Bytes) {
		t.Fail()
	}
	if bytesHave0xPrefix(Address_100_Hex_without_0x_as_Bytes) {
		t.Fail()
	}
}
