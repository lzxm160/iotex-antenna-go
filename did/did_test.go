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

	"github.com/iotexproject/iotex-address/address"

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
	didPrefix        = "did:io:ucam:"
	testuid          = "E7UAA51MKZVCSH6GYHRJ"
	testdid          = didPrefix + testuid
	usingABI         = abibin.UCamDIDManagerABI
	sender           = "io1ph0u2psnd7muq5xv9623rmxdsxc4uapxhzpg02"
	signsender       = "io1vdtfpzkwpyngzvx7u2mauepnzja7kd5rryp0sg"
	privateKey       = "414efa99dfac6f4095d6954713fb0085268d400d6a05a8ae8a69b5b1c10b4bed"
	signPrivateKey   = "0d4d9b248110257c575ef2e8d93dd53471d9178984482817dcbd6edb607f8cc5"
	endpoint         = "api.testnet.iotex.one:443"
	IoTeXDID_address = "io1zkaqhq6q2c3efj78kwtnh8nns2muczrtcjwyy9"
)

var (
	gasPrice = big.NewInt(0).SetUint64(1e12)
	gasLimit = uint64(10000000)
	uisngBIN = abibin.UCamDIDManagerBin[2:]
)

func TestDidDeployContract(t *testing.T) {
	require := require.New(t)
	conn, err := iotex.NewDefaultGRPCConn(endpoint)
	require.NoError(err)
	defer conn.Close()

	acc, err := account.HexStringToAccount(privateKey)
	require.NoError(err)
	c := iotex.NewAuthedClient(iotexapi.NewAPIServiceClient(conn), acc)

	data, err := hex.DecodeString(uisngBIN)
	require.NoError(err)
	abi, err := abi.JSON(strings.NewReader(usingABI))
	require.NoError(err)
	hash, err := c.DeployContract(data).SetArgs(abi, common.Address{0}).SetGasPrice(big.NewInt(int64(unit.Qev))).SetGasLimit(10000000).Call(context.Background())
	require.NoError(err)
	require.NotNil(hash)
	fmt.Println("hash", hex.EncodeToString(hash[:]))
	time.Sleep(20 * time.Second)
	receiptResponse, err := c.GetReceipt(hash).Call(context.Background())
	require.NoError(err)
	contractAddress := receiptResponse.GetReceiptInfo().GetReceipt().GetContractAddress()
	fmt.Println("Status:", receiptResponse.GetReceiptInfo().GetReceipt().Status)
	fmt.Println("Contract Address:", contractAddress)
}

func TestDidCreateDid(t *testing.T) {
	require := require.New(t)
	d, err := NewDID(endpoint, privateKey, IoTeXDID_address, usingABI, gasPrice, gasLimit)
	require.NoError(err)
	testDIDHash := "9c22ff5f21f0b81b113e63f7db6da94fedef11b2119b4088b89664fb9a3cb658"
	testDIDHashBytes, err := hex.DecodeString(testDIDHash)
	require.NoError(err)
	var testDidHashArray [32]byte
	copy(testDidHashArray[:], testDIDHashBytes)
	sender, err := account.HexStringToAccount(privateKey)
	require.NoError(err)
	senderethAddress := common.HexToAddress(hex.EncodeToString(sender.Address().Bytes()))
	signer, err := account.HexStringToAccount(signPrivateKey)
	require.NoError(err)
	signerethAddress := common.HexToAddress(hex.EncodeToString(signer.Address().Bytes()))
	//did := didPrefix + strings.ToLower(signerethAddress.String())
	uri := "s3://iotex-did/documents"
	contractAddress, err := address.FromString(IoTeXDID_address)
	require.NoError(err)

	hb, err := hex.DecodeString(testDIDHash)
	require.NoError(err)
	ha := string(hb)
	authMsg := "I authorize " + strings.ToLower(senderethAddress.String()) + " to create DID " + testdid + " in contract with " + strings.ToLower(common.HexToAddress(
		hex.EncodeToString(contractAddress.Bytes())).String()) + " (" +
		ha + ", " + uri + ")"
	fmt.Println("authMsg", authMsg)
	prefix := fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(authMsg))
	msg := []byte(prefix + authMsg)

	fmt.Println("final packed", string(msg))
	signedMsg, err := signer.Sign(msg)
	require.NoError(err)
	fmt.Println("signed", hex.EncodeToString(signedMsg))
	h, err := d.RegisterDID(testuid, testDIDHash, []byte(uri), signerethAddress, signedMsg)
	require.NoError(err)
	checkHash(h, t)
}

