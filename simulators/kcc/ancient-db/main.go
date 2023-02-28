package main

import (
	"context"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/hive/hivesim"
)

var (
	MinerKey  = "0x9c647b8b7c4e7c3490668fb6c11473619db80c93704c70893d3813af4090c39c"
	MinerAddr = common.HexToAddress("0x658bdf435d810c91414ec09147daa6db62406379")
)

func main() {
	suite := hivesim.Suite{
		Name:        "Ancient db",
		Description: "Testcase for ancient db",
	}
	suite.Add(hivesim.ClientTestSpec{
		Role:        "eth1",
		Name:        "Run the chain",
		Description: "Run a chain to generate ancient db",
		Files: map[string]string{
			"/genesis.json": "genesis.json",
			"/chain.tar.gz": "chain.tar.gz", // Add our pre exported chain
			"/geth.sh":      "geth.sh",      // override the default geth.sh
		},
		Parameters: hivesim.Params{
			"HIVE_CLIQUE_PRIVATEKEY": "9c647b8b7c4e7c3490668fb6c11473619db80c93704c70893d3813af4090c39c",
			"HIVE_MINER":             "658bdf435d810c91414ec09147daa6db62406379",
			"HIVE_CHAIN_ID":          "321",

			// block interval: 1s
			"HIVE_KCC_POSA_BLOCK_INTERVAL": "1",
			// epoch : 5
			"HIVE_KCC_POSA_EPOCH": "5",
			// initial valiators
			"HIVE_KCC_POSA_ISHIKARI_INIT_VALIDATORS": "0x658bdf435d810c91414ec09147daa6db62406379",
			// admin
			"HIVE_KCC_POSA_ADMIN": "0x658bdf435d810c91414ec09147daa6db62406379",
			// KCC Ishikari  fork number
			"HIVE_FORK_KCC_ISHIKARI": "9",
			// KCC Ishikari Patch001 fork number
			"HIVE_FORK_KCC_ISHIKARI_PATCH001": "10",
			// KCC Ishikari Patch002 fork number
			"HIVE_FORK_KCC_ISHIKARI_PATCH002": "11",
		},
		Run: runChain,
	})

	hivesim.MustRunSuite(hivesim.New(), suite)
}

func HexWithPrefixToUint64(hex string) (uint64, error) {
	return strconv.ParseUint(hex[2:], 16, 64)
}

func runChain(t *hivesim.T, c *hivesim.Client) {

	// We need to import exported blocks into the chain before starting it.
	// However, to import more than 110000 blocks into the chain and compact them would cost more than 3 minutes.
	// Remember to start `hive` with an extra option of `--client.checktimelimit 10m` to avoid timeout.

	// All the 110000 blocks are empty. So there are no receipts in the ancient db.

	// Why 110000? Because we will only have ancient items if the block number is greater than 90000.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()

	var number hexutil.Uint64

	// get the block number
	if err := c.RPC().CallContext(ctx, &number, "eth_blockNumber"); err != nil {
		t.Fatalf("failed to get block number: %v", err)
	}

	// our chain already has >110000 blocks
	if number < 110000 {
		t.Fatalf("block number is too small: %d < 110000, failed to import chain?", number)
	}

}
