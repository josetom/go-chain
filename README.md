# Go Chain !
A simple blockchain built with inspiration from geth

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
### Todo
- Transactions
    - Block txn from zero address
- Persist state
    - TBD : Txn Data should not be publicly accessible. Make it public and add a different struct to handle persistance
    - Persist to leveldb
- Node
    - Sync should happen only after the blocks are synced first time
- Websockets
- Coinbase (miner account) should be set in genesis file
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
cd ~/Libaray/go-chain
ln -s ../../pet-project/blockchain/.database/database database
```

# Demo
```sh
./go-chain  tx add --from="0xdd6b4d532aad2814bf5ea2bcc5e8939294857e6c" --to="0x054b08ac0c3233efe965a6f24071de1353955e59" --value=50 --data="test"
```