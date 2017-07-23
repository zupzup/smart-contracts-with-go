# smart-contracts-with-go
A simple example of how to deploy and interact with ETH smart contracts using Go on a simulated Blockchain.

## Prerequisites

* [solc](http://solidity.readthedocs.io/en/develop/installing-solidity.html)
* geth (go-ethereum)

```bash
go get github.com/ethereum/go-ethereum
cd $GOPATH/src/github.com/ethereum/go-ethereum/
make
make devtools
```

## Generating contract.go

```bash
abigen --sol=Contract.sol --pkg=main --out=contract.go
```

## Running

```bash
go build . && ./smart-contracts-with-go
```

