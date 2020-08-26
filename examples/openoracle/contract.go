// Copyright (c) 2020 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

// This example shows how to programmatically deploy a contract to IoTeX blockchain
// To run:
// go build; ./chaininfo

package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
)

var (
	gasPrice, _ = big.NewInt(0).SetString("1000000000000", 10)
	gasLimit    = uint64(1000000)
)

func main() {
	PrivateKey := os.Getenv("PrivateKey")
	if PrivateKey == "" {
		return
	}
	bin, err := ioutil.ReadFile("OpenOraclePriceData.bin")
	if err != nil {
		return
	}
	abi, err := ioutil.ReadFile("OpenOraclePriceData.abi")
	if err != nil {
		return
	}
	//"583aa7b02dbba44d257cad116e7e427d4b2040a0079c348d83636e100a4a4039"
	//s, err := NewOpenOracleService(PrivateKey, string(abi), string(bin), "", gasPrice, gasLimit, "api.testnet.iotex.one:80", false)
	//if err != nil {
	//	return
	//}
	//
	//r, err := s.Deploy(context.Background(), true)
	//fmt.Println("hash", r, err)

	authClient, err := NewOpenOracleService(PrivateKey, string(abi), string(bin), "io1wzrexqvzmevvkeyyzwyg6t6mw4696qnqfgz5jn", gasPrice, gasLimit, "api.testnet.iotex.one:80", false)
	if err != nil {
		return
	}
	testMessage, _ := hex.DecodeString("0x0000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000005ee4137c00000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000000233e65e300000000000000000000000000000000000000000000000000000000000000006707269636573000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000034254430000000000000000000000000000000000000000000000000000000000")
	testSignature, _ := hex.DecodeString("0x4e5d4cd04bd498da0e9d689aaf642e730e5135f9f96c3ca430f85cc5ab196698dcf9f816da11fe8308ec6c20c77762eba9a1583644fe61d4043ea7db7d8a79b6000000000000000000000000000000000000000000000000000000000000001c")
	hash, err := authClient.Put(context.Background(), testMessage, testSignature)
	fmt.Println("hash", hash, err)
}
