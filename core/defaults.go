package core

import (
	"github.com/josetom/go-chain/common"
	"github.com/josetom/go-chain/constants"
)

func Defaults() CoreConfig {
	return CoreConfig{
		State: StateConfig{
			DbFile: "state.db",
		},
		Block: BlockConfig{
			Reward:     100,
			Complexity: 1,
		},
	}
}

var defaultGenesis = Genesis{
	Config: GenesisConfig{
		ChainId: constants.BlockChainName,
	},
	Timestamp: 1620745200000000000,
	Balances:  getDefaultGenesisBalances(),
}

func getDefaultGenesisBalances() map[common.Address]uint {
	var defaultGenesisBalances = make(map[common.Address]uint)
	defaultGenesisBalances[common.ZeroAddress] = 1000000000
	return defaultGenesisBalances
}
