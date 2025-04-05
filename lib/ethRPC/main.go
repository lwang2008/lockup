// Copyright 2021, Lucas Wang
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package ethRPC

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	ERC721 "github.com/lwang2008/lockup/interfaces"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func VerifyTokenOwner(_tokenId, _contractAddress, _rpcURL string) (string, string, string, error) {
	// Prepare the JSON-RPC request
	requestBody := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_call",
		"params": []interface{}{
			map[string]interface{}{
				"to":   _contractAddress,
				"data": fmt.Sprintf("0x6352211e00000000000000000000000000000000000000000000000000000000%s", _tokenId),
			},
			"latest",
		},
		"id": 1,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", "", "", fmt.Errorf("error marshaling request: %v", err)
	}

	// Send the request
	resp, err := http.Post(_rpcURL, "application/json", strings.NewReader(string(jsonData)))
	if err != nil {
		return "", "", "", fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", "", fmt.Errorf("error reading response: %v", err)
	}

	// Parse the response
	var result struct {
		Result string `json:"result"`
		Error  struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return "", "", "", fmt.Errorf("error unmarshaling response: %v", err)
	}

	if result.Error.Code != 0 {
		return "", "", "", fmt.Errorf("RPC error: %s", result.Error.Message)
	}

	// Remove the "0x" prefix and pad with zeros if necessary
	owner := result.Result
	if strings.HasPrefix(owner, "0x") {
		owner = owner[2:]
	}
	owner = strings.TrimPrefix(owner, "0x")
	for len(owner) < 40 {
		owner = "0" + owner
	}

	// Create the hash data
	hashData := fmt.Sprintf("0x%s", owner)
	hash := fmt.Sprintf("0x%x", sha256.Sum256([]byte(hashData)))

	return hashData, hash, owner, nil
}

func VerifySignature(_input string, _hash []byte, _owner string) bool {
	signature := hexutil.MustDecode(_input)
	fmt.Println(hexutil.Encode(signature))
	if hexutil.Encode(signature[64:]) == "0x1b" {
		signature[64] = hexutil.MustDecode("0x00")[0]
	}
	if hexutil.Encode(signature[64:]) == "0x1c" {
		signature[64] = hexutil.MustDecode("0x01")[0]
	}

	signer, err := crypto.SigToPub(_hash, signature)
	if err != nil {
		panic(err)
	}

	return crypto.PubkeyToAddress(*signer).String() == _owner

}
