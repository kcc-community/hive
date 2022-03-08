package main

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/hive/hivesim"
)

func main() {
	suite := hivesim.Suite{
		Name:        "KCC v2 multinode hardfork test",
		Description: "",
	}
	suite.Add(hivesim.TestSpec{
		Name:        "multi node hardfork",
		Description: "run 3 nodes to see if they both work well",
		Run:         multiNodesHardFork,
	})
	hivesim.MustRunSuite(hivesim.New(), suite)
}

func multiNodesHardFork(t *hivesim.T) {

	// 3 validators
	// And the first is the validator in genesis.json
	validators := []*Validator{
		{
			miner:      "0x658bdf435d810c91414ec09147daa6db62406379",
			privateKey: "9c647b8b7c4e7c3490668fb6c11473619db80c93704c70893d3813af4090c39c",
		},
		{
			miner:      "0xa885d3767358B3ad7A5aD9dA5d4508580b1D2480",
			privateKey: "6f7bf5805c4da2645f4f53fa1650ac7e0e69c0393910d139668ab6c2528dafbb",
		},
		{
			miner:      "0x2fE42368F0b91f87f5bfab781Cb89520F3ec78aC",
			privateKey: "ea6024bdb4a933d816de47b0bf5ba6372ac63a7a4e7a080e54a2300594007f79",
		},
	}

	// start 3 clients
	for _, v := range validators {
		v.client = t.StartClient("kcc", hivesim.Params{
			"HIVE_CLIQUE_PRIVATEKEY": v.privateKey,
			"HIVE_MINER":             strings.TrimPrefix(v.miner, "0x"),
			"HIVE_CHAIN_ID":          "321",
			// block interval: 1s
			"HIVE_KCC_POSA_BLOCK_INTERVAL": "3",
			// epoch : 5
			"HIVE_KCC_POSA_EPOCH": "5",
			// initial valiators
			"HIVE_KCC_POSA_V2_INIT_VALIDATORS": "0x658bdf435d810c91414ec09147daa6db62406379,0xa885d3767358B3ad7A5aD9dA5d4508580b1D2480,0x2fE42368F0b91f87f5bfab781Cb89520F3ec78aC",
			// admin
			"HIVE_KCC_POSA_ADMIN": "0x658bdf435d810c91414ec09147daa6db62406379",
			// kcc-v2 fork number
			"HIVE_FORK_KCC_V2": "9",
			// sync mode
			"HIVE_NODETYPE": "full",
		}, hivesim.WithStaticFiles(
			map[string]string{
				"/genesis.json": "genesis.json",
			}))

		for {
			// FIXME: timeout
			enode, err := v.client.EnodeURL()

			if err != nil {
				t.Fatalf("failed to get enode: %v", err)
			}

			if enode == "" {
				fmt.Printf("Waiting for the node (validator: %v) to get ready...", v.miner)
				time.Sleep(time.Second)
				continue
			}

			v.enode = enode
			fmt.Printf("enode of %v is %v \n", v.miner, v.enode)
			break
		}

	}

	// connect nodes
	for i := range validators {
		for j := range validators {
			if i != j {
				err := validators[j].client.RPC().Call(nil, "admin_addPeer", validators[i].enode)
				if err != nil {
					t.Fatalf("failed to add %v as a peer of %v", validators[i].miner, validators[j].miner)
				}
			}
		}
	}

	fmt.Printf("nodes are connected..\n")

	// get a client from any of the 3 nodes
	rpcClient := validators[0].client.RPC()
	ethClient := ethclient.NewClient(rpcClient)

	// wait for at least 13 blocks produced
	for {
		// TODO: use context with timeout
		number, err := ethClient.BlockNumber(context.Background())
		if err != nil {
			t.Fatalf("failed to get block number: %v", err)
		}
		fmt.Printf("current block number is %v\n", number)
		if number >= 13 {
			break
		}

		time.Sleep(time.Second * 2)
	}

	// block 10 should be produced by validator[0]
	// block 11 should be produced by validator[1]
	// block 12 should be produced by validator[2]

	// block #10
	block, err := ethClient.BlockByNumber(context.Background(), big.NewInt(10))
	if err != nil {
		t.Fatalf("failed to get block #10: %v", err)
	}
	dumpBlockHeader(block)
	expectedCoinBase := common.HexToAddress(validators[0].miner).Hex()
	actualCoinBase := block.Header().Coinbase.Hex()
	if actualCoinBase != expectedCoinBase {
		t.Fatalf("block #10 should be produced by %v rather than %v", expectedCoinBase, actualCoinBase)
	}

	// block #11
	block, err = ethClient.BlockByNumber(context.Background(), big.NewInt(11))
	if err != nil {
		t.Fatalf("failed to get block #11: %v", err)
	}
	dumpBlockHeader(block)
	expectedCoinBase = common.HexToAddress(validators[1].miner).Hex()
	actualCoinBase = block.Header().Coinbase.Hex()
	if actualCoinBase != expectedCoinBase {
		t.Fatalf("block #11 should be produced by %v rather than %v", expectedCoinBase, actualCoinBase)
	}

	// block #12
	block, err = ethClient.BlockByNumber(context.Background(), big.NewInt(12))
	if err != nil {
		t.Fatalf("failed to get block #12: %v", err)
	}
	dumpBlockHeader(block)
	expectedCoinBase = common.HexToAddress(validators[2].miner).Hex()
	actualCoinBase = block.Header().Coinbase.Hex()
	if actualCoinBase != expectedCoinBase {
		t.Fatalf("block #12 should be produced by %v rather than %v", expectedCoinBase, actualCoinBase)
	}

}

type Validator struct {
	client     *hivesim.Client // client
	enode      string          // enode
	miner      string          // address of validator
	privateKey string          // privatekey of validator
}

func dumpBlockHeader(block *types.Block) {

	bytes, err := block.Header().MarshalJSON()

	if err == nil {
		fmt.Printf("block #%v:\n%v", block.Header().Number.Uint64(), string(bytes))
	}
}
