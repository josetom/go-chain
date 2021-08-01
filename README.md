# Go Chain !
A simple blockchain built with inspiration from geth

[![API Reference](
https://camo.githubusercontent.com/915b7be44ada53c290eb157634330494ebe3e30a/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f676f6c616e672f6764646f3f7374617475732e737667
)](https://pkg.go.dev/github.com/josetom/go-chain?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/josetom/go-chain)](https://goreportcard.com/report/github.com/josetom/go-chain)
[![CI](https://github.com/josetom/go-chain/actions/workflows/ci.yml/badge.svg)](https://github.com/josetom/go-chain/actions/workflows/ci.yml)

## Story 
### Supported
- Genesis
    - Initializes blockchain from genesis.json
    - creates block 0 from genesis.json data
- State
    - Holds the account balances
- Transactions
- Hashing
- Blocks
    - Validate blocks
    - Mine blocks
- Server
    - Http
- Node
    - Peers
    - Sync
- Wallets
- LevelDB
    - Stores the blocks against hash
    - Has an index to point block number against hash. Iterator is run on this index
- Consensus
    - Proof of work
### Todo
- Server
    - Websockets
    - GRPC
- State
    - Persist State
    - Incoming longer block
    - Isolate block, index databases
- Consensus
    - Proof of stake
- Transactions
    - Block txn from zero address
    - Increment txn nonce for signer
    - Txns should be atomic
- Wallet
    - Import private key
- Signer
    - Implement better way to sign the transactions
- Node
    - Node Version
    - Sync should happen only after the blocks are synced first time
- Miner
    - Isolate miner from regular nodes
- Smart contracts
    - Wasm ?
- ...??
    - Fix TODO comments
# Development
```sh
go build ./cmd/go-chain/...
./go-chain --help
```

# Testing
```sh
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
go tool cover -func coverage.out | grep total | awk '{print substr($3, 1, length($3)-1)}'
```

# Quick hacks
## If you want to view database files
```sh
mkdir .database
cd ~/Libaray
ln -s ../../pet-project/blockchain/.database go-chain
```

# Demo
## Run console
```sh
./go-chain run
```
## Create Wallet
```sh
./go-chain wallet new-account
./go-chain wallet print-pk --address="0xdd6b4d532aad2814bf5ea2bcc5e8939294857e6c"
```
## Add Txn
Chain console should be running when transaction commands are called
```sh
./go-chain  tx add --from="0xdd6b4d532aad2814bf5ea2bcc5e8939294857e6c" --to="0x054b08ac0c3233efe965a6f24071de1353955e59" --value=50 --data="test"
```