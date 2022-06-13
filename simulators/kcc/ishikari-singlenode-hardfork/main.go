package main

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/hive/hivesim"
)

var (
	slotRemoveThreshhold = common.BigToHash(big.NewInt(58))
)

func main() {
	suite := hivesim.Suite{
		Name:        "Ishikari single node hardfork",
		Description: "Single node hardfork ",
	}

	suite.Add(hivesim.ClientTestSpec{
		Role:        "eth1",
		Name:        "kcc Ishikari hardfork",
		Description: "wait for the hardfork",
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
			// KCC Ishikari Patch001 fork number
			"HIVE_FORK_KCC_ISHIKARI_PATCH001": "10",
		},
		Run: hardforkTest,
	})

	hivesim.MustRunSuite(hivesim.New(), suite)

}

type block struct {
	Number     string `json:"number"`
	Hash       string `json:"hash"`
	ParentHash string `json:"parentHash"`
}

func hardforkTest(t *hivesim.T, c *hivesim.Client) {

	if c.Type == "kcc_v1.0.3" {
		t.Logf("kcc v1.0.3 is not used in this case")
		return
	}

	start := time.Now()
	timeout := 6000 * time.Second

	for {
		var b block
		if err := c.RPC().Call(&b, "eth_getBlockByNumber", "latest", false); err != nil {
			t.Fatal("eth_getBlockByNumber call failed:", err)
		}

		blockNumber := new(big.Int).SetBytes(common.FromHex(b.Number))

		if blockNumber.Uint64() < 9 {
			t.Logf("block #%v mined", blockNumber)
			if time.Since(start) > timeout {
				t.Fatalf("timeout when waiting for #10 to be mined")
			}
			time.Sleep(time.Second)
		} else if blockNumber.Uint64() == 9 {
			t.Logf("kcc Ishikari hardfork block #9 mined")
			var removeThreshAsHash common.Hash

			if err := c.RPC().Call(&removeThreshAsHash, "eth_getStorageAt",
				common.HexToAddress("0x000000000000000000000000000000000000f444"),
				slotRemoveThreshhold, rpc.BlockNumber(9)); err != nil {

				t.Fatalf("failed to get 'removeThresh' from punish: %v", err)
			}

			removeThresh := new(big.Int).SetBytes(removeThreshAsHash[:])

			if big.NewInt(48).Cmp(removeThresh) != 0 {
				t.Fatalf("removeThresh should be %v , actually is %v", 48, removeThresh.Int64())
			}

		} else if blockNumber.Uint64() == 10 {
			t.Logf("kcc Ishikari hardfork block #10 mined, patch001 should have been applied")

			var removeThreshAsHash common.Hash

			if err := c.RPC().Call(&removeThreshAsHash, "eth_getStorageAt",
				common.HexToAddress("0x000000000000000000000000000000000000f444"),
				slotRemoveThreshhold, rpc.BlockNumber(10)); err != nil {

				t.Fatalf("failed to get 'removeThresh' from punish: %v", err)
			}

			removeThresh := new(big.Int).SetBytes(removeThreshAsHash[:])

			if big.NewInt(600).Cmp(removeThresh) != 0 {
				t.Fatalf("removeThresh should be %v , actually is %v", 600, removeThresh.Uint64())
			}
		} else {
			t.Logf("block #%v mined", blockNumber)
			break
		}
	}
}
