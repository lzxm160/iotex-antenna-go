// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package contract

import (
	"math/big"

	"github.com/iotexproject/iotex-core/action"
)

type CustomOptions struct {
	address  string
	abi      string
	data     string
	from     string
	gasPrice *big.Int
	gasLimit uint64
}
type contractOptions struct {
	CustomOptions
}
type Contract struct {
	options CustomOptions
}

func NewContract(options CustomOptions) *Contract {
	return &Contract{options}
}
func (c *Contract) ABI() string {
	return c.options.abi
}
func (c *Contract) Address() string {
	return c.options.address
}
func (c *Contract) Deploy() (*action.Execution, error) {
	execution, err := action.NewExecution("", 0, big.NewInt(0), c.gasLimit, c.gasPrice, c.options.Data())
	return execution, err
}
