// Copyright (c) 2020 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package iotex

import (
	"context"
	"math/big"

	"github.com/iotexproject/go-pkgs/hash"
	"github.com/iotexproject/iotex-proto/golang/iotexapi"
	"github.com/iotexproject/iotex-proto/golang/iotextypes"
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

// StakeCreate
type createStakeCaller struct {
	stakingBase
	candidateName string
	amount        string
	duration      uint32
	autoStake     bool
}

func (c *createStakeCaller) SetGasLimit(g uint64) StakingCaller {
	c.gasLimit = &g
	return c
}

func (c *createStakeCaller) SetGasPrice(g *big.Int) StakingCaller {
	c.gasPrice = g
	return c
}

func (c *createStakeCaller) SetNonce(n uint64) StakingCaller {
	c.nonce = &n
	return c
}

func (c *createStakeCaller) API() iotexapi.APIServiceClient {
	return c.api
}

func (c *createStakeCaller) Call(ctx context.Context, opts ...grpc.CallOption) (hash.Hash256, error) {
	tx := &iotextypes.StakeCreate{
		CandidateName:  c.candidateName,
		StakedDuration: c.duration,
		AutoStake:      c.autoStake,
		StakedAmount:   c.amount,
	}

	if len(c.payload) > 0 {
		tx.Payload = make([]byte, len(c.payload))
		copy(tx.Payload, c.payload)
	}
	sc := &sendActionCaller{
		account:  c.account,
		api:      c.api,
		gasLimit: c.gasLimit,
		gasPrice: c.gasPrice,
		nonce:    c.nonce,
		action:   tx,
	}
	return sc.Call(ctx, opts...)
}
