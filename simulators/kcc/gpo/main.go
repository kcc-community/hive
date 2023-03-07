package main

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
	"time"

	"github.com/ethereum/hive/hivesim"
)

func main() {
	suite := hivesim.Suite{
		Name:        "minimum gas of gpo",
		Description: "Check GPO is compatible",
	}

	suite.Add(hivesim.ClientTestSpec{
		Role:        "eth1",
		Name:        "min gas of gpo",
		Description: "check min gas price is compatible",
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
		Run: minGasPrice,
	})

	hivesim.MustRunSuite(hivesim.New(), suite)

}


func minGasPrice(t *hivesim.T, c *hivesim.Client) {
	time.Sleep(3*time.Second)
	i := 0
	for {
		i++
		time.Sleep(3* time.Second)
		var g string
		err := c.RPC().Call(&g, "eth_gasPrice")
		if err != nil {
			t.Fatal("failed to call eth_gasPrice: ", err)
			return
		}

		t.Logf("call method, gas: %v", new(big.Int).SetBytes(common.FromHex(g)))
		if i == 3 {
			break
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 3)
	defer cancel()
	gas, err := ethclient.NewClient(c.RPC()).SuggestGasPrice(ctx)
	if err != nil {
		t.Fatal("failed to call eth_gasPrice: ", err)
		return
	}

	t.Logf("client way, gas: %v", gas)
	if gas.Cmp(new(big.Int).SetUint64(3*params.GWei)) == 0 {
		return
	} else {
		t.Fatal("failed to get fine gas price, gas: %v", gas)
	}
}
