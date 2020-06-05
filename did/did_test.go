// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package main

import "C"
import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/iotexproject/iotex-antenna-go/v2/account"

	"github.com/iotexproject/iotex-antenna-go/v2/did/contract/IoTeXDID"
)

const (
	sender                = "io1ph0u2psnd7muq5xv9623rmxdsxc4uapxhzpg02"
	privateKey            = "414efa99dfac6f4095d6954713fb0085268d400d6a05a8ae8a69b5b1c10b4bed"
	endpoint              = "api.testnet.iotex.one:443"
	IoTeXDID_address      = "io1eurq3lx4lzx9wdj56plw5rm59f5qzanacr3raz"
	IoTeXDIDProxy_address = "io1zgs5gqjl679qlj4gqqpa9t329r8f5gr8xc9lr0"
)

func TestDidCreateDID(t *testing.T) {
	//require := require.New(t)
	h := CeateDID(endpoint, privateKey, IoTeXDIDProxy_address, IoTeXDID.IoTeXDIDABI, "1000000000000", uint64(1000000), "", "414efa99dfac6f4095d6954713fb0085268d400d6a05a8ae8a69b5b1c10b4bed", "url")
	fmt.Println(h)
}

func TestDidDeleteDID(t *testing.T) {
	//require := require.New(t)
	h := DeleteDID(endpoint, privateKey, IoTeXDIDProxy_address, IoTeXDID.IoTeXDIDABI, "1000000000000", uint64(1000000), "did:io:0x0ddfC506136fb7c050Cc2E9511eccD81b15e7426")
	fmt.Println(h)
}

func TestDidUpdateHash(t *testing.T) {
	//require := require.New(t)
	h := UpdateHash(endpoint, privateKey, IoTeXDIDProxy_address, IoTeXDID.IoTeXDIDABI, "1000000000000", uint64(1000000), "did:io:0x0ddfC506136fb7c050Cc2E9511eccD81b15e7426", "414efa99dfac6f4095d6954713fb0085268d400d6a05a8ae8a69b5b1c10ddddd")
	fmt.Println(h)
}

func TestDidUpdateUri(t *testing.T) {
	h := UpdateUri(endpoint, privateKey, IoTeXDIDProxy_address, IoTeXDID.IoTeXDIDABI, "1000000000000", uint64(1000000), "did:io:0x0ddfC506136fb7c050Cc2E9511eccD81b15e7426", "url2")
	fmt.Println(h)
}

func TestDidGetHash(t *testing.T) {
	acc, _ := account.HexStringToAccount(privateKey)
	fmt.Println(hex.EncodeToString(acc.Address().Bytes()))
	h := GetHash(endpoint, IoTeXDIDProxy_address, IoTeXDID.IoTeXDIDABI, "did:io:0x0ddfc506136fb7c050cc2e9511eccd81b15e7426")
	fmt.Printf("%s\n", C.GoString(h))
}

func TestDidGetUri(t *testing.T) {
	//require := require.New(t)
	h := GetUri(endpoint, IoTeXDIDProxy_address, IoTeXDID.IoTeXDIDABI, "did:io:0x0ddfC506136fb7c050Cc2E9511eccD81b15e7426")
	fmt.Println(h)
}
