package main

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/hive/hivesim"
)

func main() {
	suite := hivesim.Suite{
		Name:        "kcc #9: crash on synchronization",
		Description: "https://github.com/kcc-community/kcc/issues/9",
	}
	suite.Add(hivesim.TestSpec{
		Name:        "kcc #9",
		Description: "Synchronization error due to CVE-2021-39137",
		Run:         distributeBlockRewardTest,
	})
	hivesim.MustRunSuite(hivesim.New(), suite)
}

func distributeBlockRewardTest(t *hivesim.T) {

	// check client types
	// We only test against the following two version of clients
	//   - v1.0.3
	//   - latest
	clientTypes, err := t.Sim.ClientTypes()
	if err != nil {
		t.Fatalf("failed to get client types:%v", err)
	}

	var hasV1_0_3, hasVLatest bool
	for _, t := range clientTypes {
		if strings.EqualFold(t.Name, "kcc_v1.0.3") {
			hasV1_0_3 = true
		}
		if strings.EqualFold(t.Name, "kcc_latest") || strings.EqualFold(t.Name, "kcc") {
			hasVLatest = true
		}
	}

	if !hasV1_0_3 || !hasVLatest {
		t.Fatalf("missing client types, this case must be run with options '--client kcc_v1.0.3,kcc'")
	}

	// the old client(v1.0.3) is the miner
	oldClient := t.StartClient("kcc_v1.0.3", hivesim.Params{
		"HIVE_CLIQUE_PRIVATEKEY": "9c647b8b7c4e7c3490668fb6c11473619db80c93704c70893d3813af4090c39c",
		"HIVE_MINER":             "658bdf435d810c91414ec09147daa6db62406379",
		"HIVE_CHAIN_ID":          "321",
		// block interval: 1s
		"HIVE_KCC_POSA_BLOCK_INTERVAL": "3",
		// epoch : 10
		"HIVE_KCC_POSA_EPOCH": "10",
		// initial valiators
		"HIVE_KCC_POSA_ISHIKARI_INIT_VALIDATORS": "0x658bdf435d810c91414ec09147daa6db62406379",
		// admin
		"HIVE_KCC_POSA_ADMIN": "0x658bdf435d810c91414ec09147daa6db62406379",
		// sync mode
		"HIVE_NODETYPE": "archive",
	}, hivesim.WithStaticFiles(
		map[string]string{
			"/genesis.json": "genesis.json",
		}))

	// replay the malicious transcation
	// https://scan.kcc.io/tx/0xe42c11a8a31f3c0a990a8264b23bf4e936b9d97f3242cfbc21e63b2b0abd09f0
	maliciousBlockNumber := replayMaliciousTX(t, ethclient.NewClient(oldClient.RPC()))

	t.Logf("the block with the malicious TX is #%v", maliciousBlockNumber)

	// the new client(latest version) is not a miner
	// It will synchronize with the old client
	newClient := t.StartClient("kcc", hivesim.Params{
		"HIVE_CHAIN_ID": "321",
		// block interval: 1s
		"HIVE_KCC_POSA_BLOCK_INTERVAL": "3",
		// epoch : 10
		"HIVE_KCC_POSA_EPOCH": "10",
		// initial valiators
		"HIVE_KCC_POSA_ISHIKARI_INIT_VALIDATORS": "0x658bdf435d810c91414ec09147daa6db62406379",
		// admin
		"HIVE_KCC_POSA_ADMIN": "0x658bdf435d810c91414ec09147daa6db62406379",
		// sync mode
		"HIVE_NODETYPE": "archive",
		//
		"HIVE_CVE_2021_39137_BLOCK": fmt.Sprintf("%v", maliciousBlockNumber),
	}, hivesim.WithStaticFiles(
		map[string]string{
			"/genesis.json": "genesis.json",
		}))

	// connect oldClient & newClient
	connectClients(t, oldClient, newClient)

	timeout := time.NewTimer(time.Minute)
	defer timeout.Stop()

	for {
		b, err := ethclient.NewClient(newClient.RPC()).BlockNumber(context.TODO())
		if err != nil {
			t.Fatalf("failed to get blocknumber of new client: %v", err)
		}

		t.Logf("Current blocknumber of new client is %v, the malicious block is #%v", b, maliciousBlockNumber)

		if b > maliciousBlockNumber {
			t.Logf("the malicious block is successfully processed by the new client")
			break
		}

		select {
		case <-timeout.C:
			t.Fatalf("timeout when waiting for new client to get synced")
		default:
		}

		time.Sleep(time.Second * 2)
	}

}

func connectClients(t *hivesim.T, c1 *hivesim.Client, c2 *hivesim.Client) {

	enodes := make([]string, 2)

WAITING_READY_LOOP:
	for {
		for i, c := range []*hivesim.Client{c1, c2} {
			enode, err := c.EnodeURL()
			if err != nil {
				t.Fatalf("failed to get enode url:%v", err)
			}

			if enode == "" {
				// any of the nodes is not ready
				time.Sleep(time.Second * 2)
				continue WAITING_READY_LOOP
			}
			enodes[i] = enode
		}

		break
	}

	// add c1 to c2
	err := c2.RPC().Call(nil, "admin_addPeer", enodes[0])

	if err != nil {
		t.Fatalf("failed to addPeer : %v", err)
	}

}

//
// replay the malicious transcation
// https://scan.kcc.io/tx/0xe42c11a8a31f3c0a990a8264b23bf4e936b9d97f3242cfbc21e63b2b0abd09f0
func replayMaliciousTX(t *hivesim.T, c *ethclient.Client) uint64 {

	t.Logf("Relaying the malicious TX")

	privateKey, err := crypto.HexToECDSA("9c647b8b7c4e7c3490668fb6c11473619db80c93704c70893d3813af4090c39c")
	if err != nil {
		t.Fatalf("failed to decode privatekey hex as private key")
	}

	address := crypto.PubkeyToAddress(privateKey.PublicKey)

	gasPrice, err := c.SuggestGasPrice(context.TODO())
	if err != nil {
		t.Fatalf("failed to get gas price: %v", err)
	}

	nonce, err := c.NonceAt(context.TODO(), address, nil)
	if err != nil {
		t.Fatalf("failed to get nonce: %v", err)
	}

	// the malicious contract init code
	maliciousContractInitCode := common.FromHex("0x3034526020600760203460045afa602034343e604034f3")

	gasLimit, err := c.EstimateGas(context.TODO(), ethereum.CallMsg{
		From: address,
		To:   nil,
		Data: maliciousContractInitCode,
	})
	if err != nil {
		t.Fatalf("failed to get gas limit: %v", err)
	}

	tx, err := types.SignTx(types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       nil,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		Data:     maliciousContractInitCode,
	}), types.LatestSignerForChainID(big.NewInt(321)), privateKey)
	if err != nil {
		t.Fatalf("failed to sign TX: %v", err)
	}

	err = c.SendTransaction(context.TODO(), tx)
	if err != nil {
		t.Fatalf("failed to send TX : %v", err)
	}

	// wait for TX to be included in some block
	for {

		tx, pending, err := c.TransactionByHash(context.TODO(), tx.Hash())
		if err != nil {
			t.Fatalf("failed to get TX: %v", err)
		}

		if pending {
			t.Logf("waiting for the malicious transaction to be included in some block ...")
			time.Sleep(time.Second * 2)
			continue
		}

		// get Transaction receipts
		receipt, err := c.TransactionReceipt(context.TODO(), tx.Hash())
		if err != nil {
			t.Fatalf("failed to get transaction receipt: %v", err)
		}

		return receipt.BlockNumber.Uint64()

	}

}
