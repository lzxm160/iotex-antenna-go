// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package contract

import (
	"errors"
	"math/big"
	"strconv"

	"github.com/iotexproject/iotex-core/action"
)

type Options struct {
	// The byte code of the contract. Used when the contract gets deployed
	data []byte
}

func (op *Options) Length() int {
	return len(op.data)
}
func (op *Options) Data() []byte {
	return op.data
}

type Contract struct {
	// contract abi
	abi string
	// deploy and call from this address
	address string
	// contract bin
	options  Options
	gasLimit uint64
	gasPrice *big.Int
}

func NewContract(abi, addr string, options Options, gasLimit, gasPrice string) (*Contract, error) {
	limit, err := strconv.ParseUint(gasLimit, 10, 64)
	if err != nil {
		return nil, err
	}
	price, ok := new(big.Int).SetString(gasPrice, 10)
	if !ok {
		return nil, errors.New("gas price convert err")
	}

	return &Contract{abi, addr, options, limit, price}, nil
}
func (c *Contract) ABI() string {
	return c.abi
}
func (c *Contract) Address() string {
	return c.address
}
func (c *Contract) Deploy() (*action.Execution, error) {
	execution, err := action.NewExecution("", 0, big.NewInt(0), c.gasLimit, c.gasPrice, c.options.Data())
	return execution, err
}
