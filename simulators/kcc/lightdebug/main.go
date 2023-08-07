package main

import (
	"context"
	"encoding/json"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/hive/hivesim"
)

var (
	MinerKey           = "0x9c647b8b7c4e7c3490668fb6c11473619db80c93704c70893d3813af4090c39c"
	MinerAddr          = common.HexToAddress("0x658bdf435d810c91414ec09147daa6db62406379")
	ValidatorsContract = common.HexToAddress("0x000000000000000000000000000000000000f333")
)

func main() {
	suite := hivesim.Suite{
		Name:        "lightdebug mode",
		Description: "Testcase for lightdebug mode",
	}

	newParameters := func(extra string) hivesim.Params {
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
		}

		base[extra] = "1"
		return base
	}

	suite.Add(hivesim.ClientTestSpec{
		Role: "eth1",
		Name: "lightdebug(only)",
		Files: map[string]string{
			"/genesis.json": "genesis.json",
		},
		Parameters: newParameters("HIVE_RPC_LIGHT_DEBUG_ONLY"),
		Run:        lightOnly,
	})

	hivesim.MustRunSuite(hivesim.New(), suite)
}

// send a new transaction and return the transaction hash
func sendTransaction(ctx context.Context, t *hivesim.T, c *hivesim.Client) (common.Hash, *big.Int) {

	// send a transaction from miner address to validator contract
	// miner address: 0x658bdf435d810c91414ec09147daa6db62406379
	// validator contract: validatorsContract
	// value: 1 ether

	cli := ethclient.NewClient(c.RPC())

	// hex to private key
	privateKey, err := crypto.HexToECDSA(MinerKey)

	// sign transaction
	nonce, err := cli.PendingNonceAt(ctx, MinerAddr)
	if err != nil {
		t.Fatalf("failed to get nonce: %v", err)
	}

	tx, err := types.SignNewTx(privateKey, types.NewEIP2930Signer(big.NewInt(321)), &types.LegacyTx{
		Nonce:    nonce,
		To:       &ValidatorsContract,
		Value:    big.NewInt(params.Ether),
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
			t.Fatalf("failed to get receipt: %v", err)
		}

		if receipt.BlockNumber != nil {
			return tx.Hash(), receipt.BlockNumber
		}

	}

}

func lightOnly(t *hivesim.T, c *hivesim.Client) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()

	_, blockNum := sendTransaction(ctx, t, c)

	var result json.RawMessage

	// debug_traceBlock
	err := c.RPC().CallContext(ctx, &result, "debug_traceBlock", blockNum)
	if err != nil {
		t.Fatalf("failed to debug_traceBlock: %v", err)
	}

	t.Logf("debug_traceBlock result: %s", result)

	// debug_traceBlockFromFile
	err = c.RPC().CallContext(ctx, &result, "debug_traceBlockFromFile", "non-exist-file")

	if err.Error() != "could not read file: non-exist-file" {
		t.Fatalf("failed to debug_traceBlockFromFile: %v", err)
	}

}
