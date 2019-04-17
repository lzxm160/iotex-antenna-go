// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package iotx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	host = "api.iotex.one:80"
)

func TestServer_GetAccount(t *testing.T) {
	require := require.New(t)
	svr, err := NewRPCMethod(host)
	testAddress := "io1ns7y0pxmklk8ceattty6n7makpw76u770u5avy"
	request := &GetAccountRequest{Address: testAddress}
	res, err := svr.GetAccount(request)
	require.NoError(err)
	accountMeta := res.AccountMeta
	require.Equal(testAddress, accountMeta.Address)
	require.Equal(true, 364937 < accountMeta.Nonce)
	require.Equal(true, 364938 < accountMeta.PendingNonce)

	// failure
	_, err = svr.GetAccount(&GetAccountRequest{})
	require.Error(err)
}
