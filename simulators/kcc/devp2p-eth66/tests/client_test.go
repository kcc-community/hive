package tests

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"testing"
)

func TestGetBlock(t *testing.T) {
	c, _ := rpc.Dial("https://rpc-mainnet.kcc.network")

	client := ethclient.NewClient(c)

	blocker, _ := client.BlockByNumber(context.TODO(), big.NewInt(10))
	m, _ := blocker.Header().MarshalJSON()
	fmt.Println(string(m), blocker.Header())
}

func TestX(t *testing.T) {
	fmt.Println(1 | 0 | 2)
}
