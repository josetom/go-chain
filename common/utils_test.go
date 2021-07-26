package common_test

import (
	"bytes"
	"testing"

	"github.com/josetom/go-chain/common"
	"github.com/josetom/go-chain/test_helper"
)

func TestBytesToHash(t *testing.T) {
	if test_helper.Hash_Address_100_as_Bytes != common.BytesToHash(test_helper.Address_100_as_Bytes) {
		t.Fail()
	}
}

func TestHexToBytes(t *testing.T) {
	// string with prefix 0x
	if !bytes.Equal(test_helper.Address_100_as_Bytes, common.Hex2Bytes(test_helper.Address_100_Hex_without_0x)) {
		t.Fail()
	}
	// string without prefix 0x
	if !bytes.Equal(test_helper.Address_100_as_Bytes, common.Hex2Bytes(test_helper.Address_100_Hex_with_0x)) {
		t.Fail()
	}
	// len of string is odd, removing the first 0 in address_100
	if !bytes.Equal(common.Hex2Bytes(test_helper.Address_100), common.Hex2Bytes(test_helper.Address_100[1:])) {
		t.Fail()
	}
}

func TestHas0xPrefix(t *testing.T) {
	if !common.Has0xPrefix(test_helper.Address_100_Hex_with_0x) {
		t.Fail()
	}
	if common.Has0xPrefix(test_helper.Address_100_Hex_without_0x) {
		t.Fail()
	}
}

func TestBytes2Hex(t *testing.T) {
	if common.Bytes2Hex(test_helper.Address_100_as_Bytes, true) != test_helper.Address_100_Hex_with_0x {
		t.Fail()
	}
	if common.Bytes2Hex(test_helper.Address_100_as_Bytes, false) != test_helper.Address_100_Hex_without_0x {
		t.Fail()
	}
}

func TestMarshalUtil(t *testing.T) {
	if !bytes.Equal(common.MarshalUtil(test_helper.Address_100_as_Bytes), test_helper.Address_100_Hex_with_0x_as_Bytes) {
		t.Fail()
	}
}

func TestUnmarshalUtil(t *testing.T) {
	u, err := common.UnmarshalUtil(test_helper.Address_100_Hex_with_0x_as_Bytes)
	if err != nil || !bytes.Equal(u, test_helper.Address_100_as_Bytes) {
		t.Fail()
	}
	u1, err := common.UnmarshalUtil(make([]byte, 0))
	if u1 != nil || err != nil {
		t.Fail()
	}
}

func TestBytesHave0xPrefix(t *testing.T) {
	if !common.BytesHave0xPrefix(test_helper.Hash_Address_100_with_0x_as_Bytes) {
		t.Fail()
	}
	if common.BytesHave0xPrefix(test_helper.Address_100_Hex_without_0x_as_Bytes) {
		t.Fail()
	}
}

func TestDeepCopy(t *testing.T) {
	var b interface{}
	common.DeepCopy(test_helper.Address_100_Hex_with_0x, &b)
	if test_helper.Address_100_Hex_with_0x != b {
		t.Fail()
	}
}
