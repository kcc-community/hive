package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/hive/hivesim"
)

var (
	MinerKey  = "0x9c647b8b7c4e7c3490668fb6c11473619db80c93704c70893d3813af4090c39c"
	MinerAddr = common.HexToAddress("0x658bdf435d810c91414ec09147daa6db62406379")
)

func main() {
	suite := hivesim.Suite{
		Name:        "pre-1559",
		Description: "Tests for the pre-1559 behaviors of the KCC client",
	}
	suite.Add(hivesim.ClientTestSpec{
		Role:        "eth1",
		Name:        "Reject Dynamic Fee Transactions",
		Description: "The pre-1559 KCC client should reject dynamic fee transactions",
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
			// KCC Ishikari Patch002 fork number
			"HIVE_FORK_KCC_ISHIKARI_PATCH002": "11",
		},
		Run: rejectDynamicTX,
	}).Add(hivesim.ClientTestSpec{
		Role:        "eth1",
		Name:        "Header should not inlcude a base fee",
		Description: "The pre-1559 KCC block header should not have a base fee field",
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
			// KCC Ishikari Patch002 fork number
			"HIVE_FORK_KCC_ISHIKARI_PATCH002": "11",
		},
		Run: headerWithNoBasefee,
	}).Add(hivesim.ClientTestSpec{
		Role:        "eth1",
		Name:        "Node should not expose EIP-1559 related API",
		Description: "The pre-1559 KCC client should not expose EIP-1559 related API",
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
			// KCC Ishikari Patch002 fork number
			"HIVE_FORK_KCC_ISHIKARI_PATCH002": "11",
		},
		Run: noEIP1559API,
	})

	hivesim.MustRunSuite(hivesim.New(), suite)
}

func rejectDynamicTX(t *hivesim.T, c *hivesim.Client) {

	// wait for the hard fork block number
	t.Log("Waiting for the hardfork block number")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	for {
		var block struct {
			Number     string `json:"number"`
			Hash       string `json:"hash"`
			ParentHash string `json:"parentHash"`
		}

		if err := c.RPC().CallContext(ctx, &block, "eth_getBlockByNumber", "latest", false); err != nil {
			t.Fatalf("Failed to get latest block: %v", err)
		}

		num, err := HexWithPrefixToUint64(block.Number)
		if err != nil {
			t.Fatalf("Failed to get latest block: %v", err)
		}

		if num >= 11 {
			break
		}

		select {
		case <-ctx.Done():
			t.Fatalf("Timeout when waiting for the hardfork block")
		default:
		}
	}

	// send a dynamic fee transaction
	t.Log("Sending a dynamic fee transaction")

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   big.NewInt(321),
		Nonce:     0,
		GasTipCap: big.NewInt(params.GWei * 50),
		GasFeeCap: big.NewInt(params.GWei * 100),
		Gas:       21000,
		To:        &MinerAddr,
		Value:     big.NewInt(params.Ether),
	})

	pkey, _ := crypto.HexToECDSA(MinerKey[2:])

	signedTx, err := types.SignTx(tx, types.NewLondonSigner(big.NewInt(321)), pkey)

	if err != nil {
		t.Fatalf("Failed to sign the transaction: %v", err)
	}

	rawTX, err := signedTx.MarshalBinary()
	if err != nil {
		t.Fatalf("Failed to marshal the transaction: %v", err)
	}

	var r json.RawMessage

	if err := c.RPC().CallContext(ctx, &r, "eth_sendRawTransaction", "0x"+common.Bytes2Hex(rawTX)); err != nil {
		if err.Error() != types.ErrTxTypeNotSupported.Error() {
			t.Fatalf("The transaction should be rejected by an error of 'transaction type not supported', but got: %v", err)
		}
	}

}

func HexWithPrefixToUint64(hex string) (uint64, error) {
	return strconv.ParseUint(hex[2:], 16, 64)
}

func headerWithNoBasefee(t *hivesim.T, c *hivesim.Client) {

	// wait for the hard fork block number
	t.Log("Waiting for the hardfork block number")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	for {
		var block struct {
			Number     string `json:"number"`
			Hash       string `json:"hash"`
			ParentHash string `json:"parentHash"`
		}

		if err := c.RPC().CallContext(ctx, &block, "eth_getBlockByNumber", "latest", false); err != nil {
			t.Fatalf("Failed to get latest block: %v", err)
		}

		num, err := HexWithPrefixToUint64(block.Number)
		if err != nil {
			t.Fatalf("Failed to get latest block: %v", err)
		}

		if num >= 11 {
			break
		}

		select {
		case <-ctx.Done():
			t.Fatalf("Timeout when waiting for the hardfork block")
		default:
		}
	}

	// send an EIP-2930 tx

	tx := types.NewTx(&types.AccessListTx{
		ChainID: big.NewInt(321),
		Nonce:   0,
		Gas:     21000,
		To:      &MinerAddr,
		Value:   big.NewInt(params.Ether),
	})

	pkey, _ := crypto.HexToECDSA(MinerKey[2:])

	signedTx, err := types.SignTx(tx, types.NewEIP2930Signer(big.NewInt(321)), pkey)

	if err != nil {
		t.Fatalf("Failed to sign the transaction: %v", err)
	}

	rawTX, err := signedTx.MarshalBinary()
	if err != nil {
		t.Fatalf("Failed to marshal the transaction: %v", err)
	}

	var r json.RawMessage

	if err := c.RPC().CallContext(ctx, &r, "eth_sendRawTransaction", "0x"+common.Bytes2Hex(rawTX)); err != nil {
		t.Fatalf("Failed to send the transaction: %v", err)
	}

	// wait for the transaction to be mined
	var txReceipt struct {
		BlockNumber string `json:"blockNumber"`
	}
	for {

		if err := c.RPC().CallContext(ctx, &txReceipt, "eth_getTransactionReceipt", signedTx.Hash().Hex()); err != nil {
			t.Fatalf("Failed to get the transaction receipt: %v", err)
		}

		if txReceipt.BlockNumber != "" {
			break
		}

		select {
		case <-ctx.Done():
			t.Fatalf("Timeout when waiting for the transaction to be mined")
		default:
		}
	}

	t.Logf("The transaction is mined in block %s", txReceipt.BlockNumber)

	// get the block header
	var header types.Header

	if err := c.RPC().Call(&header, "eth_getBlockByNumber", txReceipt.BlockNumber, false); err != nil {
		t.Fatalf("Failed to get the block header: %v", err)
	}

	if marshalledHeader, err := header.MarshalJSON(); err != nil {
		t.Fatalf("Failed to marshal the block header: %v", err)
	} else {
		t.Logf("The block header is: %s", marshalledHeader)
	}

	if header.BaseFee != nil {
		t.Fatalf("The block header should not have a basefee field")
	}

}

func noEIP1559API(t *hivesim.T, c *hivesim.Client) {

	disabledMethods := []string{
		"eth_maxPriorityFeePerGas",
		"eth_feeHistory",
	}

	expectedError := func(method string) string {
		return fmt.Sprintf("the method %s does not exist/is not available", method)
	}

	for _, method := range disabledMethods {
		var result interface{}
		if err := c.RPC().Call(&result, method); err != nil {
			if err.Error() != expectedError(method) {
				t.Fatalf("The method %s should be disabled, but got error: %v", method, err)
			}
		} else {
			t.Fatalf("The method %s should be disabled", method)
		}
	}
}
