// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abibin

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AgentableABI is the input ABI used to generate the binding from.
const AgentableABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"did\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"h\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"uri\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"getCreateAuthMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"did\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"getDeleteAuthMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"did\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"h\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"uri\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"getUpdateAuthMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AgentableFuncSigs maps the 4-byte function signature to its string representation.
var AgentableFuncSigs = map[string]string{
	"cc3b41dc": "getCreateAuthMessage(bytes,bytes32,bytes,address)",
	"3bdcb26c": "getDeleteAuthMessage(bytes,address)",
	"dd120153": "getUpdateAuthMessage(bytes,bytes32,bytes,address)",
}

// AgentableBin is the compiled bytecode used for deploying new contracts.
var AgentableBin = "0x608060405234801561001057600080fd5b506109fb806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80633bdcb26c14610046578063cc3b41dc1461016a578063dd120153146102a6575b600080fd5b6100f56004803603604081101561005c57600080fd5b810190602081018135600160201b81111561007657600080fd5b82018360208201111561008857600080fd5b803590602001918460018302840111600160201b831117156100a957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160a01b031691506103e29050565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561012f578181015183820152602001610117565b50505050905090810190601f16801561015c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6100f56004803603608081101561018057600080fd5b810190602081018135600160201b81111561019a57600080fd5b8201836020820111156101ac57600080fd5b803590602001918460018302840111600160201b831117156101cd57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092958435959094909350604081019250602001359050600160201b81111561022757600080fd5b82018360208201111561023957600080fd5b803590602001918460018302840111600160201b8311171561025a57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160a01b031691506105429050565b6100f5600480360360808110156102bc57600080fd5b810190602081018135600160201b8111156102d657600080fd5b8201836020820111156102e857600080fd5b803590602001918460018302840111600160201b8311171561030957600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092958435959094909350604081019250602001359050600160201b81111561036357600080fd5b82018360208201111561037557600080fd5b803590602001918460018302840111600160201b8311171561039657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160a01b031691506107219050565b60606103ed8261084a565b836103f73061084a565b60405160200180806b0249030baba3437b934bd32960a51b815250600c0184805190602001908083835b602083106104405780518252601f199092019160209182019101610421565b51815160209384036101000a60001901801990921691161790526e0103a37903232b632ba32902224a21608d1b919093019081528551600f90910192860191508083835b602083106104a35780518252601f199092019160209182019101610484565b51815160209384036101000a60001901801990921691161790526c01034b71031b7b73a3930b1ba1609d1b919093019081528451600d90910192850191508083835b602083106105045780518252601f1990920191602091820191016104e5565b6001836020036101000a0380198251168184511680821785525050505050509050019350505050604051602081830303815290604052905092915050565b606061054d8261084a565b856105573061084a565b868660405160200180806b0249030baba3437b934bd32960a51b815250600c0186805190602001908083835b602083106105a25780518252601f199092019160209182019101610583565b51815160209384036101000a60001901801990921691161790526e0103a379031b932b0ba32902224a21608d1b919093019081528751600f90910192880191508083835b602083106106055780518252601f1990920191602091820191016105e6565b51815160209384036101000a60001901801990921691161790527101034b71031b7b73a3930b1ba103bb4ba34160751b919093019081528651601290910192870191508083835b6020831061066b5780518252601f19909201916020918201910161064c565b51815160209384036101000a600019018019909216911617905261040560f31b919093019081526002810186905261016160f51b60228201528451602490910192850191508083835b602083106106d35780518252601f1990920191602091820191016106b4565b6001836020036101000a03801982511681845116808217855250505050505090500180602960f81b815250600101955050505050506040516020818303038152906040529050949350505050565b606061072c8261084a565b856107363061084a565b868660405160200180806b0249030baba3437b934bd32960a51b815250600c0186805190602001908083835b602083106107815780518252601f199092019160209182019101610762565b51815160209384036101000a60001901801990921691161790526e0103a37903ab83230ba32902224a21608d1b919093019081528751600f90910192880191508083835b602083106107e45780518252601f1990920191602091820191016107c5565b51815160001960209485036101000a019081169019919091161790526f01034b71031b7b73a3930b1ba103a37960851b93909101928352865160109093019290870191508083836020831061066b5780518252601f19909201916020918201910161064c565b604080518082018252601081526f181899199a1a9b1b9c1cb0b131b232b360811b60208201528151602a80825260608281019094526001600160a01b03851692918491602082018180388339019050509050600360fc1b816000815181106108ae57fe5b60200101906001600160f81b031916908160001a905350600f60fb1b816001815181106108d757fe5b60200101906001600160f81b031916908160001a90535060005b60148110156109bd578260048583600c016020811061090c57fe5b1a60f81b6001600160f81b031916901c60f81c60ff168151811061092c57fe5b602001015160f81c60f81b82826002026002018151811061094957fe5b60200101906001600160f81b031916908160001a905350828482600c016020811061097057fe5b825191901a600f1690811061098157fe5b602001015160f81c60f81b82826002026003018151811061099e57fe5b60200101906001600160f81b031916908160001a9053506001016108f1565b5094935050505056fea265627a7a72315820ac5745f74a257ff0c363534bf8afa2e47449a5b3d4d36ad341fa38992610b93f64736f6c63430005100032"

// DeployAgentable deploys a new Ethereum contract, binding an instance of Agentable to it.
func DeployAgentable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Agentable, error) {
	parsed, err := abi.JSON(strings.NewReader(AgentableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AgentableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Agentable{AgentableCaller: AgentableCaller{contract: contract}, AgentableTransactor: AgentableTransactor{contract: contract}, AgentableFilterer: AgentableFilterer{contract: contract}}, nil
}

// Agentable is an auto generated Go binding around an Ethereum contract.
type Agentable struct {
	AgentableCaller     // Read-only binding to the contract
	AgentableTransactor // Write-only binding to the contract
	AgentableFilterer   // Log filterer for contract events
}

// AgentableCaller is an auto generated read-only Go binding around an Ethereum contract.
type AgentableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AgentableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AgentableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AgentableSession struct {
	Contract     *Agentable        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AgentableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AgentableCallerSession struct {
	Contract *AgentableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// AgentableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AgentableTransactorSession struct {
	Contract     *AgentableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AgentableRaw is an auto generated low-level Go binding around an Ethereum contract.
type AgentableRaw struct {
	Contract *Agentable // Generic contract binding to access the raw methods on
}

// AgentableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AgentableCallerRaw struct {
	Contract *AgentableCaller // Generic read-only contract binding to access the raw methods on
}

// AgentableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AgentableTransactorRaw struct {
	Contract *AgentableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAgentable creates a new instance of Agentable, bound to a specific deployed contract.
func NewAgentable(address common.Address, backend bind.ContractBackend) (*Agentable, error) {
	contract, err := bindAgentable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Agentable{AgentableCaller: AgentableCaller{contract: contract}, AgentableTransactor: AgentableTransactor{contract: contract}, AgentableFilterer: AgentableFilterer{contract: contract}}, nil
}

// NewAgentableCaller creates a new read-only instance of Agentable, bound to a specific deployed contract.
func NewAgentableCaller(address common.Address, caller bind.ContractCaller) (*AgentableCaller, error) {
	contract, err := bindAgentable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AgentableCaller{contract: contract}, nil
}

// NewAgentableTransactor creates a new write-only instance of Agentable, bound to a specific deployed contract.
func NewAgentableTransactor(address common.Address, transactor bind.ContractTransactor) (*AgentableTransactor, error) {
	contract, err := bindAgentable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AgentableTransactor{contract: contract}, nil
}

// NewAgentableFilterer creates a new log filterer instance of Agentable, bound to a specific deployed contract.
func NewAgentableFilterer(address common.Address, filterer bind.ContractFilterer) (*AgentableFilterer, error) {
	contract, err := bindAgentable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AgentableFilterer{contract: contract}, nil
}

// bindAgentable binds a generic wrapper to an already deployed contract.
func bindAgentable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AgentableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Agentable *AgentableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Agentable.Contract.AgentableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Agentable *AgentableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agentable.Contract.AgentableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Agentable *AgentableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Agentable.Contract.AgentableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Agentable *AgentableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Agentable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Agentable *AgentableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agentable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Agentable *AgentableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Agentable.Contract.contract.Transact(opts, method, params...)
}

// GetCreateAuthMessage is a free data retrieval call binding the contract method 0xcc3b41dc.
//
// Solidity: function getCreateAuthMessage(bytes did, bytes32 h, bytes uri, address agent) view returns(bytes)
func (_Agentable *AgentableCaller) GetCreateAuthMessage(opts *bind.CallOpts, did []byte, h [32]byte, uri []byte, agent common.Address) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Agentable.contract.Call(opts, out, "getCreateAuthMessage", did, h, uri, agent)
	return *ret0, err
}

// GetCreateAuthMessage is a free data retrieval call binding the contract method 0xcc3b41dc.
//
// Solidity: function getCreateAuthMessage(bytes did, bytes32 h, bytes uri, address agent) view returns(bytes)
func (_Agentable *AgentableSession) GetCreateAuthMessage(did []byte, h [32]byte, uri []byte, agent common.Address) ([]byte, error) {
	return _Agentable.Contract.GetCreateAuthMessage(&_Agentable.CallOpts, did, h, uri, agent)
}

// GetCreateAuthMessage is a free data retrieval call binding the contract method 0xcc3b41dc.
//
// Solidity: function getCreateAuthMessage(bytes did, bytes32 h, bytes uri, address agent) view returns(bytes)
func (_Agentable *AgentableCallerSession) GetCreateAuthMessage(did []byte, h [32]byte, uri []byte, agent common.Address) ([]byte, error) {
	return _Agentable.Contract.GetCreateAuthMessage(&_Agentable.CallOpts, did, h, uri, agent)
}

// GetDeleteAuthMessage is a free data retrieval call binding the contract method 0x3bdcb26c.
//
// Solidity: function getDeleteAuthMessage(bytes did, address agent) view returns(bytes)
func (_Agentable *AgentableCaller) GetDeleteAuthMessage(opts *bind.CallOpts, did []byte, agent common.Address) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Agentable.contract.Call(opts, out, "getDeleteAuthMessage", did, agent)
	return *ret0, err
}

// GetDeleteAuthMessage is a free data retrieval call binding the contract method 0x3bdcb26c.
//
// Solidity: function getDeleteAuthMessage(bytes did, address agent) view returns(bytes)
func (_Agentable *AgentableSession) GetDeleteAuthMessage(did []byte, agent common.Address) ([]byte, error) {
	return _Agentable.Contract.GetDeleteAuthMessage(&_Agentable.CallOpts, did, agent)
}

// GetDeleteAuthMessage is a free data retrieval call binding the contract method 0x3bdcb26c.
//
// Solidity: function getDeleteAuthMessage(bytes did, address agent) view returns(bytes)
func (_Agentable *AgentableCallerSession) GetDeleteAuthMessage(did []byte, agent common.Address) ([]byte, error) {
	return _Agentable.Contract.GetDeleteAuthMessage(&_Agentable.CallOpts, did, agent)
}

// GetUpdateAuthMessage is a free data retrieval call binding the contract method 0xdd120153.
//
// Solidity: function getUpdateAuthMessage(bytes did, bytes32 h, bytes uri, address agent) view returns(bytes)
func (_Agentable *AgentableCaller) GetUpdateAuthMessage(opts *bind.CallOpts, did []byte, h [32]byte, uri []byte, agent common.Address) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Agentable.contract.Call(opts, out, "getUpdateAuthMessage", did, h, uri, agent)
	return *ret0, err
}

// GetUpdateAuthMessage is a free data retrieval call binding the contract method 0xdd120153.
//
// Solidity: function getUpdateAuthMessage(bytes did, bytes32 h, bytes uri, address agent) view returns(bytes)
func (_Agentable *AgentableSession) GetUpdateAuthMessage(did []byte, h [32]byte, uri []byte, agent common.Address) ([]byte, error) {
	return _Agentable.Contract.GetUpdateAuthMessage(&_Agentable.CallOpts, did, h, uri, agent)
}

// GetUpdateAuthMessage is a free data retrieval call binding the contract method 0xdd120153.
//
// Solidity: function getUpdateAuthMessage(bytes did, bytes32 h, bytes uri, address agent) view returns(bytes)
func (_Agentable *AgentableCallerSession) GetUpdateAuthMessage(did []byte, h [32]byte, uri []byte, agent common.Address) ([]byte, error) {
	return _Agentable.Contract.GetUpdateAuthMessage(&_Agentable.CallOpts, did, h, uri, agent)
}

// DIDBaseABI is the input ABI used to generate the binding from.
const DIDBaseABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dbAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"DIDCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"DIDDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"DIDUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"db\",\"outputs\":[{\"internalType\":\"contractDIDStorage\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"did\",\"type\":\"bytes\"}],\"name\":\"decodeInternalKey\",\"outputs\":[{\"internalType\":\"bytes20\",\"name\":\"\",\"type\":\"bytes20\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"did\",\"type\":\"bytes\"}],\"name\":\"getHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"did\",\"type\":\"bytes\"}],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"did\",\"type\":\"bytes\"}],\"name\":\"getURI\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferDBOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DIDBaseFuncSigs maps the 4-byte function signature to its string representation.
var DIDBaseFuncSigs = map[string]string{
	"4d655aff": "db()",
	"bf6d3139": "decodeInternalKey(bytes)",
	"b00140aa": "getHash(bytes)",
	"b102bfbe": "getOwner(bytes)",
	"8626dea9": "getURI(bytes)",
	"8da5cb5b": "owner()",
	"32aad017": "transferDBOwnership(address)",
	"f2fde38b": "transferOwnership(address)",
}

// DIDBase is an auto generated Go binding around an Ethereum contract.
type DIDBase struct {
	DIDBaseCaller     // Read-only binding to the contract
	DIDBaseTransactor // Write-only binding to the contract
	DIDBaseFilterer   // Log filterer for contract events
}

// DIDBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type DIDBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIDBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DIDBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIDBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DIDBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIDBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DIDBaseSession struct {
	Contract     *DIDBase          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DIDBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DIDBaseCallerSession struct {
	Contract *DIDBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// DIDBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DIDBaseTransactorSession struct {
	Contract     *DIDBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DIDBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type DIDBaseRaw struct {
	Contract *DIDBase // Generic contract binding to access the raw methods on
}

// DIDBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DIDBaseCallerRaw struct {
	Contract *DIDBaseCaller // Generic read-only contract binding to access the raw methods on
}

// DIDBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DIDBaseTransactorRaw struct {
	Contract *DIDBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDIDBase creates a new instance of DIDBase, bound to a specific deployed contract.
func NewDIDBase(address common.Address, backend bind.ContractBackend) (*DIDBase, error) {
	contract, err := bindDIDBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DIDBase{DIDBaseCaller: DIDBaseCaller{contract: contract}, DIDBaseTransactor: DIDBaseTransactor{contract: contract}, DIDBaseFilterer: DIDBaseFilterer{contract: contract}}, nil
}

// NewDIDBaseCaller creates a new read-only instance of DIDBase, bound to a specific deployed contract.
func NewDIDBaseCaller(address common.Address, caller bind.ContractCaller) (*DIDBaseCaller, error) {
	contract, err := bindDIDBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DIDBaseCaller{contract: contract}, nil
}

// NewDIDBaseTransactor creates a new write-only instance of DIDBase, bound to a specific deployed contract.
func NewDIDBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*DIDBaseTransactor, error) {
	contract, err := bindDIDBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DIDBaseTransactor{contract: contract}, nil
}

// NewDIDBaseFilterer creates a new log filterer instance of DIDBase, bound to a specific deployed contract.
func NewDIDBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*DIDBaseFilterer, error) {
	contract, err := bindDIDBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DIDBaseFilterer{contract: contract}, nil
}

// bindDIDBase binds a generic wrapper to an already deployed contract.
func bindDIDBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DIDBaseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DIDBase *DIDBaseRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DIDBase.Contract.DIDBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DIDBase *DIDBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DIDBase.Contract.DIDBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DIDBase *DIDBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DIDBase.Contract.DIDBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DIDBase *DIDBaseCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DIDBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DIDBase *DIDBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DIDBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DIDBase *DIDBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DIDBase.Contract.contract.Transact(opts, method, params...)
}

// Db is a free data retrieval call binding the contract method 0x4d655aff.
//
// Solidity: function db() view returns(address)
func (_DIDBase *DIDBaseCaller) Db(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DIDBase.contract.Call(opts, out, "db")
	return *ret0, err
}

// Db is a free data retrieval call binding the contract method 0x4d655aff.
//
// Solidity: function db() view returns(address)
func (_DIDBase *DIDBaseSession) Db() (common.Address, error) {
	return _DIDBase.Contract.Db(&_DIDBase.CallOpts)
}

// Db is a free data retrieval call binding the contract method 0x4d655aff.
//
// Solidity: function db() view returns(address)
func (_DIDBase *DIDBaseCallerSession) Db() (common.Address, error) {
	return _DIDBase.Contract.Db(&_DIDBase.CallOpts)
}

// DecodeInternalKey is a free data retrieval call binding the contract method 0xbf6d3139.
//
// Solidity: function decodeInternalKey(bytes did) view returns(bytes20)
func (_DIDBase *DIDBaseCaller) DecodeInternalKey(opts *bind.CallOpts, did []byte) ([20]byte, error) {
	var (
		ret0 = new([20]byte)
	)
	out := ret0
	err := _DIDBase.contract.Call(opts, out, "decodeInternalKey", did)
	return *ret0, err
}

// DecodeInternalKey is a free data retrieval call binding the contract method 0xbf6d3139.
//
// Solidity: function decodeInternalKey(bytes did) view returns(bytes20)
func (_DIDBase *DIDBaseSession) DecodeInternalKey(did []byte) ([20]byte, error) {
	return _DIDBase.Contract.DecodeInternalKey(&_DIDBase.CallOpts, did)
}

// DecodeInternalKey is a free data retrieval call binding the contract method 0xbf6d3139.
//
// Solidity: function decodeInternalKey(bytes did) view returns(bytes20)
func (_DIDBase *DIDBaseCallerSession) DecodeInternalKey(did []byte) ([20]byte, error) {
	return _DIDBase.Contract.DecodeInternalKey(&_DIDBase.CallOpts, did)
}

