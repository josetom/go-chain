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

	Hash_Block_0 = "0x0cbad312fcd82cecf212a891c5b5393eb575bce14d94ed06fd75e2969407946e"
	Hash_Block_1 = "0x0e50d543ae455db1510d73c27a168bcf0042947cc6e24e9f8fdda4f6e194d97d" //"0xe30cee4ceb3785a2d96f0e00583303b605d4ef653da1e47c9ff9c61fa5567854"

	Hash_Txn_100_Reward   = "0xca72344a9b5b336e8f6c000715d890ee5e976fe8d09de9658dc213699c05375d"
	Hash_Block_100_Reward = "0x2d9b23342e3f76c4126d8e3529a578513fc7394bd326cf0f74cde785ab5ca564"

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
