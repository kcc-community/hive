package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/hive/simulators/kcc/ishikari-punishment/reservepool"
	"github.com/ethereum/hive/simulators/kcc/ishikari-punishment/validators"

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
	PunishAddress      = common.HexToAddress("0x000000000000000000000000000000000000f444")
	ONE_KCS            = new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	punishThreshold    = 24
	removeThreshold    = 48
)

func main() {
	suite := hivesim.Suite{
		Name:        "Ishikari hardfork punishment",
		Description: "",
	}

	suite.Add(hivesim.TestSpec{
		Name:        "punishment",
		Description: "punish validator's reward",
		AlwaysRun:   false,
		Run:         punishRewards,
	})

	suite.Add(hivesim.TestSpec{
		Name:        "multi node punishment",
		Description: "random punish one validator of 3 to testify the networking whether working",
		AlwaysRun:   false,
		Run:         MultiNodesPunishment,
	})

	hivesim.MustRunSuite(hivesim.New(), suite)
}

func punishRewards(t *hivesim.T) {
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
			"HIVE_KCC_POSA_BLOCK_INTERVAL": "1",
			// epoch : 5
			"HIVE_KCC_POSA_EPOCH": "5",
			// initial valiators
			"HIVE_KCC_POSA_ISHIKARI_INIT_VALIDATORS": "0x658bdf435d810c91414ec09147daa6db62406379,0xa885d3767358B3ad7A5aD9dA5d4508580b1D2480,0x2fE42368F0b91f87f5bfab781Cb89520F3ec78aC",
			// admin
			"HIVE_KCC_POSA_ADMIN": "0x658bdf435d810c91414ec09147daa6db62406379",
			// KCC Ishikari  fork number
			"HIVE_FORK_KCC_ISHIKARI": "9",
			// sync mode
			"HIVE_NODETYPE": "archive",
		}, hivesim.WithStaticFiles(
			map[string]string{
				"/genesis.json": "genesis.json",
			}))

		for {
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
	for i := range vals {
		for j := range vals {
			if i != j {
				err := vals[j].client.RPC().Call(nil, "admin_addPeer", vals[i].enode)
				if err != nil {
					t.Fatalf("failed to add %v as a peer of %v", vals[i].miner, vals[j].miner)
				}
			}
		}
	}

	fmt.Printf("nodes are connected..\n")

	ethClient := ethclient.NewClient(vals[0].client.RPC())

	t.Logf("waiting for the hard fork to come...")
	for {
		number, err := ethClient.BlockNumber(context.TODO())
		if err != nil {
			t.Fatalf("failed to get block number:  %v\n", err)
			return
		}

		if number < 9 {
			t.Logf("current Block number is #%v", number)
			time.Sleep(time.Second)
			continue
		}
		break
	}

	// Set blockReward of reservePool
	//
	reservePool, err := reservepool.NewReservepool(ReservePoolAddress, ethClient)
	if err != nil {
		t.Fatalf("failed to create reservesPool: %v", err)
		return
	}

	privateKey, err := crypto.HexToECDSA("9c647b8b7c4e7c3490668fb6c11473619db80c93704c70893d3813af4090c39c")
	if err != nil {
		t.Fatalf("failed to decode privatekey hex as private key")
		return
	}
	transactor, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(321))
	if err != nil {
		t.Fatalf("failed to create transactor with private key: %v", err)
		return
	}

	// Set the blockReward amount to 1 KCS
	tx, err := reservePool.SetBlockRewardAmount(transactor, ONE_KCS)
	if err != nil {
		log.Fatalf("failed to setBlockReward: %v", err)
		return
	}
	// wait for the tx to be included in a block
	for {
		_, pending, err := ethClient.TransactionByHash(context.TODO(), tx.Hash())
		if err != nil {
			t.Fatalf("failed to get TX: %v", err)
			return
		}

		if pending {
			t.Logf("waiting for the transaction of setting blockReward to be included in some block ...")
			time.Sleep(time.Second)
			continue
		}
		break
	}

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
			"HIVE_KCC_POSA_BLOCK_INTERVAL": "1",
			// epoch : 30
			"HIVE_KCC_POSA_EPOCH": "30",
			// initial valiators
			"HIVE_KCC_POSA_ISHIKARI_INIT_VALIDATORS": "0x658bdf435d810c91414ec09147daa6db62406379,0xa885d3767358B3ad7A5aD9dA5d4508580b1D2480,0x2fE42368F0b91f87f5bfab781Cb89520F3ec78aC",
			// admin
			"HIVE_KCC_POSA_ADMIN": "0x658bdf435d810c91414ec09147daa6db62406379",
			// KCC Ishikari  fork number
			"HIVE_FORK_KCC_ISHIKARI": "29",
			// sync mode
			"HIVE_NODETYPE": "archive",
		}, hivesim.WithStaticFiles(
			map[string]string{
				"/genesis.json": "genesis.json",
			}))

		for {
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
	for i := range vals {
		for j := range vals {
			if i != j {
				err := vals[j].client.RPC().Call(nil, "admin_addPeer", vals[i].enode)
				if err != nil {
					t.Fatalf("failed to add %v as a peer of %v", vals[i].miner, vals[j].miner)
				}
			}
		}
	}

	t.Logf("nodes are connected...\n")

	ethClient := ethclient.NewClient(vals[0].client.RPC())

	t.Logf("waiting for the hard fork to come...")
	for {
		number, err := ethClient.BlockNumber(context.TODO())
		if err != nil {
			t.Fatalf("failed to get block number:  %v\n", err)
			return
		}

		if number < 29 {
			t.Logf("current Block number is #%v", number)
			time.Sleep(time.Second)
			continue
		}
		break
	}

	validatorCli, err := validators.NewValidators(ValidatorsAddress, ethClient)
	if err != nil {
		t.Fatalf("failed to new validator, err: %+v", err)
		return
	}

	address := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")

	t.Logf("after hard fork, before punishment: \n")
	beforeValidatorsSlice, err := validatorCli.GetActiveValidators(&bind.CallOpts{
		Pending:     false,
		From:        address,
		BlockNumber: nil,
		Context:     nil,
	})

	t.Logf("GetActiveValidators length: %+v\n", len(beforeValidatorsSlice))
	for i := range beforeValidatorsSlice {
		t.Logf("GetActiveValidators: %+v\n", beforeValidatorsSlice[i])
	}

	// Notes: will to get error: no contract code at given address, when call GetTopValidators at height less than 9
	top, err := validatorCli.GetTopValidators(&bind.CallOpts{
		Pending:     false,
		From:        address,
		BlockNumber: big.NewInt(9),
		Context:     nil,
	})
	if err != nil {
		t.Logf("failed to get top validators, err: %+v\n", err)
		return
	}
	t.Logf("top validators length: %+v at height: %d\n", len(top), 9)

	t.Logf("...waiting 10 seconds...")
	time.Sleep(time.Second * 10)

	h, _ := ethClient.BlockNumber(context.TODO())
	t.Logf("ready to stop a client: %+v at height: %d to simulate punish a validator...", vals[2].miner, h)

	if err = t.Sim.StopClient(t.SuiteID, t.TestID, vals[2].client.Container); err != nil {
		t.Logf("failed to stop specified container, err: %+v\n", err)
		return
	}

	h, err = ethClient.BlockNumber(context.TODO())
	if err != nil {
		t.Fatalf("failed to get block number:  %v\n", err)
		return
	}

	t.Logf("stop a client at height: %d", h)
	t.Logf("waiting for the punishment to be done...\n")

	times := 0
	for {
		number, err := ethClient.BlockNumber(context.TODO())
		if err != nil {
			t.Fatalf("failed to get block number:  %v\n", err)
			return
		}

		afterValidatorsSlice, err := validatorCli.GetActiveValidators(&bind.CallOpts{
			Pending:     false,
			From:        address,
			BlockNumber: big.NewInt(int64(number)),
			Context:     nil,
		})
		if err != nil {
			t.Logf("failed to get active validators, err: %+v\n", err)
			return
		}
		if len(afterValidatorsSlice) < 3 {
			times++
			t.Logf("## only %d validators are still working, confirmed times: %d\n", len(afterValidatorsSlice), times)
		} else {
			t.Logf("current Block number is #%v", number)
			time.Sleep(time.Second * 3)
			continue
		}

		if times == 10 {
			break
		}
	}

	top, err = validatorCli.GetTopValidators(&bind.CallOpts{
		Pending:     false,
		From:        address,
		BlockNumber: nil,
		Context:     nil,
	})
	if err != nil {
		t.Logf("failed to get top validators, err: %+v\n", err)
		return
	}
	t.Logf("after punishment, top validators length: %+v at height: %d\n", len(top), 9)

	afterValidatorsSlice, err := validatorCli.GetActiveValidators(&bind.CallOpts{
		Pending:     false,
		From:        address,
		BlockNumber: nil,
		Context:     nil,
	})

	t.Logf("after punishment, GetActiveValidators length: %+v\n", len(afterValidatorsSlice))
	for i := range afterValidatorsSlice {
		t.Logf("after punishment GetActiveValidators: %+v\n", afterValidatorsSlice[i])
	}
}