func TestDidGetHash(t *testing.T) {
	require := require.New(t)
	d, err := NewDID(endpoint, "", IoTeXDID_address, usingABI, nil, 0)
	require.NoError(err)
	//acc, err := account.HexStringToAccount(signPrivateKey)
	//require.NoError(err)

	//ethAddress := common.HexToAddress(hex.EncodeToString(acc.Address().Bytes()))

	//didString := didPrefix + ethAddress.String()
	//didString = "did:io:0x6356908ace09268130dee2b7de643314bbeb3683"
	fmt.Println(testdid)
	h, err := d.GetHash(testdid)
	require.NoError(err)
	fmt.Println(h)
}

func TestDidGetURI(t *testing.T) {
	require := require.New(t)
	d, err := NewDID(endpoint, "", IoTeXDID_address, usingABI, nil, 0)
	require.NoError(err)
	//acc, err := account.HexStringToAccount(signPrivateKey)
	//require.NoError(err)
	//
	//ethAddress := common.HexToAddress(hex.EncodeToString(acc.Address().Bytes()))
	//
	//didString := didPrefix + ethAddress.String()
	fmt.Println(testdid)
	h, err := d.GetUri(testdid)
	require.NoError(err)
	fmt.Println(h)
}

func TestDidUpdate(t *testing.T) {
	testuri := "https://gateway.pinata.cloud/ipfs/QmWUzz4EoBVGX9WJM3tsjyU4CzABReD8akrZqpc6f94Ba4/ddo-test.json"
	require := require.New(t)
	d, err := NewDID(endpoint, privateKey, IoTeXDID_address, usingABI, gasPrice, gasLimit)
	require.NoError(err)
	sender, err := account.HexStringToAccount(privateKey)
	require.NoError(err)
	senderethAddress := common.HexToAddress(hex.EncodeToString(sender.Address().Bytes()))
	signer, err := account.HexStringToAccount(signPrivateKey)
	require.NoError(err)
	//signerethAddress := common.HexToAddress(hex.EncodeToString(signer.Address().Bytes()))
	contractAddress, err := address.FromString(IoTeXDID_address)
	require.NoError(err)

	//didString := didPrefix + strings.ToLower(signerethAddress.String())
	//fmt.Println(didString)
	testDIDHash := "4420e4869750c98a56ac621854d2d00e598698ac87193cdfcbb6ed1164e9cbcd"
	hb, err := hex.DecodeString(testDIDHash)
	require.NoError(err)
	ha := string(hb)
	authMsg := "I authorize " + strings.ToLower(senderethAddress.String()) + " to update DID " + testdid + " in contract to " + strings.ToLower(common.HexToAddress(hex.EncodeToString(contractAddress.Bytes())).String()) + " (" + ha + ", " + testuri + ")"
	//"I authorize ", addrToString(agent), " to update DID ", did, " in contract to ", addrToString(address(this)), " (", h, ", ", uri, ")");
	fmt.Println(authMsg)
	prefix := fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(authMsg))
	msgSum := []byte(prefix + authMsg)

	signedMsg, err := signer.Sign(msgSum)
	require.NoError(err)
	var uid [20]byte
	copy(uid[:], []byte(testuid))
	h, err := d.UpdateDID(uid, testDIDHash, testuri, signedMsg)
	require.NoError(err)
	checkHash(h, t)
}

func TestDeregister(t *testing.T) {
	require := require.New(t)
	d, err := NewDID(endpoint, privateKey, IoTeXDID_address, usingABI, gasPrice, gasLimit)
	require.NoError(err)
	sender, err := account.HexStringToAccount(privateKey)
	require.NoError(err)
	senderethAddress := common.HexToAddress(hex.EncodeToString(sender.Address().Bytes()))
	signer, err := account.HexStringToAccount(signPrivateKey)
	require.NoError(err)
	//signerethAddress := common.HexToAddress(hex.EncodeToString(signer.Address().Bytes()))
	contractAddress, err := address.FromString(IoTeXDID_address)
	require.NoError(err)

	//didString := didPrefix + strings.ToLower(signerethAddress.String())
	//fmt.Println(didString)
	authMsg := "I authorize " + strings.ToLower(senderethAddress.String()) + " to delete DID " + testdid + " in contract " + strings.ToLower(common.HexToAddress(hex.EncodeToString(contractAddress.Bytes())).String())
	prefix := fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(authMsg))
	msgSum := []byte(prefix + authMsg)

	signedMsg, err := signer.Sign(msgSum)
	require.NoError(err)
	var uid [20]byte
	copy(uid[:], []byte(testuid))
	h, err := d.DeregisterDID(uid, signedMsg)
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

	log := receiptResponse.GetReceiptInfo().GetReceipt().GetLogs()
	for _, l := range log {
		fmt.Println("data packed", string(l.Data))
		fmt.Println("data", hex.EncodeToString(l.Data))
		for _, top := range l.Topics {
			fmt.Println("topics", hex.EncodeToString(top))
		}
	}
}

