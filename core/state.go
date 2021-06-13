package core

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/josetom/go-chain/common"
)

type State struct {
	Balances  map[Address]uint
	txMemPool []Transaction

	dbFile *os.File

	latestBlock     Block
	latestBlockHash common.Hash
}

// 1. Initializes empty state
// 2. Load Genesis block
// 3. Loads the state from local db
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
		latestBlock:     Block{},
		latestBlockHash: common.Hash{},
	}

	return state.loadStateFromDisk()
}

// Load blocks from local db and validate
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

		blockFsJson := scanner.Bytes()
		if len(blockFsJson) == 0 {
			break
		}

		var blockFS BlockFS
		json.Unmarshal(blockFsJson, &blockFS)

		if err := s.applyBlock(blockFS.Block); err != nil {
			return nil, err
		}

		s.latestBlock = blockFS.Block
		s.latestBlockHash = blockFS.Hash

	}

	return s, nil
}

// verifies if block can be added to the blockchain.
// Block metadata are verified as well as transactions within (sufficient balances, etc).
func (s *State) applyBlock(b Block) error {

	if b.Header.Number != s.NextBlockNumber() {
		return fmt.Errorf("next expected block must be '%d' not '%d'", s.NextBlockNumber(), b.Header.Number)
	}

	if s.latestBlock.Header.Number > 0 && !reflect.DeepEqual(b.Header.ParentHash, s.latestBlockHash) {
		return fmt.Errorf("next block parent hash must be '%x' not '%x'", s.latestBlockHash, b.Header.ParentHash)
	}

	return s.applyTransactions(b.Transactions)
}

// Validate current balances and update the balances
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

func (s *State) applyTransactions(txs []Transaction) error {
	for _, tx := range txs {
		if err := s.applyTransaction(tx); err != nil {
			return err
		}
	}
	return nil
}

func (s *State) AddTransaction(tx Transaction) error {
	if err := s.applyTransaction(tx); err != nil {
		return err
	}
	s.txMemPool = append(s.txMemPool, tx)

	return nil
}

func (s *State) AddBlocks(blocks []Block) error {
	for _, block := range blocks {
		if _, err := s.AddBlock(block); err != nil {
			return err
		}
	}
	return nil
}

func (s *State) AddBlock(block Block) (common.Hash, error) {

	pendingState := s.copy()
	err := pendingState.applyBlock(block)
	if err != nil {
		return common.Hash{}, err
	}

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
	log.Println("Block added", blockHash)

	// Update the balances, block & hash to the current one
	s.Balances = pendingState.Balances
	s.latestBlock = block
	s.latestBlockHash = blockHash

	//TODO : this is not there in the tutorial
	s.txMemPool = s.txMemPool[len(pendingState.txMemPool):]

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

func (s *State) copy() State {
	c := State{}
	c.latestBlock = s.latestBlock
	c.latestBlockHash = s.latestBlockHash
	c.txMemPool = make([]Transaction, 0)
	c.Balances = make(map[Address]uint)

	for acc, balance := range s.Balances {
		c.Balances[acc] = balance
	}

	c.txMemPool = append(c.txMemPool, s.txMemPool...)

	return c
}

// TODO : This needs to be changed
func (s *State) Persist() (common.Hash, error) {
	// Create a new block
	block := NewBlock(s.latestBlockHash, s.NextBlockNumber(), uint64(time.Now().UnixNano()), s.txMemPool)
	return s.AddBlock(block)
}
