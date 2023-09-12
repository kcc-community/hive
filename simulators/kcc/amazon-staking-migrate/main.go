package main

import (
	"context"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/hive/hivesim"
)

var (
	NonMinerKey        = "0xdaca75e8c7d5c6fa9a9f7a04ef0a41e928e266fe96ea38c9283b925e9d249ff5"
	NonMinerAddr       = common.HexToAddress("0xF466aB15e342440ec358b2f8548c668704163891")
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
		Run:        MigrateStaking,
	})

	hivesim.MustRunSuite(hivesim.New(), suite)
}

func sendTransaction(ctx context.Context, t *hivesim.T, c *hivesim.Client, key string, to common.Address, value *big.Int) {

	cli := ethclient.NewClient(c.RPC())

	// hex to private key
	privateKey, err := crypto.HexToECDSA(key[2:])
	if err != nil {
		t.Fatalf("failed to convert key: %v", err)
	}

	addr := crypto.PubkeyToAddress(privateKey.PublicKey)

	// sign transaction
	nonce, err := cli.PendingNonceAt(ctx, addr)
	if err != nil {
		t.Fatalf("failed to get nonce: %v", err)
	}

	tx, err := types.SignNewTx(privateKey, types.NewEIP2930Signer(big.NewInt(321)), &types.LegacyTx{
		Nonce:    nonce,
		To:       &to,
		Value:    value,
		Gas:      500000,
		GasPrice: big.NewInt(params.GWei * 100),
	})
	if err != nil {
		t.Fatalf("failed to sign tx: %v", err)
	}

	err = ethclient.NewClient(c.RPC()).SendTransaction(ctx, tx)
	if err != nil {
		t.Fatalf("failed to send tx: %v", err)
	}

	for {

		select {
		case <-ctx.Done():
			t.Fatalf("timeout when getting receipt: %v", err)
		default:
		}

		// wait for the transaction to be mined
		receipt, err := cli.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			if err.Error() != "not found" {
				t.Fatalf("failed to get receipt: %v", err)
			}
			time.Sleep(time.Second)
			continue
		}

		if receipt.BlockNumber != nil {
			return
		}

	}

}

func MigrateStaking(t *hivesim.T, c *hivesim.Client) {

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

	//
	// 1. directly send 990000 KCS to ValidatorsContract
	//    - from: NonMinerAddr
	//    - key: NonMinerKey
	//

	sendTransaction(ctx, t, c, NonMinerKey, ValidatorsContract, (new(big.Int)).Mul(big.NewInt(params.Ether), big.NewInt(990000)))

	//
	// 2. The non-miner cannot migrate KCS
	//
	if err := c.RPC().CallContext(ctx, nil, "eth_call", map[string]interface{}{
		"from": NonMinerAddr.Hex(),
		"to":   ValidatorsContract.Hex(),
		"data": "0xc53aabf7", // signature of "recoverGoDaofunds()"
	}, "latest"); err == nil {
		t.Fatalf("non-miner should not be able to migrate KCS")
	} else {
		// err string should contain "must be admin"
		if !strings.Contains(err.Error(), "must be admin") {
			t.Fatalf("unexpected error: %v", err)
		}
	}

	//
	// 3. In our simulation, the admin is the miner
	//
	//
	if err := c.RPC().CallContext(ctx, nil, "eth_call", map[string]interface{}{
		"from": MinerAddr.Hex(),
		"to":   ValidatorsContract.Hex(),
		"data": "0xc53aabf7", // signature of "recoverGoDaofunds()"
	}, "latest"); err != nil {
		t.Fatalf("failed to migrate KCS: %v", err)
	}

}
