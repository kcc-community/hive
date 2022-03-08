package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/hive/simulators/kcc/punishment/punish"
	"github.com/ethereum/hive/simulators/kcc/punishment/reservepool"
	"github.com/ethereum/hive/simulators/kcc/punishment/validators"
	"log"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/hive/hivesim"
)

type Validator struct {
	client     *hivesim.Client // client
	enode      string          // enode
	miner      string          // address of validator
	privateKey string          // privatekey of validator
}

var (
	ValidatorsAddress  = common.HexToAddress("0x000000000000000000000000000000000000f333")
	ReservePoolAddress = common.HexToAddress("0x000000000000000000000000000000000000f999")
	PunishAddress = common.HexToAddress("0x000000000000000000000000000000000000f444")
	ONE_KCS            = new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	punishThreshold = 24
	removeThreshold = 48
)

func main() {
	suite := hivesim.Suite{
		Name: "kcc-punishment-smoke-test",
		Description: "PoSA: punishment smoke test",
	}

	suite.Add(hivesim.TestSpec{
		Name:        "punishment",
		Description: "single node punishment",
		AlwaysRun:   false,
		Run:         singleNodePunishment,
	})
	
	suite.Add(hivesim.TestSpec{
		Name:        "multi node punishment",
		Description: "random punish one validator of 3 to testify the networking whether working",
		AlwaysRun:   false,
		Run:         MultiNodesPunishment,
	})

	hivesim.MustRunSuite(hivesim.New(), suite)
}

func singleNodePunishment(t *hivesim.T) {

}

func MultiNodesPunishment(t *hivesim.T) {
	// 3 validators
	// And the first is the validator in genesis.json
	vals := []*Validator{
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
	for _, v := range vals {
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



	ethClient := ethclient.NewClient(vals[0].client.RPC())


	validatorCli, err := validators.NewValidators(ValidatorsAddress, ethClient)
	if err != nil {
		t.Fatalf("failed to new validator, err: %+v", err)
	}

	valPrivateKey, err := crypto.HexToECDSA("9c647b8b7c4e7c3490668fb6c11473619db80c93704c70893d3813af4090c39c")
	if err != nil {
		t.Fatalf("failed to parse private key for punish, err: %+v", err)
	}
	valTransactor, err := bind.NewKeyedTransactorWithChainID(valPrivateKey, big.NewInt(321))
	if err != nil {
		t.Fatalf("failed to create a singer for punish, err: %+v", err)
	}

	validatorCli.GetActiveValidators(valTransactor)

	t.Logf("waiting for the hard fork to come...")
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



	punishCli, err := punish.NewPunish(PunishAddress, ethClient)
	if err != nil {
		t.Fatalf("failed to new punish, err: %+v", err)
	}
	privateKey, err := crypto.HexToECDSA("9c647b8b7c4e7c3490668fb6c11473619db80c93704c70893d3813af4090c39c")
	if err != nil {
		t.Fatalf("failed to parse private key for punish, err: %+v", err)
	}
	transactor, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(321))
	if err != nil {
		t.Fatalf("failed to create a singer for punish, err: %+v", err)
	}



	punishedHeight := make(map[uint64]struct{})
	bn := uint64(0)
	for i := 0; i < punishThreshold; i++ {
		for {
			currentBlockNumber, err := ethClient.BlockNumber(context.TODO())
			if err != nil {
				t.Fatalf("failed to get block number: %v", err)
			}

			if _, ok := punishedHeight[currentBlockNumber]; !ok {
				punishCli.Punish(transactor, common.HexToAddress(vals[2].miner))
			}
			time.Sleep(time.Second * 3)
			bn = currentBlockNumber
		}
	}

	// balance of validator contract before punishment
	balanceBefore, err := ethClient.BalanceAt(context.TODO(), ValidatorsAddress, big.NewInt(int64(bn)))
	if err != nil {
		t.Logf("failed to get balance:%v", err)
	}

	// balance of validator contract after punishment
	balanceAfter, err := ethClient.BalanceAt(context.TODO(), ValidatorsAddress, big.NewInt(int64(bn+1)))
	if err != nil {
		t.Logf("failed to get balance:%v", err)
	}

	diff := new(big.Int).Sub(balanceAfter, balanceBefore)
	t.Logf("balance of validator contract, before punishment: %+v, after punishment: %+v, diff: %+v", balanceBefore.String(), balanceAfter.String(), diff.String())


	tx := new(types.Transaction)
	new(sync.Once).Do(func() {
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
		tx, err = reservePool.SetBlockRewardAmount(transactor, ONE_KCS)
		if err != nil {
			log.Fatalf("failed to setBlockReward: %v", err)
		}
	})



	// wait for the tx to be included in a block
	for {
		_, pending, err := ethClient.TransactionByHash(context.TODO(), tx.Hash())
		if err != nil {
			t.Fatalf("failed to get TX: %v", err)
		}

		if pending {
			t.Logf("waiting for the transaction of setting blockReward to be included in some block ...")
			time.Sleep(time.Second)
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

}