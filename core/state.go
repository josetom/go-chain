package core

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/josetom/go-chain/common"
)

type State struct {
	Balances  map[Address]uint
	txMemPool []Transaction

	dbFile *os.File
	time   time.Time

	latestBlock     Block
	latestBlockHash common.Hash
}

func (s *State) loadStateFromDisk() (*State, error) {
	dbPath := GetBlocksDbPath()
	f, err := os.OpenFile(dbPath, os.O_APPEND|os.O_RDWR, 0600)
	if err != nil {
		log.Print("unable to open db ", dbPath)
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	s.dbFile = f

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Print("error while scanning", err)
			return nil, err
		}

		var blockFS BlockFS
		json.Unmarshal(scanner.Bytes(), &blockFS)

		// TODO : Validate blocks against block hash
		if err := s.applyBlock(blockFS.Block, blockFS.Hash); err != nil {
			return nil, err
		}
	}

	return s, nil
}

func LoadState() (*State, error) {
	genesisContent, err := loadGenesis()
	if err != nil {
		return nil, err
	}

	balances := make(map[Address]uint)

	for address, balance := range genesisContent.Balances {
		balances[address] = balance
	}

	state := &State{
		Balances:        balances,
		txMemPool:       make([]Transaction, 0),
		dbFile:          nil,
		time:            time.Now(),
		latestBlock:     Block{},
		latestBlockHash: common.BytesToHash(nil),
	}

	return state.loadStateFromDisk()
}

func (s *State) applyBlock(b Block, h common.Hash) error {
	for _, tx := range b.Transactions {
		if err := s.applyTransaction(tx); err != nil {
			return err
		}
	}

	s.latestBlock = b
	s.latestBlockHash = h
	return nil
}

func (s *State) applyTransaction(tx Transaction) error {
	if tx.IsReward() {
		s.Balances[tx.To()] += tx.Value()
		return nil
	}
	if s.Balances[tx.From()] < tx.Value() {
		return fmt.Errorf("insufficient_balance")
	}
	s.Balances[tx.From()] -= tx.Value()
	s.Balances[tx.To()] += tx.Value()

	return nil
}

func (s *State) AddTransaction(tx Transaction) error {
	if err := s.applyTransaction(tx); err != nil {
		return err
	}
	s.txMemPool = append(s.txMemPool, tx)

	return nil
}

func (s *State) Persist() (common.Hash, error) {
	// TODO : will multiple goroutines access this and result in loss of txns ?
	// If not, the below 3 lines can be removed and optimised
	mempoolLength := len(s.txMemPool)
	mempool := make([]Transaction, mempoolLength)
	copy(mempool, s.txMemPool)

	// Create a new block
	block := NewBlock(s.latestBlockHash, s.NextBlockNumber(), uint64(time.Now().UnixNano()), mempool)
	blockHash, err := block.Hash()
	if err != nil {
		return common.Hash{}, err
	}

	// Persist the new block to file system
	blockFS := BlockFS{blockHash, block}
	blockFsJson, err := json.Marshal(blockFS)
	if err != nil {
		return common.Hash{}, err
	}
	if _, err = s.dbFile.Write(append(blockFsJson, '\n')); err != nil {
		return common.Hash{}, err
	}
	log.Println("Block created", blockHash)

	// Update the latesh block & hash to the current one
	s.latestBlock = block
	s.latestBlockHash = blockHash

	// reset the mempool
	// TODO : Can be reset to empty if this is thread safe and optimised along with first 3 lines
	s.txMemPool = s.txMemPool[mempoolLength:]

	return blockHash, nil
}

func (s *State) Close() {
	s.dbFile.Close()
}

func (s *State) LatestBlockHash() common.Hash {
	return s.latestBlockHash
}

func (s *State) LatestBlock() Block {
	return s.latestBlock
}

func (s *State) NextBlockNumber() uint64 {
	return s.latestBlock.Header.Number + 1
}
