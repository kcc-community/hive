package main

import (
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/hive/hivesim"
	"math/big"
	"strconv"
	"time"
)

func main() {
	suite := hivesim.Suite{
		Name:        "kcc-new-txpool",
		Description: "A general new-txpool test case for KCC",
	}
	suite.Add(hivesim.ClientTestSpec{
		Role:        "eth1",
		Name:        "kcc-new-txpool",
		Description: "new-txpool sort in different account after ishikari",
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

	address1 := common.HexToAddress("0x658bdf435d810c91414ec09147daa6db62406379")
	address2 := common.HexToAddress("0x7B42a4573ECb800bD404042F518a59CfBE3b582B")
	accounts := []*account{
		{
			privateKey: "9c647b8b7c4e7c3490668fb6c11473619db80c93704c70893d3813af4090c39c",
			address:    address1.String(),
			txs: []*types.LegacyTx{
				{
					Nonce:    uint64(0),
					GasPrice: big.NewInt(2 * 1e9),
					Gas:      8 * 1e5,
					To:       &address1,
					Value:    big.NewInt(0),
					Data:     nil,
				},
				{
					Nonce:    uint64(1),
					GasPrice: big.NewInt(2 * 1e9),
					Gas:      8 * 1e5,
					To:       &address1,
					Value:    big.NewInt(0),
					Data:     nil,
				},
				{
					Nonce:    uint64(2),
					GasPrice: big.NewInt(1.5 * 1e9),
					Gas:      8 * 1e5,
					To:       &address1,
					Value:    big.NewInt(0),
					Data:     nil,
				},
				{
					Nonce:    uint64(3),
					GasPrice: big.NewInt(1.2 * 1e9),
					Gas:      8 * 1e5,
					To:       &address1,
					Value:    big.NewInt(0),
					Data:     nil,
				},
			},
		},
		{
			privateKey: "7d38039b2367def23d26b092ee66a5c2c2c9be4972ba05ca326f71f8c783f44a",
			address:    address2.String(),
			txs: []*types.LegacyTx{
				{
					Nonce:    uint64(0),
					GasPrice: big.NewInt(3 * 1e9),
					Gas:      8 * 1e5,
					To:       &address2,
					Value:    big.NewInt(0),
					Data:     nil,
				},
				{
					Nonce:    uint64(1),
					GasPrice: big.NewInt(1.5 * 1e9),
					Gas:      8 * 1e5,
					To:       &address2,
					Value:    big.NewInt(0),
					Data:     nil,
				},
				{
					Nonce:    uint64(2),
					GasPrice: big.NewInt(3 * 1e9),
					Gas:      8 * 1e5,
					To:       &address2,
					Value:    big.NewInt(0),
					Data:     nil,
				},
			},
		},
	}

	wants := []common.Hash{
		common.HexToHash("0x1b0aa90aed807b772193f7200cfba272972a730fa70c8d19fc7d7438096fdbdd"),
		common.HexToHash("0xb2eaf381f63cdc2bddb62816558fafa9e40d0ba1328a4be5fc889b54a3f8324c"),
		common.HexToHash("0xad5118ff07ed360437dc3eee7ff76a62fcb03d9af9e204149f8aee7863d9ab54"),
		common.HexToHash("0x2affb932d441e398725f3c56b84738cbb5f75d2a13c3f3cb128882285c30537d"),
		common.HexToHash("0x130c83b212d528fce4d9033908374ebe69b4a8594095e2fb707d50eb9b1df6fa"),
		common.HexToHash("0x63b5feb7418de43bf2e4c886b7be1c079fbfa36763c7fb3c4435f30192f2c4fb"),
		common.HexToHash("0x1d020a72c02195408a711f4dc47ca7effde8fd1dfd080de8525c0833dd311a09"),
	}

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

	client := ethclient.NewClient(c.RPC())
	var latestTx *types.Transaction

	bes := make([]rpc.BatchElem, 0, 7)
	for _, account := range accounts {
		privateKey, err := crypto.HexToECDSA(account.privateKey)
		if err != nil {
			t.Fatal("account1 HexToECDSA err:", err)
		}

		for _, legacyTx := range account.txs {
			t.Log("transaction:", legacyTx)
			tx, err := types.SignNewTx(privateKey, signer, legacyTx)
			if err != nil {
				t.Fatal("account0 signNewTx err:", err)
			}

			data, err := tx.MarshalBinary()
			if err != nil {
				t.Fatal("MarshalBinary signNewTx err:", err)
			}

			latestTx = tx
			bes = append(bes, rpc.BatchElem{
				Method: "eth_sendRawTransaction",
				Args:   []interface{}{hexutil.Encode(data)},
				Result: nil,
				Error:  nil,
			})
		}

	}

	err := c.RPC().BatchCall(bes)
	if err != nil {
		t.Fatal("BatchCall err:", err)
	}

	var blockHash *common.Hash
	for i := 0; i < 10; i++ {
		var transaction *rpcTransaction
		err := c.RPC().Call(&transaction, "eth_getTransactionByHash", latestTx.Hash().String())
		if err != nil {
			t.Fatal("account1 call eth_getTransactionByHash err:", err)
		}

		if transaction.BlockHash != nil {
			blockHash = transaction.BlockHash
			break
		}

		t.Log("block not confirm request after 1s:", latestTx.Hash().String())
		time.Sleep(time.Second)
	}

	if blockHash == nil {
		t.Fatal("block not confirm after 10s", latestTx.Hash().String())
	}

	t.Log("block hash:", blockHash)

	blockInfo, err := client.BlockByHash(context.Background(), *blockHash)
	if err != nil {
		t.Fatal("account1 call eth_getBlockByHash err:", err)
	}

	transactions := blockInfo.Transactions()
	for index, transaction := range transactions {
		tx, _ := transaction.MarshalJSON()
		t.Log("transaction:", string(tx))

		if index+1 > len(wants) {
			t.Fatal("transaction fatal")
		}
		if wants[index] != transaction.Hash() {
			t.Fatal("transaction sort fatal", transaction.Hash().String(), "want:", wants[index])
		}
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
