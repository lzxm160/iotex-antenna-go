// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package iotx

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"

	"github.com/iotexproject/iotex-core/testutil"

	"github.com/iotexproject/iotex-antenna-go/account"
	"github.com/iotexproject/iotex-antenna-go/rpcmethod"
	"github.com/iotexproject/iotex-core/pkg/keypair"
)

// RPCMethod provides simple interface tp invoke rpc method
type Iotx struct {
	Rpc      *rpcmethod.RPCMethod
	Accounts account.Accounts
}

func NewIotx(host string) (*Iotx, error) {
	iotx := &Iotx{}
	rpc, err := rpcmethod.NewRPCMethod(host)
	if err != nil {
		return nil, err
	}
	iotx.Rpc = rpc
	iotx.Accounts = account.Accounts{}
	return iotx, nil
}
func (this *Iotx) SendTransfer(request *TransferRequest) error {
	accountPrivateKey, exist := this.Accounts.GetAccount(request.From)
	if !exist {
		return errors.New(fmt.Sprintf("account:%s not exist", request.From))
	}
	priKey, err := keypair.HexStringToPrivateKey(accountPrivateKey)
	if err != nil {
		return err
	}
	// get account nonce
	accountReq := &rpcmethod.GetAccountRequest{Address: request.From}
	res, err := this.Rpc.GetAccount(accountReq)
	if err != nil {
		return err
	}
	nonce := res.AccountMeta.Nonce
	amount, ok := new(big.Int).SetString(request.Value, 10)
	if !ok {
		return errors.New(fmt.Sprintf("amount:%s error", request.Value))
	}
	gasLimit, err := strconv.ParseUint(request.GasLimit, 10, 64)
	gasPrice, ok := new(big.Int).SetString(request.GasPrice, 10)
	if !ok {
		return errors.New(fmt.Sprintf("gasPrice:%s error", request.GasPrice))
	}
	testTransfer, err := testutil.SignedTransfer(request.To,
		priKey, nonce, amount, []byte(request.Payload), gasLimit,
		gasPrice)

	testTransferPb := testTransfer.Proto()
	finalAction := &rpcmethod.SendActionRequest{Action: testTransferPb}
	_, err = this.Rpc.SendAction(finalAction)
	return err
}
