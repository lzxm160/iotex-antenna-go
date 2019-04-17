// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package iotx

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFromRau(t *testing.T) {
	require := require.New(t)
	raw := big.NewInt(12 * 1e18)
	convert := FromRau(raw)
	require.Equal(12, convert)
}
func TestToRau(t *testing.T) {
	require := require.New(t)
	raw := int64(10)
	convert := ToRau(raw)
	require.Equal(0, convert.Cmp(big.NewInt(10*1e18)))
}
