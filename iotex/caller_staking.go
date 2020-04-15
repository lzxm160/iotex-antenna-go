// Copyright (c) 2020 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package iotex

import (
	"context"
	"math/big"

	"google.golang.org/grpc"

	"github.com/iotexproject/go-pkgs/hash"
	"github.com/iotexproject/iotex-proto/golang/iotexapi"
	"github.com/iotexproject/iotex-proto/golang/iotextypes"

	"github.com/iotexproject/iotex-antenna-go/v2/account"
)

type stakingCaller struct {
	account       account.Account
	api           iotexapi.APIServiceClient
	payload       []byte
	gasLimit      *uint64
	gasPrice      *big.Int
	nonce         *uint64
	stakingAction interface{}
}

func (c *stakingCaller) SetGasLimit(g uint64) StakingCaller {
	c.gasLimit = &g
	return c
}

func (c *stakingCaller) SetGasPrice(g *big.Int) StakingCaller {
	c.gasPrice = g
	return c
}

func (c *stakingCaller) SetNonce(n uint64) StakingCaller {
	c.nonce = &n
	return c
}

func (c *stakingCaller) API() iotexapi.APIServiceClient {
	return c.api
}

func (c *stakingCaller) Call(ctx context.Context, opts ...grpc.CallOption) (hash.Hash256, error) {
	sc := &sendActionCaller{
		account:  c.account,
		api:      c.api,
		gasLimit: c.gasLimit,
		gasPrice: c.gasPrice,
		nonce:    c.nonce,
		action:   c.stakingAction,
	}
	return sc.Call(ctx, opts...)
}

// NewStakeCreate return stake create action
func NewStakeCreate(candidateName string, amount string, duration uint32, autoStake bool, payload []byte) interface{} {
	tx := &iotextypes.StakeCreate{
		CandidateName:  candidateName,
		StakedDuration: duration,
		AutoStake:      autoStake,
		StakedAmount:   amount,
	}

	if len(payload) > 0 {
		tx.Payload = make([]byte, len(payload))
		copy(tx.Payload, payload)
	}
	return tx
}

// NewStakeUnstake return stake unstake action
func NewStakeUnstake(bucketIndex uint64, payload []byte) interface{} {
	tx := &iotextypes.StakeReclaim{
		BucketIndex: bucketIndex,
	}
	if len(payload) > 0 {
		tx.Payload = make([]byte, len(payload))
		copy(tx.Payload, payload)
	}
	return &reclaim{tx, false}
}

// NewStakeWithdraw return stake withdraw action
func NewStakeWithdraw(bucketIndex uint64, payload []byte) interface{} {
	tx := &iotextypes.StakeReclaim{
		BucketIndex: bucketIndex,
	}
	if len(payload) > 0 {
		tx.Payload = make([]byte, len(payload))
		copy(tx.Payload, payload)
	}
	return &reclaim{tx, true}
}

// NewStakeAddDeposit return stake add deposit action
func NewStakeAddDeposit(index uint64, amount string, payload []byte) interface{} {
	tx := &iotextypes.StakeAddDeposit{
		BucketIndex: index,
		Amount:      amount,
	}
	if len(payload) > 0 {
		tx.Payload = make([]byte, len(payload))
		copy(tx.Payload, payload)
	}
	return tx
}

// NewStakeChangeCandidate return stake change candidate action
func NewStakeChangeCandidate(candName string, bucketIndex uint64, payload []byte) interface{} {
	tx := &iotextypes.StakeChangeCandidate{
		CandidateName: candName,
		BucketIndex:   bucketIndex,
	}
	if len(payload) > 0 {
		tx.Payload = make([]byte, len(payload))
		copy(tx.Payload, payload)
	}
	return tx
}

// NewStakeTransfer return stake transfer action
func NewStakeTransfer(voterAddress string, bucketIndex uint64, payload []byte) interface{} {
	tx := &iotextypes.StakeTransferOwnership{
		VoterAddress: voterAddress,
		BucketIndex:  bucketIndex,
	}
	if len(payload) > 0 {
		tx.Payload = make([]byte, len(payload))
		copy(tx.Payload, payload)
	}
	return tx
}

// NewStakeRestake return restake action
func NewStakeRestake(index uint64, duration uint32, autoStake bool, payload []byte) interface{} {
	tx := &iotextypes.StakeRestake{
		BucketIndex:    index,
		StakedDuration: duration,
		AutoStake:      autoStake,
	}
	if len(payload) > 0 {
		tx.Payload = make([]byte, len(payload))
		copy(tx.Payload, payload)
	}
	return tx
}

// NewCandidateRegister return candidate register action
func NewCandidateRegister(
	name, operatorAddrStr, rewardAddrStr, ownerAddrStr, amountStr string,
	duration uint32,
	autoStake bool,
	payload []byte,
) interface{} {
	basic := &iotextypes.CandidateBasicInfo{
		Name:            name,
		OperatorAddress: operatorAddrStr,
		RewardAddress:   rewardAddrStr,
	}
	tx := &iotextypes.CandidateRegister{
		Candidate:      basic,
		StakedAmount:   amountStr,
		StakedDuration: duration,
		AutoStake:      autoStake,
		OwnerAddress:   ownerAddrStr,
	}
	if len(payload) > 0 {
		tx.Payload = make([]byte, len(payload))
		copy(tx.Payload, payload)
	}
	return tx
}

// NewCandidateUpdate return candidate update action
func NewCandidateUpdate(name, operatorAddrStr, rewardAddrStr string) interface{} {
	return &iotextypes.CandidateBasicInfo{
		Name:            name,
		OperatorAddress: operatorAddrStr,
		RewardAddress:   rewardAddrStr,
	}
}
