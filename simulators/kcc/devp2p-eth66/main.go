package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/hive/hivesim"
	"math/big"
	"strings"
	"time"
)

func main() {
	suite := hivesim.Suite{
		Name:        "devp2p-eth66",
		Description: "#10 the devp2p eth65 is not supported anymore ",
	}
	suite.Add(hivesim.TestSpec{
		Name:        "kcc #10",
		Description: "the devp2p eth65 is not supported anymore",
		Run:         devP2pEth66,
	})
	hivesim.MustRunSuite(hivesim.New(), suite)
}

type validator struct {
	client     *hivesim.Client // client
	enode      string          // enode
	miner      string          // address of validator
	privateKey string          // privatekey of validator
	params     hivesim.Params
}

var validations = []validator{
	{
		miner:      "0x658bdf435d810c91414ec09147daa6db62406379",
		privateKey: "9c647b8b7c4e7c3490668fb6c11473619db80c93704c70893d3813af4090c39c",
		params: map[string]string{
			"HIVE_CLIQUE_PRIVATEKEY": "9c647b8b7c4e7c3490668fb6c11473619db80c93704c70893d3813af4090c39c",
			"HIVE_MINER":             "658bdf435d810c91414ec09147daa6db62406379",
			"HIVE_CHAIN_ID":          "321",
			// block interval: 1s
			"HIVE_KCC_POSA_BLOCK_INTERVAL": "1",
			// epoch : 5
			"HIVE_KCC_POSA_EPOCH": "5",
			// initial valiators
			"HIVE_KCC_POSA_ISHIKARI_INIT_VALIDATORS": "0x658bdf435d810c91414ec09147daa6db62406379,0x7B42a4573ECb800bD404042F518a59CfBE3b582B,0x7f0dd90F50E4E18D58aD7780a5a642acd5f44068",
			// admin
			"HIVE_KCC_POSA_ADMIN": "0x658bdf435d810c91414ec09147daa6db62406379",
			// KCC Ishikari  fork number
			"HIVE_FORK_KCC_ISHIKARI": "9",
			// KCC Ishikari Patch001 fork number
			"HIVE_FORK_KCC_ISHIKARI_PATCH001": "10",
			// KCC Ishikari Patch002 fork number
			"HIVE_FORK_KCC_ISHIKARI_PATCH002": "11",
		},
	},
	{
		miner:      "0x7B42a4573ECb800bD404042F518a59CfBE3b582B",
		privateKey: "7d38039b2367def23d26b092ee66a5c2c2c9be4972ba05ca326f71f8c783f44a",
		params: map[string]string{
			"HIVE_CLIQUE_PRIVATEKEY": "7d38039b2367def23d26b092ee66a5c2c2c9be4972ba05ca326f71f8c783f44a",
			"HIVE_MINER":             "0x7B42a4573ECb800bD404042F518a59CfBE3b582B",
			"HIVE_CHAIN_ID":          "321",
			// block interval: 1s
			"HIVE_KCC_POSA_BLOCK_INTERVAL": "1",
			// epoch : 5
			"HIVE_KCC_POSA_EPOCH": "5",
			// initial valiators
			"HIVE_KCC_POSA_ISHIKARI_INIT_VALIDATORS": "0x658bdf435d810c91414ec09147daa6db62406379,0x7B42a4573ECb800bD404042F518a59CfBE3b582B,0x7f0dd90F50E4E18D58aD7780a5a642acd5f44068",
			// admin
			"HIVE_KCC_POSA_ADMIN": "0x658bdf435d810c91414ec09147daa6db62406379",
			// KCC Ishikari  fork number
			"HIVE_FORK_KCC_ISHIKARI": "9",
			// KCC Ishikari Patch001 fork number
			"HIVE_FORK_KCC_ISHIKARI_PATCH001": "10",
			// KCC Ishikari Patch002 fork number
			"HIVE_FORK_KCC_ISHIKARI_PATCH002": "11",
		},
	},
	{
		miner:      "0x7f0dd90f50e4e18d58ad7780a5a642acd5f44068",
		privateKey: "bfc63d285be9156d434b4b46e40c68cec7e84cfa7b9d599d5939fa6135eb878a",
		params: map[string]string{
			"HIVE_CLIQUE_PRIVATEKEY": "bfc63d285be9156d434b4b46e40c68cec7e84cfa7b9d599d5939fa6135eb878a",
			"HIVE_MINER":             "0x7f0dd90F50E4E18D58aD7780a5a642acd5f44068",
			"HIVE_CHAIN_ID":          "321",
			// block interval: 1s
			"HIVE_KCC_POSA_BLOCK_INTERVAL": "1",
			// epoch : 5
			"HIVE_KCC_POSA_EPOCH": "5",
			// initial valiators
			"HIVE_KCC_POSA_ISHIKARI_INIT_VALIDATORS": "0x658bdf435d810c91414ec09147daa6db62406379,0x7B42a4573ECb800bD404042F518a59CfBE3b582B,0x7f0dd90F50E4E18D58aD7780a5a642acd5f44068",
			// admin
			"HIVE_KCC_POSA_ADMIN": "0x658bdf435d810c91414ec09147daa6db62406379",
			// KCC Ishikari  fork number
			"HIVE_FORK_KCC_ISHIKARI": "9",
			// KCC Ishikari Patch001 fork number
			"HIVE_FORK_KCC_ISHIKARI_PATCH001": "10",
			// KCC Ishikari Patch002 fork number
			"HIVE_FORK_KCC_ISHIKARI_PATCH002": "11",
		},
	},
}

