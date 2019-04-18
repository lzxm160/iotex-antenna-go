// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package iotx

import (
	"github.com/iotexproject/iotex-antenna-go/account"
	"github.com/iotexproject/iotex-antenna-go/rpcmethod"
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
}
func (iotx *Iotx) SendTransfer() {

}

//public sendTransfer(req: TransferRequest): Promise<string> {
//const sender = this.accounts.getAccount(req.from);
//if (!sender) {
//throw new Error(`can not found account: ${req.from}`);
//}
//
//const price = req.gasPrice ? toRau(String(req.gasPrice), "Qev") : undefined;
//const payload = req.payload || "";
//return new TransferMethod(this, sender, {
//gasLimit: req.gasLimit,
//gasPrice: price,
//amount: req.value,
//recipient: req.to,
//payload: payload
//}).execute();
//}
