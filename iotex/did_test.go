// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package iotex

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
	"github.com/stretchr/testify/require"

	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-antenna-go/v2/account"
	"github.com/iotexproject/iotex-antenna-go/v2/utils/unit"
	"github.com/iotexproject/iotex-proto/golang/iotexapi"
)

const (
	sender                = "io1ph0u2psnd7muq5xv9623rmxdsxc4uapxhzpg02"
	privateKey            = "414efa99dfac6f4095d6954713fb0085268d400d6a05a8ae8a69b5b1c10b4bed"
	endpoint              = "api.testnet.iotex.one:443"
	IoTeXDID_address      = "io1eurq3lx4lzx9wdj56plw5rm59f5qzanacr3raz"
	IoTeXDIDProxy_address = "io1zgs5gqjl679qlj4gqqpa9t329r8f5gr8xc9lr0"
)

func TestDidDeployContract(t *testing.T) {
	require := require.New(t)
	conn, err := NewDefaultGRPCConn(endpoint)
	require.NoError(err)
	defer conn.Close()

	acc, err := account.HexStringToAccount(privateKey)
	require.NoError(err)
	c := NewAuthedClient(iotexapi.NewAPIServiceClient(conn), acc)

	data, err := hex.DecodeString(IoTeXDID_bin)
	require.NoError(err)

	hash, err := c.DeployContract(data).SetGasPrice(big.NewInt(int64(unit.Qev))).SetGasLimit(10000000).Call(context.Background())
	require.NoError(err)
	require.NotNil(hash)

	time.Sleep(20 * time.Second)
	receiptResponse, err := c.GetReceipt(hash).Call(context.Background())
	contractAddress := receiptResponse.GetReceiptInfo().GetReceipt().GetContractAddress()
	fmt.Println("Contract Address:", contractAddress)
}

func TestDidProxyDeployContract(t *testing.T) {
	require := require.New(t)
	conn, err := NewDefaultGRPCConn(endpoint)
	require.NoError(err)
	defer conn.Close()

	acc, err := account.HexStringToAccount(privateKey)
	require.NoError(err)
	c := NewAuthedClient(iotexapi.NewAPIServiceClient(conn), acc)

	abi, err := abi.JSON(strings.NewReader(IoTeXDIDProxy_abi))
	require.NoError(err)
	data, err := hex.DecodeString(IoTeXDIDProxy_bin)
	require.NoError(err)
	DIDcontract, err := address.FromString(IoTeXDID_address)
	require.NoError(err)

	hash, err := c.DeployContract(data).SetGasPrice(big.NewInt(int64(unit.Qev))).SetGasLimit(10000000).SetArgs(abi, DIDcontract).Call(context.Background())
	require.NoError(err)
	require.NotNil(hash)

	time.Sleep(20 * time.Second)
	receiptResponse, err := c.GetReceipt(hash).Call(context.Background())
	contractAddress := receiptResponse.GetReceiptInfo().GetReceipt().GetContractAddress()
	fmt.Println("Contract Address:", contractAddress)
}

func TestDidProxyDeployUpdateAddress(t *testing.T) {
	require := require.New(t)
	conn, err := NewDefaultGRPCConn(endpoint)
	require.NoError(err)
	defer conn.Close()

	acc, err := account.HexStringToAccount(privateKey)
	require.NoError(err)
	c := NewAuthedClient(iotexapi.NewAPIServiceClient(conn), acc)
	abi, err := abi.JSON(strings.NewReader(IoTeXDIDProxy_abi))
	require.NoError(err)
	contract, err := address.FromString(IoTeXDIDProxy_address)
	require.NoError(err)
	IoTeXDIDAddress, err := address.FromString(IoTeXDID_address)
	require.NoError(err)
	ethAddress := common.HexToAddress(hex.EncodeToString(IoTeXDIDAddress.Bytes()))
	hash, err := c.Contract(contract, abi).Execute("upgradeTo", "1.1.1", ethAddress).SetGasPrice(big.NewInt(int64(unit.Qev))).SetGasLimit(1000000).Call(context.Background())
	require.NoError(err)
	require.NotNil(hash)
	fmt.Println(hex.EncodeToString(hash[:]))

	time.Sleep(20 * time.Second)
	receiptResponse, err := c.GetReceipt(hash).Call(context.Background())
	s := receiptResponse.GetReceiptInfo().GetReceipt().GetStatus()
	fmt.Println("status:", s)
}

