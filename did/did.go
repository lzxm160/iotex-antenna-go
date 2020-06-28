package did

import (
	"context"
	"encoding/hex"
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-antenna-go/v2/account"
	"github.com/iotexproject/iotex-antenna-go/v2/iotex"
	"github.com/iotexproject/iotex-proto/golang/iotexapi"
)

const (
	createDID = "registerDID"
	deleteDID = "deregisterDID"
	update    = "updateDID"
	getHash   = "getHash"
	getURI    = "getURI"
)

type DID interface {
	RegisterDID(didHash string, uri []byte) (hash string, err error)
	DeregisterDID() (hash string, err error)
	UpdateDID(h, uri string) (hash string, err error)
	GetHash(did string) (hash string, err error)
	GetUri(did string) (uri string, err error)
}

type did struct {
	endpoint string
	account  account.Account
	contract address.Address
	abi      abi.ABI
	gasPrice *big.Int
	gasLimit uint64
}

func NewDID(endpoint, privateKey, contract, abiString string, gasPrice *big.Int, gasLimit uint64) (d DID, err error) {
	abi, err := abi.JSON(strings.NewReader(abiString)) // note,this is IoTeXDID_abi
	if err != nil {
		return
	}
	var acc account.Account
	if privateKey != "" {
		acc, err = account.HexStringToAccount(privateKey)
		if err != nil {
			return
		}
	}

	addr, err := address.FromString(contract)
	if err != nil {
		return
	}
	d = &did{endpoint, acc, addr, abi, gasPrice, gasLimit}
	return
}

func (d *did) RegisterDID(didHash string, uri []byte) (hash string, err error) {
	if len(didHash) != 64 {
		err = errors.New("hash should be 32 bytes")
		return
	}
	conn, err := iotex.NewDefaultGRPCConn(d.endpoint)
	if err != nil {
		return
	}
	defer conn.Close()
	cli := iotex.NewAuthedClient(iotexapi.NewAPIServiceClient(conn), d.account)
	hashSlice, err := hex.DecodeString(didHash)
	if err != nil {
		return
	}
	var hashArray [32]byte
	copy(hashArray[:], hashSlice)

	h, err := cli.Contract(d.contract, d.abi).Execute(createDID, hashArray,
		uri).SetGasPrice(d.gasPrice).SetGasLimit(d.gasLimit).Call(context.Background())
	if err != nil {
		return
	}
	hash = hex.EncodeToString(h[:])
	return
}

func (d *did) DeregisterDID() (hash string, err error) {
	conn, err := iotex.NewDefaultGRPCConn(d.endpoint)
	if err != nil {
		return
	}
	defer conn.Close()
	cli := iotex.NewAuthedClient(iotexapi.NewAPIServiceClient(conn), d.account)
	h, err := cli.Contract(d.contract, d.abi).Execute(deleteDID).SetGasPrice(d.gasPrice).SetGasLimit(d.gasLimit).Call(context.Background())
	if err != nil {
		return
	}
	hash = hex.EncodeToString(h[:])
	return
}

func (d *did) UpdateDID(h, uri string) (hash string, err error) {
	conn, err := iotex.NewDefaultGRPCConn(d.endpoint)
	if err != nil {
		return
	}
	defer conn.Close()
	hashSlice, err := hex.DecodeString(h)
	if err != nil {
		return
	}
	var hashArray [32]byte
	copy(hashArray[:], hashSlice)
	cli := iotex.NewAuthedClient(iotexapi.NewAPIServiceClient(conn), d.account)
	ret, err := cli.Contract(d.contract, d.abi).Execute(update, hashArray,
		[]byte(uri)).SetGasPrice(d.gasPrice).SetGasLimit(d.
		gasLimit).Call(context.Background())
	if err != nil {
		return
	}
	hash = hex.EncodeToString(ret[:])
	return
}

func (d *did) GetHash(did string) (hash string, err error) {
	conn, err := iotex.NewDefaultGRPCConn(d.endpoint)
	if err != nil {
		return
	}
	defer conn.Close()
	cli := iotex.NewReadOnlyClient(iotexapi.NewAPIServiceClient(conn))
	ret, err := cli.ReadOnlyContract(d.contract, d.abi).Read(getHash, []byte(did)).Call(context.Background())
	if err != nil {
		return
	}
	hashBytes := [32]uint8{}
	err = ret.Unmarshal(&hashBytes)
	if err != nil {
		return
	}
	hash = hex.EncodeToString(hashBytes[:])
	return
}

func (d *did) GetUri(did string) (uri string, err error) {
	conn, err := iotex.NewDefaultGRPCConn(d.endpoint)
	if err != nil {
		return
	}
	defer conn.Close()
	cli := iotex.NewReadOnlyClient(iotexapi.NewAPIServiceClient(conn))
	ret, err := cli.ReadOnlyContract(d.contract, d.abi).Read(getURI, []byte(did)).Call(context.Background())
	if err != nil {
		return
	}
	var uriBytes []byte
	err = ret.Unmarshal(&uriBytes)
	if err != nil {
		return
	}
	uri = string(uriBytes)
	return
}