func devP2pEth66(t *hivesim.T) {
	clientTypes, err := t.Sim.ClientTypes()
	if err != nil {
		t.Fatalf("failed to get client types:%v", err)
	}

	var hasV1_2_0, hasVLatest bool
	for _, t := range clientTypes {
		if strings.EqualFold(t.Name, "kcc_v1.2.0") {
			hasV1_2_0 = true
		}
		if strings.EqualFold(t.Name, "kcc_latest") || strings.EqualFold(t.Name, "kcc") {
			hasVLatest = true
		}
	}

	if !hasV1_2_0 || !hasVLatest {
		t.Fatalf("missing client types, this case must be run with options '--client kcc_v1.2.0,kcc'")
	}

	// start old version node
	oldClient := t.StartClient("kcc_v1.2.0", validations[0].params,
		hivesim.WithStaticFiles(
			map[string]string{"/genesis.json": "genesis.json"}),
	)

	// start new version node
	newClient1 := t.StartClient("kcc", validations[1].params,
		hivesim.WithStaticFiles(
			map[string]string{"/genesis.json": "genesis.json"}),
	)

	newClient2 := t.StartClient("kcc", validations[2].params,
		hivesim.WithStaticFiles(
			map[string]string{"/genesis.json": "genesis.json"}),
	)

	connectClients(t, []*hivesim.Client{oldClient, newClient1, newClient2})

	// wait ishikari block
	isIshikari := false
	for i := 0; i < 30; i++ {
		client := ethclient.NewClient(oldClient.RPC())
		block := latestBlock(t, client)
		if block.Number().Uint64() > 10 {
			isIshikari = true
			break
		}
		time.Sleep(time.Second)
	}
	if !isIshikari {
		t.Fatal("get block fatal after 30s")
	}

	oldEthClient := ethclient.NewClient(oldClient.RPC())
	newEthClient1 := ethclient.NewClient(newClient1.RPC())
	newEthClient2 := ethclient.NewClient(newClient2.RPC())

	// ##################### block sync test #####################
	isSyncBlock := 0
	for i := 0; i < 10; i++ {

		block := latestBlock(t, oldEthClient)
		if strings.ToLower(block.Header().Coinbase.String()) == strings.ToLower(validations[1].miner) ||
			strings.ToLower(block.Header().Coinbase.String()) == strings.ToLower(validations[2].miner) {
			isSyncBlock = isSyncBlock | 1
		}
		t.Logf("[client0] get block, %d, validator: %s", block.Number().Uint64(), block.Header().Coinbase.String())

		block = latestBlock(t, newEthClient1)
		if strings.ToLower(block.Header().Coinbase.String()) == strings.ToLower(validations[0].miner) {
			//isSyncToNewNode = true
			isSyncBlock = isSyncBlock | 2
		}
		t.Logf("[client1] get block, %d, validator: %s", block.Number().Uint64(), block.Header().Coinbase.String())

		if isSyncBlock == 3 {
			break
		}

		time.Sleep(time.Second)
	}

	if isSyncBlock != 3 {
		t.Fatalf("sync node error, %d", isSyncBlock)
	}

	t.Log("sync block success, between old and new node")

	// ################# transaction sync test #####################
	validatorAddress := []string{"0x658bdf435d810c91414ec09147daa6db62406379", "0x7B42a4573ECb800bD404042F518a59CfBE3b582B", "0x7f0dd90F50E4E18D58aD7780a5a642acd5f44068"}

	newHead := make(chan *types.Header, 16)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	clients := []*ethclient.Client{oldEthClient, newEthClient1, newEthClient2}
	url := fmt.Sprintf("ws://%v:8546", oldClient.IP)
	wsClient, err := rpc.Dial(url)
	if err != nil {
		t.Fatalf("new ws client error, url:%s, err:%v", url, err)
	}

	opt, err := ethclient.NewClient(wsClient).SubscribeNewHead(ctx, newHead)
	if err != nil {
		t.Fatalf("sub new head err, err:%v", url, err)
	}

	isSync := 0
LOOP:
	for {
		select {
		case <-ctx.Done():
			t.Fatalf("watch new head time out")

		case header := <-newHead:
			b, _ := header.MarshalJSON()
			t.Log("new head", string(b))

			for i, address := range validatorAddress {
				if strings.ToLower(address) != strings.ToLower(header.Coinbase.String()) {
					continue
				}

				switch i {
				case 0: // old node send transaction sync to new node
					sendTransaction(t, validations[0].privateKey, oldEthClient, clients)
					isSync = isSync | 1
				//case 1: // new node send transaction sync to new node and old node
				//	sendTransaction(t, validations[1].privateKey, newEthClient1, clients)
				//	isSync = isSync | 2
				case 2: // new node send transaction sync to new node and old node
					sendTransaction(t, validations[2].privateKey, newEthClient2, clients)
					isSync = isSync | 2
				}
			}

			if isSync == 3 {
				break LOOP
			}
		}
	}
	opt.Unsubscribe()

	t.Log("sync transaction success, between old and new node")
}

