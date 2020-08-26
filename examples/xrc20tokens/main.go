// Copyright (c) 2020 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package main

import (
	"context"
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
	bin, err := ioutil.ReadFile("XRC20.bin")
	if err != nil {
		return
	}
	abi, err := ioutil.ReadFile("XRC20.abi")
	if err != nil {
		return
	}
	s, err := NewXrc20Service(PrivateKey, string(abi), string(bin), "", gasPrice, gasLimit, "api.testnet.iotex.one:80", false)
	if err != nil {
		return
	}
	r, err := s.Deploy(context.Background(), true, 2000000, "IOTX", "IOTX")
	fmt.Println("hash", r, err)

	r, err = s.Transfer(context.Background(), "io1zk6gqq0m2z9ytlu77t76e3632ezy39fa83xjnn", big.NewInt(10))
	fmt.Println("transfer", r, err)

	b, err := s.BalanceOf(context.Background(), "io1zk6gqq0m2z9ytlu77t76e3632ezy39fa83xjnn")
	fmt.Println("balance", b, err)
}
