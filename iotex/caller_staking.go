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
	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-proto/golang/iotexapi"
	"github.com/iotexproject/iotex-proto/golang/iotextypes"

	"github.com/iotexproject/iotex-antenna-go/v2/account"
)

type (
	stakingBase struct {
		account  account.Account
		api      iotexapi.APIServiceClient
		payload  []byte
		gasLimit *uint64
		gasPrice *big.Int
		nonce    *uint64
		action   interface{}
	}
	stakingCaller struct {
		stakingBase
	}
	candidateCaller struct {
		stakingBase
	}
)

//Create Staking
func (c *stakingCaller) Create(candidateName string, amount *big.Int, duration uint32, autoStake bool) StakingCaller {
	tx := &iotextypes.StakeCreate{
		CandidateName:  candidateName,
		StakedDuration: duration,
		AutoStake:      autoStake,
		StakedAmount:   amount.String(),
	}
	return &stakingCaller{
		stakingBase{
			account: c.account,
			api:     c.api,
			action:  tx,
		}}
}

//Unstake Staking
func (c *stakingCaller) Unstake(bucketIndex uint64) StakingCaller {
	tx := &iotextypes.StakeReclaim{
		BucketIndex: bucketIndex,
	}
	unstake := &reclaim{tx, false}
	return &stakingCaller{
		stakingBase{
			account: c.account,
			api:     c.api,
			action:  unstake,
		}}
}

//Withdraw Staking
func (c *stakingCaller) Withdraw(bucketIndex uint64) StakingCaller {
	tx := &iotextypes.StakeReclaim{
		BucketIndex: bucketIndex,
	}
	withdraw := &reclaim{tx, true}
	return &stakingCaller{
		stakingBase{
			account: c.account,
			api:     c.api,
			action:  withdraw,
		}}
}

//AddDeposit Staking
func (c *stakingCaller) AddDeposit(index uint64, amount *big.Int) StakingCaller {
	tx := &iotextypes.StakeAddDeposit{
		BucketIndex: index,
		Amount:      amount.String(),
	}
	return &stakingCaller{
		stakingBase{
			account: c.account,
			api:     c.api,
			action:  tx,
		}}
}

//ChangeCandidate Staking
func (c *stakingCaller) ChangeCandidate(candName string, bucketIndex uint64) StakingCaller {
	tx := &iotextypes.StakeChangeCandidate{
		CandidateName: candName,
		BucketIndex:   bucketIndex,
	}
	return &stakingCaller{
		stakingBase{
			account: c.account,
			api:     c.api,
			action:  tx,
		}}
}

//StakingTransfer Staking
func (c *stakingCaller) StakingTransfer(voterAddress address.Address, bucketIndex uint64) StakingCaller {
	tx := &iotextypes.StakeTransferOwnership{
		VoterAddress: voterAddress.String(),
		BucketIndex:  bucketIndex,
	}
	return &stakingCaller{
		stakingBase{
			account: c.account,
			api:     c.api,
			action:  tx,
		}}
}

//Restake Staking
func (c *stakingCaller) Restake(index uint64, duration uint32, autoStake bool) StakingCaller {
	tx := &iotextypes.StakeRestake{
		BucketIndex:    index,
		StakedDuration: duration,
		AutoStake:      autoStake,
	}
	return &stakingCaller{
		stakingBase{
			account: c.account,
			api:     c.api,
			action:  tx,
		}}
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

func (c *stakingCaller) SetPayload(pl []byte) StakingCaller {
	if pl == nil {
		return c
	}
	c.payload = make([]byte, len(pl))
	copy(c.payload, pl)
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
		action:   c.action,
	}
	return sc.Call(ctx, opts...)
}

//Register Staking
func (c *candidateCaller) Register(name, operatorAddr, rewardAddr address.Address, amount *big.Int, duration uint32, autoStake bool) CandidateCaller {
	basic := &iotextypes.CandidateBasicInfo{
		Name:            name.String(),
		OperatorAddress: operatorAddr.String(),
		RewardAddress:   rewardAddr.String(),
	}
	tx := &iotextypes.CandidateRegister{
		Candidate:      basic,
		StakedAmount:   amount.String(),
		StakedDuration: duration,
		AutoStake:      autoStake,
	}
	return &candidateCaller{
		stakingBase{
			account: c.account,
			api:     c.api,
			action:  tx,
		}}
}

//Update Staking
func (c *candidateCaller) Update(name string, operatorAddr, rewardAddr address.Address) CandidateCaller {
	tx := &iotextypes.CandidateBasicInfo{
		Name:            name,
		OperatorAddress: operatorAddr.String(),
		RewardAddress:   rewardAddr.String(),
	}
	return &candidateCaller{
		stakingBase{
			account: c.account,
			api:     c.api,
			action:  tx,
		}}
}

func (c *candidateCaller) SetGasLimit(g uint64) CandidateCaller {
	c.gasLimit = &g
	return c
}

func (c *candidateCaller) SetGasPrice(g *big.Int) CandidateCaller {
	c.gasPrice = g
	return c
}

func (c *candidateCaller) SetNonce(n uint64) CandidateCaller {
	c.nonce = &n
	return c
}

func (c *candidateCaller) SetPayload(pl []byte) CandidateCaller {
	if pl == nil {
		return c
	}
	c.payload = make([]byte, len(pl))
	copy(c.payload, pl)
	return c
}

func (c *candidateCaller) API() iotexapi.APIServiceClient {
	return c.api
}

func (c *candidateCaller) Call(ctx context.Context, opts ...grpc.CallOption) (hash.Hash256, error) {
	sc := &sendActionCaller{
		account:  c.account,
		api:      c.api,
		gasLimit: c.gasLimit,
		gasPrice: c.gasPrice,
		nonce:    c.nonce,
		action:   c.action,
	}
	return sc.Call(ctx, opts...)
}