func sendTransaction(t *hivesim.T, privateKey string, client *ethclient.Client, clients []*ethclient.Client) {
	pk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		t.Fatal("HexToECDSA err:", err)
	}

	signer := types.LatestSignerForChainID(big.NewInt(321))
	address := common.HexToAddress(validations[0].miner)
	transaction := &types.LegacyTx{
		GasPrice: big.NewInt(2 * 1e9),
		Gas:      8 * 1e5,
		To:       &address,
		Value:    big.NewInt(0),
	}

	tx, err := types.SignNewTx(pk, signer, transaction)
	if err != nil {
		t.Fatal("sign new tx error", err)
	}

	err = client.SendTransaction(context.TODO(), tx)
	if err != nil {
		t.Fatal("send transaction error", err)
	}

	//after 50ms to sync
	time.Sleep(time.Millisecond * 50)

	for _, client := range clients {
		transaction, isPending, err := client.TransactionByHash(context.Background(), tx.Hash())
		if err != nil {
			t.Fatal("get transaction error", err)
		}

		b, _ := transaction.MarshalJSON()
		t.Log("transaction info: ", string(b))

		if transaction != nil && !isPending {
			t.Fatal("get transaction fatal", transaction.Hash(), isPending)
		}
	}
}

func latestBlock(t *hivesim.T, c *ethclient.Client) *types.Block {
	blockNumber, err := c.BlockNumber(context.TODO())
	if err != nil {
		t.Fatal("get block number error", err)
	}

	block, err := c.BlockByNumber(context.TODO(), big.NewInt(int64(blockNumber)))
	if err != nil {
		t.Fatalf("get block error: %d, err:%v", blockNumber, err)
	}

	return block
}

func connectClients(t *hivesim.T, clients []*hivesim.Client) {
	enodes := make([]string, 3)

WAITING_READY_LOOP:
	for {
		for i, c := range clients {
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
	err := clients[0].RPC().Call(nil, "admin_addPeer", enodes[1])
	if err != nil {
		t.Fatalf("[0->1]failed to addPeer : %v", err)
	}

	// add c1 to c2
	err = clients[0].RPC().Call(nil, "admin_addPeer", enodes[2])
	if err != nil {
		t.Fatalf("[0->2]failed to addPeer : %v", err)
	}

	// add c1 to c2
	err = clients[1].RPC().Call(nil, "admin_addPeer", enodes[2])
	if err != nil {
		t.Fatalf("[1->2]failed to addPeer : %v", err)
	}
}
