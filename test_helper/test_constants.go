package test_helper

import "crypto/sha256"

const (
	Address_0_with_0x = "0x0000000000000000000000000000000000000000"

	Address_100 = "00000000000000000100"
	Address_200 = "00000000000000000200"

	Address_100_Hex_with_0x = "0x3030303030303030303030303030303030313030"
	Address_200_Hex_with_0x = "0x3030303030303030303030303030303030323030"

	Address_100_Hex_without_0x = "3030303030303030303030303030303030313030"
	Address_200_Hex_without_0x = "3030303030303030303030303030303030323030"

	Address_100_with_extra_2_zeroes_prefixed  = "0000000000000000000100"
	Address_100_with_extra_2_numbers_prefixed = "1200000000000000000100"

	Hash_0x = "0x0000000000000000000000000000000000000000000000000000000000000000"

	Hash_Block_0 = "0x49fcfda04bf1fa758e3bd1419c6d41af962c785982af36b6ffa8c0959c51efb0"
	Hash_Block_1 = "0x4ad74d1283563d60fde33eefa8b08bcaf9c3fef1813d0c71007147d9c3dd12c0"

	Hash_Txn_100_Reward   = "0x81d2c4a516ec000d4080a36e4abbbdeb2c2d47952f32be97e8316530ab569e9c"
	Hash_Block_100_Reward = "0x198ad94ce7b309e120977398a8f205645e1939020ad65c10a2b72dce0005ee9f"

	REWARD     = "reward"
	DUMMY_DATA = "something else"
)

var Address_100_as_Bytes = []byte(Address_100)
var Address_200_as_Bytes = []byte(Address_200)

var Address_100_Hex_with_0x_as_Bytes = []byte(Address_100_Hex_with_0x)
var Address_100_Hex_without_0x_as_Bytes = []byte(Address_100_Hex_without_0x)

var Hash_Address_100_as_Bytes = sha256.Sum256(Address_100_as_Bytes)

var Hash_Address_100_with_0x = "0xf572455bfe4edc8964b3197d07d1f27c6dc16cfaf250fbdc7eaa36e2f6304864"
var Hash_Address_100_with_0x_as_Bytes = []byte(Hash_Address_100_with_0x)

var Address_100_with_extra_2_zeroes_prefixed_as_bytes = []byte(Address_100_with_extra_2_zeroes_prefixed)
var Address_100_with_extra_2_numbers_prefixed_as_bytes = []byte(Address_100_with_extra_2_numbers_prefixed)