func TestConver(t *testing.T) {
	fmt.Println(uint2str(100))
	//b := []byte("0x0ddfc506136fb7c050cc2e9511eccd81b15e7426")
	xx, ok := big.NewInt(0).SetString("0ddfc506136fb7c050cc2e9511eccd81b15e7426", 16)
	r := require.New(t)
	r.True(ok)

	var arr [32]byte
	copy(arr[2:], xx.Bytes())
	arr[0] = 0
	arr[1] = 11
	fmt.Println(addrToString(arr))

	xxx := []byte{1, 2, 3}
	fmt.Println(xxx)
	fmt.Println([]byte(hex.EncodeToString(xxx)))
}

func addrToString(value [32]byte) string {
	alphabet := []byte("0123456789abcdef")

	var str [42]byte
	str[0] = '0'
	str[1] = 'x'
	for i := uint8(0); i < 20; i++ {
		str[2+i*2] = alphabet[uint8(value[i+2]>>4)]
		str[3+i*2] = alphabet[uint8(value[i+2]&0x0f)]
	}
	return string(str[:])
}

func uint2str(i uint) string {
	if i == 0 {
		return "0"
	}
	var length uint
	for j := i; j != 0; {
		length++
		j /= 10
	}
	b := make([]byte, length)
	k := length - 1
	for i != 0 {
		b[k] = uint8(48 + i%10)
		i /= 10
		k--
	}
	return string(b)
}

func TestTe(t *testing.T) {
	b, _ := hex.DecodeString("4920617574686f72697a652030783863303365306432303031316431653738656230623666633565333865313834613430353165653420746f2063726561746520444944206469643a696f3a7563616d3a30783462353661333666363462653231333663363337333366363435383161323433366266313930303520696e20636f6e747261637420776974682030786439303539343661313138393563316237383633623166636564343534393562383136613631313020289c22ff5f21f0b81b113e63f7db6da94fedef11b2119b4088b89664fb9a3cb6582c2073333a2f2f696f7465782d6469642f646f63756d656e747329")
	fmt.Println(string(b))
}
func TestTTT(t *testing.T) {
	uidstr := "E7UAA51MKZVCSH6GYHRJ"
	var domainID [20]byte
	copy(domainID[:], []byte(uidstr))
	fmt.Println([]byte(uidstr))
	uidInt := new(big.Int)
	int256 := big.NewInt(256)
	for i := 0; i < 20; i++ {
		uidInt = uidInt.Mul(uidInt, int256)
		//b1 := uint8(domainID[i])
		//b2 := uint8(domainID[i-1])

		//if b1 >= 97 && b1 <= 102 {
		//	b1 -= 87
		//} else if b1 >= 48 && b1 <= 57 {
		//	b1 -= 48
		//} else if b1 >= 65 && b1 <= 70 {
		//	b1 -= 55
		//}
		//if b2 >= 97 && b2 <= 102 {
		//	b2 -= 87
		//} else if b2 >= 48 && b2 <= 57 {
		//	b2 -= 48
		//} else if b2 >= 65 && b2 <= 70 {
		//	b2 -= 55
		//}
		fmt.Println("b1", string(domainID[i]))
		uidInt = uidInt.Add(uidInt, big.NewInt(int64(domainID[i])))
	}
	fmt.Println(string(uidInt.Bytes()))
}

