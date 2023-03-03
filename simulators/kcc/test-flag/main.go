package main

import (
	"context"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/hive/hivesim"
)

var (
	TestnetGenesisHash = common.HexToHash("0x3e98a8f1704be5accc0ce9abf3e852e24ec030a9dda6c1beaf72c600f85cd417")
)

func main() {
	suite := hivesim.Suite{
		Name:        "testnet flag",
		Description: "Does the 'testnet' flag still work?",
	}
	suite.Add(hivesim.ClientTestSpec{
		Role:        "eth1",
		Name:        "run the testnet",
		Description: "start with '--testnet' flag, and the genesis block should be the KCC-testnet genesis block",
		Files: map[string]string{
			"/geth.sh": "geth.sh",
		}, // no files
		Parameters: hivesim.Params{
			"HIVE_TESTNET": "1", // add --testnet flag
		},
		Run: runChain,
	})

	hivesim.MustRunSuite(hivesim.New(), suite)
}

func HexWithPrefixToUint64(hex string) (uint64, error) {
	return strconv.ParseUint(hex[2:], 16, 64)
}

func runChain(t *hivesim.T, c *hivesim.Client) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := ethclient.NewClient(c.RPC())

	if block, err := client.BlockByNumber(ctx, big.NewInt(0)); err != nil {
		t.Fatalf("failed to get block: %v", err)
	} else {
		if block.Hash() != TestnetGenesisHash {
			t.Fatalf("genesis block hash is not correct, expected %s, got %s",
				TestnetGenesisHash.Hex(),
				block.Hash().String())
		}
	}

}
