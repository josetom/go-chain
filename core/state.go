package core

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/josetom/go-chain/common"
)

type State struct {
	Balances map[common.Address]uint

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
	genesisHash, err := genesisContent.Hash()
	if err != nil {
		return nil, err
	}

	balances := make(map[common.Address]uint)

	for address, balance := range genesisContent.Balances {
		balances[address] = balance
	}

	state := &State{
		Balances:        balances,
		dbFile:          nil,
		latestBlock:     Block{},
		latestBlockHash: genesisHash,
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

	blockHash, err := b.Hash()
	if err != nil {
		return err
	}

	isBlockValid, err := b.IsBlockHashValid()
	if err != nil {
		return err
	}

	if !isBlockValid {
		return fmt.Errorf("invalid block hash %x", blockHash)
	}

	err = s.applyTransactions(b.Transactions)
	if err != nil {
		return err
	}

	s.Balances[b.Header.Miner] += uint(b.Header.Reward)

	return nil
}

// Validate current balances and update the balances
func (s *State) applyTransaction(tx Transaction) error {
	isAuthentic, err := tx.IsAuthentic()
	if !isAuthentic {
		return fmt.Errorf("not_authentic")
	}
	if err != nil {
		return err
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

	return nil
}

func (s *State) AddBlock(block Block) (common.Hash, error) {

	pendingState := s.Copy()
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

func (s *State) Copy() State {
	c := State{}
	c.latestBlock = s.latestBlock
	c.latestBlockHash = s.latestBlockHash
	c.Balances = make(map[common.Address]uint)

	for acc, balance := range s.Balances {
		c.Balances[acc] = balance
	}

	return c
}
