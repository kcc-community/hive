package main

import (
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/hive/hivesim"
	"math/big"
	"time"
)

func main() {
	suite := hivesim.Suite{
		Name:        "kcc-smoke",
		Description: "A general somke test case for KCC",
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
		},
		Run: newTxPoolTest,
	})
	hivesim.MustRunSuite(hivesim.New(), suite)
}

type account struct {
	privateKey string
	address    string
	txs        []*types.LegacyTx
}

func newTxPoolTest(t *hivesim.T, c *hivesim.Client) {
	signer := types.LatestSignerForChainID(big.NewInt(321))

	//address1 := common.HexToAddress("0x658bdf435d810c91414ec09147daa6db62406379")
	address2 := common.HexToAddress("0x7B42a4573ECb800bD404042F518a59CfBE3b582B")
	accounts := []*account{
		//{
		//	privateKey: "9c647b8b7c4e7c3490668fb6c11473619db80c93704c70893d3813af4090c39c",
		//	address:    address1.String(),
		//	txs: []*types.LegacyTx{
		//		{
		//			Nonce:    uint64(0),
		//			GasPrice: big.NewInt(1e9),
		//			Gas:      8 * 1e5,
		//			To:       &address1,
		//			Value:    big.NewInt(0),
		//			Data:     nil,
		//		},
		//	},
		//},
		{
			privateKey: "7d38039b2367def23d26b092ee66a5c2c2c9be4972ba05ca326f71f8c783f44a",
			address:    address2.String(),
			txs: []*types.LegacyTx{
				{
					Nonce:    uint64(0),
					GasPrice: big.NewInt(1e9),
					Gas:      8 * 1e5,
					To:       &address2,
					Value:    big.NewInt(0),
					Data:     nil,
				},
			},
		},
	}

	privateKey1, err := crypto.HexToECDSA(accounts[0].privateKey)
	if err != nil {
		t.Fatal("account1 HexToECDSA err:", err)
	}

	//privateKey2, err := crypto.HexToECDSA(accounts[1].privateKey)
	//if err != nil {
	//	t.Fatal("account2 HexToECDSA err:", err)
	//}

	t.Log("transaction:", accounts[0].txs[0])
	tx, err := types.SignNewTx(privateKey1, signer, accounts[0].txs[0])
	if err != nil {
		t.Fatal("account0 signNewTx err:", err)
	}

	client := ethclient.NewClient(c.RPC())

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		t.Fatal("account1 call eth_sendRawTransaction err:", err)
	}

	t.Log("txID:", tx.Hash().String())

	var blockHash *common.Hash
	for i := 0; i < 10; i++ {
		var transaction *rpcTransaction
		err = c.RPC().Call(&transaction, "eth_getTransactionByHash", tx.Hash().String())
		if err != nil {
			t.Fatal("account1 call eth_getTransactionByHash err:", err)
		}

		if transaction.BlockHash != nil {
			blockHash = transaction.BlockHash
			break
		}

		t.Log("block not confirm request after 1s:", tx.Hash().String())
		time.Sleep(time.Second)
	}

	if blockHash == nil {
		t.Fatal("block not confirm after 10s", tx.Hash().String())
	}

	t.Log("block hash:", blockHash)
	blockInfo, err := client.BlockByHash(context.Background(), *blockHash)
	if err != nil {
		t.Fatal("account1 call eth_getBlockByHash err:", err)
	}

	transactions := blockInfo.Transactions()
	for _, transaction := range transactions {
		tx, _ := transaction.MarshalJSON()
		t.Log("transaction:", string(tx))
	}
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
