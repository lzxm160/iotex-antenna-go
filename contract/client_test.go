// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package contract

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	host = "api.iotex.one:80"
)

func TestServer_Deploy(t *testing.T) {
	require := require.New(t)
	bin := "608060405234801561001057600080fd5b50610177806100206000396000f30060806040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630423a1328114610050578063bfe43b4c1461007a575b600080fd5b34801561005c57600080fd5b50610068600435610148565b60408051918252519081900360200190f35b34801561008657600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100d39436949293602493928401919081908401838280828437509497506101489650505050505050565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561010d5781810151838201526020016100f5565b50505050905090810190601f16801561013a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b905600a165627a7a72305820528569d402fc6b27038b283a0d119270ba29d17ef484052f3a3021c33db1e06a0029"
	abi := `[{"constant":true,"inputs":[{"name":"x","type":"uint256"}],"name":"bar","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"y","type":"string"}],"name":"barstring","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"inputs":[],"payable":false,"stateMutability":"nonpayable","type":"constructor"}]`

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
	sct.SetContractAddress(receipt.ContractAddress)
	sct.SetExecutor(accountAddress, accountPrivateKey)

	ret, err := sct.CallMethod("bar", big.NewInt(10))
	require.NoError(err)
	require.Equal("*big.Int", reflect.TypeOf(ret).String())
	require.Equal(0, ret.(*big.Int).Cmp(big.NewInt(10)))

	ret, err = sct.CallMethod("barstring", "foobar")
	require.NoError(err)
	require.Equal("string", reflect.TypeOf(ret).String())
	require.Equal("foobar", ret.(string))
}
