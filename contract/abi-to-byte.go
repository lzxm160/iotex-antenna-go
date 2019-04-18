// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package contract

import (
	"encoding/hex"
	"strings"
	"tokenxx-algorithm-c/accounts/abi"

	"github.com/ethereum/go-ethereum/crypto"
)

func GetFuncHash(fun string) string {
	return hex.EncodeToString(crypto.Keccak256([]byte(fun))[:4])
}
func (c *Contract) EncodeArguments(method string, args ...interface{}) ([]byte, error) {
	reader := strings.NewReader(c.abi)
	abiParam, err := abi.JSON(reader)
	if err != nil {
		return nil, err
	}
	return abiParam.Pack(method, args...)
}
