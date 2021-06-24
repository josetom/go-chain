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

	currentState  *core.State
	snapshotState *core.State

	isMining bool
}

func InitMiner(state *core.State) Miner {
	snapshotState := state.Copy()

	miner := Miner{
		txnsCh:        make(chan core.Transaction, txnChSize),
		syncBlockCh:   make(chan core.Block, txnChSize),
		currentState:  state,
		snapshotState: &snapshotState,
		pendingTxns:   make(map[string]core.Transaction),
		archivedTxns:  make(map[string]core.Transaction),
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
				m.removeMinedTxns(block)
			}

		case <-ctx.Done():
			ticker.Stop()
		}
	}
}

func (m *Miner) addPendingTxn(txn core.Transaction) {
	_, isAlreadyPending := m.pendingTxns[txn.Hash().String()]
	_, isArchived := m.archivedTxns[txn.Hash().String()]

	if !isAlreadyPending || isArchived {
		m.pendingTxns[txn.Hash().String()] = txn
	}
}

func (m *Miner) mine(ctx context.Context) (core.Block, error) {
	lenPendingTxns := len(m.pendingTxns)
	if lenPendingTxns == 0 {
		return core.Block{}, nil
	}

	snapshotState := m.currentState.Copy()
	m.snapshotState = &snapshotState

	txnArr := make([]core.Transaction, lenPendingTxns)
	common.DeepCopy(txnMapToArray(m.pendingTxns), &txnArr)

	pendingBlock := core.NewBlock(
		m.currentState.LatestBlockHash(),
		m.currentState.NextBlockNumber(),
		uint64(time.Now().UnixNano()),
		common.GenNonce(),
		Config.Miner.Address,
		txnArr,
	)

	block, err := m.mineBlock(ctx, pendingBlock)
	if err != nil {
		return core.Block{}, err
	}

	_, err = m.currentState.AddBlock(block)
	if err != nil {
		return core.Block{}, err
	}

	m.removeMinedTxns(block)

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
	hash, err := pendingBlock.Hash()
	if err != nil {
		return core.Block{}, err
	}
	if core.IsBlockHashValid(hash) {
		return pendingBlock, nil
	}
	pendingBlock.Header.Nonce = common.GenNonce()
	return m.mineBlockHelper(ctx, pendingBlock)
}

func (m *Miner) mineBlock(ctx context.Context, pendingBlock core.Block) (core.Block, error) {
	// Add a reward transaction for the block that is to be mined
	txn := core.NewTransaction(
		common.ZeroAddress,
		Config.Miner.Address,
		uint(core.Config.Block.Reward),
		"reward",
	)
	txn.TxnContent.IsReward = true
	pendingBlock.Transactions = append(pendingBlock.Transactions, txn)

	return m.mineBlockHelper(ctx, pendingBlock)
}
