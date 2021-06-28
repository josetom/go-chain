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

	Hash_Block_0 = "0xbb20959c70b009f933e62e3f433b6bacdb96814aee07326defe0fa6a45999fa0"
	Hash_Block_1 = "0x3e6f52df27a0fc35927ef6708eb8c3f2784fa00aebb3785f6d4f3b27db3741c0"

	Hash_Txn_100_Reward   = "0xc1d4f07a2f2140a25a0cdc1e4a38f43cdd73474311312da15286450745fb9d2c"
	Hash_Block_100_Reward = "0xcb01eb91540a43f8afa46b94b9a23ad3ba87fe90f9b7ffd60758a1b5efa1f2ba"

	REWARD     = "reward"
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
