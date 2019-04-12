// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package contract

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	//export accountPrivateKey="9cdf22c5caa8a4d99eb674da27756b438c05c6b1e8995f4a0586745e2071b115"
	//export accountAddress="io14gnqxf9dpkn05g337rl7eyt2nxasphf5m6n0rd"
	//export accountBalance="99994712164399999990848350"
	//export accountNonce="337529"
	//export accountPendingNonce="337530"
	//export accountNumActions="506609"
	//export actionHash="b9a938e1f249d3c7ab8152e377132989535e25ea9ee376323179d1943dc15b4e"
	//export actionActionInfoLen="1"
	//export actionActionNonce="1"
	//export getActionsByAddressActionHash="10efdacf68be2fa0afdc2a46786b3caf0a59fe2386485a9f075acf2a41c93d78"
	//export blk60801Hash="a331841e6b29becdeeb65cb0948d896076b514e5f8d69d560b4424282a7882d7"
	//export blk60801HashNumActions="1960"
	//export blk60801HashTransferAmount="2389000000000000000000"
	//export getServerMetaPackageCommitID="e24aadf90f98d800b9f117354ddd5b3dbe58dde9"
	//export accountAddressUnclaimedBalance="0"
	//export getReceiptByActionBlkHeight="517"
	//
	//export epochDataHeight="1"
	//export epochGravityChainStartHeight="7502300"
	//export readContractActionHash="63c74277bcbfcef195f57713131b05cb54c47461f7e64b8f32fb58f9b8445265"
	host = "api.iotex.one:80"
)

func TestServer_Deploy(t *testing.T) {
	require := require.New(t)
	sct, err := NewSmartContract("testdata/reentry-attack.json", host)
	require.NoError(err)
	err = sct.DeployContracts()
	require.NoError(err)
	contractAddress := sct.GetContractAddresses()
	for _, v := range contractAddress {
		fmt.Println(v)
	}
}
