// Copyright (c) 2020 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package main

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-proto/golang/iotextypes"

	"github.com/iotexproject/iotex-antenna-go/v2/examples/util"
)

// MultiSendService is the MultiSendService interface
type MultiSendService interface {
	// Deploy is the Deploy interface
	Deploy(ctx context.Context, waitContractAddress bool, args ...interface{}) (string, error)
	// MultiSend is the MultiSend interface
	MultiSend(ctx context.Context, to []string, amount []*big.Int) (string, error)
}

type multiSendService struct {
	util.IotexService

	contract address.Address
	abi      abi.ABI
	bin      string
	gasPrice *big.Int
	gasLimit uint64
}

// NewMultiSendService returns MultiSendService
func NewMultiSendService(accountPrivate, abiString, binString, contract string, gasPrice *big.Int, gasLimit uint64, endpoint string, secure bool) (MultiSendService, error) {
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
	return &multiSendService{
		util.NewIotexService(accountPrivate, endpoint, secure),
		addr, abi, binString, gasPrice, gasLimit,
	}, nil
}

// Deploy is the Deploy interface
func (s *multiSendService) Deploy(ctx context.Context, waitContractAddress bool, args ...interface{}) (hash string, err error) {
	err = s.Connect()
	if err != nil {
		return
	}
	data, err := hex.DecodeString(s.bin)
	if err != nil {
		return
	}
	h, err := s.AuthClient().DeployContract(data).SetGasPrice(s.gasPrice).SetGasLimit(s.gasLimit).SetArgs(s.abi, args...).Call(ctx)
	if err != nil {
		return
	}
	hash = hex.EncodeToString(h[:])
	if waitContractAddress {
		time.Sleep(time.Second * 10)
		receiptResponse, err := s.AuthClient().GetReceipt(h).Call(ctx)
		if err != nil {
			return "", err
		}
		status := receiptResponse.GetReceiptInfo().GetReceipt().GetStatus()
		if status != uint64(iotextypes.ReceiptStatus_Success) {
			return "", errors.New("deploy error,status:" + fmt.Sprintf("%d", status))
		}
		addr := receiptResponse.GetReceiptInfo().GetReceipt().GetContractAddress()
		s.contract, err = address.FromString(addr)
		if err != nil {
			return "", err
		}
		fmt.Println("s.contract", s.contract)
	}
	return
}

// MultiSend is the MultiSend interface
func (s *multiSendService) MultiSend(ctx context.Context, to []string, amount []*big.Int) (hash string, err error) {
	err = s.Connect()
	if err != nil {
		return
	}
	h, err := s.AuthClient().Contract(s.contract, s.abi).Execute("multiSend", to, amount, "").SetGasPrice(s.gasPrice).SetGasLimit(s.gasLimit).Call(ctx)
	if err != nil {
		return
	}
	hash = hex.EncodeToString(h[:])
	return
}
