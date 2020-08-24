package main

import (
	"context"
	"encoding/hex"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/iotexproject/iotex-address/address"

	"github.com/iotexproject/iotex-antenna-go/v2/examples/service"
)

type contractExample interface {
	Deploy(ctx context.Context, waitContractAddress bool, args ...interface{}) (string, error)
	BalanceOf(ctx context.Context, addre string) (balance string, err error)
}

type iotexService struct {
	service.IotexService

	contract address.Address
	abi      abi.ABI
	bin      string
	gasPrice *big.Int
	gasLimit uint64
}

func NewIotexService(accountPrivate, abiString, binString, contract string, gasPrice *big.Int, gasLimit uint64, endpoint string, secure bool) (contractExample, error) {
	abi, err := abi.JSON(strings.NewReader(abiString))
	if err != nil {
		return nil, err
	}
	var addr address.Address
	if contract != "" {
		addr, err = address.FromString(contract)
		if err != nil {
			return nil, err
		}
	}

	return &iotexService{
		service.NewIotexService(accountPrivate, endpoint, secure),
		addr, abi, binString, gasPrice, gasLimit,
	}, nil
}

func (s *iotexService) Deploy(ctx context.Context, waitContractAddress bool, args ...interface{}) (hash string, err error) {
	err = s.Connect()
	if err != nil {
		return
	}
	h, err := s.AuthClient().DeployContract([]byte(s.bin)).SetGasPrice(s.gasPrice).SetGasLimit(s.gasLimit).SetArgs(s.abi, args...).Call(ctx)
	if err != nil {
		return
	}
	hash = hex.EncodeToString(h[:])
	if waitContractAddress {
		time.Sleep(time.Second * 5)
		receiptResponse, err := s.AuthClient().GetReceipt(h).Call(ctx)
		if err != nil {
			return "", err
		}
		addr := receiptResponse.GetReceiptInfo().GetReceipt().GetContractAddress()
		s.contract, err = address.FromString(addr)
		if err != nil {
			return "", err
		}
	}
	return
}

func (s *iotexService) BalanceOf(ctx context.Context, addre string) (balance string, err error) {
	err = s.Connect()
	if err != nil {
		return
	}
	addr, err := address.FromString(addre)
	if err != nil {
		return
	}
	ethAddr := common.HexToAddress(hex.EncodeToString(addr.Bytes()))
	ret, err := s.ReadOnlyClient().ReadOnlyContract(s.contract, s.abi).Read("balanceof", ethAddr).Call(ctx)
	if err != nil {
		return
	}
	err = ret.Unmarshal(&balance)
	return
}