// GetHash is a free data retrieval call binding the contract method 0xb00140aa.
//
// Solidity: function getHash(bytes did) view returns(bytes32)
func (_DIDBase *DIDBaseCaller) GetHash(opts *bind.CallOpts, did []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _DIDBase.contract.Call(opts, out, "getHash", did)
	return *ret0, err
}

// GetHash is a free data retrieval call binding the contract method 0xb00140aa.
//
// Solidity: function getHash(bytes did) view returns(bytes32)
func (_DIDBase *DIDBaseSession) GetHash(did []byte) ([32]byte, error) {
	return _DIDBase.Contract.GetHash(&_DIDBase.CallOpts, did)
}

// GetHash is a free data retrieval call binding the contract method 0xb00140aa.
//
// Solidity: function getHash(bytes did) view returns(bytes32)
func (_DIDBase *DIDBaseCallerSession) GetHash(did []byte) ([32]byte, error) {
	return _DIDBase.Contract.GetHash(&_DIDBase.CallOpts, did)
}

// GetOwner is a free data retrieval call binding the contract method 0xb102bfbe.
//
// Solidity: function getOwner(bytes did) view returns(address)
func (_DIDBase *DIDBaseCaller) GetOwner(opts *bind.CallOpts, did []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DIDBase.contract.Call(opts, out, "getOwner", did)
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0xb102bfbe.
//
// Solidity: function getOwner(bytes did) view returns(address)
func (_DIDBase *DIDBaseSession) GetOwner(did []byte) (common.Address, error) {
	return _DIDBase.Contract.GetOwner(&_DIDBase.CallOpts, did)
}

// GetOwner is a free data retrieval call binding the contract method 0xb102bfbe.
//
// Solidity: function getOwner(bytes did) view returns(address)
func (_DIDBase *DIDBaseCallerSession) GetOwner(did []byte) (common.Address, error) {
	return _DIDBase.Contract.GetOwner(&_DIDBase.CallOpts, did)
}

// GetURI is a free data retrieval call binding the contract method 0x8626dea9.
//
// Solidity: function getURI(bytes did) view returns(bytes)
func (_DIDBase *DIDBaseCaller) GetURI(opts *bind.CallOpts, did []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _DIDBase.contract.Call(opts, out, "getURI", did)
	return *ret0, err
}

// GetURI is a free data retrieval call binding the contract method 0x8626dea9.
//
// Solidity: function getURI(bytes did) view returns(bytes)
func (_DIDBase *DIDBaseSession) GetURI(did []byte) ([]byte, error) {
	return _DIDBase.Contract.GetURI(&_DIDBase.CallOpts, did)
}

// GetURI is a free data retrieval call binding the contract method 0x8626dea9.
//
// Solidity: function getURI(bytes did) view returns(bytes)
func (_DIDBase *DIDBaseCallerSession) GetURI(did []byte) ([]byte, error) {
	return _DIDBase.Contract.GetURI(&_DIDBase.CallOpts, did)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DIDBase *DIDBaseCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DIDBase.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DIDBase *DIDBaseSession) Owner() (common.Address, error) {
	return _DIDBase.Contract.Owner(&_DIDBase.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DIDBase *DIDBaseCallerSession) Owner() (common.Address, error) {
	return _DIDBase.Contract.Owner(&_DIDBase.CallOpts)
}

// TransferDBOwnership is a paid mutator transaction binding the contract method 0x32aad017.
//
// Solidity: function transferDBOwnership(address newOwner) returns()
func (_DIDBase *DIDBaseTransactor) TransferDBOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DIDBase.contract.Transact(opts, "transferDBOwnership", newOwner)
}

// TransferDBOwnership is a paid mutator transaction binding the contract method 0x32aad017.
//
// Solidity: function transferDBOwnership(address newOwner) returns()
func (_DIDBase *DIDBaseSession) TransferDBOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DIDBase.Contract.TransferDBOwnership(&_DIDBase.TransactOpts, newOwner)
}

