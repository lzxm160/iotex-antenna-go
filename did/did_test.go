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

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/iotexproject/iotex-address/address"

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
	IoTeXDID_address = "io1pq4zg4mf2wheah33e9d4yq8fujk7k83s3crzxz"
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
	DIDcontract, err := address.FromString(address.ZeroAddress)
	require.NoError(err)

	hash, err := c.DeployContract(data).SetArgs(abi, []byte("did:io"), DIDcontract).SetGasPrice(big.NewInt(int64(unit.
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
	h, err := d.CreateDID("", testDIDHash, []byte("uri1"))
	require.NoError(err)
	checkHash(h, t)
}

//func TestDidProxyDeployContract(t *testing.T) {
//	require := require.New(t)
//	conn, err := iotex.NewDefaultGRPCConn(endpoint)
//	require.NoError(err)
//	defer conn.Close()
//
//	acc, err := account.HexStringToAccount(privateKey)
//	require.NoError(err)
//	c := iotex.NewAuthedClient(iotexapi.NewAPIServiceClient(conn), acc)
//
//	abi, err := abi.JSON(strings.NewReader(IoTeXDIDProxy.IoTeXDIDProxyABI))
//	require.NoError(err)
//	data, err := hex.DecodeString(IoTeXDIDProxy.IoTeXDIDProxyBin[2:])
//	require.NoError(err)
//	DIDcontract, err := address.FromString(IoTeXDID_address)
//	require.NoError(err)
//
//	hash, err := c.DeployContract(data).SetGasPrice(big.NewInt(int64(unit.Qev))).SetGasLimit(10000000).SetArgs(abi, DIDcontract).Call(context.Background())
//	require.NoError(err)
//	require.NotNil(hash)
//
//	time.Sleep(20 * time.Second)
//	receiptResponse, err := c.GetReceipt(hash).Call(context.Background())
//	contractAddress := receiptResponse.GetReceiptInfo().GetReceipt().GetContractAddress()
//	fmt.Println("Contract Address:", contractAddress)
//}
//
//func TestDidProxyDeployUpdateAddress(t *testing.T) {
//	require := require.New(t)
//	conn, err := iotex.NewDefaultGRPCConn(endpoint)
//	require.NoError(err)
//	defer conn.Close()
//
//	acc, err := account.HexStringToAccount(privateKey)
//	require.NoError(err)
//	c := iotex.NewAuthedClient(iotexapi.NewAPIServiceClient(conn), acc)
//	abi, err := abi.JSON(strings.NewReader(IoTeXDIDProxy.IoTeXDIDProxyABI))
//	require.NoError(err)
//	contract, err := address.FromString(IoTeXDIDProxy_address)
//	require.NoError(err)
//	IoTeXDIDAddress, err := address.FromString(IoTeXDID_address)
//	require.NoError(err)
//	ethAddress := common.HexToAddress(hex.EncodeToString(IoTeXDIDAddress.Bytes()))
//	hash, err := c.Contract(contract, abi).Execute("upgradeTo", "1.1.1", ethAddress).SetGasPrice(big.NewInt(int64(unit.Qev))).SetGasLimit(1000000).Call(context.Background())
//	require.NoError(err)
//	require.NotNil(hash)
//	fmt.Println(hex.EncodeToString(hash[:]))
//
//	time.Sleep(20 * time.Second)
//	receiptResponse, err := c.GetReceipt(hash).Call(context.Background())
//	s := receiptResponse.GetReceiptInfo().GetReceipt().GetStatus()
//	fmt.Println("status:", s)
//}
//
//func TestDidProxyReadVersion(t *testing.T) {
//	require := require.New(t)
//	conn, err := iotex.NewDefaultGRPCConn(endpoint)
//	require.NoError(err)
//	defer conn.Close()
//
//	c := iotex.NewReadOnlyClient(iotexapi.NewAPIServiceClient(conn))
//
//	abi, err := abi.JSON(strings.NewReader(IoTeXDIDProxy.IoTeXDIDProxyABI))
//	require.NoError(err)
//	contract, err := address.FromString(IoTeXDIDProxy_address)
//	require.NoError(err)
//
//	version := "0.0.1"
//	//version := "1.1.1"
//	ret, err := c.ReadOnlyContract(contract, abi).Read("getImplFromVersion", version).Call(context.Background())
//	require.NoError(err)
//	fmt.Println(hex.EncodeToString(ret.Raw))
//	var addr common.Address
//	require.NoError(ret.Unmarshal(&addr))
//	fmt.Println(addr.String())
//
//	iotexAddr, err := address.FromBytes(addr.Bytes())
//	require.NoError(err)
//	fmt.Println(iotexAddr.String())
//}
//
//func TestDidProxyReadLatestVersion(t *testing.T) {
//	require := require.New(t)
//	conn, err := iotex.NewDefaultGRPCConn(endpoint)
//	require.NoError(err)
//	defer conn.Close()
//
//	c := iotex.NewReadOnlyClient(iotexapi.NewAPIServiceClient(conn))
//
//	abi, err := abi.JSON(strings.NewReader(IoTeXDIDProxy.IoTeXDIDProxyABI))
//	require.NoError(err)
//	contract, err := address.FromString(IoTeXDIDProxy_address)
//	require.NoError(err)
//
//	ret, err := c.ReadOnlyContract(contract, abi).Read("implementation").Call(context.Background())
//	require.NoError(err)
//	fmt.Println(hex.EncodeToString(ret.Raw))
//	var addr common.Address
//	require.NoError(ret.Unmarshal(&addr))
//	fmt.Println(addr.String())
//
//	iotexAddr, err := address.FromBytes(addr.Bytes())
//	require.NoError(err)
//	fmt.Println(iotexAddr.String())
//}

//func TestDidCreateDidSigned(t *testing.T) {
//	require := require.New(t)
//	d, err := NewDID(endpoint, privateKey, IoTeXDIDProxy_address, IoTeXDID.IoTeXDIDABI, gasPrice, gasLimit)
//	require.NoError(err)
//	//acc, err := account.HexStringToAccount(signPrivateKey)
//	//require.NoError(err)
//	testDIDHash := "1111111111111111111111111111111111111111111111111111111111111112"
//	hash, err := hex.DecodeString(testDIDHash)
//	require.NoError(err)
//	var hash32 [32]byte
//	copy(hash32[:], hash)
//	fmt.Println("hash32", hex.EncodeToString(hash32[:]))
//	signed, err := encodeParamsCreate("", "createDID", hash32, "urisigned")
//	require.NoError(err)
//	fmt.Println(hex.EncodeToString(signed))
//	var v uint8
//	v = signed[64:][0]
//	fmt.Println(v)
//	var r [32]byte
//	copy(r[:], signed[:32])
//	fmt.Println(hex.EncodeToString(r[:]))
//	var s [32]byte
//	copy(s[:], signed[32:64])
//	fmt.Println(hex.EncodeToString(s[:]))
//	h, err := d.CreateDIDSigned("", v+27, r, s, testDIDHash, "urisigned")
//	require.NoError(err)
//	checkHash(h, t)
//}
//
//func TestDidDeleteSigned(t *testing.T) {
//	require := require.New(t)
//	d, err := NewDID(endpoint, privateKey, IoTeXDIDProxy_address, IoTeXDID.IoTeXDIDABI, gasPrice, gasLimit)
//	require.NoError(err)
//	acc, err := account.HexStringToAccount(signPrivateKey)
//	require.NoError(err)
//	ethAddress := common.HexToAddress(hex.EncodeToString(acc.Address().Bytes()))
//
//	didString := "did:io:" + ethAddress.String()
//	fmt.Println(didString)
//	//dataB := encodeParamsDelete(didString, "deleteDID")
//	dataB := []byte(didString)
//	require.NoError(err)
//	fmt.Println("data", hex.EncodeToString(dataB))
//	signed, err := acc.Sign(dataB)
//	require.NoError(err)
//	fmt.Println(hex.EncodeToString(signed))
//	var v uint8
//	v = signed[64:][0]
//	fmt.Println(v)
//	var r [32]byte
//	copy(r[:], signed[:32])
//	fmt.Println(hex.EncodeToString(r[:]))
//	var s [32]byte
//	copy(s[:], signed[32:64])
//	fmt.Println(hex.EncodeToString(s[:]))
//	h, err := d.DeleteDIDSigned(didString, v, r, s)
//	require.NoError(err)
//	checkHash(h, t)
//}
//
//func TestDidUpdateHash(t *testing.T) {
//}
//
//func TestDidUpdateUri(t *testing.T) {
//	testuri := "https://gateway.pinata.cloud/ipfs/QmWUzz4EoBVGX9WJM3tsjyU4CzABReD8akrZqpc6f94Ba4/ddo-test.json"
//	require := require.New(t)
//	d, err := NewDID(endpoint, privateKey, "io1j2af3s4f7qjk8eudzx6a6kdhekr7zt2k5y5qlk", IoTeXDID.IoTeXDIDABI, gasPrice, gasLimit)
//	require.NoError(err)
//	acc, err := account.HexStringToAccount(privateKey)
//	require.NoError(err)
//
//	ethAddress := common.HexToAddress(hex.EncodeToString(acc.Address().Bytes()))
//
//	didString := "did:io:" + ethAddress.String()
//	fmt.Println(didString)
//	h, err := d.UpdateUri(didString, testuri)
//	require.NoError(err)
//	checkHash(h, t)
//}
//
//func TestDidGetHash(t *testing.T) {
//	require := require.New(t)
//	d, err := NewDID(endpoint, "", "io1j2af3s4f7qjk8eudzx6a6kdhekr7zt2k5y5qlk", IoTeXDID.IoTeXDIDABI, nil, 0)
//	require.NoError(err)
//	acc, err := account.HexStringToAccount(privateKey)
//	require.NoError(err)
//
//	ethAddress := common.HexToAddress(hex.EncodeToString(acc.Address().Bytes()))
//
//	didString := "did:io:" + ethAddress.String()
//	fmt.Println(didString)
//	h, err := d.GetHash(didString)
//	require.NoError(err)
//	fmt.Println(h)
//}
//
//func TestDidGetURI(t *testing.T) {
//	require := require.New(t)
//	d, err := NewDID(endpoint, "", "io1j2af3s4f7qjk8eudzx6a6kdhekr7zt2k5y5qlk", IoTeXDID.IoTeXDIDABI, nil, 0)
//	require.NoError(err)
//	acc, err := account.HexStringToAccount(privateKey)
//	require.NoError(err)
//
//	ethAddress := common.HexToAddress(hex.EncodeToString(acc.Address().Bytes()))
//
//	didString := "did:io:" + ethAddress.String()
//	fmt.Println(didString)
//	h, err := d.GetUri(didString)
//	require.NoError(err)
//	fmt.Println(h)
//}

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

//func TestGetLog(t *testing.T) {
//	require := require.New(t)
//	conn, err := iotex.NewDefaultGRPCConn(endpoint)
//	require.NoError(err)
//	defer conn.Close()
//	chainClient := iotexapi.NewAPIServiceClient(conn)
//	require.NoError(err)
//	response, err := chainClient.GetLogs(context.Background(), &iotexapi.GetLogsRequest{
//		Filter: &iotexapi.LogsFilter{
//			Address: []string{IoTeXDIDProxy_address},
//			Topics:  nil,
//		},
//		Lookup: &iotexapi.GetLogsRequest_ByRange{
//			ByRange: &iotexapi.GetLogsByRange{
//				//FromBlock: 3814660,
//				FromBlock: 3814818,
//				Count:     1,
//			},
//		},
//	})
//	if err != nil {
//		return
//	}
//	for _, l := range response.Logs {
//		fmt.Println(hex.EncodeToString(l.Data))
//		for _, topic := range l.Topics {
//			fmt.Println(hex.EncodeToString(topic))
//		}
//	}
//}

//func encodeParamsCreate(id, operate string, h [32]byte, uri string) ([]byte, error) {
//	//hash := crypto.Keccak256Hash(common.LeftPadBytes([]byte(id), 32),
//	//	common.LeftPadBytes([]byte(operate), 32), h[:],
//	//	[]byte(uri),
//	//)
//
//	// normally we sign prefixed hash
//	// as in solidity with `ECDSA.toEthSignedMessageHash`
//
//	//prefixedHash := crypto.Keccak256Hash(
//	//	[]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%v", len(hash))),
//	//	hash.Bytes(),
//	//)
//
//	// sign hash to validate later in Solidity
//	//pri, err := crypto.HexToECDSA(signPrivateKey)
//	//if err != nil {
//	//	return nil, err
//	//}
//	//sig, err := crypto.Sign(hash.Bytes(), pri)
//	//return sig, err
//	//arguments := abi.Arguments{
//	//	{
//	//		Type: stringTy,
//	//	},
//	//	{
//	//		Type: stringTy,
//	//	},
//	//	{
//	//		Type: bytes32Ty,
//	//	},
//	//	{
//	//		Type: stringTy,
//	//	},
//	//}
//	//
//	//bytes, _ := arguments.Pack(
//	//	id,
//	//	operate,
//	//	h,
//	//	uri,
//	//)
//	var bytes []byte
//	bytes = append(bytes, []byte(id)...)
//	bytes = append(bytes, []byte(operate)...)
//	bytes = append(bytes, h[:]...)
//	bytes = append(bytes, []byte(uri)...)
//	fmt.Println("what:", hex.EncodeToString(bytes))
//	//var buf []byte
//	//hashforcom := sha3.NewKeccak256()
//	//hashforcom.Write(bytes)
//	//buf = hashforcom.Sum(buf)
//
//	//fmt.Println(hexutil.Encode(buf))
//	hash := crypto.Keccak256Hash(bytes)
//	fmt.Println("crypto.Keccak256Hash(bytes):", hexutil.Encode(hash.Bytes()))
//
//	pri, err := crypto.HexToECDSA(signPrivateKey)
//	if err != nil {
//		return nil, err
//	}
//
//	sig, err := crypto.Sign(hash.Bytes(), pri)
//	fmt.Println("ethereum sig hash:", hex.EncodeToString(sig))
//	//sigxx, err := crypto.Sign(buf, pri)
//	//fmt.Println("ethereum sig buf:", hex.EncodeToString(sigxx))
//
//	acc, _ := account.HexStringToAccount(signPrivateKey)
//	sig2, _ := acc.Sign(bytes)
//	fmt.Println("iotex sig:", hex.EncodeToString(sig2))
//	return sig, err
//}

//func encodeParamsDelete(did, operate string) []byte {
//	arguments := abi.Arguments{
//		{
//			Type: stringTy,
//		},
//		{
//			Type: stringTy,
//		},
//	}
//
//	bytes, _ := arguments.Pack(
//		did,
//		operate,
//	)
//	fmt.Println("what:", hex.EncodeToString(bytes))
//	var buf []byte
//	hash := sha3.NewKeccak256()
//	hash.Write(bytes)
//	buf = hash.Sum(buf)
//
//	fmt.Println(hexutil.Encode(buf))
//
//	return buf
//}

//func TestEncode(t *testing.T) {
//	encodeParamsCreate("", "xx", [32]byte{1, 2, 3}, "uri")
//}
//
//func TestEcrecover(t *testing.T) {
//	require := require.New(t)
//	bytes := []byte("xxxx")
//
//	prefixedHash := crypto.Keccak256Hash(bytes)
//	acc, err := account.HexStringToAccount(signPrivateKey)
//	fmt.Println(acc.Address().String())
//
//	pri, err := crypto.HexToECDSA(signPrivateKey)
//	require.NoError(err)
//
//	sig, err := crypto.Sign(prefixedHash.Bytes(), pri)
//	require.NoError(err)
//
//	pub, err := crypto.Ecrecover(prefixedHash.Bytes(), sig)
//	require.NoError(err)
//	var addr common.Address
//	copy(addr[:], crypto.Keccak256(pub[1:])[12:])
//	fmt.Println(addr.String())
//	add, err := address.FromBytes(addr.Bytes())
//	require.NoError(err)
//	fmt.Println(add.String())
//}
