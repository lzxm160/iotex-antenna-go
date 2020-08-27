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
		fmt.Println("Environment Variable [PrivateKey] not defined")
		return
	}

	bin, err := ioutil.ReadFile("multisend.bin")
	if err != nil {
		fmt.Println("multiSend.bin not found")
		return
	}

	abi, err := ioutil.ReadFile("multisend.abi")
	if err != nil {
		fmt.Println("multiSend.abi not found")
		return
	}

	//s, err := NewMultiSendService(PrivateKey, string(abi), string(bin), "", gasPrice, gasLimit, "api.testnet.iotex.one:80", false)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//r, err := s.Deploy(context.Background(), true)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("Contract deployed, action hash:", r)

	s, err := NewMultiSendService(PrivateKey, string(abi), string(bin), "io17hc9cdtxc58mhfnwrjf0mvsgr5zvm3mdgflzp7", gasPrice, gasLimit, "api.testnet.iotex.one:80", false)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 0.1iotex once
	wasteTime := make([]uint64, 0)
	for i := 0; i < 30; i++ {
		r, err := s.MultiSend(context.Background(), []string{"io1vdtfpzkwpyngzvx7u2mauepnzja7kd5rryp0sg"}, []*big.Int{big.NewInt(10)})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Token MultiSend completed: ", r)
		t, err := s.CheckTime(context.Background(), r)
		if err != nil {
			fmt.Println("failed one", r, err)
			continue
		}
		fmt.Println("ttttttttt", t)
		wasteTime = append(wasteTime, t)
	}
	var tt uint64
	var max uint64 = wasteTime[0]
	var min uint64 = wasteTime[0]
	for _, t := range wasteTime {
		if t > max {
			max = t
		}
		if t < min {
			min = t
		}
		tt += t
	}
	fmt.Println("float64(tt)", float64(tt))
	fmt.Println("float64(len(wasteTime))", float64(len(wasteTime)))
	fmt.Println("everage time:", float64(tt)/float64(len(wasteTime)))
	fmt.Println("max", max)
	fmt.Println("min", min)
}
