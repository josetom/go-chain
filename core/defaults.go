package core

import "github.com/josetom/go-chain/constants"

func Defaults() CoreConfig {
	return CoreConfig{
		State: StateConfig{
			DbFile: "state.db",
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

func getDefaultGenesisBalances() map[Address]uint {
	var defaultGenesisBalances = make(map[Address]uint)
	defaultGenesisBalances[NewAddress("0x0000000000000000000000000000000000000000")] = 1000000000
	return defaultGenesisBalances
}
