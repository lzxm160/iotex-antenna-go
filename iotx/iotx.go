// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package iotx

import (
	"github.com/iotexproject/iotex-antenna-go/rpcmethod"
	"github.com/iotexproject/iotex-core/protogen/iotexapi"
)

// RPCMethod provides simple interface tp invoke rpc method
type Iotx struct {
	*rpcmethod.RPCMethod
}

// NewRPCMethod returns RPCMethod interacting with endpoint
func NewRPCMethod(endpoint string) (*Iotx, error) {
	r, err := rpcmethod.NewRPCMethod(endpoint)
	if err != nil {
		return nil, err
	}
	return &Iotx{r}, nil
}

// Close closes the underlaying connection, after Close, no method should be
// invoked anymore
func (r *Iotx) Close() {
	r.RPCMethod.Close()
}

// GetAccount gets the address detail of an address
func (r *Iotx) GetAccount(in *iotexapi.GetAccountRequest) (*GetAccountResponse, error) {
	ret, err := r.RPCMethod.GetAccount(in)
	if err != nil {
		return nil, err
	}
	return &GetAccountResponse{ret}, nil
}

// GetActions gets action(s) by:
// 1. start index and action count
// 2. action hash
// 3. address with start index and action count
// 4. get unconfirmed actions by address with start index and action count
// 5. block hash with start index and action count
func (r *Iotx) GetActions(in *iotexapi.GetActionsRequest) (*iotexapi.GetActionsResponse, error) {
	return r.RPCMethod.GetActions(in)
}

// GetBlockMetas gets block metadata(s) by:
// 1. start index and block count
// 2. block hash
func (r *Iotx) GetBlockMetas(in *iotexapi.GetBlockMetasRequest) (*iotexapi.GetBlockMetasResponse, error) {
	return r.RPCMethod.GetBlockMetas(in)
}

// GetChainMeta gets chain metadata
func (r *Iotx) GetChainMeta(in *iotexapi.GetChainMetaRequest) (*iotexapi.GetChainMetaResponse, error) {
	return r.RPCMethod.GetChainMeta(in)
}

// GetServerMeta gets server metadata
func (r *Iotx) GetServerMeta(in *iotexapi.GetServerMetaRequest) (*iotexapi.GetServerMetaResponse, error) {
	return r.RPCMethod.GetServerMeta(in)
}

// SendAction sends atcion to svr
func (r *Iotx) SendAction(in *iotexapi.SendActionRequest) (*iotexapi.SendActionResponse, error) {
	return r.RPCMethod.SendAction(in)
}

// GetReceiptByAction gets receipt by action hash
func (r *Iotx) GetReceiptByAction(in *iotexapi.GetReceiptByActionRequest) (*iotexapi.GetReceiptByActionResponse, error) {
	return r.RPCMethod.GetReceiptByAction(in)
}

// ReadContract reads contract
func (r *Iotx) ReadContract(in *iotexapi.ReadContractRequest) (*iotexapi.ReadContractResponse, error) {
	return r.RPCMethod.ReadContract(in)
}

// SuggestGasPrice suggests gas price
func (r *Iotx) SuggestGasPrice(in *iotexapi.SuggestGasPriceRequest) (*iotexapi.SuggestGasPriceResponse, error) {
	return r.RPCMethod.SuggestGasPrice(in)
}

// EstimateGasForAction estimates gas for action
func (r *Iotx) EstimateGasForAction(in *iotexapi.EstimateGasForActionRequest) (*iotexapi.EstimateGasForActionResponse, error) {
	return r.RPCMethod.EstimateGasForAction(in)
}

// ReadState reads state from blockchain
func (r *Iotx) ReadState(in *iotexapi.ReadStateRequest) (*iotexapi.ReadStateResponse, error) {
	return r.RPCMethod.ReadState(in)
}

// GetEpochMeta get epoch meta
func (r *Iotx) GetEpochMeta(in *iotexapi.GetEpochMetaRequest) (*iotexapi.GetEpochMetaResponse, error) {
	return r.RPCMethod.GetEpochMeta(in)
}
