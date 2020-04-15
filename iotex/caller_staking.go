// Copyright (c) 2020 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package iotex

import (
	"context"
	"math/big"

	"github.com/iotexproject/iotex-proto/golang/iotextypes"

	"github.com/iotexproject/go-pkgs/hash"
	"github.com/iotexproject/iotex-proto/golang/iotexapi"
	"google.golang.org/grpc"

	"github.com/iotexproject/iotex-antenna-go/v2/account"
)

type stakingBase struct {
	account  account.Account
	api      iotexapi.APIServiceClient
	payload  []byte
	gasLimit *uint64
	gasPrice *big.Int
	nonce    *uint64
}

func (c *stakingBase) SetGasLimit(g uint64) StakingCaller {
	c.gasLimit = &g
	return c
}

func (c *stakingBase) SetGasPrice(g *big.Int) StakingCaller {
	c.gasPrice = g
	return c
}

func (c *stakingBase) SetNonce(n uint64) StakingCaller {
	c.nonce = &n
	return c
}

func (c *stakingBase) API() iotexapi.APIServiceClient {
	return c.api
}

func (c *stakingBase) Call(ctx context.Context, action interface{}, opts ...grpc.CallOption) (hash.Hash256, error) {
	sc := &sendActionCaller{
		account:  c.account,
		api:      c.api,
		gasLimit: c.gasLimit,
		gasPrice: c.gasPrice,
		nonce:    c.nonce,
		action:   action,
	}
	return sc.Call(ctx, opts...)
}

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
