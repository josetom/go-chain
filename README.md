# Go Chain !

## Story 
### Supported
- Genesis
- Load persisted state
- Add Transactions
- Hashing
- Persist state
    - To a file in json format
### Todo
- Persist state
    - TBD : Txn Data should not be publicly accessible. Make it public and add a different struct to handle persistance
    - Persist to leveldb
- Blocks
- Validate blocks 
- Mine blocks
- Create Wallet
- Synchronize new nodes
- ...??

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