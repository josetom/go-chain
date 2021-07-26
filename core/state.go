package core

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/josetom/go-chain/common"
	"github.com/josetom/go-chain/db"
	"github.com/josetom/go-chain/db/types"
)

type State struct {
	Balances map[common.Address]uint
	db       types.Database

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
		latestBlock:     Block{},
		latestBlockHash: genesisHash,
	}

	return state.loadStateFromDisk()
}

// Load blocks from local db and validate
func (s *State) loadStateFromDisk() (*State, error) {
	database, err := db.NewDatabase(GetBlocksDbPath())
	if err != nil {
		return nil, err
	}
	s.db = database

	iter := s.db.NewIterator([]byte(INDEX_BLOCK_NUMBER), getBlockNumberAsIndexBytes(1), nil)

	for iter.Next() {
		blockFS, err := s.GetBlockWithHashBytes(iter.Value())
		if err != nil {
			return nil, err
		}

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

func (s *State) ValidateTxn(tx Transaction) error {
	isAuthentic, err := tx.IsAuthentic()
	if !isAuthentic {
		return fmt.Errorf("not_authentic")
	}
	if err != nil {
		return err
	}
	if s.Balances[tx.From()] < tx.Cost() {
		return fmt.Errorf("insufficient_balance")
	}
	return nil
}

// Validate current balances and update the balances
func (s *State) applyTransaction(tx Transaction) error {
	if err := s.ValidateTxn(tx); err != nil {
		return err
	}
	s.Balances[tx.From()] -= tx.Cost()
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

	blockKey := append([]byte(INDEX_BLOCK_HASH), blockHash.Bytes()...)
	blockNumberIndex := append([]byte(INDEX_BLOCK_NUMBER), getBlockNumberAsIndexBytes(block.Header.Number)...)

	// add block and index to db
	batch := s.db.NewBatch()
	if err = batch.Put(blockNumberIndex, blockHash.Bytes()); err != nil {
		return common.Hash{}, err
	}

	if err = batch.Put(blockKey, blockFsJson); err != nil {
		return common.Hash{}, err
	}
	if err = batch.Write(); err != nil {
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
	s.db.Close()
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

func (s *State) GetBlock(blockHash common.Hash) (BlockFS, error) {
	key := append([]byte(INDEX_BLOCK_HASH), blockHash.Bytes()...)
	content, err := s.db.Get(key)
	if err != nil {
		return BlockFS{}, err
	}

	var blockFS *BlockFS
	err = json.Unmarshal(content, &blockFS)
	if err != nil {
		return BlockFS{}, err
	}
	return *blockFS, nil
}

func (s *State) GetBlockWithHashBytes(hashBytes []byte) (BlockFS, error) {
	var h common.Hash
	h.SetBytes(hashBytes)
	return s.GetBlock(h)
}