func TestDidProxyReadVersion(t *testing.T) {
	require := require.New(t)
	conn, err := NewDefaultGRPCConn(endpoint)
	require.NoError(err)
	defer conn.Close()

	c := NewReadOnlyClient(iotexapi.NewAPIServiceClient(conn))

	abi, err := abi.JSON(strings.NewReader(IoTeXDIDProxy_abi))
	require.NoError(err)
	contract, err := address.FromString(IoTeXDIDProxy_address)
	require.NoError(err)

	version := "0.0.1"
	//version := "1.1.1"
	ret, err := c.ReadOnlyContract(contract, abi).Read("getImplFromVersion", version).Call(context.Background())
	require.NoError(err)
	fmt.Println(ret.method)
	fmt.Println(hex.EncodeToString(ret.Raw))
	var addr common.Address
	require.NoError(ret.Unmarshal(&addr))
	fmt.Println(addr.String())

	iotexAddr, err := address.FromBytes(addr.Bytes())
	require.NoError(err)
	fmt.Println(iotexAddr.String())
}

func TestDidProxyReadLatestVersion(t *testing.T) {
	require := require.New(t)
	conn, err := NewDefaultGRPCConn(endpoint)
	require.NoError(err)
	defer conn.Close()

	c := NewReadOnlyClient(iotexapi.NewAPIServiceClient(conn))

	abi, err := abi.JSON(strings.NewReader(IoTeXDIDProxy_abi))
	require.NoError(err)
	contract, err := address.FromString(IoTeXDIDProxy_address)
	require.NoError(err)

	ret, err := c.ReadOnlyContract(contract, abi).Read("implementation").Call(context.Background())
	require.NoError(err)
	fmt.Println(ret.method)
	fmt.Println(hex.EncodeToString(ret.Raw))
	var addr common.Address
	require.NoError(ret.Unmarshal(&addr))
	fmt.Println(addr.String())

	iotexAddr, err := address.FromBytes(addr.Bytes())
	require.NoError(err)
	fmt.Println(iotexAddr.String())
}

func TestDidCreateDid(t *testing.T) {
	require := require.New(t)
	conn, err := NewDefaultGRPCConn(endpoint)
	require.NoError(err)
	defer conn.Close()

	acc, err := account.HexStringToAccount(privateKey)
	require.NoError(err)
	c := NewAuthedClient(iotexapi.NewAPIServiceClient(conn), acc)
	abi, err := abi.JSON(strings.NewReader(IoTeXDID_abi)) // note,this is IoTeXDID_abi
	require.NoError(err)
	contract, err := address.FromString(IoTeXDIDProxy_address)
	require.NoError(err)
	hash1 := [32]uint8{1}
	hash, err := c.Contract(contract, abi).Execute("createDID", "", hash1, "uri1").SetGasPrice(big.NewInt(int64(unit.Qev))).SetGasLimit(1000000).Call(context.Background())
	require.NoError(err)
	require.NotNil(hash)
	fmt.Println(hex.EncodeToString(hash[:]))

	time.Sleep(20 * time.Second)
	receiptResponse, err := c.GetReceipt(hash).Call(context.Background())
	s := receiptResponse.GetReceiptInfo().GetReceipt().GetStatus()
	fmt.Println("status:", s)
}

func TestDidDelete(t *testing.T) {
	require := require.New(t)
	conn, err := NewDefaultGRPCConn(endpoint)
	require.NoError(err)
	defer conn.Close()

	acc, err := account.HexStringToAccount(privateKey)
	require.NoError(err)
	c := NewAuthedClient(iotexapi.NewAPIServiceClient(conn), acc)
	abi, err := abi.JSON(strings.NewReader(IoTeXDID_abi)) // note,this is IoTeXDID_abi
	require.NoError(err)
	contract, err := address.FromString(IoTeXDIDProxy_address)
	require.NoError(err)

	didAddress, err := address.FromString(sender)
	require.NoError(err)
	ethAddress := common.HexToAddress(hex.EncodeToString(didAddress.Bytes()))
	didString := "did:io:" + ethAddress.String()
	fmt.Println(didString)

	hash, err := c.Contract(contract, abi).Execute("deleteDID", didString).SetGasPrice(big.NewInt(int64(unit.Qev))).SetGasLimit(1000000).Call(context.Background())
	require.NoError(err)
	require.NotNil(hash)
	fmt.Println(hex.EncodeToString(hash[:]))

	time.Sleep(20 * time.Second)
	receiptResponse, err := c.GetReceipt(hash).Call(context.Background())
	s := receiptResponse.GetReceiptInfo().GetReceipt().GetStatus()
	fmt.Println("status:", s)
}

