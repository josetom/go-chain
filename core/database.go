package core

import (
	"bufio"
	"encoding/json"
	"os"
	"reflect"

	"github.com/josetom/go-chain/common"
)

// TODO : this needs to be fixed and changed to query from leveldb
func (s *State) GetBlocksAfter(hash common.Hash) ([]Block, error) {
	f, err := os.OpenFile(GetBlocksDbPath(), os.O_RDONLY, 0600)
	if err != nil {
		return nil, err
	}

	blocks := make([]Block, 0)
	shouldStartCollecting := false

	if reflect.DeepEqual(hash, common.Hash{}) {
		shouldStartCollecting = true
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		var blockFs BlockFS
		err = json.Unmarshal(scanner.Bytes(), &blockFs)
		if err != nil {
			return nil, err
		}

		if shouldStartCollecting {
			blocks = append(blocks, blockFs.Block)
			continue
		}

		if hash == blockFs.Hash {
			shouldStartCollecting = true
		}
	}

	return blocks, nil
}
