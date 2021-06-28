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

	Hash_Block_0 = "0xc0c869808bc0aaf964341f9f59b49b9b67ccc52997a6a5bc72c0fd6b648a26f0"
	Hash_Block_1 = "0xe8db397923f5d4cdc605308b5c5a4097fce898b71359a7f61beb74e784d12580"

	Hash_Txn_100_Reward   = "0x6e6cdf6ae97854dcf5111f3aab286aed4efd49355b5d8f46c76472f1a03abbd7"
	Hash_Block_100_Reward = "0xb9172dc5d46b4c8c3bab36bad278f96130d8d61ccba19506865fad00229b54d6"

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
