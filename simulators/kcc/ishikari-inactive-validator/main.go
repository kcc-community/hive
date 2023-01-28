package main

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"

	"github.com/ethereum/hive/hivesim"

	validatorsContract "github.com/ethereum/hive/simulators/kcc/ishikari-inactive-validator/validators"
)

func main() {
	suite := hivesim.Suite{
		Name:        "Ishikari Hardfork: Active validator becomes inactive",
		Description: "Ishikari Hardfork: Active validator becomes inactive",
	}
	suite.Add(hivesim.TestSpec{
		Name:        "Ishikari Hardfork: Active validator becomes inactive",
		Description: "Ishikari Hardfork: Active validator becomes inactive",
		Run:         activeBecomeInactive,
	})
	hivesim.MustRunSuite(hivesim.New(), suite)
}

func activeBecomeInactive(t *hivesim.T) {

	// The admin user
	// And he will also be our vote for testing
	admin := getUser(t)

	// create 30 private key for each validator
	validators := createAddress(t, 30)

	// We will start 31 validators
	for _, v := range validators {

		v.client = t.StartClient("kcc", hivesim.Params{
			"HIVE_CLIQUE_PRIVATEKEY": common.Bytes2Hex(v.key.D.Bytes()),
			"HIVE_MINER":             v.address.Hex()[2:],
			"HIVE_CHAIN_ID":          "321",
			// block interval: 1s
			"HIVE_KCC_POSA_BLOCK_INTERVAL": "1",
			// epoch : 10
			"HIVE_KCC_POSA_EPOCH": "40",
			// initial valiators
			"HIVE_KCC_POSA_ISHIKARI_INIT_VALIDATORS": ToAddressCSList(validators),

			// the manager of each validator is the validator itself

			"HIVE_FORK_KCC_ISHIKARI": "39",
			// KCC Ishikari Patch001 fork number
			"HIVE_FORK_KCC_ISHIKARI_PATCH001": "39",
			// KCC Ishikari Patch002 fork number
			"HIVE_FORK_KCC_ISHIKARI_PATCH002": "39",

			// admin
			"HIVE_KCC_POSA_ADMIN": admin.address.Hex(),
			// sync mode
			"HIVE_NODETYPE": "archive",
		}, hivesim.WithStaticFiles(
			map[string]string{
				"/genesis.json": "genesis.json",
			}))
	}

	// collecting enode of each validator
	timer := time.NewTimer(time.Minute * 2)
	for _, v := range validators {

		// inner retry loop
		for {
			select {
			case <-timer.C:
				t.Fatalf("time out when collecting all the enodes")
			default:
			}

			enode, err := v.client.EnodeURL()
			if err != nil {
				t.Logf("cannot get the enode of %v:  %v, retrying..", enode, err)
				continue
			}

			v.enode = enode
			break
		}
	}
	timer.Stop()

	// connect all nodes
	timer.Reset(time.Minute * 2)
	for i := range validators {
		for j := range validators {
			if j > i {
				// inner retry loop
				for {
					select {
					case <-timer.C:
						t.Fatalf("time out when connection all nodes")
					default:
					}
					err := connectClients(t, validators[i], validators[j])
					if err != nil {
						t.Logf("failed to connect %v to %v: %v",
							validators[i].address, validators[j].address, err)
						continue
					}
					break
				}
			}
		}
	}
	timer.Stop()

	// Wait for all the nodes to get synced
	for {

		time.Sleep(time.Second * 10)

		isEveryNodeUptodate := true

		t.Logf("==== Sync Status ====> ")
		for _, c := range validators {
			n, err := ethclient.NewClient(c.client.RPC()).BlockNumber(context.Background())
			isEveryNodeUptodate = isEveryNodeUptodate && (n > 80)
			if err != nil {
				t.Logf("%v: %v", c.address.Hex(), err.Error())
			} else {
				t.Logf("%v: #%v", c.address.Hex(), n)
			}
		}

		// if every node has synced up to block #80,
		// we will proceed
		if isEveryNodeUptodate {
			t.Logf("Nodes are all synced.")
			break
		}
	}

	t.Logf("finished setting up fixtures..")

	client := ethclient.NewClient(validators[0].client.RPC())

	// Get current active validators
	vContract, err := validatorsContract.NewValidators(
		common.HexToAddress("0x000000000000000000000000000000000000f333"),
		client)
	if err != nil {
		t.Fatalf("failed to create validator contract instance: %v", err)
	}

	//
	// First: Take a snapshot of the active validators
	//

	number, err := client.BlockNumber(context.Background())
	if err != nil {
		t.Fatalf("failed to get current block number: %v", err)
	}

	// Take a snapshot of the info of all validators
	// Return the last found inactive validator
	snapshot := func(n uint64) *AddressAndKey {

		// last found inactive validator
		var inactive *AddressAndKey

		t.Logf("====> Snapshot at block #%v", n)
		activeValidators, err := vContract.GetActiveValidators(&bind.CallOpts{
			BlockNumber: new(big.Int).SetUint64(n),
		})
		if err != nil {
			t.Fatalf("failed to get active validators")
		}

		isActive := BuildAddressMap(activeValidators)

		for i, v := range validators {
			ballots, err := vContract.GetPoolsuppliedBallot(&bind.CallOpts{
				BlockNumber: new(big.Int).SetUint64(n),
			}, v.address)
			if err != nil {
				t.Fatalf("failed to get supplied ballot of %v: %v", v.address, err)
			}

			if isActive[v.address] {
				t.Logf("[%v] %v %v active", i, v.address, ballots.String())
			} else {
				inactive = v
				t.Logf("[%v] %v %v inactive", i, v.address, ballots.String())
			}
		}
		t.Logf("=====================================")

		return inactive
	}

	lastInactive := snapshot(number)

	if lastInactive == nil {
		t.Fatalf("no inactive validator found")
	}

	//
	// Vote for the inactive validator
	//
	t.Logf("voting for validator %v ...", lastInactive.address)
	transactOpt, err := bind.NewKeyedTransactorWithChainID(admin.key, big.NewInt(321))
	if err != nil {
		t.Fatalf("failed to create transact options: %v", err)
	}

	transactOpt.Value = new(big.Int).SetUint64(params.Ether)

	tx, err := vContract.Vote(transactOpt, lastInactive.address)
	if err != nil {
		t.Fatalf("failed to vote for the inactive validator: %v", err)
	}

	t.Logf("waiting for tx to be included in some block")

	// in which block, this tx has been included
	var includedBlock uint64

	timer.Reset(time.Minute * 2)
	for {
		select {
		case <-timer.C:
			t.Fatalf("timeout when wait for tx to be included")
		default:
		}

		time.Sleep(time.Second * 3)

		n, err := client.BlockNumber(context.Background())
		if err != nil {
			t.Logf("failed to get current block number: %v", err)
		}

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			t.Logf("failed to get tx receipts: %v", err)
			continue
		}

		if receipt.Status == 0 {
			t.Fatalf("tx reverted ")
		}

		// wait for 12 block confirm
		if receipt.BlockNumber.Uint64()+12 < n {
			t.Logf("tx has being included in block #%v", receipt.BlockNumber.Uint64())
			includedBlock = receipt.BlockNumber.Uint64()
			break
		}
	}
	timer.Stop()

	snapshot(includedBlock)

	t.Logf("waiting for the next epoch")

	targetBlock := (includedBlock/40 + 1) * 40

	timer.Reset(time.Minute * 2)
	for {
		select {
		case <-timer.C:
			t.Fatalf("timeout when waiting for targetBlock")
		default:
		}
		time.Sleep(time.Second)

		n, err := client.BlockNumber(context.Background())
		if err != nil {
			t.Logf("failed to get current block number: %v", err)
			continue
		}

		t.Logf("current block number: %v", n)

		if n > targetBlock+12 {
			break
		}
	}

	nowInvalidValidator := snapshot(targetBlock)

	if nowInvalidValidator.address == lastInactive.address {
		t.Fatalf("the inactive validator should not be the same")
	}

}

