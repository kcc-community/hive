package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/hive/hivesim"
	Store "github.com/ethereum/hive/simulators/kcc/deploy-contract/contract"
	"math/big"
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
	rpc := ethclient.NewClient(c.RPC())

	pk, err := crypto.HexToECDSA("9c647b8b7c4e7c3490668fb6c11473619db80c93704c70893d3813af4090c39c")
	if err != nil {
		t.Fatal("HexToECDSA err", err)
	}

	transactOpts, err := bind.NewKeyedTransactorWithChainID(pk, big.NewInt(321))
	if err != nil {
		t.Fatal("NewKeyedTransactorWithChainID err", err)
	}

	address, transaction, store, err := Store.DeployStore(transactOpts, rpc)
	if err != nil {
		t.Fatal("DeployStore err:", err)
	}

	txJson, _ := transaction.MarshalJSON()
	t.Log("transaction", string(txJson), address.String())

	tx, err := store.SetItem(transactOpts, common.HexToHash("0x0"), common.HexToHash("0x1"))
	if err != nil {
		t.Fatal("setItem err:", err)
	}
	txJson, _ = tx.MarshalJSON()
	t.Log("SetItem transaction", string(txJson))

	opt := &bind.CallOpts{
		From: address,
	}
	data, err := store.Items(opt, common.HexToHash("0x0"))
	if err != nil {
		t.Fatal("setItem err:", err)
	}

	t.Log(data == common.HexToHash("0x1"), data, common.HexToHash("0x1"))
}
