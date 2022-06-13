package main

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/hive/hivesim"
)

func main() {
	suite := hivesim.Suite{
		Name:        "Large Block",
		Description: "A Large Block",
	}

	suite.Add(hivesim.ClientTestSpec{
		Role:        "eth1",
		Name:        "Test with large block",
		Description: "Test if the gas block limit is expected",
		Files: map[string]string{
			"/genesis.json": "genesis.json",
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
			// Gas ceil
			"HIVE_GAS_TARGET": "13500000",
			// Gas floor
			"HIVE_GAS_LIMIT": "15000000",
		},
		Run: LargeBlock,
	})

	hivesim.MustRunSuite(hivesim.New(), suite)

}

type block struct {
	Number     string `json:"number"`
	Hash       string `json:"hash"`
	ParentHash string `json:"parentHash"`
	GasLimit   string `json:"gasLimit"`
}

func LargeBlock(t *hivesim.T, c *hivesim.Client) {

	if c.Type == "kcc_v1.0.3" {
		t.Logf("kcc v1.0.3 is not used in this case")
		return
	}

	start := time.Now()
	timeout := 120 * time.Second

	for {
		var b block
		if err := c.RPC().Call(&b, "eth_getBlockByNumber", "latest", false); err != nil {
			t.Fatal("eth_getBlockByNumber call failed:", err)
		}

		blockNumber := new(big.Int).SetBytes(common.FromHex(b.Number))
		gasLimit := new(big.Int).SetBytes(common.FromHex(b.GasLimit))

		if blockNumber.Uint64() < 9 {
			t.Logf("block #%v mined, gasLimit: %v", blockNumber, gasLimit.String())
			if time.Since(start) > timeout {
				t.Fatalf("timeout when waiting for #10 to be mined")
			}
			time.Sleep(time.Second)
		} else if blockNumber.Uint64() == 9 {
			t.Logf("kcc Ishikari hardfork block #9 mined, gasLimit: %v", gasLimit.String())

			if gasLimit.Cmp(big.NewInt(8_000_000)) <= 0 {
				t.Fatalf("The gas block limit of block #9 should be greater than 8M")
			}
		} else {
			t.Logf("block #%v mined, gas limit: %v", blockNumber, gasLimit.String())
			break
		}
	}
}
