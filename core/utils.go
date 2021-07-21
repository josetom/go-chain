package core

import "fmt"

// Returns with padded 0s upto length 10
// Will work for the next century
func getBlockNumberAsIndexBytes(blockNumber uint64) []byte {
	return []byte(fmt.Sprintf("%010d", blockNumber))
}
