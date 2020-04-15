package iotex

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/iotexproject/go-pkgs/hash"
	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-antenna-go/v2/account"
	"github.com/iotexproject/iotex-proto/golang/iotexapi"
	"google.golang.org/grpc"
)

// SendActionCaller is used to perform a send action call.
type SendActionCaller interface {
	API() iotexapi.APIServiceClient
	Call(ctx context.Context, opts ...grpc.CallOption) (hash.Hash256, error)
}

// TransferCaller is used to perform a transfer call.
type TransferCaller interface {
	SendActionCaller

	SetGasPrice(*big.Int) TransferCaller
	SetGasLimit(uint64) TransferCaller
	SetPayload([]byte) TransferCaller
	SetNonce(uint64) TransferCaller
}

// ClaimRewardCaller is used to perform a claim reward call.
type ClaimRewardCaller interface {
	SendActionCaller

	SetGasPrice(*big.Int) ClaimRewardCaller
	SetGasLimit(uint64) ClaimRewardCaller
	SetData([]byte) ClaimRewardCaller
	SetNonce(uint64) ClaimRewardCaller
}

// DeployContractCaller is used to perform a deploy contract call.
type DeployContractCaller interface {
	SendActionCaller

	SetArgs(abi abi.ABI, args ...interface{}) DeployContractCaller
	SetGasPrice(*big.Int) DeployContractCaller
	SetGasLimit(uint64) DeployContractCaller
	SetNonce(uint64) DeployContractCaller
}

// GetReceiptCaller is used to perform a get receipt call.
type GetReceiptCaller interface {
	Call(ctx context.Context, opts ...grpc.CallOption) (*iotexapi.GetReceiptByActionResponse, error)
}

// GetLogsCaller is used to get logs
type GetLogsCaller interface {
	Call(ctx context.Context, opts ...grpc.CallOption) (*iotexapi.GetLogsResponse, error)
}

// AuthedClient is an iotex client which associate with an account credentials, so it can perform write actions.
type AuthedClient interface {
	ReadOnlyClient

	Contract(contract address.Address, abi abi.ABI) Contract
	Transfer(to address.Address, value *big.Int) TransferCaller
	ClaimReward(value *big.Int) ClaimRewardCaller
	DeployContract(data []byte) DeployContractCaller
	// staking related
	Staking() StakingCaller
	Candidate() CandidateCaller
	Account() account.Account
}

// ReadOnlyClient is an iotex client which can perform read actions.
type ReadOnlyClient interface {
	ReadOnlyContract(contract address.Address, abi abi.ABI) ReadOnlyContract
	GetReceipt(actionHash hash.Hash256) GetReceiptCaller
	GetLogs(request *iotexapi.GetLogsRequest) GetLogsCaller
	API() iotexapi.APIServiceClient
}

// ReadContractCaller is used to perform a read contract call.
type ReadContractCaller interface {
	Call(ctx context.Context, opts ...grpc.CallOption) (Data, error)
}

// ExecuteContractCaller is used to perform an execute contract call.
type ExecuteContractCaller interface {
	SendActionCaller

	SetGasPrice(*big.Int) ExecuteContractCaller
	SetGasLimit(uint64) ExecuteContractCaller
	SetAmount(*big.Int) ExecuteContractCaller
	SetNonce(uint64) ExecuteContractCaller
}

// Contract allows to read or execute on this contract's methods.
type Contract interface {
	ReadOnlyContract

	Execute(method string, args ...interface{}) ExecuteContractCaller
}

// ReadOnlyContract allows to read on this contract's methods.
type ReadOnlyContract interface {
	Read(method string, args ...interface{}) ReadContractCaller
}

// StakingCaller is used to perform a staking call.
type StakingCaller interface {
	Create(candidateName string, amount *big.Int, duration uint32, autoStake bool) StakeCreateCaller
	Unstake(bucketIndex uint64) StakeUnstakeCaller
	Withdraw(bucketIndex uint64) StakeWithdrawCaller
	AddDeposit(index uint64, amount *big.Int) StakeAddDepositCaller
	ChangeCandidate(candName string, bucketIndex uint64) StakeChangeCandidateCaller
	StakingTransfer(voterAddress address.Address, bucketIndex uint64) StakeTransferCaller
	Restake(index uint64, duration uint32, autoStake bool) RestakeCaller
}

// CandidateCaller is used to perform a candidate call.
type CandidateCaller interface {
	Register(name, operatorAddr, rewardAddr address.Address, amount *big.Int, duration uint32, autoStake bool) CandidateRegisterCaller
	Update(name string, operatorAddr, rewardAddr address.Address) CandidateUpdateCaller
}

// StakingBase is for staking
type StakingBase interface {
	SendActionCaller
	SetGasPrice(*big.Int) StakingBase
	SetGasLimit(uint64) StakingBase
	SetNonce(uint64) StakingBase
	SetPayload([]byte) StakingBase
}

// StakeCreateCaller is for stake create call.
type StakeCreateCaller interface {
	StakingBase
}

// StakeUnstakeCaller is for stake unstake call.
type StakeUnstakeCaller interface {
	StakingBase
}

// StakeWithdrawCaller is for stake withdraw call.
type StakeWithdrawCaller interface {
	StakingBase
}

// StakeAddDepositCaller is for stake add deposit call.
type StakeAddDepositCaller interface {
	StakingBase
}

// StakeChangeCandidateCaller is for stake change candidate call.
type StakeChangeCandidateCaller interface {
	StakingBase
}

// StakeTransferCaller is for stake transfer call.
type StakeTransferCaller interface {
	StakingBase
}

// RestakeCaller is for stake restake call.
type RestakeCaller interface {
	StakingBase
}

// CandidateRegisterCaller is for candidate register call.
type CandidateRegisterCaller interface {
	StakingBase
}

// CandidateUpdateCaller is for candidate update call.
type CandidateUpdateCaller interface {
	StakingBase
}
