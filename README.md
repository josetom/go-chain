# Go Chain !
A simple blockchain built with inspiration from geth

## Story 
### Supported
- Genesis
- Load persisted state
- Add Transactions
- Hashing
- Persist state
    - To a file in json format
- Blocks
    - Persist blocks in state.db
- Http Server
- Node
    - Peers
    - Sync
### Todo
- Persist state
    - TBD : Txn Data should not be publicly accessible. Make it public and add a different struct to handle persistance
    - Persist to leveldb
- Blocks
    - Validate blocks
    - Mine blocks
    - Change the current state persistance to mining based approach
- Node
    - Sync should happen only after the blocks are synced first time
- Create Wallet
- Websockets
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
```

# Quick hacks
## If you want to view database files
```sh
mkdir .database
cd ~/Libaray/go-chain
ln -s ../../pet-project/blockchain/.database/database database
```