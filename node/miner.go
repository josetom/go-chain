package node

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/josetom/go-chain/common"
	"github.com/josetom/go-chain/core"
)

const (
	txnChSize = 4096
)

type Miner struct {
	txnsCh      chan core.Transaction
	syncBlockCh chan core.Block

	archivedTxns map[string]core.Transaction
	pendingTxns  map[string]core.Transaction

	state        *core.State
	pendingState *core.State

	isMining bool
}

func InitMiner(state *core.State) Miner {
	pendingState := state.Copy()

	miner := Miner{
		txnsCh:       make(chan core.Transaction, txnChSize),
		syncBlockCh:  make(chan core.Block, txnChSize),
		state:        state,
		pendingState: &pendingState,
		pendingTxns:  make(map[string]core.Transaction),
		archivedTxns: make(map[string]core.Transaction),
	}

	return miner
}

func (m *Miner) mainLoop(ctx context.Context) {

	var miningCtx context.Context
	var stopCurrentMining context.CancelFunc

	tickerDuration := time.Duration(Config.Miner.TickerDuration)
	ticker := time.NewTicker(tickerDuration * time.Second)

	for {
		select {
		case txn := <-m.txnsCh:
			m.addPendingTxn(txn)

		case <-ticker.C:
			go func() {
				m.isMining = true

				miningCtx, stopCurrentMining = context.WithCancel(ctx)
				block, err := m.mine(miningCtx)
				if err != nil {
					log.Println(err)
				} else {
					if !block.IsEmpty() {
						log.Printf("Mined block %d", block.Header.Number)
					}
				}

				m.isMining = false
			}()

		case block := <-m.syncBlockCh:
			if m.isMining {
				stopCurrentMining()
			}
			m.removeMinedTxns(block)
			m.resetPendingState()

		case <-ctx.Done():
			ticker.Stop()
		}
	}
}

func (m *Miner) addPendingTxn(txn core.Transaction) {
	_, isAlreadyPending := m.pendingTxns[txn.Hash().String()]
	_, isArchived := m.archivedTxns[txn.Hash().String()]

	if !isAlreadyPending || isArchived {
		m.pendingState.AddTransaction(txn)
		m.pendingTxns[txn.Hash().String()] = txn
	}
}

func (m *Miner) mine(ctx context.Context) (core.Block, error) {
	lenPendingTxns := len(m.pendingTxns)
	if lenPendingTxns == 0 {
		return core.Block{}, nil
	}

	txnArr := make([]core.Transaction, lenPendingTxns)
	common.DeepCopy(txnMapToArray(m.pendingTxns), &txnArr)

	pendingBlock := core.NewBlock(
		m.pendingState.LatestBlockHash(),
		m.pendingState.NextBlockNumber(),
		uint64(time.Now().UnixNano()),
		common.GenNonce(),
		Config.Miner.Address,
		core.MINING_ALGO_POW,
		uint64(core.Config.Block.Reward),
		txnArr,
		nil,
	)

	block, err := m.mineBlock(ctx, pendingBlock)
	if err != nil {
		return core.Block{}, err
	}

	m.removeMinedTxns(block)

	_, err = m.state.AddBlock(block)
	if err != nil {
		return core.Block{}, err
	}

	m.resetPendingState()

	return block, err
}

func (m *Miner) removeMinedTxns(block core.Block) {
	for _, txn := range block.Transactions {
		m.archivedTxns[txn.Hash().String()] = txn
		delete(m.pendingTxns, txn.Hash().String())
	}
}

func (m *Miner) mineBlockHelper(ctx context.Context, pendingBlock core.Block) (core.Block, error) {
	select {
	case <-ctx.Done():
		return core.Block{}, fmt.Errorf("mining cancelled for block : %d", pendingBlock.Header.Number)
	default:
	}
	isBlockValid, err := pendingBlock.IsBlockHashValid()
	if err != nil {
		return core.Block{}, err
	}
	if isBlockValid {
		return pendingBlock, nil
	}
	pendingBlock.Header.Nonce = common.GenNonce()
	return m.mineBlockHelper(ctx, pendingBlock)
}

func (m *Miner) mineBlock(ctx context.Context, pendingBlock core.Block) (core.Block, error) {
	return m.mineBlockHelper(ctx, pendingBlock)
}

func (m *Miner) resetPendingState() {
	pendingState := m.state.Copy()
	m.pendingState = &pendingState
}
