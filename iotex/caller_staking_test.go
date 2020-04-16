package iotex

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/iotexproject/go-pkgs/hash"
	"github.com/iotexproject/iotex-proto/golang/iotexapi"
	"github.com/iotexproject/iotex-proto/golang/iotextypes"

	"github.com/iotexproject/iotex-antenna-go/v2/account"
	"github.com/iotexproject/iotex-antenna-go/v2/utils/unit"
)

func TestStake(t *testing.T) {
	require := require.New(t)
	conn, err := NewDefaultGRPCConn(_testnet)
	require.NoError(err)
	defer conn.Close()

	acc, err := account.HexStringToAccount(_accountPrivateKey)
	require.NoError(err)
	c := NewAuthedClient(iotexapi.NewAPIServiceClient(conn), acc)

	// CandidateRegister
	sc := NewCandidateRegister("test", "io10a298zmzvrt4guq79a9f4x7qedj59y7ery84he", "io13sj9mzpewn25ymheukte4v39hvjdtrfp00mlyv", "io19d0p3ah4g8ww9d7kcxfq87yxe7fnr8rpth5shj", "100", uint32(10000), false, []byte("payload"))
	ret, err := c.Staking(sc).SetGasPrice(big.NewInt(int64(unit.Qev))).SetGasLimit(1000000).Call(context.Background())
	require.Error(err)
	require.Equal(hash.ZeroHash256, ret)

	// need to fix when testnet ready
	//time.Sleep(time.Second * 10)
	//receipt, err := c.GetReceipt(hash).Call(context.Background())
	//require.NoError(err)
	//require.Equal(iotextypes.ReceiptStatus_Success, receipt.ReceiptInfo.Receipt.Status)

	// StakeCreate
	sc = NewStakeCreate("io19d0p3ah4g8ww9d7kcxfq87yxe7fnr8rpth5shj", "100", uint32(10000), true, []byte("payload"))
	ret, err = c.Staking(sc).SetGasPrice(big.NewInt(int64(unit.Qev))).SetGasLimit(1000000).Call(context.Background())
	require.Error(err)
	require.Equal(hash.ZeroHash256, ret)

	// need to fix when testnet ready
	//time.Sleep(time.Second * 10)
	//receipt, err = c.GetReceipt(hash).Call(context.Background())
	//require.NoError(err)
	//require.Equal(iotextypes.ReceiptStatus_Success, receipt.ReceiptInfo.Receipt.Status)

	// StakeUnstake
	sc = NewStakeUnstake(1, []byte("payload"))
	ret, err = c.Staking(sc).SetGasPrice(big.NewInt(int64(unit.Qev))).SetGasLimit(1000000).Call(context.Background())
	require.NoError(err)
	require.NotEqual(hash.ZeroHash256, ret)
	time.Sleep(time.Second * 10)
	receipt, err := c.GetReceipt(ret).Call(context.Background())
	require.NoError(err)
	require.NotEqual(iotextypes.ReceiptStatus_Success, receipt.ReceiptInfo.Receipt.Status)
}
