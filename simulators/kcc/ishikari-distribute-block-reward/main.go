package main

import (
	"context"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/hive/hivesim"
	"github.com/ethereum/hive/simulators/kcc/ishikari-distribute-block-reward/reservepool"
)

var (
	ValidatorsAddress  = common.HexToAddress("0x000000000000000000000000000000000000f333")
	ReservePoolAddress = common.HexToAddress("0x000000000000000000000000000000000000f999")
	ONE_KCS            = new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
)

func main() {
	suite := hivesim.Suite{
		Name:        "KCC Ishikari distribute block reward",
		Description: "",
	}
	suite.Add(hivesim.TestSpec{
		Name:        "distributeBlockReward",
		Description: "Compare the balances in validatorsContract before and after calling 'distributeBlockReward'.",
		Run:         distributeBlockRewardTest,
	})
	hivesim.MustRunSuite(hivesim.New(), suite)
}

func distributeBlockRewardTest(t *hivesim.T) {

	// In genesis.json, we have preallocated some KCS to the reservePool

	// create a kcc client
	// note: We are using an archive node, as we are going
	// to check the balances at different blocks.
	client := t.StartClient("kcc", hivesim.Params{
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
		// KCC Ishikari fork number
		"HIVE_FORK_KCC_ISHIKARI": "9",
		// sync mode
		"HIVE_NODETYPE": "archive",
	}, hivesim.WithStaticFiles(
		map[string]string{
			"/genesis.json": "genesis.json",
		}))

	ethClient := ethclient.NewClient(client.RPC())

	//
	// Wait for the hardfork
	// i.e block #9
	//
	t.Logf("waiting for the hardfork to come...")
	for {
		number, err := ethClient.BlockNumber(context.TODO())
		if err != nil {
			t.Fatalf("failed to get block number:  %v", err)
		}

		if number < 9 {
			t.Logf("current Block number is #%v", number)
			time.Sleep(time.Second * 3)
			continue
		}

		break
	}

	//
	// Set blockReward of reservePool
	//
	reservePool, err := reservepool.NewReservepool(ReservePoolAddress, ethClient)
	if err != nil {
		t.Fatalf("failed to create reservessPool: %v", err)
	}

	privateKey, err := crypto.HexToECDSA("9c647b8b7c4e7c3490668fb6c11473619db80c93704c70893d3813af4090c39c")
	if err != nil {
		t.Fatalf("failed to decode privatekey hex as private key")
	}
	transactor, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(321))
	if err != nil {
		t.Fatalf("failed to create transactor with private key: %v", err)
	}

	// Set the blockReward amount to 1 KCS
	tx, err := reservePool.SetBlockRewardAmount(transactor, ONE_KCS)
	if err != nil {
		log.Fatalf("failed to setBlockReward: %v", err)
	}

	// wait for the tx to be included in a block
	for {

		_, pending, err := ethClient.TransactionByHash(context.TODO(), tx.Hash())
		if err != nil {
			t.Fatalf("failed to get TX: %v", err)
		}

		if pending {
			t.Logf("waiting for the transaction of setting blockReward to be included in some block ...")
			time.Sleep(time.Second * 2)
			continue
		}

		break

	}

	currentBlockNumber, err := ethClient.BlockNumber(context.TODO())
	if err != nil {
		t.Fatalf("failed to get block number: %v", err)
	}

	// the first block of the next epoch
	nextEpochStartBlock := (currentBlockNumber + 10) - (currentBlockNumber+10)%10

	// wait for block "nextEpochStartBlock"
	for {
		number, err := ethClient.BlockNumber(context.TODO())
		if err != nil {
			t.Fatalf("failed to get block number:  %v", err)
		}

		if number <= nextEpochStartBlock {
			t.Logf("current Block number is #%v", number)
			time.Sleep(time.Second * 3)
			continue
		}

		break
	}

	// check balance change

	// balance in validators Contract at block (nextEpochStartBlock-1)
	balanceBefore, err := ethClient.BalanceAt(context.TODO(), ValidatorsAddress, big.NewInt(int64(nextEpochStartBlock-1)))
	if err != nil {
		t.Logf("failed to get balance:%v", err)
	}
	// balance in validators Contract at block (nextEpochStartBlock)
	balanceAfter, err := ethClient.BalanceAt(context.TODO(), ValidatorsAddress, big.NewInt(int64(nextEpochStartBlock)))
	if err != nil {
		t.Logf("failed to get balance:%v", err)
	}

	// balance diff
	diff := new(big.Int).Sub(balanceAfter, balanceBefore)

	t.Logf("The difference of balances in validators Contract at block #%v and block #%v is %v wei", nextEpochStartBlock, nextEpochStartBlock-1, diff.String())

	if diff.Cmp(ONE_KCS) != 0 {
		t.Fatalf("expected balance difference is : %v, actual: %v", ONE_KCS.String(), diff.String())
	}
}
