# Go Chain !

## Story 
### Supported
- Genesis
- Load persisted state
- Add Transactions
- Persist state

### Todo
- Blocks
- Validate blocks 
- Mine blocks
- Create Wallet
- Synchronize new nodes
- ...??

# Development
```sh
go build ./cmd/gbc/...
./gbc --help
```

# Testing
```sh
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```