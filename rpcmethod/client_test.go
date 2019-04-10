// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package rpcmethod

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/iotexproject/iotex-core/protogen/iotexapi"
)

const(
	host="http://127.0.0.1:14004"
)
func TestGetAccount(t *testing.T) {
	require := require.New(t)
	rpc,err:=NewRPCMethod(host)
	require.NoError(err)
	req:=iotexapi.GetAccountRequest{}
	req.Address="io1mflp9m6hcgm2qcghchsdqj3z3eccrnekx9p0ms"
	res,err:=rpc.GetAccount(&req)
	require.NoError(err)
	fmt.Println(res)
}
