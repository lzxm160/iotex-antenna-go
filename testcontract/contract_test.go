// Copyright (c) 2020 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package testcontract

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/iotexproject/iotex-address/address"

	"github.com/iotexproject/iotex-antenna-go/v2/account"
	"github.com/iotexproject/iotex-antenna-go/v2/iotex"
	"github.com/iotexproject/iotex-proto/golang/iotexapi"

	"github.com/stretchr/testify/require"
)

const (
	sender         = "io1ph0u2psnd7muq5xv9623rmxdsxc4uapxhzpg02"
	signsender     = "io1vdtfpzkwpyngzvx7u2mauepnzja7kd5rryp0sg"
	privateKey     = "414efa99dfac6f4095d6954713fb0085268d400d6a05a8ae8a69b5b1c10b4bed"
	signPrivateKey = "0d4d9b248110257c575ef2e8d93dd53471d9178984482817dcbd6edb607f8cc5"
	endpoint       = "api.testnet.iotex.one:443"
	contract       = "io108u2u0e3x2kt2j0q0k5tpekvvudk6ywqfuxp37"
)

var (
	gasPrice = big.NewInt(0).SetUint64(1e12)
	gasLimit = uint64(10000000)
)

func TestDidCreateDid(t *testing.T) {
	require := require.New(t)
	conn, err := iotex.NewDefaultGRPCConn(endpoint)
	if err != nil {
		return
	}
	defer conn.Close()
	var acc account.Account
	acc, err = account.HexStringToAccount(privateKey)
	require.NoError(err)
	addr, err := address.FromString(contract)
	require.NoError(err)
	cli := iotex.NewAuthedClient(iotexapi.NewAPIServiceClient(conn), acc)
	//recipients []common.Address, amounts []*big.Int
	recipients := []common.Address{
		common.HexToAddress("f959d06fB374ABE5Fc6166B2B4c42e03E0F84E00"),
		common.HexToAddress("Fec1FB152c43c0460e6A153f4aAbE29E933cd74B"),
	}
	//io1l9vaqmanwj47tlrpv6etf3pwq0s0snsq4vxke2 - 0xf959d06fB374ABE5Fc6166B2B4c42e03E0F84E00
	//contractAddr := common.BytesToAddress(contract.Bytes())
	//	io1lmqlk9fvg0qyvrn2z5l542lzn6fne46tza8a3j - 0xFec1FB152c43c0460e6A153f4aAbE29E933cd74B
	amounts := []*big.Int{big.NewInt(11111111), big.NewInt(22222222)}
	h, err := cli.Contract(addr, MultisendABI).Execute("multiSend", recipients, amounts, "").SetGasPrice(gasPrice).SetGasLimit(gasLimit).
		Call(context.Background())
	require.NoError(err)
	fmt.Println(hex.EncodeToString(h[:]))
}

//
//func TestDidDeployContract(t *testing.T) {
//	require := require.New(t)
//	conn, err := iotex.NewDefaultGRPCConn(endpoint)
//	require.NoError(err)
//	defer conn.Close()
//
//	acc, err := account.HexStringToAccount(privateKey)
//	require.NoError(err)
//	c := iotex.NewAuthedClient(iotexapi.NewAPIServiceClient(conn), acc)
//
//	data, err := hex.DecodeString(IoTeXDID.IoTeXDIDBin[2:])
//	require.NoError(err)
//
//	hash, err := c.DeployContract(data).SetGasPrice(big.NewInt(int64(unit.Qev))).SetGasLimit(10000000).Call(context.Background())
//	require.NoError(err)
//	require.NotNil(hash)
//
//	time.Sleep(20 * time.Second)
//	receiptResponse, err := c.GetReceipt(hash).Call(context.Background())
//	contractAddress := receiptResponse.GetReceiptInfo().GetReceipt().GetContractAddress()
//	fmt.Println("Contract Address:", contractAddress)
//}