// TransferDBOwnership is a paid mutator transaction binding the contract method 0x32aad017.
//
// Solidity: function transferDBOwnership(address newOwner) returns()
func (_DIDBase *DIDBaseTransactorSession) TransferDBOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DIDBase.Contract.TransferDBOwnership(&_DIDBase.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DIDBase *DIDBaseTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DIDBase.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DIDBase *DIDBaseSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DIDBase.Contract.TransferOwnership(&_DIDBase.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DIDBase *DIDBaseTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DIDBase.Contract.TransferOwnership(&_DIDBase.TransactOpts, newOwner)
}

// DIDBaseDIDCreatedIterator is returned from FilterDIDCreated and is used to iterate over the raw logs and unpacked data for DIDCreated events raised by the DIDBase contract.
type DIDBaseDIDCreatedIterator struct {
	Event *DIDBaseDIDCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DIDBaseDIDCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DIDBaseDIDCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DIDBaseDIDCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DIDBaseDIDCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DIDBaseDIDCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DIDBaseDIDCreated represents a DIDCreated event raised by the DIDBase contract.
type DIDBaseDIDCreated struct {
	Operator common.Address
	Did      string
	Hash     [32]byte
	Uri      string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDIDCreated is a free log retrieval operation binding the contract event 0x0a36151ef7283a91a6a8d052cd44a0c20f8c7207decf22d985eba8eb0a5a268b.
//
// Solidity: event DIDCreated(address indexed operator, string did, bytes32 hash, string uri)
func (_DIDBase *DIDBaseFilterer) FilterDIDCreated(opts *bind.FilterOpts, operator []common.Address) (*DIDBaseDIDCreatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DIDBase.contract.FilterLogs(opts, "DIDCreated", operatorRule)
	if err != nil {
		return nil, err
	}
	return &DIDBaseDIDCreatedIterator{contract: _DIDBase.contract, event: "DIDCreated", logs: logs, sub: sub}, nil
}

// WatchDIDCreated is a free log subscription operation binding the contract event 0x0a36151ef7283a91a6a8d052cd44a0c20f8c7207decf22d985eba8eb0a5a268b.
//
// Solidity: event DIDCreated(address indexed operator, string did, bytes32 hash, string uri)
func (_DIDBase *DIDBaseFilterer) WatchDIDCreated(opts *bind.WatchOpts, sink chan<- *DIDBaseDIDCreated, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DIDBase.contract.WatchLogs(opts, "DIDCreated", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DIDBaseDIDCreated)
				if err := _DIDBase.contract.UnpackLog(event, "DIDCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDIDCreated is a log parse operation binding the contract event 0x0a36151ef7283a91a6a8d052cd44a0c20f8c7207decf22d985eba8eb0a5a268b.
//
// Solidity: event DIDCreated(address indexed operator, string did, bytes32 hash, string uri)
func (_DIDBase *DIDBaseFilterer) ParseDIDCreated(log types.Log) (*DIDBaseDIDCreated, error) {
	event := new(DIDBaseDIDCreated)
	if err := _DIDBase.contract.UnpackLog(event, "DIDCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DIDBaseDIDDeletedIterator is returned from FilterDIDDeleted and is used to iterate over the raw logs and unpacked data for DIDDeleted events raised by the DIDBase contract.
type DIDBaseDIDDeletedIterator struct {
	Event *DIDBaseDIDDeleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DIDBaseDIDDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DIDBaseDIDDeleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DIDBaseDIDDeleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DIDBaseDIDDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DIDBaseDIDDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DIDBaseDIDDeleted represents a DIDDeleted event raised by the DIDBase contract.
type DIDBaseDIDDeleted struct {
	Operator common.Address
	Did      common.Hash
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDIDDeleted is a free log retrieval operation binding the contract event 0x56c5173d83675a7f2b174113df9205152eb36c65ab39e41d5d1b8d9cb88eabd2.
//
// Solidity: event DIDDeleted(address indexed operator, string indexed did)
func (_DIDBase *DIDBaseFilterer) FilterDIDDeleted(opts *bind.FilterOpts, operator []common.Address, did []string) (*DIDBaseDIDDeletedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _DIDBase.contract.FilterLogs(opts, "DIDDeleted", operatorRule, didRule)
	if err != nil {
		return nil, err
	}
	return &DIDBaseDIDDeletedIterator{contract: _DIDBase.contract, event: "DIDDeleted", logs: logs, sub: sub}, nil
}

// WatchDIDDeleted is a free log subscription operation binding the contract event 0x56c5173d83675a7f2b174113df9205152eb36c65ab39e41d5d1b8d9cb88eabd2.
//
// Solidity: event DIDDeleted(address indexed operator, string indexed did)
func (_DIDBase *DIDBaseFilterer) WatchDIDDeleted(opts *bind.WatchOpts, sink chan<- *DIDBaseDIDDeleted, operator []common.Address, did []string) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _DIDBase.contract.WatchLogs(opts, "DIDDeleted", operatorRule, didRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DIDBaseDIDDeleted)
				if err := _DIDBase.contract.UnpackLog(event, "DIDDeleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDIDDeleted is a log parse operation binding the contract event 0x56c5173d83675a7f2b174113df9205152eb36c65ab39e41d5d1b8d9cb88eabd2.
//
// Solidity: event DIDDeleted(address indexed operator, string indexed did)
func (_DIDBase *DIDBaseFilterer) ParseDIDDeleted(log types.Log) (*DIDBaseDIDDeleted, error) {
	event := new(DIDBaseDIDDeleted)
	if err := _DIDBase.contract.UnpackLog(event, "DIDDeleted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DIDBaseDIDUpdatedIterator is returned from FilterDIDUpdated and is used to iterate over the raw logs and unpacked data for DIDUpdated events raised by the DIDBase contract.
type DIDBaseDIDUpdatedIterator struct {
	Event *DIDBaseDIDUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DIDBaseDIDUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DIDBaseDIDUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DIDBaseDIDUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DIDBaseDIDUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DIDBaseDIDUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DIDBaseDIDUpdated represents a DIDUpdated event raised by the DIDBase contract.
type DIDBaseDIDUpdated struct {
	Operator common.Address
	Did      common.Hash
	Hash     [32]byte
	Uri      string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDIDUpdated is a free log retrieval operation binding the contract event 0x30665598c4da8be2cccefb7fe32a379decd1eae407564569a8e95b51e3afccd7.
//
// Solidity: event DIDUpdated(address indexed operator, string indexed did, bytes32 hash, string uri)
func (_DIDBase *DIDBaseFilterer) FilterDIDUpdated(opts *bind.FilterOpts, operator []common.Address, did []string) (*DIDBaseDIDUpdatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _DIDBase.contract.FilterLogs(opts, "DIDUpdated", operatorRule, didRule)
	if err != nil {
		return nil, err
	}
	return &DIDBaseDIDUpdatedIterator{contract: _DIDBase.contract, event: "DIDUpdated", logs: logs, sub: sub}, nil
}

// WatchDIDUpdated is a free log subscription operation binding the contract event 0x30665598c4da8be2cccefb7fe32a379decd1eae407564569a8e95b51e3afccd7.
//
// Solidity: event DIDUpdated(address indexed operator, string indexed did, bytes32 hash, string uri)
func (_DIDBase *DIDBaseFilterer) WatchDIDUpdated(opts *bind.WatchOpts, sink chan<- *DIDBaseDIDUpdated, operator []common.Address, did []string) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _DIDBase.contract.WatchLogs(opts, "DIDUpdated", operatorRule, didRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DIDBaseDIDUpdated)
				if err := _DIDBase.contract.UnpackLog(event, "DIDUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDIDUpdated is a log parse operation binding the contract event 0x30665598c4da8be2cccefb7fe32a379decd1eae407564569a8e95b51e3afccd7.
//
// Solidity: event DIDUpdated(address indexed operator, string indexed did, bytes32 hash, string uri)
func (_DIDBase *DIDBaseFilterer) ParseDIDUpdated(log types.Log) (*DIDBaseDIDUpdated, error) {
	event := new(DIDBaseDIDUpdated)
	if err := _DIDBase.contract.UnpackLog(event, "DIDUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DIDBaseOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DIDBase contract.
type DIDBaseOwnershipTransferredIterator struct {
	Event *DIDBaseOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DIDBaseOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DIDBaseOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DIDBaseOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DIDBaseOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DIDBaseOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DIDBaseOwnershipTransferred represents a OwnershipTransferred event raised by the DIDBase contract.
type DIDBaseOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DIDBase *DIDBaseFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DIDBaseOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DIDBase.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DIDBaseOwnershipTransferredIterator{contract: _DIDBase.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DIDBase *DIDBaseFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DIDBaseOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DIDBase.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DIDBaseOwnershipTransferred)
				if err := _DIDBase.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DIDBase *DIDBaseFilterer) ParseOwnershipTransferred(log types.Log) (*DIDBaseOwnershipTransferred, error) {
	event := new(DIDBaseOwnershipTransferred)
	if err := _DIDBase.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DIDManagerBaseABI is the input ABI used to generate the binding from.
const DIDManagerBaseABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"DIDCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"DIDDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"DIDUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"db\",\"outputs\":[{\"internalType\":\"contractDIDStorage\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"did\",\"type\":\"bytes\"}],\"name\":\"decodeInternalKey\",\"outputs\":[{\"internalType\":\"bytes20\",\"name\":\"\",\"type\":\"bytes20\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"did\",\"type\":\"bytes\"}],\"name\":\"getHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"did\",\"type\":\"bytes\"}],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"did\",\"type\":\"bytes\"}],\"name\":\"getURI\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferDBOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DIDManagerBaseFuncSigs maps the 4-byte function signature to its string representation.
var DIDManagerBaseFuncSigs = map[string]string{
	"4d655aff": "db()",
	"bf6d3139": "decodeInternalKey(bytes)",
	"b00140aa": "getHash(bytes)",
	"b102bfbe": "getOwner(bytes)",
	"8626dea9": "getURI(bytes)",
	"8da5cb5b": "owner()",
	"32aad017": "transferDBOwnership(address)",
	"f2fde38b": "transferOwnership(address)",
}

// DIDManagerBase is an auto generated Go binding around an Ethereum contract.
type DIDManagerBase struct {
	DIDManagerBaseCaller     // Read-only binding to the contract
	DIDManagerBaseTransactor // Write-only binding to the contract
	DIDManagerBaseFilterer   // Log filterer for contract events
}

// DIDManagerBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type DIDManagerBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIDManagerBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DIDManagerBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIDManagerBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DIDManagerBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIDManagerBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DIDManagerBaseSession struct {
	Contract     *DIDManagerBase   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DIDManagerBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DIDManagerBaseCallerSession struct {
	Contract *DIDManagerBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// DIDManagerBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DIDManagerBaseTransactorSession struct {
	Contract     *DIDManagerBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// DIDManagerBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type DIDManagerBaseRaw struct {
	Contract *DIDManagerBase // Generic contract binding to access the raw methods on
}

// DIDManagerBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DIDManagerBaseCallerRaw struct {
	Contract *DIDManagerBaseCaller // Generic read-only contract binding to access the raw methods on
}

// DIDManagerBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DIDManagerBaseTransactorRaw struct {
	Contract *DIDManagerBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDIDManagerBase creates a new instance of DIDManagerBase, bound to a specific deployed contract.
func NewDIDManagerBase(address common.Address, backend bind.ContractBackend) (*DIDManagerBase, error) {
	contract, err := bindDIDManagerBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DIDManagerBase{DIDManagerBaseCaller: DIDManagerBaseCaller{contract: contract}, DIDManagerBaseTransactor: DIDManagerBaseTransactor{contract: contract}, DIDManagerBaseFilterer: DIDManagerBaseFilterer{contract: contract}}, nil
}

// NewDIDManagerBaseCaller creates a new read-only instance of DIDManagerBase, bound to a specific deployed contract.
func NewDIDManagerBaseCaller(address common.Address, caller bind.ContractCaller) (*DIDManagerBaseCaller, error) {
	contract, err := bindDIDManagerBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DIDManagerBaseCaller{contract: contract}, nil
}

// NewDIDManagerBaseTransactor creates a new write-only instance of DIDManagerBase, bound to a specific deployed contract.
func NewDIDManagerBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*DIDManagerBaseTransactor, error) {
	contract, err := bindDIDManagerBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DIDManagerBaseTransactor{contract: contract}, nil
}

// NewDIDManagerBaseFilterer creates a new log filterer instance of DIDManagerBase, bound to a specific deployed contract.
func NewDIDManagerBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*DIDManagerBaseFilterer, error) {
	contract, err := bindDIDManagerBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DIDManagerBaseFilterer{contract: contract}, nil
}

// bindDIDManagerBase binds a generic wrapper to an already deployed contract.
func bindDIDManagerBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DIDManagerBaseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DIDManagerBase *DIDManagerBaseRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DIDManagerBase.Contract.DIDManagerBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DIDManagerBase *DIDManagerBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DIDManagerBase.Contract.DIDManagerBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DIDManagerBase *DIDManagerBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DIDManagerBase.Contract.DIDManagerBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DIDManagerBase *DIDManagerBaseCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DIDManagerBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DIDManagerBase *DIDManagerBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DIDManagerBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DIDManagerBase *DIDManagerBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DIDManagerBase.Contract.contract.Transact(opts, method, params...)
}

// Db is a free data retrieval call binding the contract method 0x4d655aff.
//
// Solidity: function db() view returns(address)
func (_DIDManagerBase *DIDManagerBaseCaller) Db(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DIDManagerBase.contract.Call(opts, out, "db")
	return *ret0, err
}

// Db is a free data retrieval call binding the contract method 0x4d655aff.
//
// Solidity: function db() view returns(address)
func (_DIDManagerBase *DIDManagerBaseSession) Db() (common.Address, error) {
	return _DIDManagerBase.Contract.Db(&_DIDManagerBase.CallOpts)
}

// Db is a free data retrieval call binding the contract method 0x4d655aff.
//
// Solidity: function db() view returns(address)
func (_DIDManagerBase *DIDManagerBaseCallerSession) Db() (common.Address, error) {
	return _DIDManagerBase.Contract.Db(&_DIDManagerBase.CallOpts)
}

// DecodeInternalKey is a free data retrieval call binding the contract method 0xbf6d3139.
//
// Solidity: function decodeInternalKey(bytes did) view returns(bytes20)
func (_DIDManagerBase *DIDManagerBaseCaller) DecodeInternalKey(opts *bind.CallOpts, did []byte) ([20]byte, error) {
	var (
		ret0 = new([20]byte)
	)
	out := ret0
	err := _DIDManagerBase.contract.Call(opts, out, "decodeInternalKey", did)
	return *ret0, err
}

// DecodeInternalKey is a free data retrieval call binding the contract method 0xbf6d3139.
//
// Solidity: function decodeInternalKey(bytes did) view returns(bytes20)
func (_DIDManagerBase *DIDManagerBaseSession) DecodeInternalKey(did []byte) ([20]byte, error) {
	return _DIDManagerBase.Contract.DecodeInternalKey(&_DIDManagerBase.CallOpts, did)
}

// DecodeInternalKey is a free data retrieval call binding the contract method 0xbf6d3139.
//
// Solidity: function decodeInternalKey(bytes did) view returns(bytes20)
func (_DIDManagerBase *DIDManagerBaseCallerSession) DecodeInternalKey(did []byte) ([20]byte, error) {
	return _DIDManagerBase.Contract.DecodeInternalKey(&_DIDManagerBase.CallOpts, did)
}

// GetHash is a free data retrieval call binding the contract method 0xb00140aa.
//
// Solidity: function getHash(bytes did) view returns(bytes32)
func (_DIDManagerBase *DIDManagerBaseCaller) GetHash(opts *bind.CallOpts, did []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _DIDManagerBase.contract.Call(opts, out, "getHash", did)
	return *ret0, err
}

// GetHash is a free data retrieval call binding the contract method 0xb00140aa.
//
// Solidity: function getHash(bytes did) view returns(bytes32)
func (_DIDManagerBase *DIDManagerBaseSession) GetHash(did []byte) ([32]byte, error) {
	return _DIDManagerBase.Contract.GetHash(&_DIDManagerBase.CallOpts, did)
}

// GetHash is a free data retrieval call binding the contract method 0xb00140aa.
//
// Solidity: function getHash(bytes did) view returns(bytes32)
func (_DIDManagerBase *DIDManagerBaseCallerSession) GetHash(did []byte) ([32]byte, error) {
	return _DIDManagerBase.Contract.GetHash(&_DIDManagerBase.CallOpts, did)
}

// GetOwner is a free data retrieval call binding the contract method 0xb102bfbe.
//
// Solidity: function getOwner(bytes did) view returns(address)
func (_DIDManagerBase *DIDManagerBaseCaller) GetOwner(opts *bind.CallOpts, did []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DIDManagerBase.contract.Call(opts, out, "getOwner", did)
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0xb102bfbe.
//
// Solidity: function getOwner(bytes did) view returns(address)
func (_DIDManagerBase *DIDManagerBaseSession) GetOwner(did []byte) (common.Address, error) {
	return _DIDManagerBase.Contract.GetOwner(&_DIDManagerBase.CallOpts, did)
}

// GetOwner is a free data retrieval call binding the contract method 0xb102bfbe.
//
// Solidity: function getOwner(bytes did) view returns(address)
func (_DIDManagerBase *DIDManagerBaseCallerSession) GetOwner(did []byte) (common.Address, error) {
	return _DIDManagerBase.Contract.GetOwner(&_DIDManagerBase.CallOpts, did)
}

// GetURI is a free data retrieval call binding the contract method 0x8626dea9.
//
// Solidity: function getURI(bytes did) view returns(bytes)
func (_DIDManagerBase *DIDManagerBaseCaller) GetURI(opts *bind.CallOpts, did []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _DIDManagerBase.contract.Call(opts, out, "getURI", did)
	return *ret0, err
}

// GetURI is a free data retrieval call binding the contract method 0x8626dea9.
//
// Solidity: function getURI(bytes did) view returns(bytes)
func (_DIDManagerBase *DIDManagerBaseSession) GetURI(did []byte) ([]byte, error) {
	return _DIDManagerBase.Contract.GetURI(&_DIDManagerBase.CallOpts, did)
}

// GetURI is a free data retrieval call binding the contract method 0x8626dea9.
//
// Solidity: function getURI(bytes did) view returns(bytes)
func (_DIDManagerBase *DIDManagerBaseCallerSession) GetURI(did []byte) ([]byte, error) {
	return _DIDManagerBase.Contract.GetURI(&_DIDManagerBase.CallOpts, did)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DIDManagerBase *DIDManagerBaseCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DIDManagerBase.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DIDManagerBase *DIDManagerBaseSession) Owner() (common.Address, error) {
	return _DIDManagerBase.Contract.Owner(&_DIDManagerBase.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DIDManagerBase *DIDManagerBaseCallerSession) Owner() (common.Address, error) {
	return _DIDManagerBase.Contract.Owner(&_DIDManagerBase.CallOpts)
}

// TransferDBOwnership is a paid mutator transaction binding the contract method 0x32aad017.
//
// Solidity: function transferDBOwnership(address newOwner) returns()
func (_DIDManagerBase *DIDManagerBaseTransactor) TransferDBOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DIDManagerBase.contract.Transact(opts, "transferDBOwnership", newOwner)
}

// TransferDBOwnership is a paid mutator transaction binding the contract method 0x32aad017.
//
// Solidity: function transferDBOwnership(address newOwner) returns()
func (_DIDManagerBase *DIDManagerBaseSession) TransferDBOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DIDManagerBase.Contract.TransferDBOwnership(&_DIDManagerBase.TransactOpts, newOwner)
}

// TransferDBOwnership is a paid mutator transaction binding the contract method 0x32aad017.
//
// Solidity: function transferDBOwnership(address newOwner) returns()
func (_DIDManagerBase *DIDManagerBaseTransactorSession) TransferDBOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DIDManagerBase.Contract.TransferDBOwnership(&_DIDManagerBase.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DIDManagerBase *DIDManagerBaseTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DIDManagerBase.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DIDManagerBase *DIDManagerBaseSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DIDManagerBase.Contract.TransferOwnership(&_DIDManagerBase.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DIDManagerBase *DIDManagerBaseTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DIDManagerBase.Contract.TransferOwnership(&_DIDManagerBase.TransactOpts, newOwner)
}

// DIDManagerBaseDIDCreatedIterator is returned from FilterDIDCreated and is used to iterate over the raw logs and unpacked data for DIDCreated events raised by the DIDManagerBase contract.
type DIDManagerBaseDIDCreatedIterator struct {
	Event *DIDManagerBaseDIDCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DIDManagerBaseDIDCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DIDManagerBaseDIDCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DIDManagerBaseDIDCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DIDManagerBaseDIDCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DIDManagerBaseDIDCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DIDManagerBaseDIDCreated represents a DIDCreated event raised by the DIDManagerBase contract.
type DIDManagerBaseDIDCreated struct {
	Operator common.Address
	Did      string
	Hash     [32]byte
	Uri      string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDIDCreated is a free log retrieval operation binding the contract event 0x0a36151ef7283a91a6a8d052cd44a0c20f8c7207decf22d985eba8eb0a5a268b.
//
// Solidity: event DIDCreated(address indexed operator, string did, bytes32 hash, string uri)
func (_DIDManagerBase *DIDManagerBaseFilterer) FilterDIDCreated(opts *bind.FilterOpts, operator []common.Address) (*DIDManagerBaseDIDCreatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DIDManagerBase.contract.FilterLogs(opts, "DIDCreated", operatorRule)
	if err != nil {
		return nil, err
	}
	return &DIDManagerBaseDIDCreatedIterator{contract: _DIDManagerBase.contract, event: "DIDCreated", logs: logs, sub: sub}, nil
}

// WatchDIDCreated is a free log subscription operation binding the contract event 0x0a36151ef7283a91a6a8d052cd44a0c20f8c7207decf22d985eba8eb0a5a268b.
//
// Solidity: event DIDCreated(address indexed operator, string did, bytes32 hash, string uri)
func (_DIDManagerBase *DIDManagerBaseFilterer) WatchDIDCreated(opts *bind.WatchOpts, sink chan<- *DIDManagerBaseDIDCreated, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DIDManagerBase.contract.WatchLogs(opts, "DIDCreated", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DIDManagerBaseDIDCreated)
				if err := _DIDManagerBase.contract.UnpackLog(event, "DIDCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDIDCreated is a log parse operation binding the contract event 0x0a36151ef7283a91a6a8d052cd44a0c20f8c7207decf22d985eba8eb0a5a268b.
//
// Solidity: event DIDCreated(address indexed operator, string did, bytes32 hash, string uri)
func (_DIDManagerBase *DIDManagerBaseFilterer) ParseDIDCreated(log types.Log) (*DIDManagerBaseDIDCreated, error) {
	event := new(DIDManagerBaseDIDCreated)
	if err := _DIDManagerBase.contract.UnpackLog(event, "DIDCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DIDManagerBaseDIDDeletedIterator is returned from FilterDIDDeleted and is used to iterate over the raw logs and unpacked data for DIDDeleted events raised by the DIDManagerBase contract.
type DIDManagerBaseDIDDeletedIterator struct {
	Event *DIDManagerBaseDIDDeleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DIDManagerBaseDIDDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DIDManagerBaseDIDDeleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DIDManagerBaseDIDDeleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DIDManagerBaseDIDDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DIDManagerBaseDIDDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DIDManagerBaseDIDDeleted represents a DIDDeleted event raised by the DIDManagerBase contract.
type DIDManagerBaseDIDDeleted struct {
	Operator common.Address
	Did      common.Hash
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDIDDeleted is a free log retrieval operation binding the contract event 0x56c5173d83675a7f2b174113df9205152eb36c65ab39e41d5d1b8d9cb88eabd2.
//
// Solidity: event DIDDeleted(address indexed operator, string indexed did)
func (_DIDManagerBase *DIDManagerBaseFilterer) FilterDIDDeleted(opts *bind.FilterOpts, operator []common.Address, did []string) (*DIDManagerBaseDIDDeletedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _DIDManagerBase.contract.FilterLogs(opts, "DIDDeleted", operatorRule, didRule)
	if err != nil {
		return nil, err
	}
	return &DIDManagerBaseDIDDeletedIterator{contract: _DIDManagerBase.contract, event: "DIDDeleted", logs: logs, sub: sub}, nil
}

// WatchDIDDeleted is a free log subscription operation binding the contract event 0x56c5173d83675a7f2b174113df9205152eb36c65ab39e41d5d1b8d9cb88eabd2.
//
// Solidity: event DIDDeleted(address indexed operator, string indexed did)
func (_DIDManagerBase *DIDManagerBaseFilterer) WatchDIDDeleted(opts *bind.WatchOpts, sink chan<- *DIDManagerBaseDIDDeleted, operator []common.Address, did []string) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _DIDManagerBase.contract.WatchLogs(opts, "DIDDeleted", operatorRule, didRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DIDManagerBaseDIDDeleted)
				if err := _DIDManagerBase.contract.UnpackLog(event, "DIDDeleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDIDDeleted is a log parse operation binding the contract event 0x56c5173d83675a7f2b174113df9205152eb36c65ab39e41d5d1b8d9cb88eabd2.
//
// Solidity: event DIDDeleted(address indexed operator, string indexed did)
func (_DIDManagerBase *DIDManagerBaseFilterer) ParseDIDDeleted(log types.Log) (*DIDManagerBaseDIDDeleted, error) {
	event := new(DIDManagerBaseDIDDeleted)
	if err := _DIDManagerBase.contract.UnpackLog(event, "DIDDeleted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DIDManagerBaseDIDUpdatedIterator is returned from FilterDIDUpdated and is used to iterate over the raw logs and unpacked data for DIDUpdated events raised by the DIDManagerBase contract.
type DIDManagerBaseDIDUpdatedIterator struct {
	Event *DIDManagerBaseDIDUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DIDManagerBaseDIDUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DIDManagerBaseDIDUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DIDManagerBaseDIDUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DIDManagerBaseDIDUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DIDManagerBaseDIDUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DIDManagerBaseDIDUpdated represents a DIDUpdated event raised by the DIDManagerBase contract.
type DIDManagerBaseDIDUpdated struct {
	Operator common.Address
	Did      common.Hash
	Hash     [32]byte
	Uri      string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDIDUpdated is a free log retrieval operation binding the contract event 0x30665598c4da8be2cccefb7fe32a379decd1eae407564569a8e95b51e3afccd7.
//
// Solidity: event DIDUpdated(address indexed operator, string indexed did, bytes32 hash, string uri)
func (_DIDManagerBase *DIDManagerBaseFilterer) FilterDIDUpdated(opts *bind.FilterOpts, operator []common.Address, did []string) (*DIDManagerBaseDIDUpdatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _DIDManagerBase.contract.FilterLogs(opts, "DIDUpdated", operatorRule, didRule)
	if err != nil {
		return nil, err
	}
	return &DIDManagerBaseDIDUpdatedIterator{contract: _DIDManagerBase.contract, event: "DIDUpdated", logs: logs, sub: sub}, nil
}

// WatchDIDUpdated is a free log subscription operation binding the contract event 0x30665598c4da8be2cccefb7fe32a379decd1eae407564569a8e95b51e3afccd7.
//
// Solidity: event DIDUpdated(address indexed operator, string indexed did, bytes32 hash, string uri)
func (_DIDManagerBase *DIDManagerBaseFilterer) WatchDIDUpdated(opts *bind.WatchOpts, sink chan<- *DIDManagerBaseDIDUpdated, operator []common.Address, did []string) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _DIDManagerBase.contract.WatchLogs(opts, "DIDUpdated", operatorRule, didRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DIDManagerBaseDIDUpdated)
				if err := _DIDManagerBase.contract.UnpackLog(event, "DIDUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDIDUpdated is a log parse operation binding the contract event 0x30665598c4da8be2cccefb7fe32a379decd1eae407564569a8e95b51e3afccd7.
//
// Solidity: event DIDUpdated(address indexed operator, string indexed did, bytes32 hash, string uri)
func (_DIDManagerBase *DIDManagerBaseFilterer) ParseDIDUpdated(log types.Log) (*DIDManagerBaseDIDUpdated, error) {
	event := new(DIDManagerBaseDIDUpdated)
	if err := _DIDManagerBase.contract.UnpackLog(event, "DIDUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DIDManagerBaseOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DIDManagerBase contract.
type DIDManagerBaseOwnershipTransferredIterator struct {
	Event *DIDManagerBaseOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DIDManagerBaseOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DIDManagerBaseOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DIDManagerBaseOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DIDManagerBaseOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DIDManagerBaseOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DIDManagerBaseOwnershipTransferred represents a OwnershipTransferred event raised by the DIDManagerBase contract.
type DIDManagerBaseOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DIDManagerBase *DIDManagerBaseFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DIDManagerBaseOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DIDManagerBase.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DIDManagerBaseOwnershipTransferredIterator{contract: _DIDManagerBase.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DIDManagerBase *DIDManagerBaseFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DIDManagerBaseOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DIDManagerBase.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DIDManagerBaseOwnershipTransferred)
				if err := _DIDManagerBase.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DIDManagerBase *DIDManagerBaseFilterer) ParseOwnershipTransferred(log types.Log) (*DIDManagerBaseOwnershipTransferred, error) {
	event := new(DIDManagerBaseOwnershipTransferred)
	if err := _DIDManagerBase.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DIDStorageABI is the input ABI used to generate the binding from.
const DIDStorageABI = "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_prefix\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"internalDID\",\"type\":\"bytes20\"}],\"name\":\"exist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"internalDID\",\"type\":\"bytes20\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPrefix\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_prefix\",\"type\":\"bytes\"}],\"name\":\"setPrefix\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"internalDID\",\"type\":\"bytes20\"},{\"internalType\":\"bool\",\"name\":\"isExist\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"uri\",\"type\":\"bytes\"}],\"name\":\"upsert\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DIDStorageFuncSigs maps the 4-byte function signature to its string representation.
var DIDStorageFuncSigs = map[string]string{
	"430e0aee": "exist(bytes20)",
	"5acecc78": "get(bytes20)",
	"12d05d3d": "getPrefix()",
	"8da5cb5b": "owner()",
	"4f153145": "setPrefix(bytes)",
	"f2fde38b": "transferOwnership(address)",
	"48cab05b": "upsert(bytes20,bool,address,bytes32,bytes)",
}

// DIDStorageBin is the compiled bytecode used for deploying new contracts.
var DIDStorageBin = "0x608060405234801561001057600080fd5b506040516109bb3803806109bb8339818101604052602081101561003357600080fd5b810190808051604051939291908464010000000082111561005357600080fd5b90830190602082018581111561006857600080fd5b825164010000000081118282018810171561008257600080fd5b82525081516020918201929091019080838360005b838110156100af578181015183820152602001610097565b50505050905090810190601f1680156100dc5780820380516001836020036101000a031916815260200191505b506040525050600080546001600160a01b0319163317905550610107816001600160e01b0361010d16565b506101d6565b6000546001600160a01b0316331461012457600080fd5b805161013790600190602084019061013b565b5050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061017c57805160ff19168380011785556101a9565b828001600101855582156101a9579182015b828111156101a957825182559160200191906001019061018e565b506101b59291506101b9565b5090565b6101d391905b808211156101b557600081556001016101bf565b90565b6107d6806101e56000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80634f1531451161005b5780634f1531451461020f5780635acecc78146102b55780638da5cb5b14610374578063f2fde38b146103985761007d565b806312d05d3d14610082578063430e0aee146100ff57806348cab05b1461013a575b600080fd5b61008a6103be565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100c45781810151838201526020016100ac565b50505050905090810190601f1680156100f15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101266004803603602081101561011557600080fd5b50356001600160601b031916610454565b604080519115158252519081900360200190f35b61020d600480360360a081101561015057600080fd5b6001600160601b03198235169160208101351515916001600160a01b036040830135169160608101359181019060a08101608082013564010000000081111561019857600080fd5b8201836020820111156101aa57600080fd5b803590602001918460018302840111640100000000831117156101cc57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610473945050505050565b005b61020d6004803603602081101561022557600080fd5b81019060208101813564010000000081111561024057600080fd5b82018360208201111561025257600080fd5b8035906020019184600183028401116401000000008311171561027457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610527945050505050565b6102dc600480360360208110156102cb57600080fd5b50356001600160601b031916610555565b60405180846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561033757818101518382015260200161031f565b50505050905090810190601f1680156103645780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b61037c610675565b604080516001600160a01b039092168252519081900360200190f35b61020d600480360360208110156103ae57600080fd5b50356001600160a01b0316610684565b60018054604080516020601f600260001961010087891615020190951694909404938401819004810282018101909252828152606093909290918301828280156104495780601f1061041e57610100808354040283529160200191610449565b820191906000526020600020905b81548152906001019060200180831161042c57829003601f168201915b505050505090505b90565b6001600160601b03191660009081526002602052604090205460ff1690565b6000546001600160a01b0316331461048a57600080fd5b6040805160808101825285151581526001600160a01b038581166020808401918252838501878152606085018781526001600160601b03198c166000908152600280855297902086518154955160ff1990961690151517610100600160a81b0319166101009590961694909402949094178355516001830155915180519394919361051d93928501929190910190610709565b5050505050505050565b6000546001600160a01b0316331461053e57600080fd5b8051610551906001906020840190610709565b5050565b600080606061056384610454565b6105a9576040805162461bcd60e51b815260206004820152601260248201527111125108191bd95cc81b9bdd08195e1a5cdd60721b604482015290519081900360640190fd5b6001600160601b031984166000908152600260208181526040928390208054600180830154928501805487516101009382161584026000190190911696909604601f810186900486028701860190975286865291046001600160a01b03169491939092909183918301828280156106615780601f1061063657610100808354040283529160200191610661565b820191906000526020600020905b81548152906001019060200180831161064457829003601f168201915b505050505090509250925092509193909250565b6000546001600160a01b031681565b6000546001600160a01b0316331461069b57600080fd5b6001600160a01b0381166106ae57600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061074a57805160ff1916838001178555610777565b82800160010185558215610777579182015b8281111561077757825182559160200191906001019061075c565b50610783929150610787565b5090565b61045191905b80821115610783576000815560010161078d56fea265627a7a72315820f3c2c8af7f5e3f24ba7076e122ee0d78744d15be9914a9df3341e14fe3054aa664736f6c63430005100032"

// DeployDIDStorage deploys a new Ethereum contract, binding an instance of DIDStorage to it.
func DeployDIDStorage(auth *bind.TransactOpts, backend bind.ContractBackend, _prefix []byte) (common.Address, *types.Transaction, *DIDStorage, error) {
	parsed, err := abi.JSON(strings.NewReader(DIDStorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DIDStorageBin), backend, _prefix)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DIDStorage{DIDStorageCaller: DIDStorageCaller{contract: contract}, DIDStorageTransactor: DIDStorageTransactor{contract: contract}, DIDStorageFilterer: DIDStorageFilterer{contract: contract}}, nil
}

// DIDStorage is an auto generated Go binding around an Ethereum contract.
type DIDStorage struct {
	DIDStorageCaller     // Read-only binding to the contract
	DIDStorageTransactor // Write-only binding to the contract
	DIDStorageFilterer   // Log filterer for contract events
}

// DIDStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type DIDStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIDStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DIDStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIDStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DIDStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIDStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DIDStorageSession struct {
	Contract     *DIDStorage       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DIDStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DIDStorageCallerSession struct {
	Contract *DIDStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DIDStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DIDStorageTransactorSession struct {
	Contract     *DIDStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DIDStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type DIDStorageRaw struct {
	Contract *DIDStorage // Generic contract binding to access the raw methods on
}

// DIDStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DIDStorageCallerRaw struct {
	Contract *DIDStorageCaller // Generic read-only contract binding to access the raw methods on
}

// DIDStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DIDStorageTransactorRaw struct {
	Contract *DIDStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDIDStorage creates a new instance of DIDStorage, bound to a specific deployed contract.
func NewDIDStorage(address common.Address, backend bind.ContractBackend) (*DIDStorage, error) {
	contract, err := bindDIDStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DIDStorage{DIDStorageCaller: DIDStorageCaller{contract: contract}, DIDStorageTransactor: DIDStorageTransactor{contract: contract}, DIDStorageFilterer: DIDStorageFilterer{contract: contract}}, nil
}

// NewDIDStorageCaller creates a new read-only instance of DIDStorage, bound to a specific deployed contract.
func NewDIDStorageCaller(address common.Address, caller bind.ContractCaller) (*DIDStorageCaller, error) {
	contract, err := bindDIDStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DIDStorageCaller{contract: contract}, nil
}

// NewDIDStorageTransactor creates a new write-only instance of DIDStorage, bound to a specific deployed contract.
func NewDIDStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*DIDStorageTransactor, error) {
	contract, err := bindDIDStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DIDStorageTransactor{contract: contract}, nil
}

// NewDIDStorageFilterer creates a new log filterer instance of DIDStorage, bound to a specific deployed contract.
func NewDIDStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*DIDStorageFilterer, error) {
	contract, err := bindDIDStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DIDStorageFilterer{contract: contract}, nil
}

// bindDIDStorage binds a generic wrapper to an already deployed contract.
func bindDIDStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DIDStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DIDStorage *DIDStorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DIDStorage.Contract.DIDStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DIDStorage *DIDStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DIDStorage.Contract.DIDStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DIDStorage *DIDStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DIDStorage.Contract.DIDStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DIDStorage *DIDStorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DIDStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DIDStorage *DIDStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DIDStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DIDStorage *DIDStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DIDStorage.Contract.contract.Transact(opts, method, params...)
}

// Exist is a free data retrieval call binding the contract method 0x430e0aee.
//
// Solidity: function exist(bytes20 internalDID) view returns(bool)
func (_DIDStorage *DIDStorageCaller) Exist(opts *bind.CallOpts, internalDID [20]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DIDStorage.contract.Call(opts, out, "exist", internalDID)
	return *ret0, err
}

// Exist is a free data retrieval call binding the contract method 0x430e0aee.
//
// Solidity: function exist(bytes20 internalDID) view returns(bool)
func (_DIDStorage *DIDStorageSession) Exist(internalDID [20]byte) (bool, error) {
	return _DIDStorage.Contract.Exist(&_DIDStorage.CallOpts, internalDID)
}

// Exist is a free data retrieval call binding the contract method 0x430e0aee.
//
// Solidity: function exist(bytes20 internalDID) view returns(bool)
func (_DIDStorage *DIDStorageCallerSession) Exist(internalDID [20]byte) (bool, error) {
	return _DIDStorage.Contract.Exist(&_DIDStorage.CallOpts, internalDID)
}

// Get is a free data retrieval call binding the contract method 0x5acecc78.
//
// Solidity: function get(bytes20 internalDID) view returns(address, bytes32, bytes)
func (_DIDStorage *DIDStorageCaller) Get(opts *bind.CallOpts, internalDID [20]byte) (common.Address, [32]byte, []byte, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new([32]byte)
		ret2 = new([]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _DIDStorage.contract.Call(opts, out, "get", internalDID)
	return *ret0, *ret1, *ret2, err
}

// Get is a free data retrieval call binding the contract method 0x5acecc78.
//
// Solidity: function get(bytes20 internalDID) view returns(address, bytes32, bytes)
func (_DIDStorage *DIDStorageSession) Get(internalDID [20]byte) (common.Address, [32]byte, []byte, error) {
	return _DIDStorage.Contract.Get(&_DIDStorage.CallOpts, internalDID)
}

// Get is a free data retrieval call binding the contract method 0x5acecc78.
//
// Solidity: function get(bytes20 internalDID) view returns(address, bytes32, bytes)
func (_DIDStorage *DIDStorageCallerSession) Get(internalDID [20]byte) (common.Address, [32]byte, []byte, error) {
	return _DIDStorage.Contract.Get(&_DIDStorage.CallOpts, internalDID)
}

// GetPrefix is a free data retrieval call binding the contract method 0x12d05d3d.
//
// Solidity: function getPrefix() view returns(bytes)
func (_DIDStorage *DIDStorageCaller) GetPrefix(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _DIDStorage.contract.Call(opts, out, "getPrefix")
	return *ret0, err
}

// GetPrefix is a free data retrieval call binding the contract method 0x12d05d3d.
//
// Solidity: function getPrefix() view returns(bytes)
func (_DIDStorage *DIDStorageSession) GetPrefix() ([]byte, error) {
	return _DIDStorage.Contract.GetPrefix(&_DIDStorage.CallOpts)
}

// GetPrefix is a free data retrieval call binding the contract method 0x12d05d3d.
//
// Solidity: function getPrefix() view returns(bytes)
func (_DIDStorage *DIDStorageCallerSession) GetPrefix() ([]byte, error) {
	return _DIDStorage.Contract.GetPrefix(&_DIDStorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DIDStorage *DIDStorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DIDStorage.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DIDStorage *DIDStorageSession) Owner() (common.Address, error) {
	return _DIDStorage.Contract.Owner(&_DIDStorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DIDStorage *DIDStorageCallerSession) Owner() (common.Address, error) {
	return _DIDStorage.Contract.Owner(&_DIDStorage.CallOpts)
}

// SetPrefix is a paid mutator transaction binding the contract method 0x4f153145.
//
// Solidity: function setPrefix(bytes _prefix) returns()
func (_DIDStorage *DIDStorageTransactor) SetPrefix(opts *bind.TransactOpts, _prefix []byte) (*types.Transaction, error) {
	return _DIDStorage.contract.Transact(opts, "setPrefix", _prefix)
}

// SetPrefix is a paid mutator transaction binding the contract method 0x4f153145.
//
// Solidity: function setPrefix(bytes _prefix) returns()
func (_DIDStorage *DIDStorageSession) SetPrefix(_prefix []byte) (*types.Transaction, error) {
	return _DIDStorage.Contract.SetPrefix(&_DIDStorage.TransactOpts, _prefix)
}

// SetPrefix is a paid mutator transaction binding the contract method 0x4f153145.
//
// Solidity: function setPrefix(bytes _prefix) returns()
func (_DIDStorage *DIDStorageTransactorSession) SetPrefix(_prefix []byte) (*types.Transaction, error) {
	return _DIDStorage.Contract.SetPrefix(&_DIDStorage.TransactOpts, _prefix)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DIDStorage *DIDStorageTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DIDStorage.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DIDStorage *DIDStorageSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DIDStorage.Contract.TransferOwnership(&_DIDStorage.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DIDStorage *DIDStorageTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DIDStorage.Contract.TransferOwnership(&_DIDStorage.TransactOpts, newOwner)
}

// Upsert is a paid mutator transaction binding the contract method 0x48cab05b.
//
// Solidity: function upsert(bytes20 internalDID, bool isExist, address owner, bytes32 hash, bytes uri) returns()
func (_DIDStorage *DIDStorageTransactor) Upsert(opts *bind.TransactOpts, internalDID [20]byte, isExist bool, owner common.Address, hash [32]byte, uri []byte) (*types.Transaction, error) {
	return _DIDStorage.contract.Transact(opts, "upsert", internalDID, isExist, owner, hash, uri)
}

// Upsert is a paid mutator transaction binding the contract method 0x48cab05b.
//
// Solidity: function upsert(bytes20 internalDID, bool isExist, address owner, bytes32 hash, bytes uri) returns()
func (_DIDStorage *DIDStorageSession) Upsert(internalDID [20]byte, isExist bool, owner common.Address, hash [32]byte, uri []byte) (*types.Transaction, error) {
	return _DIDStorage.Contract.Upsert(&_DIDStorage.TransactOpts, internalDID, isExist, owner, hash, uri)
}

// Upsert is a paid mutator transaction binding the contract method 0x48cab05b.
//
// Solidity: function upsert(bytes20 internalDID, bool isExist, address owner, bytes32 hash, bytes uri) returns()
func (_DIDStorage *DIDStorageTransactorSession) Upsert(internalDID [20]byte, isExist bool, owner common.Address, hash [32]byte, uri []byte) (*types.Transaction, error) {
	return _DIDStorage.Contract.Upsert(&_DIDStorage.TransactOpts, internalDID, isExist, owner, hash, uri)
}

// DIDStorageOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DIDStorage contract.
type DIDStorageOwnershipTransferredIterator struct {
	Event *DIDStorageOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DIDStorageOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DIDStorageOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DIDStorageOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DIDStorageOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DIDStorageOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DIDStorageOwnershipTransferred represents a OwnershipTransferred event raised by the DIDStorage contract.
type DIDStorageOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DIDStorage *DIDStorageFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DIDStorageOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DIDStorage.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DIDStorageOwnershipTransferredIterator{contract: _DIDStorage.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DIDStorage *DIDStorageFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DIDStorageOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DIDStorage.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DIDStorageOwnershipTransferred)
				if err := _DIDStorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DIDStorage *DIDStorageFilterer) ParseOwnershipTransferred(log types.Log) (*DIDStorageOwnershipTransferred, error) {
	event := new(DIDStorageOwnershipTransferred)
	if err := _DIDStorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DIDVerifierABI is the input ABI used to generate the binding from.
const DIDVerifierABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DIDVerifierFuncSigs maps the 4-byte function signature to its string representation.
var DIDVerifierFuncSigs = map[string]string{
	"bb9c6c3e": "verify(string)",
}

// DIDVerifier is an auto generated Go binding around an Ethereum contract.
type DIDVerifier struct {
	DIDVerifierCaller     // Read-only binding to the contract
	DIDVerifierTransactor // Write-only binding to the contract
	DIDVerifierFilterer   // Log filterer for contract events
}

// DIDVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type DIDVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIDVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DIDVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIDVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DIDVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIDVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DIDVerifierSession struct {
	Contract     *DIDVerifier      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DIDVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DIDVerifierCallerSession struct {
	Contract *DIDVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DIDVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DIDVerifierTransactorSession struct {
	Contract     *DIDVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DIDVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type DIDVerifierRaw struct {
	Contract *DIDVerifier // Generic contract binding to access the raw methods on
}

// DIDVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DIDVerifierCallerRaw struct {
	Contract *DIDVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// DIDVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DIDVerifierTransactorRaw struct {
	Contract *DIDVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDIDVerifier creates a new instance of DIDVerifier, bound to a specific deployed contract.
func NewDIDVerifier(address common.Address, backend bind.ContractBackend) (*DIDVerifier, error) {
	contract, err := bindDIDVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DIDVerifier{DIDVerifierCaller: DIDVerifierCaller{contract: contract}, DIDVerifierTransactor: DIDVerifierTransactor{contract: contract}, DIDVerifierFilterer: DIDVerifierFilterer{contract: contract}}, nil
}

// NewDIDVerifierCaller creates a new read-only instance of DIDVerifier, bound to a specific deployed contract.
func NewDIDVerifierCaller(address common.Address, caller bind.ContractCaller) (*DIDVerifierCaller, error) {
	contract, err := bindDIDVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DIDVerifierCaller{contract: contract}, nil
}

// NewDIDVerifierTransactor creates a new write-only instance of DIDVerifier, bound to a specific deployed contract.
func NewDIDVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*DIDVerifierTransactor, error) {
	contract, err := bindDIDVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DIDVerifierTransactor{contract: contract}, nil
}

// NewDIDVerifierFilterer creates a new log filterer instance of DIDVerifier, bound to a specific deployed contract.
func NewDIDVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*DIDVerifierFilterer, error) {
	contract, err := bindDIDVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DIDVerifierFilterer{contract: contract}, nil
}

// bindDIDVerifier binds a generic wrapper to an already deployed contract.
func bindDIDVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DIDVerifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DIDVerifier *DIDVerifierRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DIDVerifier.Contract.DIDVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DIDVerifier *DIDVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DIDVerifier.Contract.DIDVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DIDVerifier *DIDVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DIDVerifier.Contract.DIDVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DIDVerifier *DIDVerifierCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DIDVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DIDVerifier *DIDVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DIDVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DIDVerifier *DIDVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DIDVerifier.Contract.contract.Transact(opts, method, params...)
}

// Verify is a free data retrieval call binding the contract method 0xbb9c6c3e.
//
// Solidity: function verify(string did) view returns(bool)
func (_DIDVerifier *DIDVerifierCaller) Verify(opts *bind.CallOpts, did string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DIDVerifier.contract.Call(opts, out, "verify", did)
	return *ret0, err
}

// Verify is a free data retrieval call binding the contract method 0xbb9c6c3e.
//
// Solidity: function verify(string did) view returns(bool)
func (_DIDVerifier *DIDVerifierSession) Verify(did string) (bool, error) {
	return _DIDVerifier.Contract.Verify(&_DIDVerifier.CallOpts, did)
}

// Verify is a free data retrieval call binding the contract method 0xbb9c6c3e.
//
// Solidity: function verify(string did) view returns(bool)
func (_DIDVerifier *DIDVerifierCallerSession) Verify(did string) (bool, error) {
	return _DIDVerifier.Contract.Verify(&_DIDVerifier.CallOpts, did)
}

// IDIDABI is the input ABI used to generate the binding from.
const IDIDABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"DIDCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"DIDDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"DIDUpdated\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"getHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"getURI\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IDIDFuncSigs maps the 4-byte function signature to its string representation.
var IDIDFuncSigs = map[string]string{
	"b00140aa": "getHash(bytes)",
	"b102bfbe": "getOwner(bytes)",
	"8626dea9": "getURI(bytes)",
}

// IDID is an auto generated Go binding around an Ethereum contract.
type IDID struct {
	IDIDCaller     // Read-only binding to the contract
	IDIDTransactor // Write-only binding to the contract
	IDIDFilterer   // Log filterer for contract events
}

// IDIDCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDIDCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDIDTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDIDTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDIDFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDIDFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDIDSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDIDSession struct {
	Contract     *IDID             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDIDCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDIDCallerSession struct {
	Contract *IDIDCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IDIDTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDIDTransactorSession struct {
	Contract     *IDIDTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDIDRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDIDRaw struct {
	Contract *IDID // Generic contract binding to access the raw methods on
}

// IDIDCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDIDCallerRaw struct {
	Contract *IDIDCaller // Generic read-only contract binding to access the raw methods on
}

// IDIDTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDIDTransactorRaw struct {
	Contract *IDIDTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDID creates a new instance of IDID, bound to a specific deployed contract.
func NewIDID(address common.Address, backend bind.ContractBackend) (*IDID, error) {
	contract, err := bindIDID(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDID{IDIDCaller: IDIDCaller{contract: contract}, IDIDTransactor: IDIDTransactor{contract: contract}, IDIDFilterer: IDIDFilterer{contract: contract}}, nil
}

// NewIDIDCaller creates a new read-only instance of IDID, bound to a specific deployed contract.
func NewIDIDCaller(address common.Address, caller bind.ContractCaller) (*IDIDCaller, error) {
	contract, err := bindIDID(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDIDCaller{contract: contract}, nil
}

// NewIDIDTransactor creates a new write-only instance of IDID, bound to a specific deployed contract.
func NewIDIDTransactor(address common.Address, transactor bind.ContractTransactor) (*IDIDTransactor, error) {
	contract, err := bindIDID(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDIDTransactor{contract: contract}, nil
}

// NewIDIDFilterer creates a new log filterer instance of IDID, bound to a specific deployed contract.
func NewIDIDFilterer(address common.Address, filterer bind.ContractFilterer) (*IDIDFilterer, error) {
	contract, err := bindIDID(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDIDFilterer{contract: contract}, nil
}

// bindIDID binds a generic wrapper to an already deployed contract.
func bindIDID(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IDIDABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDID *IDIDRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IDID.Contract.IDIDCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDID *IDIDRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDID.Contract.IDIDTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDID *IDIDRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDID.Contract.IDIDTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDID *IDIDCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IDID.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDID *IDIDTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDID.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDID *IDIDTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDID.Contract.contract.Transact(opts, method, params...)
}

// GetHash is a free data retrieval call binding the contract method 0xb00140aa.
//
// Solidity: function getHash(bytes ) view returns(bytes32)
func (_IDID *IDIDCaller) GetHash(opts *bind.CallOpts, arg0 []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IDID.contract.Call(opts, out, "getHash", arg0)
	return *ret0, err
}

// GetHash is a free data retrieval call binding the contract method 0xb00140aa.
//
// Solidity: function getHash(bytes ) view returns(bytes32)
func (_IDID *IDIDSession) GetHash(arg0 []byte) ([32]byte, error) {
	return _IDID.Contract.GetHash(&_IDID.CallOpts, arg0)
}

// GetHash is a free data retrieval call binding the contract method 0xb00140aa.
//
// Solidity: function getHash(bytes ) view returns(bytes32)
func (_IDID *IDIDCallerSession) GetHash(arg0 []byte) ([32]byte, error) {
	return _IDID.Contract.GetHash(&_IDID.CallOpts, arg0)
}

// GetOwner is a free data retrieval call binding the contract method 0xb102bfbe.
//
// Solidity: function getOwner(bytes ) view returns(address)
func (_IDID *IDIDCaller) GetOwner(opts *bind.CallOpts, arg0 []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IDID.contract.Call(opts, out, "getOwner", arg0)
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0xb102bfbe.
//
// Solidity: function getOwner(bytes ) view returns(address)
func (_IDID *IDIDSession) GetOwner(arg0 []byte) (common.Address, error) {
	return _IDID.Contract.GetOwner(&_IDID.CallOpts, arg0)
}

// GetOwner is a free data retrieval call binding the contract method 0xb102bfbe.
//
// Solidity: function getOwner(bytes ) view returns(address)
func (_IDID *IDIDCallerSession) GetOwner(arg0 []byte) (common.Address, error) {
	return _IDID.Contract.GetOwner(&_IDID.CallOpts, arg0)
}

// GetURI is a free data retrieval call binding the contract method 0x8626dea9.
//
// Solidity: function getURI(bytes ) view returns(bytes)
func (_IDID *IDIDCaller) GetURI(opts *bind.CallOpts, arg0 []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _IDID.contract.Call(opts, out, "getURI", arg0)
	return *ret0, err
}

// GetURI is a free data retrieval call binding the contract method 0x8626dea9.
//
// Solidity: function getURI(bytes ) view returns(bytes)
func (_IDID *IDIDSession) GetURI(arg0 []byte) ([]byte, error) {
	return _IDID.Contract.GetURI(&_IDID.CallOpts, arg0)
}

// GetURI is a free data retrieval call binding the contract method 0x8626dea9.
//
// Solidity: function getURI(bytes ) view returns(bytes)
func (_IDID *IDIDCallerSession) GetURI(arg0 []byte) ([]byte, error) {
	return _IDID.Contract.GetURI(&_IDID.CallOpts, arg0)
}

// IDIDDIDCreatedIterator is returned from FilterDIDCreated and is used to iterate over the raw logs and unpacked data for DIDCreated events raised by the IDID contract.
type IDIDDIDCreatedIterator struct {
	Event *IDIDDIDCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IDIDDIDCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDIDDIDCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IDIDDIDCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IDIDDIDCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDIDDIDCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDIDDIDCreated represents a DIDCreated event raised by the IDID contract.
type IDIDDIDCreated struct {
	Operator common.Address
	Did      string
	Hash     [32]byte
	Uri      string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDIDCreated is a free log retrieval operation binding the contract event 0x0a36151ef7283a91a6a8d052cd44a0c20f8c7207decf22d985eba8eb0a5a268b.
//
// Solidity: event DIDCreated(address indexed operator, string did, bytes32 hash, string uri)
func (_IDID *IDIDFilterer) FilterDIDCreated(opts *bind.FilterOpts, operator []common.Address) (*IDIDDIDCreatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IDID.contract.FilterLogs(opts, "DIDCreated", operatorRule)
	if err != nil {
		return nil, err
	}
	return &IDIDDIDCreatedIterator{contract: _IDID.contract, event: "DIDCreated", logs: logs, sub: sub}, nil
}

// WatchDIDCreated is a free log subscription operation binding the contract event 0x0a36151ef7283a91a6a8d052cd44a0c20f8c7207decf22d985eba8eb0a5a268b.
//
// Solidity: event DIDCreated(address indexed operator, string did, bytes32 hash, string uri)
func (_IDID *IDIDFilterer) WatchDIDCreated(opts *bind.WatchOpts, sink chan<- *IDIDDIDCreated, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IDID.contract.WatchLogs(opts, "DIDCreated", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDIDDIDCreated)
				if err := _IDID.contract.UnpackLog(event, "DIDCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDIDCreated is a log parse operation binding the contract event 0x0a36151ef7283a91a6a8d052cd44a0c20f8c7207decf22d985eba8eb0a5a268b.
//
// Solidity: event DIDCreated(address indexed operator, string did, bytes32 hash, string uri)
func (_IDID *IDIDFilterer) ParseDIDCreated(log types.Log) (*IDIDDIDCreated, error) {
	event := new(IDIDDIDCreated)
	if err := _IDID.contract.UnpackLog(event, "DIDCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IDIDDIDDeletedIterator is returned from FilterDIDDeleted and is used to iterate over the raw logs and unpacked data for DIDDeleted events raised by the IDID contract.
type IDIDDIDDeletedIterator struct {
	Event *IDIDDIDDeleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IDIDDIDDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDIDDIDDeleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IDIDDIDDeleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IDIDDIDDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDIDDIDDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDIDDIDDeleted represents a DIDDeleted event raised by the IDID contract.
type IDIDDIDDeleted struct {
	Operator common.Address
	Did      common.Hash
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDIDDeleted is a free log retrieval operation binding the contract event 0x56c5173d83675a7f2b174113df9205152eb36c65ab39e41d5d1b8d9cb88eabd2.
//
// Solidity: event DIDDeleted(address indexed operator, string indexed did)
func (_IDID *IDIDFilterer) FilterDIDDeleted(opts *bind.FilterOpts, operator []common.Address, did []string) (*IDIDDIDDeletedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _IDID.contract.FilterLogs(opts, "DIDDeleted", operatorRule, didRule)
	if err != nil {
		return nil, err
	}
	return &IDIDDIDDeletedIterator{contract: _IDID.contract, event: "DIDDeleted", logs: logs, sub: sub}, nil
}

// WatchDIDDeleted is a free log subscription operation binding the contract event 0x56c5173d83675a7f2b174113df9205152eb36c65ab39e41d5d1b8d9cb88eabd2.
//
// Solidity: event DIDDeleted(address indexed operator, string indexed did)
func (_IDID *IDIDFilterer) WatchDIDDeleted(opts *bind.WatchOpts, sink chan<- *IDIDDIDDeleted, operator []common.Address, did []string) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _IDID.contract.WatchLogs(opts, "DIDDeleted", operatorRule, didRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDIDDIDDeleted)
				if err := _IDID.contract.UnpackLog(event, "DIDDeleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDIDDeleted is a log parse operation binding the contract event 0x56c5173d83675a7f2b174113df9205152eb36c65ab39e41d5d1b8d9cb88eabd2.
//
// Solidity: event DIDDeleted(address indexed operator, string indexed did)
func (_IDID *IDIDFilterer) ParseDIDDeleted(log types.Log) (*IDIDDIDDeleted, error) {
	event := new(IDIDDIDDeleted)
	if err := _IDID.contract.UnpackLog(event, "DIDDeleted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IDIDDIDUpdatedIterator is returned from FilterDIDUpdated and is used to iterate over the raw logs and unpacked data for DIDUpdated events raised by the IDID contract.
type IDIDDIDUpdatedIterator struct {
	Event *IDIDDIDUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IDIDDIDUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDIDDIDUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IDIDDIDUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IDIDDIDUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDIDDIDUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDIDDIDUpdated represents a DIDUpdated event raised by the IDID contract.
type IDIDDIDUpdated struct {
	Operator common.Address
	Did      common.Hash
	Hash     [32]byte
	Uri      string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDIDUpdated is a free log retrieval operation binding the contract event 0x30665598c4da8be2cccefb7fe32a379decd1eae407564569a8e95b51e3afccd7.
//
// Solidity: event DIDUpdated(address indexed operator, string indexed did, bytes32 hash, string uri)
func (_IDID *IDIDFilterer) FilterDIDUpdated(opts *bind.FilterOpts, operator []common.Address, did []string) (*IDIDDIDUpdatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _IDID.contract.FilterLogs(opts, "DIDUpdated", operatorRule, didRule)
	if err != nil {
		return nil, err
	}
	return &IDIDDIDUpdatedIterator{contract: _IDID.contract, event: "DIDUpdated", logs: logs, sub: sub}, nil
}

// WatchDIDUpdated is a free log subscription operation binding the contract event 0x30665598c4da8be2cccefb7fe32a379decd1eae407564569a8e95b51e3afccd7.
//
// Solidity: event DIDUpdated(address indexed operator, string indexed did, bytes32 hash, string uri)
func (_IDID *IDIDFilterer) WatchDIDUpdated(opts *bind.WatchOpts, sink chan<- *IDIDDIDUpdated, operator []common.Address, did []string) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _IDID.contract.WatchLogs(opts, "DIDUpdated", operatorRule, didRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDIDDIDUpdated)
				if err := _IDID.contract.UnpackLog(event, "DIDUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDIDUpdated is a log parse operation binding the contract event 0x30665598c4da8be2cccefb7fe32a379decd1eae407564569a8e95b51e3afccd7.
//
// Solidity: event DIDUpdated(address indexed operator, string indexed did, bytes32 hash, string uri)
func (_IDID *IDIDFilterer) ParseDIDUpdated(log types.Log) (*IDIDDIDUpdated, error) {
	event := new(IDIDDIDUpdated)
	if err := _IDID.contract.UnpackLog(event, "DIDUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = map[string]string{
	"8da5cb5b": "owner()",
	"f2fde38b": "transferOwnership(address)",
}

// OwnableBin is the compiled bytecode used for deploying new contracts.
var OwnableBin = "0x608060405234801561001057600080fd5b50600080546001600160a01b03191633179055610150806100326000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80638da5cb5b1461003b578063f2fde38b1461005f575b600080fd5b610043610087565b604080516001600160a01b039092168252519081900360200190f35b6100856004803603602081101561007557600080fd5b50356001600160a01b0316610096565b005b6000546001600160a01b031681565b6000546001600160a01b031633146100ad57600080fd5b6001600160a01b0381166100c057600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b039290921691909117905556fea265627a7a723158205f7ceb1d0a9c3dbe921137faeb9d09d890873f2bd492d41e4642341fbefeca4764736f6c63430005100032"

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OwnableOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// UCamDIDManagerABI is the input ABI used to generate the binding from.
const UCamDIDManagerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dbAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"DIDCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"DIDDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"DIDUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"uid\",\"type\":\"bytes20\"},{\"internalType\":\"bytes32\",\"name\":\"h\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"uri\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"authorizer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"auth\",\"type\":\"bytes\"}],\"name\":\"createDIDByAgent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"db\",\"outputs\":[{\"internalType\":\"contractDIDStorage\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"did\",\"type\":\"bytes\"}],\"name\":\"decodeInternalKey\",\"outputs\":[{\"internalType\":\"bytes20\",\"name\":\"\",\"type\":\"bytes20\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"uid\",\"type\":\"bytes20\"},{\"internalType\":\"bytes\",\"name\":\"auth\",\"type\":\"bytes\"}],\"name\":\"deleteDIDByAgent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"did\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"h\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"uri\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"getCreateAuthMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"did\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"getDeleteAuthMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"did\",\"type\":\"bytes\"}],\"name\":\"getHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"did\",\"type\":\"bytes\"}],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"did\",\"type\":\"bytes\"}],\"name\":\"getURI\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"did\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"h\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"uri\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"getUpdateAuthMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferDBOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"uid\",\"type\":\"bytes20\"},{\"internalType\":\"bytes32\",\"name\":\"h\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"uri\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"auth\",\"type\":\"bytes\"}],\"name\":\"updateDIDByAgent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// UCamDIDManagerFuncSigs maps the 4-byte function signature to its string representation.
var UCamDIDManagerFuncSigs = map[string]string{
	"3dc7ad02": "createDIDByAgent(bytes20,bytes32,bytes,address,bytes)",
	"4d655aff": "db()",
	"bf6d3139": "decodeInternalKey(bytes)",
	"d9d50607": "deleteDIDByAgent(bytes20,bytes)",
	"cc3b41dc": "getCreateAuthMessage(bytes,bytes32,bytes,address)",
	"3bdcb26c": "getDeleteAuthMessage(bytes,address)",
	"b00140aa": "getHash(bytes)",
	"b102bfbe": "getOwner(bytes)",
	"8626dea9": "getURI(bytes)",
	"dd120153": "getUpdateAuthMessage(bytes,bytes32,bytes,address)",
	"8da5cb5b": "owner()",
	"32aad017": "transferDBOwnership(address)",
	"f2fde38b": "transferOwnership(address)",
	"98793984": "updateDIDByAgent(bytes20,bytes32,bytes,bytes)",
}

// UCamDIDManagerBin is the compiled bytecode used for deploying new contracts.
var UCamDIDManagerBin = "0x60806040523480156200001157600080fd5b5060405162003c4a38038062003c4a833981810160405260208110156200003757600080fd5b505160408051808201909152600c81526b3234b21d34b79d3ab1b0b69d60a11b6020820152600080546001600160a01b0319163317905581906001600160a01b0382166200013d57806040516200008e9062000162565b60208082528251818301528251829160408301919085019080838360005b83811015620000c6578181015183820152602001620000ac565b50505050905090810190601f168015620000f45780820380516001836020036101000a031916815260200191505b5092505050604051809103906000f08015801562000116573d6000803e3d6000fd5b50600180546001600160a01b0319166001600160a01b039290921691909117905562000159565b600180546001600160a01b0319166001600160a01b0384161790555b50505062000170565b6109bb806200328f83390190565b61310f80620001806000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c8063b00140aa1161008c578063cc3b41dc11610066578063cc3b41dc146107b5578063d9d50607146108f1578063dd120153146109a6578063f2fde38b14610ae2576100ea565b8063b00140aa1461059a578063b102bfbe14610650578063bf6d3139146106f4576100ea565b80634d655aff116100c85780634d655aff1461038b5780638626dea9146103af5780638da5cb5b14610453578063987939841461045b576100ea565b806332aad017146100ef5780633bdcb26c146101175780633dc7ad021461023b575b600080fd5b6101156004803603602081101561010557600080fd5b50356001600160a01b0316610b08565b005b6101c66004803603604081101561012d57600080fd5b810190602081018135600160201b81111561014757600080fd5b82018360208201111561015957600080fd5b803590602001918460018302840111600160201b8311171561017a57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160a01b03169150610b889050565b6040805160208082528351818301528351919283929083019185019080838360005b838110156102005781810151838201526020016101e8565b50505050905090810190601f16801561022d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610115600480360360a081101561025157600080fd5b6001600160601b031982351691602081013591810190606081016040820135600160201b81111561028157600080fd5b82018360208201111561029357600080fd5b803590602001918460018302840111600160201b831117156102b457600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092956001600160a01b03853516959094909350604081019250602001359050600160201b81111561031757600080fd5b82018360208201111561032957600080fd5b803590602001918460018302840111600160201b8311171561034a57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610ce9945050505050565b610393610d79565b604080516001600160a01b039092168252519081900360200190f35b6101c6600480360360208110156103c557600080fd5b810190602081018135600160201b8111156103df57600080fd5b8201836020820111156103f157600080fd5b803590602001918460018302840111600160201b8311171561041257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610d88945050505050565b610393610fb7565b6101156004803603608081101561047157600080fd5b6001600160601b031982351691602081013591810190606081016040820135600160201b8111156104a157600080fd5b8201836020820111156104b357600080fd5b803590602001918460018302840111600160201b831117156104d457600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561052657600080fd5b82018360208201111561053857600080fd5b803590602001918460018302840111600160201b8311171561055957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610fc6945050505050565b61063e600480360360208110156105b057600080fd5b810190602081018135600160201b8111156105ca57600080fd5b8201836020820111156105dc57600080fd5b803590602001918460018302840111600160201b831117156105fd57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506111a0945050505050565b60408051918252519081900360200190f35b6103936004803603602081101561066657600080fd5b810190602081018135600160201b81111561068057600080fd5b82018360208201111561069257600080fd5b803590602001918460018302840111600160201b831117156106b357600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506113cb945050505050565b6107986004803603602081101561070a57600080fd5b810190602081018135600160201b81111561072457600080fd5b82018360208201111561073657600080fd5b803590602001918460018302840111600160201b8311171561075757600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506115f6945050505050565b604080516001600160601b03199092168252519081900360200190f35b6101c6600480360360808110156107cb57600080fd5b810190602081018135600160201b8111156107e557600080fd5b8201836020820111156107f757600080fd5b803590602001918460018302840111600160201b8311171561081857600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092958435959094909350604081019250602001359050600160201b81111561087257600080fd5b82018360208201111561088457600080fd5b803590602001918460018302840111600160201b831117156108a557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160a01b031691506119449050565b6101156004803603604081101561090757600080fd5b6001600160601b03198235169190810190604081016020820135600160201b81111561093257600080fd5b82018360208201111561094457600080fd5b803590602001918460018302840111600160201b8311171561096557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611b23945050505050565b6101c6600480360360808110156109bc57600080fd5b810190602081018135600160201b8111156109d657600080fd5b8201836020820111156109e857600080fd5b803590602001918460018302840111600160201b83111715610a0957600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092958435959094909350604081019250602001359050600160201b811115610a6357600080fd5b820183602082011115610a7557600080fd5b803590602001918460018302840111600160201b83111715610a9657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160a01b03169150611cf99050565b61011560048036036020811015610af857600080fd5b50356001600160a01b0316611e22565b6000546001600160a01b03163314610b1f57600080fd5b6001546040805163f2fde38b60e01b81526001600160a01b0384811660048301529151919092169163f2fde38b91602480830192600092919082900301818387803b158015610b6d57600080fd5b505af1158015610b81573d6000803e3d6000fd5b5050505050565b6060610b9382611ea7565b83610b9d30611ea7565b60405160200180806b0249030baba3437b934bd32960a51b815250600c0184805190602001908083835b60208310610be65780518252601f199092019160209182019101610bc7565b51815160209384036101000a60001901801990921691161790526e0103a37903232b632ba32902224a21608d1b919093019081528551600f90910192860191508083835b60208310610c495780518252601f199092019160209182019101610c2a565b51815160209384036101000a60001901801990921691161790526c01034b71031b7b73a3930b1ba1609d1b919093019081528451600d90910192850191508083835b60208310610caa5780518252601f199092019160209182019101610c8b565b6001836020036101000a038019825116818451168082178552505050505050905001935050505060405160208183030381529060405290505b92915050565b6060610cf486612023565b9050610d0b610d0582878733611944565b836121d5565b6001600160a01b0316836001600160a01b031614610d64576040805162461bcd60e51b8152602060048201526011602482015270696e76616c6964207369676e617475726560781b604482015290519081900360640190fd5b610d7181878588886122d0565b505050505050565b6001546001600160a01b031681565b60606000610d95836115f6565b60015460408051632187057760e11b81526001600160601b03198416600482015290519293506001600160a01b039091169163430e0aee91602480820192602092909190829003018186803b158015610ded57600080fd5b505afa158015610e01573d6000803e3d6000fd5b505050506040513d6020811015610e1757600080fd5b5051610e5f576040805162461bcd60e51b815260206004820152601260248201527111125108191bd95cc81b9bdd08195e1a5cdd60721b604482015290519081900360640190fd5b60015460408051630b59d98f60e31b81526001600160601b03198416600482015290516060926001600160a01b031691635acecc78916024808301926000929190829003018186803b158015610eb457600080fd5b505afa158015610ec8573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526060811015610ef157600080fd5b81516020830151604080850180519151939592948301929184600160201b821115610f1b57600080fd5b908301906020820185811115610f3057600080fd5b8251600160201b811182820188101715610f4957600080fd5b82525081516020918201929091019080838360005b83811015610f76578181015183820152602001610f5e565b50505050905090810190601f168015610fa35780820380516001836020036101000a031916815260200191505b50604052509196505050505050505b919050565b6000546001600160a01b031681565b6060610fd185612023565b60015460408051630b59d98f60e31b81526001600160601b03198916600482015290519293506000926001600160a01b0390921691635acecc78916024808201928692909190829003018186803b15801561102b57600080fd5b505afa15801561103f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052606081101561106857600080fd5b81516020830151604080850180519151939592948301929184600160201b82111561109257600080fd5b9083019060208201858111156110a757600080fd5b8251600160201b8111828201881017156110c057600080fd5b82525081516020918201929091019080838360005b838110156110ed5781810151838201526020016110d5565b50505050905090810190601f16801561111a5780820380516001836020036101000a031916815260200191505b506040525050505050905061113a61113483878733611cf9565b846121d5565b6001600160a01b0316816001600160a01b031614611193576040805162461bcd60e51b8152602060048201526011602482015270696e76616c6964207369676e617475726560781b604482015290519081900360640190fd5b610d7182878388886125a6565b6000806111ac836115f6565b60015460408051632187057760e11b81526001600160601b03198416600482015290519293506001600160a01b039091169163430e0aee91602480820192602092909190829003018186803b15801561120457600080fd5b505afa158015611218573d6000803e3d6000fd5b505050506040513d602081101561122e57600080fd5b5051611276576040805162461bcd60e51b815260206004820152601260248201527111125108191bd95cc81b9bdd08195e1a5cdd60721b604482015290519081900360640190fd5b60015460408051630b59d98f60e31b81526001600160601b03198416600482015290516000926001600160a01b031691635acecc789160248083019286929190829003018186803b1580156112ca57600080fd5b505afa1580156112de573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052606081101561130757600080fd5b81516020830151604080850180519151939592948301929184600160201b82111561133157600080fd5b90830190602082018581111561134657600080fd5b8251600160201b81118282018810171561135f57600080fd5b82525081516020918201929091019080838360005b8381101561138c578181015183820152602001611374565b50505050905090810190601f1680156113b95780820380516001836020036101000a031916815260200191505b50604052509298975050505050505050565b6000806113d7836115f6565b60015460408051632187057760e11b81526001600160601b03198416600482015290519293506001600160a01b039091169163430e0aee91602480820192602092909190829003018186803b15801561142f57600080fd5b505afa158015611443573d6000803e3d6000fd5b505050506040513d602081101561145957600080fd5b50516114a1576040805162461bcd60e51b815260206004820152601260248201527111125108191bd95cc81b9bdd08195e1a5cdd60721b604482015290519081900360640190fd5b60015460408051630b59d98f60e31b81526001600160601b03198416600482015290516000926001600160a01b031691635acecc789160248083019286929190829003018186803b1580156114f557600080fd5b505afa158015611509573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052606081101561153257600080fd5b81516020830151604080850180519151939592948301929184600160201b82111561155c57600080fd5b90830190602082018581111561157157600080fd5b8251600160201b81118282018810171561158a57600080fd5b82525081516020918201929091019080838360005b838110156115b757818101518382015260200161159f565b50505050905090810190601f1680156115e45780820380516001836020036101000a031916815260200191505b50604052509398975050505050505050565b600061173a82600160009054906101000a90046001600160a01b03166001600160a01b03166312d05d3d6040518163ffffffff1660e01b815260040160006040518083038186803b15801561164a57600080fd5b505afa15801561165e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561168757600080fd5b8101908080516040519392919084600160201b8211156116a657600080fd5b9083019060208201858111156116bb57600080fd5b8251600160201b8111828201881017156116d457600080fd5b82525081516020918201929091019080838360005b838110156117015781810151838201526020016116e9565b50505050905090810190601f16801561172e5780820380516001836020036101000a031916815260200191505b50604052505050612a1b565b611779576040805162461bcd60e51b815260206004820152600b60248201526a1a5b9d985b1a590811125160aa1b604482015290519081900360640190fd5b60606118be83600160009054906101000a90046001600160a01b03166001600160a01b03166312d05d3d6040518163ffffffff1660e01b815260040160006040518083038186803b1580156117cd57600080fd5b505afa1580156117e1573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561180a57600080fd5b8101908080516040519392919084600160201b82111561182957600080fd5b90830190602082018581111561183e57600080fd5b8251600160201b81118282018810171561185757600080fd5b82525081516020918201929091019080838360005b8381101561188457818101518382015260200161186c565b50505050905090810190601f1680156118b15780820380516001836020036101000a031916815260200191505b5060405250505051612a97565b90508051601414611904576040805162461bcd60e51b815260206004820152600b60248201526a1a5b9d985b1a590811125160aa1b604482015290519081900360640190fd5b6000805b6014811015611939576101008202915082818151811061192457fe5b016020015160f81c9190910190600101611908565b5060601b9392505050565b606061194f82611ea7565b8561195930611ea7565b868660405160200180806b0249030baba3437b934bd32960a51b815250600c0186805190602001908083835b602083106119a45780518252601f199092019160209182019101611985565b51815160209384036101000a60001901801990921691161790526e0103a379031b932b0ba32902224a21608d1b919093019081528751600f90910192880191508083835b60208310611a075780518252601f1990920191602091820191016119e8565b51815160209384036101000a60001901801990921691161790527101034b71031b7b73a3930b1ba103bb4ba34160751b919093019081528651601290910192870191508083835b60208310611a6d5780518252601f199092019160209182019101611a4e565b51815160209384036101000a600019018019909216911617905261040560f31b919093019081526002810186905261016160f51b60228201528451602490910192850191508083835b60208310611ad55780518252601f199092019160209182019101611ab6565b6001836020036101000a03801982511681845116808217855250505050505090500180602960f81b815250600101955050505050506040516020818303038152906040529050949350505050565b6060611b2e83612023565b60015460408051630b59d98f60e31b81526001600160601b03198716600482015290519293506000926001600160a01b0390921691635acecc78916024808201928692909190829003018186803b158015611b8857600080fd5b505afa158015611b9c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526060811015611bc557600080fd5b81516020830151604080850180519151939592948301929184600160201b821115611bef57600080fd5b908301906020820185811115611c0457600080fd5b8251600160201b811182820188101715611c1d57600080fd5b82525081516020918201929091019080838360005b83811015611c4a578181015183820152602001611c32565b50505050905090810190601f168015611c775780820380516001836020036101000a031916815260200191505b5060405250505050509050611c8f6111348333610b88565b6001600160a01b0316816001600160a01b031614611ce8576040805162461bcd60e51b8152602060048201526011602482015270696e76616c6964207369676e617475726560781b604482015290519081900360640190fd5b611cf3828583612b56565b50505050565b6060611d0482611ea7565b85611d0e30611ea7565b868660405160200180806b0249030baba3437b934bd32960a51b815250600c0186805190602001908083835b60208310611d595780518252601f199092019160209182019101611d3a565b51815160209384036101000a60001901801990921691161790526e0103a37903ab83230ba32902224a21608d1b919093019081528751600f90910192880191508083835b60208310611dbc5780518252601f199092019160209182019101611d9d565b51815160001960209485036101000a019081169019919091161790526f01034b71031b7b73a3930b1ba103a37960851b939091019283528651601090930192908701915080838360208310611a6d5780518252601f199092019160209182019101611a4e565b6000546001600160a01b03163314611e3957600080fd5b6001600160a01b038116611e4c57600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b604080518082018252601081526f181899199a1a9b1b9c1cb0b131b232b360811b60208201528151602a80825260608281019094526001600160a01b03851692918491602082018180388339019050509050600360fc1b81600081518110611f0b57fe5b60200101906001600160f81b031916908160001a905350600f60fb1b81600181518110611f3457fe5b60200101906001600160f81b031916908160001a90535060005b601481101561201a578260048583600c0160208110611f6957fe5b1a60f81b6001600160f81b031916901c60f81c60ff1681518110611f8957fe5b602001015160f81c60f81b828260020260020181518110611fa657fe5b60200101906001600160f81b031916908160001a905350828482600c0160208110611fcd57fe5b825191901a600f16908110611fde57fe5b602001015160f81c60f81b828260020260030181518110611ffb57fe5b60200101906001600160f81b031916908160001a905350600101611f4e565b50949350505050565b600154604080516312d05d3d60e01b815290516060926001600160a01b0316916312d05d3d916004808301926000929190829003018186803b15801561206857600080fd5b505afa15801561207c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260208110156120a557600080fd5b8101908080516040519392919084600160201b8211156120c457600080fd5b9083019060208201858111156120d957600080fd5b8251600160201b8111828201881017156120f257600080fd5b82525081516020918201929091019080838360005b8381101561211f578181015183820152602001612107565b50505050905090810190601f16801561214c5780820380516001836020036101000a031916815260200191505b50604052505050826040516020018083805190602001908083835b602083106121865780518252601f199092019160209182019101612167565b5181516020939093036101000a60001901801990911692169190911790526001600160601b0319949094169190930190815260408051808303600b190181526014909201905295945050505050565b60006122c96121e48451612f50565b8460405160200180807f19457468657265756d205369676e6564204d6573736167653a0a000000000000815250601a0183805190602001908083835b6020831061223f5780518252601f199092019160209182019101612220565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106122875780518252601f199092019160209182019101612268565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040528051906020012083613008565b9392505050565b60015460408051632187057760e11b81526001600160601b03198716600482015290516001600160a01b039092169163430e0aee91602480820192602092909190829003018186803b15801561232557600080fd5b505afa158015612339573d6000803e3d6000fd5b505050506040513d602081101561234f57600080fd5b505115612393576040805162461bcd60e51b815260206004820152600d60248201526c191d5c1b1a58d85d1948111251609a1b604482015290519081900360640190fd5b600180546040516348cab05b60e01b81526001600160601b0319871660048201908152602482018490526001600160a01b0387811660448401526064830187905260a060848401908152865160a4850152865191909416946348cab05b948a9491938a938a938a939160c490910190602085019080838360005b8381101561242557818101518382015260200161240d565b50505050905090810190601f1680156124525780820380516001836020036101000a031916815260200191505b509650505050505050600060405180830381600087803b15801561247557600080fd5b505af1158015612489573d6000803e3d6000fd5b50505050826001600160a01b03167f0a36151ef7283a91a6a8d052cd44a0c20f8c7207decf22d985eba8eb0a5a268b868484604051808060200184815260200180602001838103835286818151815260200191508051906020019080838360005b838110156125025781810151838201526020016124ea565b50505050905090810190601f16801561252f5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b8381101561256257818101518382015260200161254a565b50505050905090810190601f16801561258f5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a25050505050565b60015460408051632187057760e11b81526001600160601b03198716600482015290516001600160a01b039092169163430e0aee91602480820192602092909190829003018186803b1580156125fb57600080fd5b505afa15801561260f573d6000803e3d6000fd5b505050506040513d602081101561262557600080fd5b505161266d576040805162461bcd60e51b815260206004820152601260248201527111125108191bd95cc81b9bdd08195e1a5cdd60721b604482015290519081900360640190fd5b60015460408051630b59d98f60e31b81526001600160601b03198716600482015290516000926001600160a01b031691635acecc789160248083019286929190829003018186803b1580156126c157600080fd5b505afa1580156126d5573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260608110156126fe57600080fd5b81516020830151604080850180519151939592948301929184600160201b82111561272857600080fd5b90830190602082018581111561273d57600080fd5b8251600160201b81118282018810171561275657600080fd5b82525081516020918201929091019080838360005b8381101561278357818101518382015260200161276b565b50505050905090810190601f1680156127b05780820380516001836020036101000a031916815260200191505b5060405250505050509050836001600160a01b0316816001600160a01b031614612811576040805162461bcd60e51b815260206004820152600d60248201526c3737903832b936b4b9b9b4b7b760991b604482015290519081900360640190fd5b600180546040516348cab05b60e01b81526001600160601b0319881660048201908152602482018490526001600160a01b0388811660448401526064830188905260a060848401908152875160a4850152875191909416946348cab05b948b9491938b938b938b939160c490910190602085019080838360005b838110156128a357818101518382015260200161288b565b50505050905090810190601f1680156128d05780820380516001836020036101000a031916815260200191505b509650505050505050600060405180830381600087803b1580156128f357600080fd5b505af1158015612907573d6000803e3d6000fd5b50505050856040518082805190602001908083835b6020831061293b5780518252601f19909201916020918201910161291c565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020846001600160a01b03167f30665598c4da8be2cccefb7fe32a379decd1eae407564569a8e95b51e3afccd785856040518083815260200180602001828103825283818151815260200191508051906020019080838360005b838110156129d85781810151838201526020016129c0565b50505050905090810190601f168015612a055780820380516001836020036101000a031916815260200191505b50935050505060405180910390a3505050505050565b60008151835111612a2e57506000610ce3565b60005b8251811015612a8d57838181518110612a4657fe5b602001015160f81c60f81b6001600160f81b031916838281518110612a6757fe5b01602001516001600160f81b03191614612a85576000915050610ce3565b600101612a31565b5060019392505050565b60608183511015612ae4576040805162461bcd60e51b815260206004820152601260248201527152656164206f7574206f6620626f756e647360701b604482015290519081900360640190fd5b8251829003606081158015612b045760405191506020820160405261201a565b6040519150601f8316801560200281840101848101878315602002848b0101015b81831015612b3d578051835260209283019201612b25565b5050848452601f01601f19166040525050949350505050565b60015460408051632187057760e11b81526001600160601b03198516600482015290516001600160a01b039092169163430e0aee91602480820192602092909190829003018186803b158015612bab57600080fd5b505afa158015612bbf573d6000803e3d6000fd5b505050506040513d6020811015612bd557600080fd5b5051612c1d576040805162461bcd60e51b815260206004820152601260248201527111125108191bd95cc81b9bdd08195e1a5cdd60721b604482015290519081900360640190fd5b60015460408051630b59d98f60e31b81526001600160601b031985166004820152905160009283926060926001600160a01b0390921691635acecc78916024808201928792909190829003018186803b158015612c7957600080fd5b505afa158015612c8d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526060811015612cb657600080fd5b81516020830151604080850180519151939592948301929184600160201b821115612ce057600080fd5b908301906020820185811115612cf557600080fd5b8251600160201b811182820188101715612d0e57600080fd5b82525081516020918201929091019080838360005b83811015612d3b578181015183820152602001612d23565b50505050905090810190601f168015612d685780820380516001836020036101000a031916815260200191505b50604052505050925092509250836001600160a01b0316836001600160a01b031614612dcb576040805162461bcd60e51b815260206004820152600d60248201526c3737903832b936b4b9b9b4b7b760991b604482015290519081900360640190fd5b6001546040516348cab05b60e01b81526001600160601b03198716600482019081526000602483018190526001600160a01b0388811660448501526064840187905260a060848501908152865160a4860152865191909516946348cab05b948b948b938a938a93909160c49091019060208501908083838b5b83811015612e5c578181015183820152602001612e44565b50505050905090810190601f168015612e895780820380516001836020036101000a031916815260200191505b509650505050505050600060405180830381600087803b158015612eac57600080fd5b505af1158015612ec0573d6000803e3d6000fd5b50505050856040518082805190602001908083835b60208310612ef45780518252601f199092019160209182019101612ed5565b5181516020939093036101000a60001901801990911692169190911790526040519201829003822093503392507f56c5173d83675a7f2b174113df9205152eb36c65ab39e41d5d1b8d9cb88eabd29160009150a3505050505050565b606081612f7557506040805180820190915260018152600360fc1b6020820152610fb2565b8160005b8115612f8d57600101600a82049150612f79565b6060816040519080825280601f01601f191660200182016040528015612fba576020820181803883390190505b50905060001982015b851561201a57600a860660300160f81b82828060019003935081518110612fe657fe5b60200101906001600160f81b031916908160001a905350600a86049550612fc3565b60008060008084516041146130235760009350505050610ce3565b50505060208201516040830151606084015160001a601b81101561304557601b015b8060ff16601b1415801561305d57508060ff16601c14155b1561306e5760009350505050610ce3565b6040805160008152602080820180845289905260ff8416828401526060820186905260808201859052915160019260a0808401939192601f1981019281900390910190855afa1580156130c5573d6000803e3d6000fd5b5050604051601f19015197965050505050505056fea265627a7a72315820e068f77dc519ef0005232b12c18d7a1bd973f80b80134d90b07bbea9471a968c64736f6c63430005100032608060405234801561001057600080fd5b506040516109bb3803806109bb8339818101604052602081101561003357600080fd5b810190808051604051939291908464010000000082111561005357600080fd5b90830190602082018581111561006857600080fd5b825164010000000081118282018810171561008257600080fd5b82525081516020918201929091019080838360005b838110156100af578181015183820152602001610097565b50505050905090810190601f1680156100dc5780820380516001836020036101000a031916815260200191505b506040525050600080546001600160a01b0319163317905550610107816001600160e01b0361010d16565b506101d6565b6000546001600160a01b0316331461012457600080fd5b805161013790600190602084019061013b565b5050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061017c57805160ff19168380011785556101a9565b828001600101855582156101a9579182015b828111156101a957825182559160200191906001019061018e565b506101b59291506101b9565b5090565b6101d391905b808211156101b557600081556001016101bf565b90565b6107d6806101e56000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80634f1531451161005b5780634f1531451461020f5780635acecc78146102b55780638da5cb5b14610374578063f2fde38b146103985761007d565b806312d05d3d14610082578063430e0aee146100ff57806348cab05b1461013a575b600080fd5b61008a6103be565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100c45781810151838201526020016100ac565b50505050905090810190601f1680156100f15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101266004803603602081101561011557600080fd5b50356001600160601b031916610454565b604080519115158252519081900360200190f35b61020d600480360360a081101561015057600080fd5b6001600160601b03198235169160208101351515916001600160a01b036040830135169160608101359181019060a08101608082013564010000000081111561019857600080fd5b8201836020820111156101aa57600080fd5b803590602001918460018302840111640100000000831117156101cc57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610473945050505050565b005b61020d6004803603602081101561022557600080fd5b81019060208101813564010000000081111561024057600080fd5b82018360208201111561025257600080fd5b8035906020019184600183028401116401000000008311171561027457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610527945050505050565b6102dc600480360360208110156102cb57600080fd5b50356001600160601b031916610555565b60405180846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561033757818101518382015260200161031f565b50505050905090810190601f1680156103645780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b61037c610675565b604080516001600160a01b039092168252519081900360200190f35b61020d600480360360208110156103ae57600080fd5b50356001600160a01b0316610684565b60018054604080516020601f600260001961010087891615020190951694909404938401819004810282018101909252828152606093909290918301828280156104495780601f1061041e57610100808354040283529160200191610449565b820191906000526020600020905b81548152906001019060200180831161042c57829003601f168201915b505050505090505b90565b6001600160601b03191660009081526002602052604090205460ff1690565b6000546001600160a01b0316331461048a57600080fd5b6040805160808101825285151581526001600160a01b038581166020808401918252838501878152606085018781526001600160601b03198c166000908152600280855297902086518154955160ff1990961690151517610100600160a81b0319166101009590961694909402949094178355516001830155915180519394919361051d93928501929190910190610709565b5050505050505050565b6000546001600160a01b0316331461053e57600080fd5b8051610551906001906020840190610709565b5050565b600080606061056384610454565b6105a9576040805162461bcd60e51b815260206004820152601260248201527111125108191bd95cc81b9bdd08195e1a5cdd60721b604482015290519081900360640190fd5b6001600160601b031984166000908152600260208181526040928390208054600180830154928501805487516101009382161584026000190190911696909604601f810186900486028701860190975286865291046001600160a01b03169491939092909183918301828280156106615780601f1061063657610100808354040283529160200191610661565b820191906000526020600020905b81548152906001019060200180831161064457829003601f168201915b505050505090509250925092509193909250565b6000546001600160a01b031681565b6000546001600160a01b0316331461069b57600080fd5b6001600160a01b0381166106ae57600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061074a57805160ff1916838001178555610777565b82800160010185558215610777579182015b8281111561077757825182559160200191906001019061075c565b50610783929150610787565b5090565b61045191905b80821115610783576000815560010161078d56fea265627a7a72315820f3c2c8af7f5e3f24ba7076e122ee0d78744d15be9914a9df3341e14fe3054aa664736f6c63430005100032"

// DeployUCamDIDManager deploys a new Ethereum contract, binding an instance of UCamDIDManager to it.
func DeployUCamDIDManager(auth *bind.TransactOpts, backend bind.ContractBackend, _dbAddr common.Address) (common.Address, *types.Transaction, *UCamDIDManager, error) {
	parsed, err := abi.JSON(strings.NewReader(UCamDIDManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UCamDIDManagerBin), backend, _dbAddr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UCamDIDManager{UCamDIDManagerCaller: UCamDIDManagerCaller{contract: contract}, UCamDIDManagerTransactor: UCamDIDManagerTransactor{contract: contract}, UCamDIDManagerFilterer: UCamDIDManagerFilterer{contract: contract}}, nil
}

// UCamDIDManager is an auto generated Go binding around an Ethereum contract.
type UCamDIDManager struct {
	UCamDIDManagerCaller     // Read-only binding to the contract
	UCamDIDManagerTransactor // Write-only binding to the contract
	UCamDIDManagerFilterer   // Log filterer for contract events
}

// UCamDIDManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type UCamDIDManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UCamDIDManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UCamDIDManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UCamDIDManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UCamDIDManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UCamDIDManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UCamDIDManagerSession struct {
	Contract     *UCamDIDManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UCamDIDManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UCamDIDManagerCallerSession struct {
	Contract *UCamDIDManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// UCamDIDManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UCamDIDManagerTransactorSession struct {
	Contract     *UCamDIDManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// UCamDIDManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type UCamDIDManagerRaw struct {
	Contract *UCamDIDManager // Generic contract binding to access the raw methods on
}

// UCamDIDManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UCamDIDManagerCallerRaw struct {
	Contract *UCamDIDManagerCaller // Generic read-only contract binding to access the raw methods on
}

// UCamDIDManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UCamDIDManagerTransactorRaw struct {
	Contract *UCamDIDManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUCamDIDManager creates a new instance of UCamDIDManager, bound to a specific deployed contract.
func NewUCamDIDManager(address common.Address, backend bind.ContractBackend) (*UCamDIDManager, error) {
	contract, err := bindUCamDIDManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UCamDIDManager{UCamDIDManagerCaller: UCamDIDManagerCaller{contract: contract}, UCamDIDManagerTransactor: UCamDIDManagerTransactor{contract: contract}, UCamDIDManagerFilterer: UCamDIDManagerFilterer{contract: contract}}, nil
}

// NewUCamDIDManagerCaller creates a new read-only instance of UCamDIDManager, bound to a specific deployed contract.
func NewUCamDIDManagerCaller(address common.Address, caller bind.ContractCaller) (*UCamDIDManagerCaller, error) {
	contract, err := bindUCamDIDManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UCamDIDManagerCaller{contract: contract}, nil
}

// NewUCamDIDManagerTransactor creates a new write-only instance of UCamDIDManager, bound to a specific deployed contract.
func NewUCamDIDManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*UCamDIDManagerTransactor, error) {
	contract, err := bindUCamDIDManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UCamDIDManagerTransactor{contract: contract}, nil
}

// NewUCamDIDManagerFilterer creates a new log filterer instance of UCamDIDManager, bound to a specific deployed contract.
func NewUCamDIDManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*UCamDIDManagerFilterer, error) {
	contract, err := bindUCamDIDManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UCamDIDManagerFilterer{contract: contract}, nil
}

// bindUCamDIDManager binds a generic wrapper to an already deployed contract.
func bindUCamDIDManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UCamDIDManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UCamDIDManager *UCamDIDManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UCamDIDManager.Contract.UCamDIDManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UCamDIDManager *UCamDIDManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UCamDIDManager.Contract.UCamDIDManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UCamDIDManager *UCamDIDManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UCamDIDManager.Contract.UCamDIDManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UCamDIDManager *UCamDIDManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UCamDIDManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UCamDIDManager *UCamDIDManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UCamDIDManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UCamDIDManager *UCamDIDManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UCamDIDManager.Contract.contract.Transact(opts, method, params...)
}

// Db is a free data retrieval call binding the contract method 0x4d655aff.
//
// Solidity: function db() view returns(address)
func (_UCamDIDManager *UCamDIDManagerCaller) Db(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UCamDIDManager.contract.Call(opts, out, "db")
	return *ret0, err
}

// Db is a free data retrieval call binding the contract method 0x4d655aff.
//
// Solidity: function db() view returns(address)
func (_UCamDIDManager *UCamDIDManagerSession) Db() (common.Address, error) {
	return _UCamDIDManager.Contract.Db(&_UCamDIDManager.CallOpts)
}

// Db is a free data retrieval call binding the contract method 0x4d655aff.
//
// Solidity: function db() view returns(address)
func (_UCamDIDManager *UCamDIDManagerCallerSession) Db() (common.Address, error) {
	return _UCamDIDManager.Contract.Db(&_UCamDIDManager.CallOpts)
}

// DecodeInternalKey is a free data retrieval call binding the contract method 0xbf6d3139.
//
// Solidity: function decodeInternalKey(bytes did) view returns(bytes20)
func (_UCamDIDManager *UCamDIDManagerCaller) DecodeInternalKey(opts *bind.CallOpts, did []byte) ([20]byte, error) {
	var (
		ret0 = new([20]byte)
	)
	out := ret0
	err := _UCamDIDManager.contract.Call(opts, out, "decodeInternalKey", did)
	return *ret0, err
}

// DecodeInternalKey is a free data retrieval call binding the contract method 0xbf6d3139.
//
// Solidity: function decodeInternalKey(bytes did) view returns(bytes20)
func (_UCamDIDManager *UCamDIDManagerSession) DecodeInternalKey(did []byte) ([20]byte, error) {
	return _UCamDIDManager.Contract.DecodeInternalKey(&_UCamDIDManager.CallOpts, did)
}

// DecodeInternalKey is a free data retrieval call binding the contract method 0xbf6d3139.
//
// Solidity: function decodeInternalKey(bytes did) view returns(bytes20)
func (_UCamDIDManager *UCamDIDManagerCallerSession) DecodeInternalKey(did []byte) ([20]byte, error) {
	return _UCamDIDManager.Contract.DecodeInternalKey(&_UCamDIDManager.CallOpts, did)
}

// GetCreateAuthMessage is a free data retrieval call binding the contract method 0xcc3b41dc.
//
// Solidity: function getCreateAuthMessage(bytes did, bytes32 h, bytes uri, address agent) view returns(bytes)
func (_UCamDIDManager *UCamDIDManagerCaller) GetCreateAuthMessage(opts *bind.CallOpts, did []byte, h [32]byte, uri []byte, agent common.Address) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _UCamDIDManager.contract.Call(opts, out, "getCreateAuthMessage", did, h, uri, agent)
	return *ret0, err
}

// GetCreateAuthMessage is a free data retrieval call binding the contract method 0xcc3b41dc.
//
// Solidity: function getCreateAuthMessage(bytes did, bytes32 h, bytes uri, address agent) view returns(bytes)
func (_UCamDIDManager *UCamDIDManagerSession) GetCreateAuthMessage(did []byte, h [32]byte, uri []byte, agent common.Address) ([]byte, error) {
	return _UCamDIDManager.Contract.GetCreateAuthMessage(&_UCamDIDManager.CallOpts, did, h, uri, agent)
}

// GetCreateAuthMessage is a free data retrieval call binding the contract method 0xcc3b41dc.
//
// Solidity: function getCreateAuthMessage(bytes did, bytes32 h, bytes uri, address agent) view returns(bytes)
func (_UCamDIDManager *UCamDIDManagerCallerSession) GetCreateAuthMessage(did []byte, h [32]byte, uri []byte, agent common.Address) ([]byte, error) {
	return _UCamDIDManager.Contract.GetCreateAuthMessage(&_UCamDIDManager.CallOpts, did, h, uri, agent)
}

// GetDeleteAuthMessage is a free data retrieval call binding the contract method 0x3bdcb26c.
//
// Solidity: function getDeleteAuthMessage(bytes did, address agent) view returns(bytes)
func (_UCamDIDManager *UCamDIDManagerCaller) GetDeleteAuthMessage(opts *bind.CallOpts, did []byte, agent common.Address) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _UCamDIDManager.contract.Call(opts, out, "getDeleteAuthMessage", did, agent)
	return *ret0, err
}

// GetDeleteAuthMessage is a free data retrieval call binding the contract method 0x3bdcb26c.
//
// Solidity: function getDeleteAuthMessage(bytes did, address agent) view returns(bytes)
func (_UCamDIDManager *UCamDIDManagerSession) GetDeleteAuthMessage(did []byte, agent common.Address) ([]byte, error) {
	return _UCamDIDManager.Contract.GetDeleteAuthMessage(&_UCamDIDManager.CallOpts, did, agent)
}

// GetDeleteAuthMessage is a free data retrieval call binding the contract method 0x3bdcb26c.
//
// Solidity: function getDeleteAuthMessage(bytes did, address agent) view returns(bytes)
func (_UCamDIDManager *UCamDIDManagerCallerSession) GetDeleteAuthMessage(did []byte, agent common.Address) ([]byte, error) {
	return _UCamDIDManager.Contract.GetDeleteAuthMessage(&_UCamDIDManager.CallOpts, did, agent)
}

// GetHash is a free data retrieval call binding the contract method 0xb00140aa.
//
// Solidity: function getHash(bytes did) view returns(bytes32)
func (_UCamDIDManager *UCamDIDManagerCaller) GetHash(opts *bind.CallOpts, did []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _UCamDIDManager.contract.Call(opts, out, "getHash", did)
	return *ret0, err
}

// GetHash is a free data retrieval call binding the contract method 0xb00140aa.
//
// Solidity: function getHash(bytes did) view returns(bytes32)
func (_UCamDIDManager *UCamDIDManagerSession) GetHash(did []byte) ([32]byte, error) {
	return _UCamDIDManager.Contract.GetHash(&_UCamDIDManager.CallOpts, did)
}

// GetHash is a free data retrieval call binding the contract method 0xb00140aa.
//
// Solidity: function getHash(bytes did) view returns(bytes32)
func (_UCamDIDManager *UCamDIDManagerCallerSession) GetHash(did []byte) ([32]byte, error) {
	return _UCamDIDManager.Contract.GetHash(&_UCamDIDManager.CallOpts, did)
}

// GetOwner is a free data retrieval call binding the contract method 0xb102bfbe.
//
// Solidity: function getOwner(bytes did) view returns(address)
func (_UCamDIDManager *UCamDIDManagerCaller) GetOwner(opts *bind.CallOpts, did []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UCamDIDManager.contract.Call(opts, out, "getOwner", did)
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0xb102bfbe.
//
// Solidity: function getOwner(bytes did) view returns(address)
func (_UCamDIDManager *UCamDIDManagerSession) GetOwner(did []byte) (common.Address, error) {
	return _UCamDIDManager.Contract.GetOwner(&_UCamDIDManager.CallOpts, did)
}

// GetOwner is a free data retrieval call binding the contract method 0xb102bfbe.
//
// Solidity: function getOwner(bytes did) view returns(address)
func (_UCamDIDManager *UCamDIDManagerCallerSession) GetOwner(did []byte) (common.Address, error) {
	return _UCamDIDManager.Contract.GetOwner(&_UCamDIDManager.CallOpts, did)
}

// GetURI is a free data retrieval call binding the contract method 0x8626dea9.
//
// Solidity: function getURI(bytes did) view returns(bytes)
func (_UCamDIDManager *UCamDIDManagerCaller) GetURI(opts *bind.CallOpts, did []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _UCamDIDManager.contract.Call(opts, out, "getURI", did)
	return *ret0, err
}

// GetURI is a free data retrieval call binding the contract method 0x8626dea9.
//
// Solidity: function getURI(bytes did) view returns(bytes)
func (_UCamDIDManager *UCamDIDManagerSession) GetURI(did []byte) ([]byte, error) {
	return _UCamDIDManager.Contract.GetURI(&_UCamDIDManager.CallOpts, did)
}

// GetURI is a free data retrieval call binding the contract method 0x8626dea9.
//
// Solidity: function getURI(bytes did) view returns(bytes)
func (_UCamDIDManager *UCamDIDManagerCallerSession) GetURI(did []byte) ([]byte, error) {
	return _UCamDIDManager.Contract.GetURI(&_UCamDIDManager.CallOpts, did)
}

// GetUpdateAuthMessage is a free data retrieval call binding the contract method 0xdd120153.
//
// Solidity: function getUpdateAuthMessage(bytes did, bytes32 h, bytes uri, address agent) view returns(bytes)
func (_UCamDIDManager *UCamDIDManagerCaller) GetUpdateAuthMessage(opts *bind.CallOpts, did []byte, h [32]byte, uri []byte, agent common.Address) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _UCamDIDManager.contract.Call(opts, out, "getUpdateAuthMessage", did, h, uri, agent)
	return *ret0, err
}

// GetUpdateAuthMessage is a free data retrieval call binding the contract method 0xdd120153.
//
// Solidity: function getUpdateAuthMessage(bytes did, bytes32 h, bytes uri, address agent) view returns(bytes)
func (_UCamDIDManager *UCamDIDManagerSession) GetUpdateAuthMessage(did []byte, h [32]byte, uri []byte, agent common.Address) ([]byte, error) {
	return _UCamDIDManager.Contract.GetUpdateAuthMessage(&_UCamDIDManager.CallOpts, did, h, uri, agent)
}

// GetUpdateAuthMessage is a free data retrieval call binding the contract method 0xdd120153.
//
// Solidity: function getUpdateAuthMessage(bytes did, bytes32 h, bytes uri, address agent) view returns(bytes)
func (_UCamDIDManager *UCamDIDManagerCallerSession) GetUpdateAuthMessage(did []byte, h [32]byte, uri []byte, agent common.Address) ([]byte, error) {
	return _UCamDIDManager.Contract.GetUpdateAuthMessage(&_UCamDIDManager.CallOpts, did, h, uri, agent)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UCamDIDManager *UCamDIDManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UCamDIDManager.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UCamDIDManager *UCamDIDManagerSession) Owner() (common.Address, error) {
	return _UCamDIDManager.Contract.Owner(&_UCamDIDManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UCamDIDManager *UCamDIDManagerCallerSession) Owner() (common.Address, error) {
	return _UCamDIDManager.Contract.Owner(&_UCamDIDManager.CallOpts)
}

// CreateDIDByAgent is a paid mutator transaction binding the contract method 0x3dc7ad02.
//
// Solidity: function createDIDByAgent(bytes20 uid, bytes32 h, bytes uri, address authorizer, bytes auth) returns()
func (_UCamDIDManager *UCamDIDManagerTransactor) CreateDIDByAgent(opts *bind.TransactOpts, uid [20]byte, h [32]byte, uri []byte, authorizer common.Address, auth []byte) (*types.Transaction, error) {
	return _UCamDIDManager.contract.Transact(opts, "createDIDByAgent", uid, h, uri, authorizer, auth)
}

// CreateDIDByAgent is a paid mutator transaction binding the contract method 0x3dc7ad02.
//
// Solidity: function createDIDByAgent(bytes20 uid, bytes32 h, bytes uri, address authorizer, bytes auth) returns()
func (_UCamDIDManager *UCamDIDManagerSession) CreateDIDByAgent(uid [20]byte, h [32]byte, uri []byte, authorizer common.Address, auth []byte) (*types.Transaction, error) {
	return _UCamDIDManager.Contract.CreateDIDByAgent(&_UCamDIDManager.TransactOpts, uid, h, uri, authorizer, auth)
}

// CreateDIDByAgent is a paid mutator transaction binding the contract method 0x3dc7ad02.
//
// Solidity: function createDIDByAgent(bytes20 uid, bytes32 h, bytes uri, address authorizer, bytes auth) returns()
func (_UCamDIDManager *UCamDIDManagerTransactorSession) CreateDIDByAgent(uid [20]byte, h [32]byte, uri []byte, authorizer common.Address, auth []byte) (*types.Transaction, error) {
	return _UCamDIDManager.Contract.CreateDIDByAgent(&_UCamDIDManager.TransactOpts, uid, h, uri, authorizer, auth)
}

// DeleteDIDByAgent is a paid mutator transaction binding the contract method 0xd9d50607.
//
// Solidity: function deleteDIDByAgent(bytes20 uid, bytes auth) returns()
func (_UCamDIDManager *UCamDIDManagerTransactor) DeleteDIDByAgent(opts *bind.TransactOpts, uid [20]byte, auth []byte) (*types.Transaction, error) {
	return _UCamDIDManager.contract.Transact(opts, "deleteDIDByAgent", uid, auth)
}

// DeleteDIDByAgent is a paid mutator transaction binding the contract method 0xd9d50607.
//
// Solidity: function deleteDIDByAgent(bytes20 uid, bytes auth) returns()
func (_UCamDIDManager *UCamDIDManagerSession) DeleteDIDByAgent(uid [20]byte, auth []byte) (*types.Transaction, error) {
	return _UCamDIDManager.Contract.DeleteDIDByAgent(&_UCamDIDManager.TransactOpts, uid, auth)
}

// DeleteDIDByAgent is a paid mutator transaction binding the contract method 0xd9d50607.
//
// Solidity: function deleteDIDByAgent(bytes20 uid, bytes auth) returns()
func (_UCamDIDManager *UCamDIDManagerTransactorSession) DeleteDIDByAgent(uid [20]byte, auth []byte) (*types.Transaction, error) {
	return _UCamDIDManager.Contract.DeleteDIDByAgent(&_UCamDIDManager.TransactOpts, uid, auth)
}

// TransferDBOwnership is a paid mutator transaction binding the contract method 0x32aad017.
//
// Solidity: function transferDBOwnership(address newOwner) returns()
func (_UCamDIDManager *UCamDIDManagerTransactor) TransferDBOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _UCamDIDManager.contract.Transact(opts, "transferDBOwnership", newOwner)
}

// TransferDBOwnership is a paid mutator transaction binding the contract method 0x32aad017.
//
// Solidity: function transferDBOwnership(address newOwner) returns()
func (_UCamDIDManager *UCamDIDManagerSession) TransferDBOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UCamDIDManager.Contract.TransferDBOwnership(&_UCamDIDManager.TransactOpts, newOwner)
}

// TransferDBOwnership is a paid mutator transaction binding the contract method 0x32aad017.
//
// Solidity: function transferDBOwnership(address newOwner) returns()
func (_UCamDIDManager *UCamDIDManagerTransactorSession) TransferDBOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UCamDIDManager.Contract.TransferDBOwnership(&_UCamDIDManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UCamDIDManager *UCamDIDManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _UCamDIDManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UCamDIDManager *UCamDIDManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UCamDIDManager.Contract.TransferOwnership(&_UCamDIDManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UCamDIDManager *UCamDIDManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UCamDIDManager.Contract.TransferOwnership(&_UCamDIDManager.TransactOpts, newOwner)
}

// UpdateDIDByAgent is a paid mutator transaction binding the contract method 0x98793984.
//
// Solidity: function updateDIDByAgent(bytes20 uid, bytes32 h, bytes uri, bytes auth) returns()
func (_UCamDIDManager *UCamDIDManagerTransactor) UpdateDIDByAgent(opts *bind.TransactOpts, uid [20]byte, h [32]byte, uri []byte, auth []byte) (*types.Transaction, error) {
	return _UCamDIDManager.contract.Transact(opts, "updateDIDByAgent", uid, h, uri, auth)
}

// UpdateDIDByAgent is a paid mutator transaction binding the contract method 0x98793984.
//
// Solidity: function updateDIDByAgent(bytes20 uid, bytes32 h, bytes uri, bytes auth) returns()
func (_UCamDIDManager *UCamDIDManagerSession) UpdateDIDByAgent(uid [20]byte, h [32]byte, uri []byte, auth []byte) (*types.Transaction, error) {
	return _UCamDIDManager.Contract.UpdateDIDByAgent(&_UCamDIDManager.TransactOpts, uid, h, uri, auth)
}

// UpdateDIDByAgent is a paid mutator transaction binding the contract method 0x98793984.
//
// Solidity: function updateDIDByAgent(bytes20 uid, bytes32 h, bytes uri, bytes auth) returns()
func (_UCamDIDManager *UCamDIDManagerTransactorSession) UpdateDIDByAgent(uid [20]byte, h [32]byte, uri []byte, auth []byte) (*types.Transaction, error) {
	return _UCamDIDManager.Contract.UpdateDIDByAgent(&_UCamDIDManager.TransactOpts, uid, h, uri, auth)
}

// UCamDIDManagerDIDCreatedIterator is returned from FilterDIDCreated and is used to iterate over the raw logs and unpacked data for DIDCreated events raised by the UCamDIDManager contract.
type UCamDIDManagerDIDCreatedIterator struct {
	Event *UCamDIDManagerDIDCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UCamDIDManagerDIDCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UCamDIDManagerDIDCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UCamDIDManagerDIDCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UCamDIDManagerDIDCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UCamDIDManagerDIDCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UCamDIDManagerDIDCreated represents a DIDCreated event raised by the UCamDIDManager contract.
type UCamDIDManagerDIDCreated struct {
	Operator common.Address
	Did      string
	Hash     [32]byte
	Uri      string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDIDCreated is a free log retrieval operation binding the contract event 0x0a36151ef7283a91a6a8d052cd44a0c20f8c7207decf22d985eba8eb0a5a268b.
//
// Solidity: event DIDCreated(address indexed operator, string did, bytes32 hash, string uri)
func (_UCamDIDManager *UCamDIDManagerFilterer) FilterDIDCreated(opts *bind.FilterOpts, operator []common.Address) (*UCamDIDManagerDIDCreatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _UCamDIDManager.contract.FilterLogs(opts, "DIDCreated", operatorRule)
	if err != nil {
		return nil, err
	}
	return &UCamDIDManagerDIDCreatedIterator{contract: _UCamDIDManager.contract, event: "DIDCreated", logs: logs, sub: sub}, nil
}

// WatchDIDCreated is a free log subscription operation binding the contract event 0x0a36151ef7283a91a6a8d052cd44a0c20f8c7207decf22d985eba8eb0a5a268b.
//
// Solidity: event DIDCreated(address indexed operator, string did, bytes32 hash, string uri)
func (_UCamDIDManager *UCamDIDManagerFilterer) WatchDIDCreated(opts *bind.WatchOpts, sink chan<- *UCamDIDManagerDIDCreated, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _UCamDIDManager.contract.WatchLogs(opts, "DIDCreated", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UCamDIDManagerDIDCreated)
				if err := _UCamDIDManager.contract.UnpackLog(event, "DIDCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDIDCreated is a log parse operation binding the contract event 0x0a36151ef7283a91a6a8d052cd44a0c20f8c7207decf22d985eba8eb0a5a268b.
//
// Solidity: event DIDCreated(address indexed operator, string did, bytes32 hash, string uri)
func (_UCamDIDManager *UCamDIDManagerFilterer) ParseDIDCreated(log types.Log) (*UCamDIDManagerDIDCreated, error) {
	event := new(UCamDIDManagerDIDCreated)
	if err := _UCamDIDManager.contract.UnpackLog(event, "DIDCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// UCamDIDManagerDIDDeletedIterator is returned from FilterDIDDeleted and is used to iterate over the raw logs and unpacked data for DIDDeleted events raised by the UCamDIDManager contract.
type UCamDIDManagerDIDDeletedIterator struct {
	Event *UCamDIDManagerDIDDeleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UCamDIDManagerDIDDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UCamDIDManagerDIDDeleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UCamDIDManagerDIDDeleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UCamDIDManagerDIDDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UCamDIDManagerDIDDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UCamDIDManagerDIDDeleted represents a DIDDeleted event raised by the UCamDIDManager contract.
type UCamDIDManagerDIDDeleted struct {
	Operator common.Address
	Did      common.Hash
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDIDDeleted is a free log retrieval operation binding the contract event 0x56c5173d83675a7f2b174113df9205152eb36c65ab39e41d5d1b8d9cb88eabd2.
//
// Solidity: event DIDDeleted(address indexed operator, string indexed did)
func (_UCamDIDManager *UCamDIDManagerFilterer) FilterDIDDeleted(opts *bind.FilterOpts, operator []common.Address, did []string) (*UCamDIDManagerDIDDeletedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _UCamDIDManager.contract.FilterLogs(opts, "DIDDeleted", operatorRule, didRule)
	if err != nil {
		return nil, err
	}
	return &UCamDIDManagerDIDDeletedIterator{contract: _UCamDIDManager.contract, event: "DIDDeleted", logs: logs, sub: sub}, nil
}

// WatchDIDDeleted is a free log subscription operation binding the contract event 0x56c5173d83675a7f2b174113df9205152eb36c65ab39e41d5d1b8d9cb88eabd2.
//
// Solidity: event DIDDeleted(address indexed operator, string indexed did)
func (_UCamDIDManager *UCamDIDManagerFilterer) WatchDIDDeleted(opts *bind.WatchOpts, sink chan<- *UCamDIDManagerDIDDeleted, operator []common.Address, did []string) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _UCamDIDManager.contract.WatchLogs(opts, "DIDDeleted", operatorRule, didRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UCamDIDManagerDIDDeleted)
				if err := _UCamDIDManager.contract.UnpackLog(event, "DIDDeleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDIDDeleted is a log parse operation binding the contract event 0x56c5173d83675a7f2b174113df9205152eb36c65ab39e41d5d1b8d9cb88eabd2.
//
// Solidity: event DIDDeleted(address indexed operator, string indexed did)
func (_UCamDIDManager *UCamDIDManagerFilterer) ParseDIDDeleted(log types.Log) (*UCamDIDManagerDIDDeleted, error) {
	event := new(UCamDIDManagerDIDDeleted)
	if err := _UCamDIDManager.contract.UnpackLog(event, "DIDDeleted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// UCamDIDManagerDIDUpdatedIterator is returned from FilterDIDUpdated and is used to iterate over the raw logs and unpacked data for DIDUpdated events raised by the UCamDIDManager contract.
type UCamDIDManagerDIDUpdatedIterator struct {
	Event *UCamDIDManagerDIDUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UCamDIDManagerDIDUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UCamDIDManagerDIDUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UCamDIDManagerDIDUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UCamDIDManagerDIDUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UCamDIDManagerDIDUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UCamDIDManagerDIDUpdated represents a DIDUpdated event raised by the UCamDIDManager contract.
type UCamDIDManagerDIDUpdated struct {
	Operator common.Address
	Did      common.Hash
	Hash     [32]byte
	Uri      string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDIDUpdated is a free log retrieval operation binding the contract event 0x30665598c4da8be2cccefb7fe32a379decd1eae407564569a8e95b51e3afccd7.
//
// Solidity: event DIDUpdated(address indexed operator, string indexed did, bytes32 hash, string uri)
func (_UCamDIDManager *UCamDIDManagerFilterer) FilterDIDUpdated(opts *bind.FilterOpts, operator []common.Address, did []string) (*UCamDIDManagerDIDUpdatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _UCamDIDManager.contract.FilterLogs(opts, "DIDUpdated", operatorRule, didRule)
	if err != nil {
		return nil, err
	}
	return &UCamDIDManagerDIDUpdatedIterator{contract: _UCamDIDManager.contract, event: "DIDUpdated", logs: logs, sub: sub}, nil
}

// WatchDIDUpdated is a free log subscription operation binding the contract event 0x30665598c4da8be2cccefb7fe32a379decd1eae407564569a8e95b51e3afccd7.
//
// Solidity: event DIDUpdated(address indexed operator, string indexed did, bytes32 hash, string uri)
func (_UCamDIDManager *UCamDIDManagerFilterer) WatchDIDUpdated(opts *bind.WatchOpts, sink chan<- *UCamDIDManagerDIDUpdated, operator []common.Address, did []string) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _UCamDIDManager.contract.WatchLogs(opts, "DIDUpdated", operatorRule, didRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UCamDIDManagerDIDUpdated)
				if err := _UCamDIDManager.contract.UnpackLog(event, "DIDUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDIDUpdated is a log parse operation binding the contract event 0x30665598c4da8be2cccefb7fe32a379decd1eae407564569a8e95b51e3afccd7.
//
// Solidity: event DIDUpdated(address indexed operator, string indexed did, bytes32 hash, string uri)
func (_UCamDIDManager *UCamDIDManagerFilterer) ParseDIDUpdated(log types.Log) (*UCamDIDManagerDIDUpdated, error) {
	event := new(UCamDIDManagerDIDUpdated)
	if err := _UCamDIDManager.contract.UnpackLog(event, "DIDUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// UCamDIDManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the UCamDIDManager contract.
type UCamDIDManagerOwnershipTransferredIterator struct {
	Event *UCamDIDManagerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UCamDIDManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UCamDIDManagerOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UCamDIDManagerOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UCamDIDManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UCamDIDManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UCamDIDManagerOwnershipTransferred represents a OwnershipTransferred event raised by the UCamDIDManager contract.
type UCamDIDManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UCamDIDManager *UCamDIDManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UCamDIDManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UCamDIDManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UCamDIDManagerOwnershipTransferredIterator{contract: _UCamDIDManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UCamDIDManager *UCamDIDManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UCamDIDManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UCamDIDManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UCamDIDManagerOwnershipTransferred)
				if err := _UCamDIDManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UCamDIDManager *UCamDIDManagerFilterer) ParseOwnershipTransferred(log types.Log) (*UCamDIDManagerOwnershipTransferred, error) {
	event := new(UCamDIDManagerOwnershipTransferred)
	if err := _UCamDIDManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
