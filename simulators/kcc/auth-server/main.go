package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/hive/hivesim"
	"github.com/golang-jwt/jwt/v4"
)

var (
	MinerKey  = "0x9c647b8b7c4e7c3490668fb6c11473619db80c93704c70893d3813af4090c39c"
	MinerAddr = common.HexToAddress("0x658bdf435d810c91414ec09147daa6db62406379")
	JWTSecret = common.Hex2Bytes("7365637265747365637265747365637265747365637265747365637265747365")
)

func main() {
	suite := hivesim.Suite{
		Name:        "Auth Server",
		Description: "Test the RPC server with JWT authentication",
	}
	suite.Add(hivesim.ClientTestSpec{
		Role:        "eth1",
		Name:        "ETH API",
		Description: "Get the latest block number from the Auth Server",
		Files: map[string]string{
			"/genesis.json": "genesis.json",
			"/geth.sh":      "geth.sh", // override the default geth.sh
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
		Run: testOpenAPIFromAuthSrv,
	}).Add(hivesim.ClientTestSpec{
		Role:        "eth1",
		Name:        "Engine API",
		Description: "The 'Engine' API namespace should be disabled",
		Files: map[string]string{
			"/genesis.json": "genesis.json",
			"/geth.sh":      "geth.sh", // override the default geth.sh
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
		Run: testNoEngineAPI,
	})

	hivesim.MustRunSuite(hivesim.New(), suite)
}

func HexWithPrefixToUint64(hex string) (uint64, error) {
	return strconv.ParseUint(hex[2:], 16, 64)
}

func testOpenAPIFromAuthSrv(t *hivesim.T, c *hivesim.Client) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	for {
		var number hexutil.Uint64

		// get the block number
		if err := c.RPC().CallContext(ctx, &number, "eth_blockNumber"); err != nil {
			t.Fatalf("failed to get block number: %v", err)
		}

		if number > 11 {
			break
		}

		t.Logf("block number: %d", number)

		select {
		case <-ctx.Done():
			t.Fatalf("timeout")
		default:
		}

		time.Sleep(time.Second)

	}

	// get the block number with the jwt auth token

	// create a JWT token
	token, err := newJTWToken()

	if err != nil {
		t.Fatalf("failed to create the JWT token: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("http://%s:8551", c.IP),
		bytes.NewReader([]byte(`{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}`)))
	if err != nil {
		t.Fatalf("failed to create the http request: %v", err)
	}

	req.Header = http.Header{
		"Authorization": []string{fmt.Sprintf("Bearer %s", token)},
		"accept":        []string{"application/json"},
		"content-type":  []string{"application/json"},
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {

		// connection refused
		// That means the auth server is completely disabled
		if strings.Contains(err.Error(), "connect: connection refused") {
			t.Logf("auth server is disabled")
			return
		}

		// other cases
		t.Fatalf("failed to send the http request: %v", err)
	}

	// If the auth server is enabled, it should handle the RPC call successfully

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read the http response body: %v", err)
	} else if resp.StatusCode != http.StatusOK {
		t.Logf("http response: %s", string(bodyBytes))
		t.Fatalf("http response status code: %d", resp.StatusCode)
	}

	var rawRPCResp struct {
		BlockNumber hexutil.Uint64 `json:"result"`
	}

	if err := json.Unmarshal(bodyBytes, &rawRPCResp); err != nil {
		t.Fatalf("failed to unmarshal the http response body: %v", err)
	}

	t.Logf("block number from auth server: %d", rawRPCResp.BlockNumber)

	if rawRPCResp.BlockNumber < 11 {
		t.Fatalf("block number from auth server is less than 11")
	}

}

func testNoEngineAPI(t *hivesim.T, c *hivesim.Client) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	// create a JWT token
	token, err := newJTWToken()
	if err != nil {
		t.Fatalf("failed to create the JWT token: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("http://%s:8551", c.IP),
		// https://github.com/ethereum/execution-apis/blob/d03c193dc317538e2a1a098030c21bacc2fd1333/src/engine/paris.md#engine_getpayloadv1
		// Let's check the engine_getPayloadV1 API
		bytes.NewReader([]byte(`{"jsonrpc":"2.0","method":"engine_getPayloadV1","params":["0xffffffffffffffff"],"id":1}`)))
	if err != nil {
		t.Fatalf("failed to create the http request: %v", err)
	}

	req.Header = http.Header{
		"Authorization": []string{fmt.Sprintf("Bearer %s", token)},
		"accept":        []string{"application/json"},
		"content-type":  []string{"application/json"},
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {

		// connection refused
		// That means the auth server is completely disabled
		if strings.Contains(err.Error(), "connect: connection refused") {
			t.Logf("auth server is disabled")
			return
		}

		t.Fatalf("failed to send the http request: %v", err)
	}

	// If the auth server is enabled, it should not accept any RPC calls from the Engine namespace,
	// and returns a "-32601" error code.

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)

	t.Logf("http response: %s", string(bodyBytes))

	if err != nil {
		t.Fatalf("failed to read the http response body: %v", err)
	} else if resp.StatusCode != http.StatusOK {
		t.Fatalf("http response status code: %d", resp.StatusCode)
	}

	var rawRPCResp struct {
		Error struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}

	if err := json.Unmarshal(bodyBytes, &rawRPCResp); err != nil {
		t.Fatalf("failed to unmarshal the http response body: %v", err)
	}

	t.Logf("error code: %d, error msg: %v", rawRPCResp.Error.Code, rawRPCResp.Error.Message)

	if rawRPCResp.Error.Code != -32601 {
		t.Fatalf("error code is not -32601")
	}

}

// create a jwt token
func newJTWToken() (string, error) {

	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "hive",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 60)),
		IssuedAt:  jwt.NewNumericDate(time.Now().Add(-time.Second)),
		NotBefore: jwt.NewNumericDate(time.Now().Add(-time.Second)),
	}).SignedString(JWTSecret)

}
