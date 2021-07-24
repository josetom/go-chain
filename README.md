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
- State
- Transactions
- Hashing
- Blocks
    - Validate blocks
    - Mine blocks
- Http Server
- Node
    - Peers
    - Sync
- Wallets
- LevelDB
    - Stores the blocks against hash
    - Has an index to point block number against hash. Iterator is run on this index
### Todo
- Persist state
    - TBD : Txn Data should not be publicly accessible. Make it public and add a different struct to handle persistance
- Websockets
- GRPC
- Concurrency checks
- Proof of stake
- Misc items
    - Transactions
        - Block txn from zero address
        - Increment txn nonce for signer
    - Wallet
        - Import private key
    - Node
        - Node Version
        - Sync should happen only after the blocks are synced first time
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