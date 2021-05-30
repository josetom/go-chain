package common

import (
	"crypto/sha256"
	"encoding/hex"
)

// Converts bytes to sha256 hash
func BytesToHash(b []byte) Hash {
	return sha256.Sum256(b)
}

// converts the hexadecimal string to bytes
// removes the 0x from the hex string if prefix is present
// if odd number of characters, prefix hex string with 0
func Hex2Bytes(s string) []byte {
	if has0xPrefix(s) {
		s = s[2:]
	}
	if len(s)%2 == 1 {
		s = "0" + s
	}
	return hex2BytesHelper(s)
}

// checks if a string is of the format 0xSomething
func has0xPrefix(s string) bool {
	return len(s) > 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X')
}

func hex2BytesHelper(s string) []byte {
	h, _ := hex.DecodeString(s)
	return h
}

func Bytes2Hex(b []byte, has0xPrefix bool) string {
	h := hex.EncodeToString(b)
	if has0xPrefix {
		return "0x" + h
	}
	return h
}

func MarshalUtil(b []byte) []byte {
	result := make([]byte, len(b)*2+2)
	copy(result, "0x")
	hex.Encode(result[2:], b)
	return result
}

// TODO : validate the logic in go-eth. it looks little different
func UnmarshalUtil(b []byte) ([]byte, error) {
	if len(b) == 0 {
		return nil, nil
	}
	if bytesHave0xPrefix(b) {
		b = b[2:]
	}
	result := make([]byte, hex.DecodedLen(len(b)))
	hex.Decode(result, b)
	return result, nil
}

func bytesHave0xPrefix(b []byte) bool {
	return len(b) >= 2 && b[0] == '0' && (b[1] == 'x' || b[1] == 'X')
}
