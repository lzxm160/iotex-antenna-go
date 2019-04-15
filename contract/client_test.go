// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package contract

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	host = "api.iotex.one:80"
)

func TestServer_Deploy(t *testing.T) {
	require := require.New(t)
	//priKey, err := keypair.HexStringToPrivateKey("4e1a9091301ca53ae505292911bf4c39f720288724226efc654f2593b56085bd")
	//require.NoError(err)
	//addr, err := address.FromBytes(priKey.PublicKey().Hash())
	//require.NoError(err)
	//fmt.Println(addr.String())
	//var evmContractAddrHash common.Address
	//copy(evmContractAddrHash[:], addr.Bytes())
	//fmt.Println(evmContractAddrHash.String())
	bin := "6080604052348015600f57600080fd5b5060998061001e6000396000f300608060405260043610603e5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630423a13281146043575b600080fd5b348015604e57600080fd5b506058600435606a565b60408051918252519081900360200190f35b905600a165627a7a72305820fdc290d7b4e9de2ed3b68502db42e0755a574973a05ab70305662785f3c621aa0029"
	abi := `[{"constant":true,"inputs":[{"name":"x","type":"uint256"}],"name":"bar","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[],"payable":false,"stateMutability":"nonpayable","type":"constructor"}]`

	gasLimit := uint64(1000000)
	gasPrice := big.NewInt(9000000000000)
	sct, err := NewContract(host, bin, abi, gasLimit, gasPrice)
	require.NoError(err)
	accountPrivateKey := "8c379a71721322d16912a88b1602c5596ca9e99a5f70777561c3029efa71a435"
	accountAddress := "io1ns7y0pxmklk8ceattty6n7makpw76u770u5avy"
	sct.SetOwner(accountAddress, accountPrivateKey)

	hash, err := sct.Deploy()
	require.NoError(err)
	receipt, err := sct.CheckCallResult(hash)
	require.NoError(err)
	fmt.Println("receipt contract:", receipt.ContractAddress)
	sct.SetContractAddress(receipt.ContractAddress)
	fmt.Println("contract:", sct.ContractAddress())

	sct.SetExecutor(accountAddress, accountPrivateKey)
	ret, err := sct.CallMethod("bar")
	require.NoError(err)
	fmt.Println(ret)
}