//func TestBase(t *testing.T) {
//	en := base64.StdEncoding.EncodeToString([]byte(`<settings xmlns="http://maven.apache.org/SETTINGS/1.0.0"
//	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
//	xsi:schemaLocation="http://maven.apache.org/SETTINGS/1.0.0
//                      http://maven.apache.org/xsd/settings-1.0.0.xsd">
//	<activeProfiles>
//		<activeProfile>github</activeProfile>
//	</activeProfiles>
//	<profiles>
//		<profile>
//			<id>github</id>
//			<repositories>
//				<repository>
//					<id>central</id>
//					<url>https://repo.maven.apache.org/maven2</url>
//				</repository>
//				<repository>
//					<id>github-did-common-java</id>
//					<name>GitHub decentralized-identity/did-common-java</name>
//					<url>https://maven.pkg.github.com/decentralized-identity/did-common-java</url>
//				</repository>
//			</repositories>
//		</profile>
//	</profiles>
//	<servers>
//		<!-- personal access token with permission: read:packages -->
//		<server>
//			<id>github-did-common-java</id>
//			<username>lzxm160</username>
//			<password>ef8161e22449b5b5e440a481d156d85c1e9ef338</password>
//		</server>
//	</servers>
//</settings>
//`))
//	fmt.Println(en)
//}

//func TestDidCreateDid(t *testing.T) {
//	require := require.New(t)
//	d, err := NewDID(endpoint, privateKey, IoTeXDID_address, abibin.AddressBasedDIDManagerWithAgentEnabledABI, gasPrice, gasLimit)
//	require.NoError(err)
//	testDIDHash := "1111111111111111111111111111111111111111111111111111111111111112"
//	testDIDHashBytes, err := hex.DecodeString(testDIDHash)
//	require.NoError(err)
//	var testDidHashArray [32]byte
//	copy(testDidHashArray[:], testDIDHashBytes)
//	sender, err := account.HexStringToAccount(privateKey)
//	require.NoError(err)
//	senderethAddress := common.HexToAddress(hex.EncodeToString(sender.Address().Bytes()))
//	signer, err := account.HexStringToAccount(signPrivateKey)
//	require.NoError(err)
//	signerethAddress := common.HexToAddress(hex.EncodeToString(signer.Address().Bytes()))
//	didPrefix := "did:io:"
//	uri := "uri1"
//	contractAddress, err := address.FromString(IoTeXDID_address)
//	require.NoError(err)
//
//	//abi.encodePacked(
//	//"I authorize ", addrToString(agent), " to create DID ", did, " in contract with ", addrToString(address(this)), " (", h, ", ", uri, ")");
//	//recover(keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n", uint2str(message.length), message)), signature);
//	stringTy, err := abi.NewType("string", "", nil)
//	require.NoError(err)
//	bytesTy, err := abi.NewType("bytes", "", nil)
//	require.NoError(err)
//	bytes32Ty, err := abi.NewType("bytes32", "", nil)
//	require.NoError(err)
//	didPacker := abi.Arguments{{Type: stringTy}, {Type: stringTy}}
//	didPacked, err := didPacker.Pack(didPrefix, strings.ToLower(signerethAddress.String()))
//	require.NoError(err)
//	arguments0 := abi.Arguments{{Type: stringTy}, {Type: stringTy}, {Type: stringTy}, {Type: bytesTy}, {Type: stringTy}, {Type: stringTy}, {Type: stringTy}, {Type: bytes32Ty}, {Type: stringTy}, {Type: bytesTy}, {Type: stringTy}}
//	authMsg, err := arguments0.Pack(
//		"I authorize ",
//		strings.ToLower(senderethAddress.String()),
//		" to create DID ",
//		didPacked,
//		" in contract with ",
//		strings.ToLower(common.HexToAddress(
//			hex.EncodeToString(contractAddress.Bytes())).String()),
//		" (",
//		testDidHashArray,
//		", ",
//		[]byte(uri),
//		")")
//	fmt.Println("string authMsg", string(authMsg))
//	//prefix := fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(authMsg))
//	arguments := abi.Arguments{{Type: stringTy}, {Type: stringTy}, {Type: bytesTy}}
//	msg, err := arguments.Pack(
//		"\x19Ethereum Signed Message:\n",
//		fmt.Sprintf("%d", len(authMsg)),
//		authMsg,
//	)
//	require.NoError(err)
//	//msg := append([]byte(prefix), []byte(message)...)
//	fmt.Println("final packed", string(msg))
//	signedMsg, err := signer.Sign(msg)
//	require.NoError(err)
//	fmt.Println("signed", hex.EncodeToString(signedMsg))
//	h, err := d.RegisterDID(testDIDHash, []byte(uri), signerethAddress, signedMsg)
//	require.NoError(err)
//	checkHash(h, t)
//}

func TestD(t *testing.T) {
	//fmt.Println(hex.DecodeString(""))
	fmt.Println([]byte("did:io:"))
}
