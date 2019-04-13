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
	//newPriKey := "8c379a71721322d16912a88b1602c5596ca9e99a5f70777561c3029efa71a435"
	//priKey, err := keypair.HexStringToPrivateKey(newPriKey)
	//require.NoError(err)
	//addr, err := address.FromBytes(priKey.PublicKey().Hash())
	//require.NoError(err)
	//fmt.Println(addr.String())
	bin := "608060405234801561001057600080fd5b50600080546001818101835582805260647f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563928301558254808201845560c890830155825490810190925561012c910155610125806100706000396000f300608060405260043610603e5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663c298557881146043575b600080fd5b348015604e57600080fd5b50605560a3565b60408051602080825283518183015283519192839290830191858101910280838360005b83811015608f5781810151838201526020016079565b505050509050019250505060405180910390f35b6060600080548060200260200160405190810160405280929190818152602001828054801560ef57602002820191906000526020600020905b81548152602001906001019080831160dc575b50505050509050905600a165627a7a723058202b92ee77c566f1cb20c90f2216e68746355bac330751bbfe5ff7799b10e2e0440029"
	gasLimit := uint64(1000000)
	gasPrice := big.NewInt(1000000000000)
	sct, err := NewContract(host, bin, gasLimit, gasPrice)
	require.NoError(err)
	accountPrivateKey := "8c379a71721322d16912a88b1602c5596ca9e99a5f70777561c3029efa71a435"
	accountAddress := "io1ns7y0pxmklk8ceattty6n7makpw76u770u5avy"
	sct.SetExecutor(accountAddress, accountPrivateKey)
	hash, err := sct.Deploy()
	require.NoError(err)
	receipt, err := sct.CheckCallResult(hash)
	require.NoError(err)
	fmt.Println(receipt.ContractAddress)
	sct.SetContractAddress(receipt.ContractAddress)
	ret, err := sct.CallMethod("febb0f7e")
	require.NoError(err)
	fmt.Println(ret)
}