func connectClients(t *hivesim.T, c1 *AddressAndKey, c2 *AddressAndKey) error {

	return c1.client.RPC().Call(nil, "admin_addPeer", c2.enode)

}

// Address and its private key
type AddressAndKey struct {
	address common.Address
	key     *ecdsa.PrivateKey
	// optional fields for validator
	client *hivesim.Client
	enode  string
}

// Get the AddressAndKey of the preset user
// We have allocated some KCS to this user in the genesis block.
func getUser(t *hivesim.T) *AddressAndKey {

	key, err := crypto.HexToECDSA("3470d1a4c35904a0fc246765b2456e04eab5bf6725a8550805e198f663a135af")
	if err != nil {
		t.Fatalf("failed to decode private: %v", err)
	}

	return &AddressAndKey{
		address: crypto.PubkeyToAddress(key.PublicKey),
		key:     key,
	}
}

// create addresses
func createAddress(t *hivesim.T, n uint64) []*AddressAndKey {

	if n == 0 {
		return nil
	}

	ret := make([]*AddressAndKey, n)

	// The first key is always the miner
	{
		// The miner for the first epoch
		miner, err := crypto.HexToECDSA("9c647b8b7c4e7c3490668fb6c11473619db80c93704c70893d3813af4090c39c")
		if err != nil {
			t.Fatalf("failed to parse miner key: %v", err)
		}

		ret[0] = &AddressAndKey{
			address: crypto.PubkeyToAddress(miner.PublicKey),
			key:     miner,
		}
	}

	for i := uint64(1); i < n; i++ {

		key, err := crypto.GenerateKey()
		if err != nil {
			t.Fatalf("failed to create a new key: %v", err)
		}

		ret[i] = &AddressAndKey{
			address: crypto.PubkeyToAddress(key.PublicKey),
			key:     key,
		}

	}

	return ret

}

func ToAddressCSList(addrs []*AddressAndKey) string {
	builder := &strings.Builder{}

	for i, a := range addrs {
		builder.WriteString(a.address.Hex())
		if i < len(addrs)-1 {
			builder.WriteString(",")
		}
	}
	return builder.String()
}

func BuildAddressMap(addrs []common.Address) map[common.Address]bool {
	m := make(map[common.Address]bool)
	for _, a := range addrs {
		m[a] = true
	}
	return m
}
