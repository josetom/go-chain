package core

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/josetom/go-chain/node"
)

type State struct {
	Balances  map[Account]uint
	txMemPool []Transaction

	dbFile *os.File
}

var state *State = &State{make(map[Account]uint), make([]Transaction, 0), nil}

func loadStateFromDisk() (*State, error) {
	txDbPath := filepath.Join(node.Config.DataDir, Config.State.DbFile)
	f, err := os.OpenFile(txDbPath, os.O_APPEND|os.O_RDWR, 0600)
	if err != nil {
		log.Print("unable to open txn file", txDbPath)
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	state.dbFile = f

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Print("error while scanning", err)
			return nil, err
		}

		var tx Transaction
		json.Unmarshal(scanner.Bytes(), &tx)

		if err := state.apply(tx); err != nil {
			return nil, err
		}
	}

	return state, nil
}

func LoadState() (*State, error) {
	genesisContent, err := LoadGenesis()
	if err != nil {
		return nil, err
	}
	for account, balance := range genesisContent.Balances {
		state.Balances[account] = balance
	}
	return loadStateFromDisk()
}

func (s *State) apply(tx Transaction) error {
	if tx.IsReward() {
		state.Balances[tx.To] += tx.Value
		return nil
	}
	if s.Balances[tx.From] <= tx.Value {
		return fmt.Errorf("insufficient_balance")
	}
	s.Balances[tx.From] -= tx.Value
	s.Balances[tx.To] += tx.Value

	return nil
}

func (s *State) Add(tx Transaction) error {
	if err := s.apply(tx); err != nil {
		return err
	}
	s.txMemPool = append(s.txMemPool, tx)

	return nil
}

func (s *State) Persist() error {
	// make a copy of txMemPool to since txMemPool can get txns added while looping
	mempool := make([]Transaction, len(s.txMemPool))
	copy(mempool, s.txMemPool)

	for i := 0; i < len(mempool); i++ {
		txJson, err := json.Marshal(mempool[i])
		if err != nil {
			return err
		}

		if _, err = s.dbFile.Write(append(txJson, '\n')); err != nil {
			return err
		}

		s.txMemPool = s.txMemPool[1:]
	}

	return nil

}

func (s *State) Close() {
	s.dbFile.Close()
}
