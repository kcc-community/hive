package main

import (
	"context"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/hive/hivesim"
)

var (
	MinerKey           = "0x9c647b8b7c4e7c3490668fb6c11473619db80c93704c70893d3813af4090c39c"
	MinerAddr          = common.HexToAddress("0x658bdf435d810c91414ec09147daa6db62406379")
	ValidatorsContract = common.HexToAddress("0x000000000000000000000000000000000000f333")
)

func main() {
	suite := hivesim.Suite{
		Name:        "migrate peg token ",
		Description: "",
	}

	newParameters := func() hivesim.Params {
		base := hivesim.Params{
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
			// KCC Amazon Hardfork block
			"HIVE_FORK_KCC_AMAZON": "12",
		}

		return base
	}

	suite.Add(hivesim.ClientTestSpec{
		Role: "eth1",
		Name: "migrate USDT",
		Files: map[string]string{
			"/genesis.json": "genesis.json",
		},
		Parameters: newParameters(),
		Run:        MigrateUSDT,
	})

	hivesim.MustRunSuite(hivesim.New(), suite)
}

func MigrateUSDT(t *hivesim.T, c *hivesim.Client) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	//
	// Wait for all the hard forks to be activated
	//   - wait until block 13

	for {

		select {
		case <-ctx.Done():
			t.Fatalf("timeout when getting receipt: %v", ctx.Err())
		default:
		}

		// wait for the block number to be greater or equal to 13
		block, err := ethclient.NewClient(c.RPC()).BlockByNumber(ctx, nil)

		if err != nil {
			t.Fatalf("failed to get block: %v", err)
		}

		if block.Number().Uint64() >= 13 {
			break
		}

		time.Sleep(time.Second)
	}

	var (
		usdtContractAddr = common.HexToAddress("0x0039f574eE5cC39bdD162E9A88e3EB1f111bAF48")
		migratorAddr     = common.HexToAddress(os.Getenv("MIGRATOR_ADDR"))
		nonMigratorAddr  = MinerAddr // the miner is not a migrator
	)

	//
	// 1. check the USDT balance of the USDT contract
	//

	balance := new(big.Int)

	if err := c.RPC().CallContext(ctx, balance, "eth_call", map[string]interface{}{
		"from": nil,
		"to":   usdtContractAddr,
		// 0x70a08231 is the signature of the balanceOf function
		"data": "0x70a08231000000000000000000000000" + migratorAddr.Hex()[2:],
	}, "latest"); err != nil {
		t.Fatalf("failed to call balanceOf: %v", err)
	}

	// print balance
	t.Logf("migrator balance: %s", balance.String())

	// the balance must be greater than 0
	if balance.Cmp(big.NewInt(0)) <= 0 {
		t.Fatal("migrator balance is not enough")
	}

	//
	// 2. The non-migrator should not be able to migrate
	//
	if err := c.RPC().CallContext(ctx, nil, "eth_call", map[string]interface{}{
		"from": nonMigratorAddr,
		"to":   usdtContractAddr,
		// The signature of the migrate function is "migrate(address, address)"
		"data": "0x1068361f000000000000000000000000" + usdtContractAddr.Hex()[2:] + "000000000000000000000000" + migratorAddr.Hex()[2:],
	}); err == nil {
		t.Fatal("non-migrator should not be able to migrate")
	} else {
		// err is not nil
		// the error should contains "only migrator role"
		if !strings.Contains(err.Error(), "only migrator role") {
			t.Fatalf("unexpected error: %v", err)
		}
	}

	//
	// 3. The migrator should be able to migrate
	//
	if err := c.RPC().CallContext(ctx, nil, "eth_call", map[string]interface{}{
		"from": migratorAddr,
		"to":   usdtContractAddr,
		// The signature of the migrate function is "migrate(address, address)"
		"data": "0x1068361f000000000000000000000000" + usdtContractAddr.Hex()[2:] + "000000000000000000000000" + migratorAddr.Hex()[2:],
	}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}

}
