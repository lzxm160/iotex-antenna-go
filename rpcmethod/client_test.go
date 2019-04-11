// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package rpcmethod

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/iotexproject/iotex-core/protogen/iotexapi"
	ta "github.com/iotexproject/iotex-core/test/testaddress"
	"github.com/iotexproject/iotex-core/testutil"
)

const (
	host           = "127.0.0.1:14014"
	testAddress    = "io187wzp08vnhjjpkydnr97qlh8kh0dpkkytfam8j"
	testPrivateKey = "0806c458b262edd333a191e92f561aff338211ee3e18ab315a074a2d82aa343f"
	testPublicKey  = "044e18306ae9ef4ec9d07bf6e705442d4d1a75e6cdf750330ca2d880f2cc54607c9c33deb9eae9c06e06e04fe9ce3d43962cc67d5aa34fbeb71270d4bad3d648d9"
)

func TestGetAccount(t *testing.T) {
	require := require.New(t)
	rpc, err := NewRPCMethod(host)
	require.NoError(err)
	req := iotexapi.GetAccountRequest{}
	req.Address = testAddress
	res, err := rpc.GetAccount(&req)
	require.NoError(err)
	fmt.Println(res)
}
func TestGetChainMeta(t *testing.T) {
	require := require.New(t)
	rpc, err := NewRPCMethod(host)
	require.NoError(err)
	req := iotexapi.GetChainMetaRequest{}
	res, err := rpc.GetChainMeta(&req)
	require.NoError(err)
	fmt.Println(res)
}
func TestGetActions(t *testing.T) {
	require := require.New(t)
	rpc, err := NewRPCMethod(host)
	require.NoError(err)
	req := &iotexapi.GetActionsRequest{
		Lookup: &iotexapi.GetActionsRequest_ByIndex{
			ByIndex: &iotexapi.GetActionsByIndexRequest{
				Start: 0,
				Count: 3,
			},
		},
	}
	res, err := rpc.GetActions(req)
	require.NoError(err)
	fmt.Println(res)
}
func TestGetBlockMetas(t *testing.T) {
	require := require.New(t)
	rpc, err := NewRPCMethod(host)
	require.NoError(err)
	request := &iotexapi.GetBlockMetasRequest{
		Lookup: &iotexapi.GetBlockMetasRequest_ByIndex{
			ByIndex: &iotexapi.GetBlockMetasByIndexRequest{
				Start: 0,
				Count: 3,
			},
		},
	}
	res, err := rpc.GetBlockMetas(request)
	require.NoError(err)
	fmt.Println(res)
}
func TestSendAction(t *testing.T) {
	require := require.New(t)
	rpc, err := NewRPCMethod(host)
	require.NoError(err)
	testTransfer, err := testutil.SignedTransfer(ta.Addrinfo["alfa"].String(),
		ta.Keyinfo["alfa"].PriKey, 3, big.NewInt(10), []byte{}, testutil.TestGasLimit,
		big.NewInt(testutil.TestGasPriceInt64))
	require.NoError(err)
	testTransferPb := testTransfer.Proto()
	request := &iotexapi.SendActionRequest{Action: testTransferPb}

	res, err := rpc.SendAction(request)
	require.NoError(err)
	fmt.Println(res)
}
