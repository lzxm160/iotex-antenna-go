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
	// get some data from coinbase,curl https://api.pro.coinbase.com/oracle, have to be latest data or it will fail
	mes := "0x0000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000005f45b26c00000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000002a036b7380000000000000000000000000000000000000000000000000000000000000006707269636573000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000034254430000000000000000000000000000000000000000000000000000000000"
	sig := "0xa8939cf38773386575dbd95ad129439825319ed34c4740c31b8337c513e601432a4a00ccb4019d5fe7e592538b69f596b68d70132ef05554036aff7d2b07e7eb000000000000000000000000000000000000000000000000000000000000001c"
	testMessage, _ := hex.DecodeString(mes)
	testSignature, _ := hex.DecodeString(sig)
	hash, err := authClient.Put(context.Background(), testMessage, testSignature)
	fmt.Println("hash", hash, err)
}
