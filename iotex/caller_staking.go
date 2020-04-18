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
	stakingCaller struct {
		account account.Account
		api     iotexapi.APIServiceClient
	}
	candidateCaller struct {
		account account.Account
		api     iotexapi.APIServiceClient
	}
	stakingBase struct {
		account  account.Account
		api      iotexapi.APIServiceClient
		payload  []byte
		gasLimit *uint64
		gasPrice *big.Int
		nonce    *uint64
		action   interface{}
	}
)

//Create Staking
func (c *stakingCaller) Create(candidateName string, amount *big.Int, duration uint32, autoStake bool) StakeCreateCaller {
	tx := &iotextypes.StakeCreate{
		CandidateName:  candidateName,
		StakedDuration: duration,
		AutoStake:      autoStake,
		StakedAmount:   amount.String(),
	}
	return &stakingBase{account: c.account, api: c.api, action: tx}
}

//Unstake Staking
func (c *stakingCaller) Unstake(bucketIndex uint64) StakeUnstakeCaller {
	tx := &iotextypes.StakeReclaim{
		BucketIndex: bucketIndex,
	}
	unstake := &reclaim{tx, false}
	return &stakingBase{account: c.account, api: c.api, action: unstake}
}

//Withdraw Staking
func (c *stakingCaller) Withdraw(bucketIndex uint64) StakeWithdrawCaller {
	tx := &iotextypes.StakeReclaim{
		BucketIndex: bucketIndex,
	}
	withdraw := &reclaim{tx, true}
	return &stakingBase{account: c.account, api: c.api, action: withdraw}
}

//AddDeposit Staking
func (c *stakingCaller) AddDeposit(index uint64, amount *big.Int) StakeAddDepositCaller {
	tx := &iotextypes.StakeAddDeposit{
		BucketIndex: index,
		Amount:      amount.String(),
	}
	return &stakingBase{account: c.account, api: c.api, action: tx}
}

//ChangeCandidate Staking
func (c *stakingCaller) ChangeCandidate(candName string, bucketIndex uint64) StakeChangeCandidateCaller {
	tx := &iotextypes.StakeChangeCandidate{
		CandidateName: candName,
		BucketIndex:   bucketIndex,
	}
	return &stakingBase{account: c.account, api: c.api, action: tx}
}

//StakingTransfer Staking
func (c *stakingCaller) StakingTransfer(voterAddress address.Address, bucketIndex uint64) StakeTransferCaller {
	tx := &iotextypes.StakeTransferOwnership{
		VoterAddress: voterAddress.String(),
		BucketIndex:  bucketIndex,
	}
	return &stakingBase{account: c.account, api: c.api, action: tx}
}

//Restake Staking
func (c *stakingCaller) Restake(index uint64, duration uint32, autoStake bool) RestakeCaller {
	tx := &iotextypes.StakeRestake{
		BucketIndex:    index,
		StakedDuration: duration,
		AutoStake:      autoStake,
	}
	return &stakingBase{account: c.account, api: c.api, action: tx}
}

//Register Staking
func (c *candidateCaller) Register(name, operatorAddr, rewardAddr address.Address, amount *big.Int, duration uint32, autoStake bool) CandidateRegisterCaller {
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
	return &stakingBase{account: c.account, api: c.api, action: tx}
}

//Update Staking
func (c *candidateCaller) Update(name string, operatorAddr, rewardAddr address.Address) CandidateUpdateCaller {
	tx := &iotextypes.CandidateBasicInfo{
		Name:            name,
		OperatorAddress: operatorAddr.String(),
		RewardAddress:   rewardAddr.String(),
	}
	return &stakingBase{account: c.account, api: c.api, action: tx}
}

func (c *stakingBase) SetGasLimit(g uint64) StakingBase {
	c.gasLimit = &g
	return c
}

func (c *stakingBase) SetGasPrice(g *big.Int) StakingBase {
	c.gasPrice = g
	return c
}

func (c *stakingBase) SetNonce(n uint64) StakingBase {
	c.nonce = &n
	return c
}

func (c *stakingBase) SetPayload(pl []byte) StakingBase {
	if pl == nil {
		return c
	}
	c.payload = make([]byte, len(pl))
	copy(c.payload, pl)
	return c
}

func (c *stakingBase) API() iotexapi.APIServiceClient {
	return c.api
}

func (c *stakingBase) Call(ctx context.Context, opts ...grpc.CallOption) (hash.Hash256, error) {
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
