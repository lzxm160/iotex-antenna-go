// Copyright (c) 2020 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package did

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/iotexproject/iotex-antenna-go/v2/did/abibin"

	//"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/iotexproject/go-pkgs/hash"

	"github.com/iotexproject/iotex-antenna-go/v2/iotex"

	"github.com/stretchr/testify/require"

	"github.com/iotexproject/iotex-antenna-go/v2/account"

	"github.com/iotexproject/iotex-antenna-go/v2/utils/unit"
	"github.com/iotexproject/iotex-proto/golang/iotexapi"
)

const (
	sender = "io1ph0u2psnd7muq5xv9623rmxdsxc4uapxhzpg02"
	//signsender            = "io1vdtfpzkwpyngzvx7u2mauepnzja7kd5rryp0sg"
	privateKey = "414efa99dfac6f4095d6954713fb0085268d400d6a05a8ae8a69b5b1c10b4bed"
	//signPrivateKey        = "0d4d9b248110257c575ef2e8d93dd53471d9178984482817dcbd6edb607f8cc5"
	endpoint = "api.testnet.iotex.one:443"
	//IoTeXDIDDB_address = "io1qswvysrfav83ehgsfsfa4v8ugw4vqfsydp8snk"
	IoTeXDID_address = "io1m3wjevwhz2s58sasq0wj4luxrnqt047s687zw8"
	//IoTeXDIDProxy_address = "io1kqwtdw7daf74l8sf0u9j3u9fgv225ua6hjn3n0"
)

var (
	gasPrice = big.NewInt(0).SetUint64(1e12)
	gasLimit = uint64(10000000)
	//stringTy, _  = abi.NewType("string")
	//bytes32Ty, _ = abi.NewType("bytes32")
)

func TestDidDeployContract(t *testing.T) {
	require := require.New(t)
	conn, err := iotex.NewDefaultGRPCConn(endpoint)
	require.NoError(err)
	defer conn.Close()

	acc, err := account.HexStringToAccount(privateKey)
	require.NoError(err)
	c := iotex.NewAuthedClient(iotexapi.NewAPIServiceClient(conn), acc)

	data, err := hex.DecodeString(abibin.AddressBasedDIDManagerBin[2:])
	require.NoError(err)
	abi, err := abi.JSON(strings.NewReader(abibin.AddressBasedDIDManagerABI))
	require.NoError(err)
	hash, err := c.DeployContract(data).SetArgs(abi, []byte("did:io:"), common.Address{0}).SetGasPrice(big.NewInt(
		int64(unit.
			Qev))).
		SetGasLimit(
			10000000).Call(context.Background())
	require.NoError(err)
	require.NotNil(hash)
	fmt.Println("hash", hex.EncodeToString(hash[:]))
	time.Sleep(20 * time.Second)
	receiptResponse, err := c.GetReceipt(hash).Call(context.Background())
	contractAddress := receiptResponse.GetReceiptInfo().GetReceipt().GetContractAddress()
	fmt.Println("Status:", receiptResponse.GetReceiptInfo().GetReceipt().Status)
	fmt.Println("Contract Address:", contractAddress)
}

func TestDidCreateDid(t *testing.T) {
	require := require.New(t)
	d, err := NewDID(endpoint, privateKey, IoTeXDID_address, abibin.AddressBasedDIDManagerABI, gasPrice, gasLimit)
	require.NoError(err)
	testDIDHash := "1111111111111111111111111111111111111111111111111111111111111112"
	h, err := d.RegisterDID(testDIDHash, []byte("uri1"))
	require.NoError(err)
	checkHash(h, t)
}

func TestDidGetHash(t *testing.T) {
	require := require.New(t)
	d, err := NewDID(endpoint, "", IoTeXDID_address, abibin.AddressBasedDIDManagerABI, nil, 0)
	require.NoError(err)
	acc, err := account.HexStringToAccount(privateKey)
	require.NoError(err)

	ethAddress := common.HexToAddress(hex.EncodeToString(acc.Address().Bytes()))

	didString := "did:io:" + ethAddress.String()
	fmt.Println(didString)
	h, err := d.GetHash(didString)
	require.NoError(err)
	fmt.Println(h)
}

func TestDidGetURI(t *testing.T) {
	require := require.New(t)
	d, err := NewDID(endpoint, "", IoTeXDID_address, abibin.AddressBasedDIDManagerABI, nil, 0)
	require.NoError(err)
	acc, err := account.HexStringToAccount(privateKey)
	require.NoError(err)

	ethAddress := common.HexToAddress(hex.EncodeToString(acc.Address().Bytes()))

	didString := "did:io:" + ethAddress.String()
	fmt.Println(didString)
	h, err := d.GetUri(didString)
	require.NoError(err)
	fmt.Println(h)
}

func TestDidUpdateUri(t *testing.T) {
	testuri := "https://gateway.pinata.cloud/ipfs/QmWUzz4EoBVGX9WJM3tsjyU4CzABReD8akrZqpc6f94Ba4/ddo-test.json"
	require := require.New(t)
	d, err := NewDID(endpoint, privateKey, IoTeXDID_address, abibin.AddressBasedDIDManagerABI, gasPrice, gasLimit)
	require.NoError(err)
	acc, err := account.HexStringToAccount(privateKey)
	require.NoError(err)

	ethAddress := common.HexToAddress(hex.EncodeToString(acc.Address().Bytes()))

	didString := "did:io:" + ethAddress.String()
	fmt.Println(didString)
	h, err := d.UpdateDID("22222222222222", testuri)
	require.NoError(err)
	checkHash(h, t)
}

func TestDeregister(t *testing.T) {
	require := require.New(t)
	d, err := NewDID(endpoint, privateKey, IoTeXDID_address, abibin.AddressBasedDIDManagerABI, gasPrice, gasLimit)
	require.NoError(err)

	h, err := d.DeregisterDID()
	require.NoError(err)
	checkHash(h, t)
}

func checkHash(h string, t *testing.T) {
	fmt.Println("check hash:", h)
	require := require.New(t)
	time.Sleep(20 * time.Second)
	conn, err := iotex.NewDefaultGRPCConn(endpoint)
	require.NoError(err)
	defer conn.Close()
	ha, err := hash.HexStringToHash256(h)
	require.NoError(err)
	c := iotex.NewReadOnlyClient(iotexapi.NewAPIServiceClient(conn))
	receiptResponse, err := c.GetReceipt(ha).Call(context.Background())
	s := receiptResponse.GetReceiptInfo().GetReceipt().GetStatus()
	fmt.Println("status:", s)
}
