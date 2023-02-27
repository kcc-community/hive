package main

import (
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/hive/hivesim"
	Store "github.com/ethereum/hive/simulators/kcc/deploy-contract/contract"
	"math/big"
	"strconv"
	"time"
)

func main() {
	suite := hivesim.Suite{
		Name:        "kcc-deploy-contract",
		Description: "A general deploy-contract test case for KCC",
	}
	suite.Add(hivesim.ClientTestSpec{
		Role:        "eth1",
		Name:        "mine one block",
		Description: "Waits for a single block to be mined by kcc",
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
		Run: deployContract,
	})
	hivesim.MustRunSuite(hivesim.New(), suite)
}

func deployContract(t *hivesim.T, c *hivesim.Client) {

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

		num, err := strconv.ParseUint(block.Number[2:], 16, 64)
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

	rpc := ethclient.NewClient(c.RPC())

	pk, err := crypto.HexToECDSA("9c647b8b7c4e7c3490668fb6c11473619db80c93704c70893d3813af4090c39c")
	if err != nil {
		t.Fatal("HexToECDSA err", err)
	}

	transactOpts, err := bind.NewKeyedTransactorWithChainID(pk, big.NewInt(321))
	if err != nil {
		t.Fatal("NewKeyedTransactorWithChainID err", err)
	}

	transactOpts.Nonce = big.NewInt(0)
	address, transaction, store, err := Store.DeployStore(transactOpts, rpc)
	if err != nil {
		t.Fatal("DeployStore err:", err)
	}

	txJson, _ := transaction.MarshalJSON()
	t.Log("transaction", string(txJson), address.String())

	time.Sleep(time.Second * 2)
	transactOpts.Nonce = big.NewInt(1)
	tx, err := store.SetItem(transactOpts, common.HexToHash("0x0"), common.HexToHash("0x1"))
	if err != nil {
		t.Fatal("setItem err:", err)
	}
	txJson, _ = tx.MarshalJSON()
	t.Log("SetItem transaction", string(txJson))

	//var blockHash *common.Hash
	for i := 0; i < 5; i++ {

		var transaction *rpcTransaction
		err := c.RPC().Call(&transaction, "eth_getTransactionByHash", tx.Hash().String())
		if err != nil {
			t.Fatal("account1 call eth_getTransactionByHash err:", err)
		}

		if transaction.BlockHash != nil {
			//blockHash = transaction.BlockHash
			break
		}
	}

	receipts, err := rpc.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		t.Fatal("account1 call eth_getBlockByHash err:", err)
	}
	t.Log("receipts,gas_used: ", receipts.GasUsed)

	data, err := store.Items(nil, common.HexToHash("0x0"))
	if err != nil {
		t.Fatal("setItem err:", err)
	}

	if data != common.HexToHash("0x1") {
		t.Fatal("contract data set err:", err)
	}

	t.Log(data, common.HexToHash("0x1"))
}

type rpcTransaction struct {
	tx *types.Transaction
	txExtraInfo
}

type txExtraInfo struct {
	BlockNumber *string         `json:"blockNumber,omitempty"`
	BlockHash   *common.Hash    `json:"blockHash,omitempty"`
	From        *common.Address `json:"from,omitempty"`
}

func (tx *rpcTransaction) UnmarshalJSON(msg []byte) error {
	if err := json.Unmarshal(msg, &tx.tx); err != nil {
		return err
	}
	return json.Unmarshal(msg, &tx.txExtraInfo)
}
