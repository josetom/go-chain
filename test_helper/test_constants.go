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

	Test_Address_1 = "0xdd6b4d532aad2814bf5ea2bcc5e8939294857e6c"
	Test_Address_2 = "0x054b08ac0c3233efe965a6f24071de1353955e59"

	Hash_0x = "0x0000000000000000000000000000000000000000000000000000000000000000"

	Hash_Block_0 = "0x062b7510997caa2ea51080f20d3f9cef93217b0a8537732e2994b8bd71c1e43c"
	Hash_Block_1 = "0x0177e796fd9e5e7d10c926ecfce6d1c1bc49644ef77dc34886b24207bfe81269"

	Hash_Txn_100_Reward   = "0x6e6cdf6ae97854dcf5111f3aab286aed4efd49355b5d8f46c76472f1a03abbd7"
	Hash_Block_100_Reward = "0x896f4b5111bbef9fac55f2851a6048c6e9004043e7cba3cb16bb7d40c15ed52d"

	DUMMY_DATA = "something else"

	WALLET_PWD = "wallet_pwd"
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