func TestDidUpdateHash(t *testing.T) {
	require := require.New(t)
	conn, err := NewDefaultGRPCConn(endpoint)
	require.NoError(err)
	defer conn.Close()

	acc, err := account.HexStringToAccount(privateKey)
	require.NoError(err)
	c := NewAuthedClient(iotexapi.NewAPIServiceClient(conn), acc)
	abi, err := abi.JSON(strings.NewReader(IoTeXDID_abi)) // note,this is IoTeXDID_abi
	require.NoError(err)
	contract, err := address.FromString(IoTeXDIDProxy_address)
	require.NoError(err)
	hash1 := [32]uint8{111}

	didAddress, err := address.FromString(sender)
	require.NoError(err)
	ethAddress := common.HexToAddress(hex.EncodeToString(didAddress.Bytes()))
	didString := "did:io:" + ethAddress.String()
	fmt.Println(didString)

	hash, err := c.Contract(contract, abi).Execute("updateHash", didString, hash1).SetGasPrice(big.NewInt(int64(unit.Qev))).SetGasLimit(1000000).Call(context.Background())
	require.NoError(err)
	require.NotNil(hash)
	fmt.Println(hex.EncodeToString(hash[:]))

	time.Sleep(20 * time.Second)
	receiptResponse, err := c.GetReceipt(hash).Call(context.Background())
	s := receiptResponse.GetReceiptInfo().GetReceipt().GetStatus()
	fmt.Println("status:", s)
}

func TestDidUpdateUri(t *testing.T) {
	require := require.New(t)
	conn, err := NewDefaultGRPCConn(endpoint)
	require.NoError(err)
	defer conn.Close()

	acc, err := account.HexStringToAccount(privateKey)
	require.NoError(err)
	c := NewAuthedClient(iotexapi.NewAPIServiceClient(conn), acc)
	abi, err := abi.JSON(strings.NewReader(IoTeXDID_abi)) // note,this is IoTeXDID_abi
	require.NoError(err)
	contract, err := address.FromString(IoTeXDIDProxy_address)
	require.NoError(err)

	didAddress, err := address.FromString(sender)
	require.NoError(err)
	ethAddress := common.HexToAddress(hex.EncodeToString(didAddress.Bytes()))
	didString := "did:io:" + ethAddress.String()
	fmt.Println(didString)

	hash, err := c.Contract(contract, abi).Execute("updateURI", didString, "urixxxx").SetGasPrice(big.NewInt(int64(unit.Qev))).SetGasLimit(1000000).Call(context.Background())
	require.NoError(err)
	require.NotNil(hash)
	fmt.Println(hex.EncodeToString(hash[:]))

	time.Sleep(20 * time.Second)
	receiptResponse, err := c.GetReceipt(hash).Call(context.Background())
	s := receiptResponse.GetReceiptInfo().GetReceipt().GetStatus()
	fmt.Println("status:", s)
}

func TestDidReadHashContract(t *testing.T) {
	require := require.New(t)
	conn, err := NewDefaultGRPCConn(endpoint)
	require.NoError(err)
	defer conn.Close()

	c := NewReadOnlyClient(iotexapi.NewAPIServiceClient(conn))

	abi, err := abi.JSON(strings.NewReader(IoTeXDID_abi))
	require.NoError(err)
	contract, err := address.FromString(IoTeXDIDProxy_address)
	require.NoError(err)
	didAddress, err := address.FromString(sender)
	require.NoError(err)
	ethAddress := common.HexToAddress(hex.EncodeToString(didAddress.Bytes()))

	didString := "did:io:" + ethAddress.String()
	fmt.Println(didString)
	ret, err := c.ReadOnlyContract(contract, abi).Read("getHash", didString).Call(context.Background())
	require.NoError(err)
	fmt.Println(hex.EncodeToString(ret.Raw))
	hash1 := [32]uint8{}
	require.NoError(ret.Unmarshal(&hash1))
	fmt.Println(hex.EncodeToString(hash1[:]))
}

func TestDidReadUriContract(t *testing.T) {
	require := require.New(t)
	conn, err := NewDefaultGRPCConn(endpoint)
	require.NoError(err)
	defer conn.Close()

	c := NewReadOnlyClient(iotexapi.NewAPIServiceClient(conn))

	abi, err := abi.JSON(strings.NewReader(IoTeXDID_abi))
	require.NoError(err)
	contract, err := address.FromString(IoTeXDIDProxy_address)
	require.NoError(err)
	didAddress, err := address.FromString(sender)
	require.NoError(err)
	ethAddress := common.HexToAddress(hex.EncodeToString(didAddress.Bytes()))

	didString := "did:io:" + ethAddress.String()
	fmt.Println(didString)
	ret, err := c.ReadOnlyContract(contract, abi).Read("getURI", didString).Call(context.Background())
	require.NoError(err)
	fmt.Println(hex.EncodeToString(ret.Raw))
	var uri string
	require.NoError(ret.Unmarshal(&uri))
	fmt.Println(uri)
}
