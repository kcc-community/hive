package tests

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"testing"
)

func TestEthGetTransaction(t *testing.T) {
	c, _ := rpc.Dial("https://rpc-mainnet.kcc.network")
	tx := new(rpcTransaction)
	err := c.Call(tx, "eth_getTransactionByHash", "0x886feb9c5262c041340b86b72ad9a9882e9d1d2cdb7dc22078d93896cdfac94a")
	if err != nil {
		panic(err)
	}

	fmt.Println(tx)
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
