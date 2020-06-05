// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/iotexproject/iotex-antenna-go/v2/did/contract/IoTeXDID"
)

const (
	sender                = "io1ph0u2psnd7muq5xv9623rmxdsxc4uapxhzpg02"
	privateKey            = "414efa99dfac6f4095d6954713fb0085268d400d6a05a8ae8a69b5b1c10b4bed"
	endpoint              = "api.testnet.iotex.one:443"
	IoTeXDID_address      = "io1eurq3lx4lzx9wdj56plw5rm59f5qzanacr3raz"
	IoTeXDIDProxy_address = "io1zgs5gqjl679qlj4gqqpa9t329r8f5gr8xc9lr0"
)

func TestDidCreateDid(t *testing.T) {
	require := require.New(t)
	h, err := CeateDID(endpoint, privateKey, IoTeXDIDProxy_address, IoTeXDID.IoTeXDIDABI, "1000000000000", uint64(1000000), "", "414efa99dfac6f4095d6954713fb0085268d400d6a05a8ae8a69b5b1c10b4bed", "url")
	require.NoError(err)
	fmt.Println(h)
}

func TestDidDelete(t *testing.T) {

}

func TestDidUpdateHash(t *testing.T) {

}

func TestDidUpdateUri(t *testing.T) {

}

func TestDidReadHashContract(t *testing.T) {

}

func TestDidReadUriContract(t *testing.T) {

}
