package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/hive/hivesim"
)

var (
	// the number of seconds before a sync is considered stalled or failed
	syncTimeout = 65 * time.Second
	params      = hivesim.Params{
		"HIVE_CLIQUE_PRIVATEKEY": "9c647b8b7c4e7c3490668fb6c11473619db80c93704c70893d3813af4090c39c",
		"HIVE_MINER":             "658bdf435d810c91414ec09147daa6db62406379",
		"HIVE_NETWORK_ID":         "321",
		"HIVE_CHAIN_ID":       "321",

		"HIVE_NODETYPE":       "",
		// block interval: 1s
		"HIVE_KCC_POSA_BLOCK_INTERVAL": "1",
		// epoch : 5
		"HIVE_KCC_POSA_EPOCH": "10",
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
	}
	sourceFiles = map[string]string{
		"genesis.json": "./simplechain/genesis.json",
		"chain.rlp":    "./simplechain/chain.rlp",
		"geth.sh": 		"./simplechain/geth.sh",
	}
	sinkFiles = map[string]string{
		"genesis.json": "./simplechain/genesis.json",
		"geth.sh": "./simplechain/geth.sh",
	}
)

var (
	testchainHeadNumber = uint64(8591)
	testchainHeadHash   = common.HexToHash("0x0520a5c756a2b225ac8e35fbd027f82c3fb21082a1f53c178540df5e94000013")
)

func main() {
	var suite = hivesim.Suite{
		Name: "sync with other clients",
		Description: `This suite of tests verifies that clients can sync from each other in different modes.
For each client, we test if it can serve as a sync source for all other clients (including itself).`,
	}
	suite.Add(hivesim.ClientTestSpec{
		Role:        "eth1",
		Name:        "CLIENT as sync source",
		Description: "This loads the test chain into the client and verifies whether it was imported correctly.",
		Parameters:  params,
		Files:       sourceFiles,
		Run:         runSourceTest,
	})
	hivesim.MustRunSuite(hivesim.New(), suite)
}

func runSourceTest(t *hivesim.T, c *hivesim.Client) {
	// Check whether the source has imported its chain.rlp correctly.
	source := &node{c}
	if err := source.checkHead(testchainHeadNumber, testchainHeadHash); err != nil {
		t.Fatal(err)
	}

	// Configure sink to connect to the source node.
	enode, err := source.EnodeURL()
	if err != nil {
		t.Fatal("can't get node peer-to-peer endpoint:", enode)
	}
	t.Logf("enode: %s\n", enode)
	sinkParams := params.Set("HIVE_BOOTNODE", enode)

	// Sync all sink nodes against the source.
	t.RunAllClients(hivesim.ClientTestSpec{
		Role:        "eth1",
		Name:        fmt.Sprintf("sync %s -> CLIENT", source.Type),
		Description: fmt.Sprintf("This test attempts to sync the chain from a %s node.", source.Type),
		Parameters:  sinkParams,
		Files:       sinkFiles,
		Run:         runSyncTest,
	})
}

func runSyncTest(t *hivesim.T, c *hivesim.Client) {
	node := &node{c}
	err := node.checkSync(t, testchainHeadNumber, testchainHeadHash)
	if err != nil {
		t.Fatal("sync failed:", err)
	}
}

type node struct {
	*hivesim.Client
}

// checkSync waits for the node to reach the head of the chain.
func (n *node) checkSync(t *hivesim.T, wantNumber uint64, wantHash common.Hash) error {
	var (
		timeout = time.After(syncTimeout)
		current = uint64(0)
	)
	for {
		select {
		case <-timeout:
			return fmt.Errorf("timeout (%v elapsed, current head is %d)", syncTimeout, current)
		default:
			block, err := n.head()
			if err != nil {
				t.Logf("error getting block from %s (%s): %v", n.Type, n.Container, err)
				return err
			}
			blockNumber := block.Number.Uint64()
			if blockNumber != current {
				t.Logf("%s has new head %d\n", n.Type, blockNumber)
			}
			if current == wantNumber {
				if block.Hash() != wantHash {
					return fmt.Errorf("wrong head hash %x, want %x", block.Hash(), wantHash)
				}
				return nil // success
			}
			// check in a little while....
			current = blockNumber
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

// head returns the node's chain head.
func (n *node) head() (*types.Header, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	return ethclient.NewClient(n.RPC()).HeaderByNumber(ctx, nil)
}

// checkHead checks whether the remote chain head matches the given values.
func (n *node) checkHead(num uint64, hash common.Hash) error {
	head, err := n.head()
	if err != nil {
		return fmt.Errorf("can't query chain head: %v", err)
	}
	if head.Hash() != hash {
		fmt.Printf("info: wrong chain head %d (%s), want %d (%s)\n", head.Number, head.Hash().String(), num, hash.String())

		return fmt.Errorf("wrong chain head %d (%s), want %d (%s)", head.Number, head.Hash().TerminalString(), num, hash.TerminalString())
	} else {
		fmt.Printf("have the same header!\n")
	}
	return nil
}
