package main

import (
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)

	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(10000000000)}
	sim := backends.NewSimulatedBackend(alloc)

	addr, _, contract, err := DeployWinnerTakesAll(auth, sim, big.NewInt(10), big.NewInt(makeTimestamp(time.Now().Add(2*time.Minute))), big.NewInt(makeTimestamp(time.Now().Add(5*time.Minute))))
	if err != nil {
		log.Fatalf("Failed to deploy new token contract: %v", err)
	}
	fmt.Printf("Contract deployed to %s\n", addr.String())
	deadlineCampaign, _ := contract.DeadlineCampaign(nil)
	fmt.Printf("Pre-mining Campaign Deadline: %d\n", deadlineCampaign)

	deadlineCampaign, _ = contract.DeadlineCampaign(&bind.CallOpts{Pending: true})
	fmt.Printf("Pre-mining pending Campaign Deadline: %d\n", deadlineCampaign)

	sim.Commit()

	postDeadlineCampaign, _ := contract.DeadlineCampaign(nil)
	fmt.Printf("Post-mining Campaign Deadline: %d\n", postDeadlineCampaign)
	postDeadlineProjects, _ := contract.DeadlineProjects(nil)
	fmt.Printf("Post-mining Projects Deadline: %d\n", postDeadlineProjects)
}

func makeTimestamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}
