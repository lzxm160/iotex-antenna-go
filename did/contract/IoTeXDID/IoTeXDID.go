// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IoTeXDID

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

// IoTeXDIDABI is the input ABI used to generate the binding from.
const IoTeXDIDABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"didString\",\"type\":\"string\"}],\"name\":\"CreateDID\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"didString\",\"type\":\"string\"}],\"name\":\"DeleteDID\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"didString\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"UpdateHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"didString\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"UpdateURI\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"createDID\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"sigV\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"sigR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sigS\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"createDIDSigned\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"deleteDID\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"deleteDID\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"sigV\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"sigR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sigS\",\"type\":\"bytes32\"}],\"name\":\"deleteDIDSigned\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"dids\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"getHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"getURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"updateHash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"sigV\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"sigR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sigS\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"updateHashSigned\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"updateURI\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"sigV\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"sigR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sigS\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"updateURISigned\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IoTeXDIDFuncSigs maps the 4-byte function signature to its string representation.
var IoTeXDIDFuncSigs = map[string]string{
	"3e05e066": "createDID(string,bytes32,string)",
	"303bc6dc": "createDIDSigned(string,uint8,bytes32,bytes32,bytes32,string)",
	"42e643bc": "deleteDID(string)",
	"7e7bad7e": "deleteDID(string,address)",
	"be0a254e": "deleteDIDSigned(string,uint8,bytes32,bytes32)",
	"f44ab516": "dids(string)",
	"5b6beeb9": "getHash(string)",
	"93ff5d3e": "getURI(string)",
	"78eab45a": "updateHash(string,bytes32)",
	"dafd7e95": "updateHashSigned(string,uint8,bytes32,bytes32,bytes32)",
	"1789e2d8": "updateURI(string,string)",
	"f1ae6261": "updateURISigned(string,uint8,bytes32,bytes32,string)",
}

// IoTeXDIDBin is the compiled bytecode used for deploying new contracts.
var IoTeXDIDBin = "0x608060405234801561001057600080fd5b50612445806100206000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80637e7bad7e116100715780637e7bad7e1461065957806393ff5d3e14610708578063be0a254e14610821578063dafd7e95146108d3578063f1ae62611461098b578063f44ab51614610ac9576100b4565b80631789e2d8146100b9578063303bc6dc146101e45780633e05e0661461032857806342e643bc146104595780635b6beeb9146104fd57806378eab45a146105b3575b600080fd5b6101e2600480360360408110156100cf57600080fd5b810190602081018135600160201b8111156100e957600080fd5b8201836020820111156100fb57600080fd5b803590602001918460018302840111600160201b8311171561011c57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561016e57600080fd5b82018360208201111561018057600080fd5b803590602001918460018302840111600160201b831117156101a157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610bf7945050505050565b005b6101e2600480360360c08110156101fa57600080fd5b810190602081018135600160201b81111561021457600080fd5b82018360208201111561022657600080fd5b803590602001918460018302840111600160201b8311171561024757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929560ff8535169560208601359560408101359550606081013594509192509060a081019060800135600160201b8111156102b457600080fd5b8201836020820111156102c657600080fd5b803590602001918460018302840111600160201b831117156102e757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610c06945050505050565b6101e26004803603606081101561033e57600080fd5b810190602081018135600160201b81111561035857600080fd5b82018360208201111561036a57600080fd5b803590602001918460018302840111600160201b8311171561038b57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092958435959094909350604081019250602001359050600160201b8111156103e557600080fd5b8201836020820111156103f757600080fd5b803590602001918460018302840111600160201b8311171561041857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610cb9945050505050565b6101e26004803603602081101561046f57600080fd5b810190602081018135600160201b81111561048957600080fd5b82018360208201111561049b57600080fd5b803590602001918460018302840111600160201b831117156104bc57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610cca945050505050565b6105a16004803603602081101561051357600080fd5b810190602081018135600160201b81111561052d57600080fd5b82018360208201111561053f57600080fd5b803590602001918460018302840111600160201b8311171561056057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610cd7945050505050565b60408051918252519081900360200190f35b6101e2600480360360408110156105c957600080fd5b810190602081018135600160201b8111156105e357600080fd5b8201836020820111156105f557600080fd5b803590602001918460018302840111600160201b8311171561061657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610dfb915050565b6101e26004803603604081101561066f57600080fd5b810190602081018135600160201b81111561068957600080fd5b82018360208201111561069b57600080fd5b803590602001918460018302840111600160201b831117156106bc57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160a01b03169150610e069050565b6107ac6004803603602081101561071e57600080fd5b810190602081018135600160201b81111561073857600080fd5b82018360208201111561074a57600080fd5b803590602001918460018302840111600160201b8311171561076b57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061101f945050505050565b6040805160208082528351818301528351919283929083019185019080838360005b838110156107e65781810151838201526020016107ce565b50505050905090810190601f1680156108135780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101e26004803603608081101561083757600080fd5b810190602081018135600160201b81111561085157600080fd5b82018360208201111561086357600080fd5b803590602001918460018302840111600160201b8311171561088457600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505060ff83351693505050602081013590604001356111ce565b6101e2600480360360a08110156108e957600080fd5b810190602081018135600160201b81111561090357600080fd5b82018360208201111561091557600080fd5b803590602001918460018302840111600160201b8311171561093657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505060ff83351693505050602081013590604081013590606001356112c3565b6101e2600480360360a08110156109a157600080fd5b810190602081018135600160201b8111156109bb57600080fd5b8201836020820111156109cd57600080fd5b803590602001918460018302840111600160201b831117156109ee57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929560ff85351695602086013595604081013595509193509150608081019060600135600160201b811115610a5557600080fd5b820183602082011115610a6757600080fd5b803590602001918460018302840111600160201b83111715610a8857600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506113cb945050505050565b610b6d60048036036020811015610adf57600080fd5b810190602081018135600160201b811115610af957600080fd5b820183602082011115610b0b57600080fd5b803590602001918460018302840111600160201b83111715610b2c57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061150d945050505050565b604051808415151515815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610bba578181015183820152602001610ba2565b50505050905090810190601f168015610be75780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b610c028233836115cb565b5050565b6000610c35876040518060400160405280600981526020016818dc99585d1951125160ba1b8152508585611857565b805190602001209050610cb08760018389898960405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015610c9f573d6000803e3d6000fd5b505050602060405103518585611a02565b50505050505050565b610cc583338484611a02565b505050565b610cd48133610e06565b50565b60006060610ce483611cda565b90506000816040518082805190602001908083835b60208310610d185780518252601f199092019160209182019101610cf9565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092205460ff169150610d919050576040805162461bcd60e51b8152602060048201526012602482015271191a5908191bd95cc81b9bdd08195e1a5cdd60721b604482015290519081900360640190fd5b6000816040518082805190602001908083835b60208310610dc35780518252601f199092019160209182019101610da4565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092206001015495945050505050565b610c02823383611ded565b81816060610e1382612006565b835190915015610e6257610e2783826120e7565b610e625760405162461bcd60e51b81526004018080602001828103825260218152602001806123f06021913960400191505060405180910390fd5b6000816040518082805190602001908083835b60208310610e945780518252601f199092019160209182019101610e75565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092205460ff169150610f149050576040805162461bcd60e51b815260206004820152601960248201527831b0b63632b91034b9903737ba1030903234b21037bbb732b960391b604482015290519081900360640190fd5b600080610f2086612006565b6040518082805190602001908083835b60208310610f4f5780518252601f199092019160209182019101610f30565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220805460ff19169315159390931790925550610f96905084612006565b6040518082805190602001908083835b60208310610fc55780518252601f199092019160209182019101610fa6565b5181516020939093036101000a60001901801990911692169190911790526040519201829003822093507ff8910781c5d425c8a634b0a60a6567ef96cae8b2aa97b0db44389fff2bbcc64f92506000919050a25050505050565b60608061102b83611cda565b90506000816040518082805190602001908083835b6020831061105f5780518252601f199092019160209182019101611040565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092205460ff1691506110d89050576040805162461bcd60e51b8152602060048201526012602482015271191a5908191bd95cc81b9bdd08195e1a5cdd60721b604482015290519081900360640190fd5b6000816040518082805190602001908083835b6020831061110a5780518252601f1990920191602091820191016110eb565b518151600019602094850361010090810a8201928316921993909316919091179092529490920196875260408051978890038201882060029081018054601f6001821615909802909501909416049485018290048202880182019052838752909450919250508301828280156111c15780601f10611196576101008083540402835291602001916111c1565b820191906000526020600020905b8154815290600101906020018083116111a457829003601f168201915b5050505050915050919050565b6000846040516020018082805190602001908083835b602083106112035780518252601f1990920191602091820191016111e4565b51815160001960209485036101000a019081169019919091161790526819195b195d1951125160ba1b93909101928352604080516016198186030181526009850180835281519184019190912060009091526029850180835281905260ff8b166049860152606985018a90526089850189905290519096506112bc95508a945060019360a98082019450601f19830192918290030190855afa1580156112ad573d6000803e3d6000fd5b50505060206040510351610e06565b5050505050565b600085826040516020018083805190602001908083835b602083106112f95780518252601f1990920191602091820191016112da565b6001836020036101000a03801982511681845116808217855250505050505090500180690eae0c8c2e8ca90c2e6d60b31b815250600a01828152602001925050506040516020818303038152906040528051906020012090506113c38660018388888860405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa1580156113b3573d6000803e3d6000fd5b5050506020604051035184611ded565b505050505050565b600085826040516020018083805190602001908083835b602083106114015780518252601f1990920191602091820191016113e2565b51815160209384036101000a60001901801990921691161790526875706461746555524960b81b919093019081528451600990910192850191508083835b6020831061145e5780518252601f19909201916020918201910161143f565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040528051906020012090506113c38660018388888860405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa1580156114fd573d6000803e3d6000fd5b50505060206040510351846115cb565b80516020818301810180516000825292820193820193909320919092528054600180830154600280850180546040805161010096831615969096026000190190911692909204601f810188900488028501880190925281845260ff90941695919493918301828280156115c15780601f10611596576101008083540402835291602001916115c1565b820191906000526020600020905b8154815290600101906020018083116115a457829003601f168201915b5050505050905083565b828260606115d882612006565b835190915015611627576115ec83826120e7565b6116275760405162461bcd60e51b81526004018080602001828103825260218152602001806123f06021913960400191505060405180910390fd5b6000816040518082805190602001908083835b602083106116595780518252601f19909201916020918201910161163a565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092205460ff1691506116d99050576040805162461bcd60e51b815260206004820152601960248201527831b0b63632b91034b9903737ba1030903234b21037bbb732b960391b604482015290519081900360640190fd5b8360006116e587612006565b6040518082805190602001908083835b602083106117145780518252601f1990920191602091820191016116f5565b51815160209384036101000a60001901801990921691161790529201948552506040519384900381019093208451611759956002909201949190910192509050612354565b5061176385612006565b6040518082805190602001908083835b602083106117925780518252601f199092019160209182019101611773565b51815160209384036101000a6000190180199092169116179052604080519290940182900382208183528a51838301528a519096507fa55ea16d9a68fded3c5b007d1fbba6ed1524d9354d44c66bcc54c721108586c495508a94929350839283019185019080838360005b838110156118155781810151838201526020016117fd565b50505050905090810190601f1680156118425780820380516001836020036101000a031916815260200191505b509250505060405180910390a2505050505050565b6060808590506060859050606084905060608151602060ff16845186510101016040519080825280601f01601f1916602001820160405280156118a1576020820181803883390190505b509050806000805b86518110156118fa578681815181106118be57fe5b602001015160f81c60f81b8383806001019450815181106118db57fe5b60200101906001600160f81b031916908160001a9053506001016118a9565b5060005b855181101561194f5785818151811061191357fe5b602001015160f81c60f81b83838060010194508151811061193057fe5b60200101906001600160f81b031916908160001a9053506001016118fe565b5060005b602081101561199d5789816020811061196857fe5b1a60f81b83838060010194508151811061197e57fe5b60200101906001600160f81b031916908160001a905350600101611953565b5060005b84518110156119f2578481815181106119b657fe5b602001015160f81c60f81b8383806001019450815181106119d357fe5b60200101906001600160f81b031916908160001a9053506001016119a1565b50909a9950505050505050505050565b835115611a6c57611a1b84611a16856121de565b6120e7565b611a6c576040805162461bcd60e51b815260206004820152601860248201527f696420646f6573206e6f74206d61746368207369676e65720000000000000000604482015290519081900360640190fd5b6060611a7784612006565b90506000816040518082805190602001908083835b60208310611aab5780518252601f199092019160209182019101611a8c565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092205460ff16159150611b259050576040805162461bcd60e51b815260206004820152601260248201527164696420616c72656164792065786973747360701b604482015290519081900360640190fd5b6040518060600160405280600115158152602001848152602001838152506000826040518082805190602001908083835b60208310611b755780518252601f199092019160209182019101611b56565b51815160209384036101000a6000190180199092169116179052920194855250604080519485900382019094208551815460ff1916901515178155858201516001820155938501518051611bd29450600286019350910190612354565b50905050611be7611be2856121de565b611cda565b6040518082805190602001908083835b60208310611c165780518252601f199092019160209182019101611bf7565b51815160209384036101000a60001901801990921691161790526040805192909401829003822081835287518383015287519096507fedcad282d2fc969322aa68caeb2b9bf9de762094ec288a9965535b8968dca17f95508794929350839283019185019080838360005b83811015611c99578181015183820152602001611c81565b50505050905090810190601f168015611cc65780820380516001836020036101000a031916815260200191505b509250505060405180910390a25050505050565b606080829050606081516040519080825280601f01601f191660200182016040528015611d0e576020820181803883390190505b50905060005b8251811015611de5576041838281518110611d2b57fe5b016020015160f81c10801590611d555750605a838281518110611d4a57fe5b016020015160f81c11155b15611da257828181518110611d6657fe5b602001015160f81c60f81b60f81c60200160f81b828281518110611d8657fe5b60200101906001600160f81b031916908160001a905350611ddd565b828181518110611dae57fe5b602001015160f81c60f81b828281518110611dc557fe5b60200101906001600160f81b031916908160001a9053505b600101611d14565b509392505050565b82826060611dfa82612006565b835190915015611e4957611e0e83826120e7565b611e495760405162461bcd60e51b81526004018080602001828103825260218152602001806123f06021913960400191505060405180910390fd5b6000816040518082805190602001908083835b60208310611e7b5780518252601f199092019160209182019101611e5c565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092205460ff169150611efb9050576040805162461bcd60e51b815260206004820152601960248201527831b0b63632b91034b9903737ba1030903234b21037bbb732b960391b604482015290519081900360640190fd5b836000611f0787612006565b6040518082805190602001908083835b60208310611f365780518252601f199092019160209182019101611f17565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092206001019290925550611f74905085612006565b6040518082805190602001908083835b60208310611fa35780518252601f199092019160209182019101611f84565b51815160209384036101000a6000190180199092169116179052604080519290940182900382208a835293519395507fad2d38b24c03f1f0db04b9345ca745dcf8904f27f8977599ec4769743d5e898d94509083900301919050a2505050505050565b6060604051806040016040528060078152602001663234b21d34b79d60c91b815250612031836121de565b6040516020018083805190602001908083835b602083106120635780518252601f199092019160209182019101612044565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106120ab5780518252601f19909201916020918201910161208c565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040529050919050565b60006120f282611cda565b6040516020018082805190602001908083835b602083106121245780518252601f199092019160209182019101612105565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040528051906020012061216884611cda565b6040516020018082805190602001908083835b6020831061219a5780518252601f19909201916020918201910161217b565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040528051906020012014905092915050565b604080518082018252601081526f181899199a1a9b1b9c1cb0b131b232b360811b60208201528151602a80825260608281019094526001600160a01b03851692918491602082018180388339019050509050600360fc1b8160008151811061224257fe5b60200101906001600160f81b031916908160001a905350600f60fb1b8160018151811061226b57fe5b60200101906001600160f81b031916908160001a90535060005b601481101561234b578260048583600c01602081106122a057fe5b835191901a90911c60ff169081106122b457fe5b602001015160f81c60f81b8282600202600201815181106122d157fe5b60200101906001600160f81b031916908160001a905350828482600c01602081106122f857fe5b1a60f81b60f81c600f1660ff168151811061230f57fe5b602001015160f81c60f81b82826002026003018151811061232c57fe5b60200101906001600160f81b031916908160001a905350600101612285565b50949350505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061239557805160ff19168380011785556123c2565b828001600101855582156123c2579182015b828111156123c25782518255916020019190600101906123a7565b506123ce9291506123d2565b5090565b6123ec91905b808211156123ce57600081556001016123d8565b9056fe63616c6c657220646f6573206e6f74206f776e2074686520676976656e20646964a265627a7a72315820d826b9e223fc519b3cbdc75f7b086317e3b2a6123ff41faad874df664a2af3a664736f6c63430005100032"

// DeployIoTeXDID deploys a new Ethereum contract, binding an instance of IoTeXDID to it.
func DeployIoTeXDID(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IoTeXDID, error) {
	parsed, err := abi.JSON(strings.NewReader(IoTeXDIDABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IoTeXDIDBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IoTeXDID{IoTeXDIDCaller: IoTeXDIDCaller{contract: contract}, IoTeXDIDTransactor: IoTeXDIDTransactor{contract: contract}, IoTeXDIDFilterer: IoTeXDIDFilterer{contract: contract}}, nil
}

// IoTeXDID is an auto generated Go binding around an Ethereum contract.
type IoTeXDID struct {
	IoTeXDIDCaller     // Read-only binding to the contract
	IoTeXDIDTransactor // Write-only binding to the contract
	IoTeXDIDFilterer   // Log filterer for contract events
}

// IoTeXDIDCaller is an auto generated read-only Go binding around an Ethereum contract.
type IoTeXDIDCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IoTeXDIDTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IoTeXDIDTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IoTeXDIDFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IoTeXDIDFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IoTeXDIDSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IoTeXDIDSession struct {
	Contract     *IoTeXDID         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IoTeXDIDCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IoTeXDIDCallerSession struct {
	Contract *IoTeXDIDCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IoTeXDIDTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IoTeXDIDTransactorSession struct {
	Contract     *IoTeXDIDTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IoTeXDIDRaw is an auto generated low-level Go binding around an Ethereum contract.
type IoTeXDIDRaw struct {
	Contract *IoTeXDID // Generic contract binding to access the raw methods on
}

// IoTeXDIDCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IoTeXDIDCallerRaw struct {
	Contract *IoTeXDIDCaller // Generic read-only contract binding to access the raw methods on
}

// IoTeXDIDTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IoTeXDIDTransactorRaw struct {
	Contract *IoTeXDIDTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIoTeXDID creates a new instance of IoTeXDID, bound to a specific deployed contract.
func NewIoTeXDID(address common.Address, backend bind.ContractBackend) (*IoTeXDID, error) {
	contract, err := bindIoTeXDID(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IoTeXDID{IoTeXDIDCaller: IoTeXDIDCaller{contract: contract}, IoTeXDIDTransactor: IoTeXDIDTransactor{contract: contract}, IoTeXDIDFilterer: IoTeXDIDFilterer{contract: contract}}, nil
}

// NewIoTeXDIDCaller creates a new read-only instance of IoTeXDID, bound to a specific deployed contract.
func NewIoTeXDIDCaller(address common.Address, caller bind.ContractCaller) (*IoTeXDIDCaller, error) {
	contract, err := bindIoTeXDID(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IoTeXDIDCaller{contract: contract}, nil
}

// NewIoTeXDIDTransactor creates a new write-only instance of IoTeXDID, bound to a specific deployed contract.
func NewIoTeXDIDTransactor(address common.Address, transactor bind.ContractTransactor) (*IoTeXDIDTransactor, error) {
	contract, err := bindIoTeXDID(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IoTeXDIDTransactor{contract: contract}, nil
}

// NewIoTeXDIDFilterer creates a new log filterer instance of IoTeXDID, bound to a specific deployed contract.
func NewIoTeXDIDFilterer(address common.Address, filterer bind.ContractFilterer) (*IoTeXDIDFilterer, error) {
	contract, err := bindIoTeXDID(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IoTeXDIDFilterer{contract: contract}, nil
}

// bindIoTeXDID binds a generic wrapper to an already deployed contract.
func bindIoTeXDID(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IoTeXDIDABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IoTeXDID *IoTeXDIDRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IoTeXDID.Contract.IoTeXDIDCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IoTeXDID *IoTeXDIDRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IoTeXDID.Contract.IoTeXDIDTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IoTeXDID *IoTeXDIDRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IoTeXDID.Contract.IoTeXDIDTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IoTeXDID *IoTeXDIDCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IoTeXDID.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IoTeXDID *IoTeXDIDTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IoTeXDID.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IoTeXDID *IoTeXDIDTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IoTeXDID.Contract.contract.Transact(opts, method, params...)
}

// Dids is a free data retrieval call binding the contract method 0xf44ab516.
//
// Solidity: function dids(string ) view returns(bool exist, bytes32 hash, string uri)
func (_IoTeXDID *IoTeXDIDCaller) Dids(opts *bind.CallOpts, arg0 string) (struct {
	Exist bool
	Hash  [32]byte
	Uri   string
}, error) {
	ret := new(struct {
		Exist bool
		Hash  [32]byte
		Uri   string
	})
	out := ret
	err := _IoTeXDID.contract.Call(opts, out, "dids", arg0)
	return *ret, err
}

// Dids is a free data retrieval call binding the contract method 0xf44ab516.
//
// Solidity: function dids(string ) view returns(bool exist, bytes32 hash, string uri)
func (_IoTeXDID *IoTeXDIDSession) Dids(arg0 string) (struct {
	Exist bool
	Hash  [32]byte
	Uri   string
}, error) {
	return _IoTeXDID.Contract.Dids(&_IoTeXDID.CallOpts, arg0)
}

// Dids is a free data retrieval call binding the contract method 0xf44ab516.
//
// Solidity: function dids(string ) view returns(bool exist, bytes32 hash, string uri)
func (_IoTeXDID *IoTeXDIDCallerSession) Dids(arg0 string) (struct {
	Exist bool
	Hash  [32]byte
	Uri   string
}, error) {
	return _IoTeXDID.Contract.Dids(&_IoTeXDID.CallOpts, arg0)
}

// GetHash is a free data retrieval call binding the contract method 0x5b6beeb9.
//
// Solidity: function getHash(string did) view returns(bytes32)
func (_IoTeXDID *IoTeXDIDCaller) GetHash(opts *bind.CallOpts, did string) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IoTeXDID.contract.Call(opts, out, "getHash", did)
	return *ret0, err
}

// GetHash is a free data retrieval call binding the contract method 0x5b6beeb9.
//
// Solidity: function getHash(string did) view returns(bytes32)
func (_IoTeXDID *IoTeXDIDSession) GetHash(did string) ([32]byte, error) {
	return _IoTeXDID.Contract.GetHash(&_IoTeXDID.CallOpts, did)
}

// GetHash is a free data retrieval call binding the contract method 0x5b6beeb9.
//
// Solidity: function getHash(string did) view returns(bytes32)
func (_IoTeXDID *IoTeXDIDCallerSession) GetHash(did string) ([32]byte, error) {
	return _IoTeXDID.Contract.GetHash(&_IoTeXDID.CallOpts, did)
}

// GetURI is a free data retrieval call binding the contract method 0x93ff5d3e.
//
// Solidity: function getURI(string did) view returns(string)
func (_IoTeXDID *IoTeXDIDCaller) GetURI(opts *bind.CallOpts, did string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IoTeXDID.contract.Call(opts, out, "getURI", did)
	return *ret0, err
}

// GetURI is a free data retrieval call binding the contract method 0x93ff5d3e.
//
// Solidity: function getURI(string did) view returns(string)
func (_IoTeXDID *IoTeXDIDSession) GetURI(did string) (string, error) {
	return _IoTeXDID.Contract.GetURI(&_IoTeXDID.CallOpts, did)
}

// GetURI is a free data retrieval call binding the contract method 0x93ff5d3e.
//
// Solidity: function getURI(string did) view returns(string)
func (_IoTeXDID *IoTeXDIDCallerSession) GetURI(did string) (string, error) {
	return _IoTeXDID.Contract.GetURI(&_IoTeXDID.CallOpts, did)
}

// CreateDID is a paid mutator transaction binding the contract method 0x3e05e066.
//
// Solidity: function createDID(string id, bytes32 hash, string uri) returns()
func (_IoTeXDID *IoTeXDIDTransactor) CreateDID(opts *bind.TransactOpts, id string, hash [32]byte, uri string) (*types.Transaction, error) {
	return _IoTeXDID.contract.Transact(opts, "createDID", id, hash, uri)
}

// CreateDID is a paid mutator transaction binding the contract method 0x3e05e066.
//
// Solidity: function createDID(string id, bytes32 hash, string uri) returns()
func (_IoTeXDID *IoTeXDIDSession) CreateDID(id string, hash [32]byte, uri string) (*types.Transaction, error) {
	return _IoTeXDID.Contract.CreateDID(&_IoTeXDID.TransactOpts, id, hash, uri)
}

// CreateDID is a paid mutator transaction binding the contract method 0x3e05e066.
//
// Solidity: function createDID(string id, bytes32 hash, string uri) returns()
func (_IoTeXDID *IoTeXDIDTransactorSession) CreateDID(id string, hash [32]byte, uri string) (*types.Transaction, error) {
	return _IoTeXDID.Contract.CreateDID(&_IoTeXDID.TransactOpts, id, hash, uri)
}

// CreateDIDSigned is a paid mutator transaction binding the contract method 0x303bc6dc.
//
// Solidity: function createDIDSigned(string id, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 hash, string uri) returns()
func (_IoTeXDID *IoTeXDIDTransactor) CreateDIDSigned(opts *bind.TransactOpts, id string, sigV uint8, sigR [32]byte, sigS [32]byte, hash [32]byte, uri string) (*types.Transaction, error) {
	return _IoTeXDID.contract.Transact(opts, "createDIDSigned", id, sigV, sigR, sigS, hash, uri)
}

// CreateDIDSigned is a paid mutator transaction binding the contract method 0x303bc6dc.
//
// Solidity: function createDIDSigned(string id, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 hash, string uri) returns()
func (_IoTeXDID *IoTeXDIDSession) CreateDIDSigned(id string, sigV uint8, sigR [32]byte, sigS [32]byte, hash [32]byte, uri string) (*types.Transaction, error) {
	return _IoTeXDID.Contract.CreateDIDSigned(&_IoTeXDID.TransactOpts, id, sigV, sigR, sigS, hash, uri)
}

// CreateDIDSigned is a paid mutator transaction binding the contract method 0x303bc6dc.
//
// Solidity: function createDIDSigned(string id, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 hash, string uri) returns()
func (_IoTeXDID *IoTeXDIDTransactorSession) CreateDIDSigned(id string, sigV uint8, sigR [32]byte, sigS [32]byte, hash [32]byte, uri string) (*types.Transaction, error) {
	return _IoTeXDID.Contract.CreateDIDSigned(&_IoTeXDID.TransactOpts, id, sigV, sigR, sigS, hash, uri)
}

// DeleteDID is a paid mutator transaction binding the contract method 0x42e643bc.
//
// Solidity: function deleteDID(string did) returns()
func (_IoTeXDID *IoTeXDIDTransactor) DeleteDID(opts *bind.TransactOpts, did string) (*types.Transaction, error) {
	return _IoTeXDID.contract.Transact(opts, "deleteDID", did)
}

// DeleteDID is a paid mutator transaction binding the contract method 0x42e643bc.
//
// Solidity: function deleteDID(string did) returns()
func (_IoTeXDID *IoTeXDIDSession) DeleteDID(did string) (*types.Transaction, error) {
	return _IoTeXDID.Contract.DeleteDID(&_IoTeXDID.TransactOpts, did)
}

// DeleteDID is a paid mutator transaction binding the contract method 0x42e643bc.
//
// Solidity: function deleteDID(string did) returns()
func (_IoTeXDID *IoTeXDIDTransactorSession) DeleteDID(did string) (*types.Transaction, error) {
	return _IoTeXDID.Contract.DeleteDID(&_IoTeXDID.TransactOpts, did)
}

// DeleteDID0 is a paid mutator transaction binding the contract method 0x7e7bad7e.
//
// Solidity: function deleteDID(string did, address signer) returns()
func (_IoTeXDID *IoTeXDIDTransactor) DeleteDID0(opts *bind.TransactOpts, did string, signer common.Address) (*types.Transaction, error) {
	return _IoTeXDID.contract.Transact(opts, "deleteDID0", did, signer)
}

// DeleteDID0 is a paid mutator transaction binding the contract method 0x7e7bad7e.
//
// Solidity: function deleteDID(string did, address signer) returns()
func (_IoTeXDID *IoTeXDIDSession) DeleteDID0(did string, signer common.Address) (*types.Transaction, error) {
	return _IoTeXDID.Contract.DeleteDID0(&_IoTeXDID.TransactOpts, did, signer)
}

// DeleteDID0 is a paid mutator transaction binding the contract method 0x7e7bad7e.
//
// Solidity: function deleteDID(string did, address signer) returns()
func (_IoTeXDID *IoTeXDIDTransactorSession) DeleteDID0(did string, signer common.Address) (*types.Transaction, error) {
	return _IoTeXDID.Contract.DeleteDID0(&_IoTeXDID.TransactOpts, did, signer)
}

// DeleteDIDSigned is a paid mutator transaction binding the contract method 0xbe0a254e.
//
// Solidity: function deleteDIDSigned(string did, uint8 sigV, bytes32 sigR, bytes32 sigS) returns()
func (_IoTeXDID *IoTeXDIDTransactor) DeleteDIDSigned(opts *bind.TransactOpts, did string, sigV uint8, sigR [32]byte, sigS [32]byte) (*types.Transaction, error) {
	return _IoTeXDID.contract.Transact(opts, "deleteDIDSigned", did, sigV, sigR, sigS)
}

// DeleteDIDSigned is a paid mutator transaction binding the contract method 0xbe0a254e.
//
// Solidity: function deleteDIDSigned(string did, uint8 sigV, bytes32 sigR, bytes32 sigS) returns()
func (_IoTeXDID *IoTeXDIDSession) DeleteDIDSigned(did string, sigV uint8, sigR [32]byte, sigS [32]byte) (*types.Transaction, error) {
	return _IoTeXDID.Contract.DeleteDIDSigned(&_IoTeXDID.TransactOpts, did, sigV, sigR, sigS)
}

// DeleteDIDSigned is a paid mutator transaction binding the contract method 0xbe0a254e.
//
// Solidity: function deleteDIDSigned(string did, uint8 sigV, bytes32 sigR, bytes32 sigS) returns()
func (_IoTeXDID *IoTeXDIDTransactorSession) DeleteDIDSigned(did string, sigV uint8, sigR [32]byte, sigS [32]byte) (*types.Transaction, error) {
	return _IoTeXDID.Contract.DeleteDIDSigned(&_IoTeXDID.TransactOpts, did, sigV, sigR, sigS)
}

// UpdateHash is a paid mutator transaction binding the contract method 0x78eab45a.
//
// Solidity: function updateHash(string did, bytes32 hash) returns()
func (_IoTeXDID *IoTeXDIDTransactor) UpdateHash(opts *bind.TransactOpts, did string, hash [32]byte) (*types.Transaction, error) {
	return _IoTeXDID.contract.Transact(opts, "updateHash", did, hash)
}

// UpdateHash is a paid mutator transaction binding the contract method 0x78eab45a.
//
// Solidity: function updateHash(string did, bytes32 hash) returns()
func (_IoTeXDID *IoTeXDIDSession) UpdateHash(did string, hash [32]byte) (*types.Transaction, error) {
	return _IoTeXDID.Contract.UpdateHash(&_IoTeXDID.TransactOpts, did, hash)
}

// UpdateHash is a paid mutator transaction binding the contract method 0x78eab45a.
//
// Solidity: function updateHash(string did, bytes32 hash) returns()
func (_IoTeXDID *IoTeXDIDTransactorSession) UpdateHash(did string, hash [32]byte) (*types.Transaction, error) {
	return _IoTeXDID.Contract.UpdateHash(&_IoTeXDID.TransactOpts, did, hash)
}

// UpdateHashSigned is a paid mutator transaction binding the contract method 0xdafd7e95.
//
// Solidity: function updateHashSigned(string did, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 hash) returns()
func (_IoTeXDID *IoTeXDIDTransactor) UpdateHashSigned(opts *bind.TransactOpts, did string, sigV uint8, sigR [32]byte, sigS [32]byte, hash [32]byte) (*types.Transaction, error) {
	return _IoTeXDID.contract.Transact(opts, "updateHashSigned", did, sigV, sigR, sigS, hash)
}

// UpdateHashSigned is a paid mutator transaction binding the contract method 0xdafd7e95.
//
// Solidity: function updateHashSigned(string did, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 hash) returns()
func (_IoTeXDID *IoTeXDIDSession) UpdateHashSigned(did string, sigV uint8, sigR [32]byte, sigS [32]byte, hash [32]byte) (*types.Transaction, error) {
	return _IoTeXDID.Contract.UpdateHashSigned(&_IoTeXDID.TransactOpts, did, sigV, sigR, sigS, hash)
}

// UpdateHashSigned is a paid mutator transaction binding the contract method 0xdafd7e95.
//
// Solidity: function updateHashSigned(string did, uint8 sigV, bytes32 sigR, bytes32 sigS, bytes32 hash) returns()
func (_IoTeXDID *IoTeXDIDTransactorSession) UpdateHashSigned(did string, sigV uint8, sigR [32]byte, sigS [32]byte, hash [32]byte) (*types.Transaction, error) {
	return _IoTeXDID.Contract.UpdateHashSigned(&_IoTeXDID.TransactOpts, did, sigV, sigR, sigS, hash)
}

// UpdateURI is a paid mutator transaction binding the contract method 0x1789e2d8.
//
// Solidity: function updateURI(string did, string uri) returns()
func (_IoTeXDID *IoTeXDIDTransactor) UpdateURI(opts *bind.TransactOpts, did string, uri string) (*types.Transaction, error) {
	return _IoTeXDID.contract.Transact(opts, "updateURI", did, uri)
}

// UpdateURI is a paid mutator transaction binding the contract method 0x1789e2d8.
//
// Solidity: function updateURI(string did, string uri) returns()
func (_IoTeXDID *IoTeXDIDSession) UpdateURI(did string, uri string) (*types.Transaction, error) {
	return _IoTeXDID.Contract.UpdateURI(&_IoTeXDID.TransactOpts, did, uri)
}

// UpdateURI is a paid mutator transaction binding the contract method 0x1789e2d8.
//
// Solidity: function updateURI(string did, string uri) returns()
func (_IoTeXDID *IoTeXDIDTransactorSession) UpdateURI(did string, uri string) (*types.Transaction, error) {
	return _IoTeXDID.Contract.UpdateURI(&_IoTeXDID.TransactOpts, did, uri)
}

// UpdateURISigned is a paid mutator transaction binding the contract method 0xf1ae6261.
//
// Solidity: function updateURISigned(string did, uint8 sigV, bytes32 sigR, bytes32 sigS, string uri) returns()
func (_IoTeXDID *IoTeXDIDTransactor) UpdateURISigned(opts *bind.TransactOpts, did string, sigV uint8, sigR [32]byte, sigS [32]byte, uri string) (*types.Transaction, error) {
	return _IoTeXDID.contract.Transact(opts, "updateURISigned", did, sigV, sigR, sigS, uri)
}

// UpdateURISigned is a paid mutator transaction binding the contract method 0xf1ae6261.
//
// Solidity: function updateURISigned(string did, uint8 sigV, bytes32 sigR, bytes32 sigS, string uri) returns()
func (_IoTeXDID *IoTeXDIDSession) UpdateURISigned(did string, sigV uint8, sigR [32]byte, sigS [32]byte, uri string) (*types.Transaction, error) {
	return _IoTeXDID.Contract.UpdateURISigned(&_IoTeXDID.TransactOpts, did, sigV, sigR, sigS, uri)
}

// UpdateURISigned is a paid mutator transaction binding the contract method 0xf1ae6261.
//
// Solidity: function updateURISigned(string did, uint8 sigV, bytes32 sigR, bytes32 sigS, string uri) returns()
func (_IoTeXDID *IoTeXDIDTransactorSession) UpdateURISigned(did string, sigV uint8, sigR [32]byte, sigS [32]byte, uri string) (*types.Transaction, error) {
	return _IoTeXDID.Contract.UpdateURISigned(&_IoTeXDID.TransactOpts, did, sigV, sigR, sigS, uri)
}

// IoTeXDIDCreateDIDIterator is returned from FilterCreateDID and is used to iterate over the raw logs and unpacked data for CreateDID events raised by the IoTeXDID contract.
type IoTeXDIDCreateDIDIterator struct {
	Event *IoTeXDIDCreateDID // Event containing the contract specifics and raw log

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
func (it *IoTeXDIDCreateDIDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoTeXDIDCreateDID)
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
		it.Event = new(IoTeXDIDCreateDID)
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
func (it *IoTeXDIDCreateDIDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoTeXDIDCreateDIDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoTeXDIDCreateDID represents a CreateDID event raised by the IoTeXDID contract.
type IoTeXDIDCreateDID struct {
	Id        common.Hash
	DidString string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCreateDID is a free log retrieval operation binding the contract event 0xedcad282d2fc969322aa68caeb2b9bf9de762094ec288a9965535b8968dca17f.
//
// Solidity: event CreateDID(string indexed id, string didString)
func (_IoTeXDID *IoTeXDIDFilterer) FilterCreateDID(opts *bind.FilterOpts, id []string) (*IoTeXDIDCreateDIDIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IoTeXDID.contract.FilterLogs(opts, "CreateDID", idRule)
	if err != nil {
		return nil, err
	}
	return &IoTeXDIDCreateDIDIterator{contract: _IoTeXDID.contract, event: "CreateDID", logs: logs, sub: sub}, nil
}

// WatchCreateDID is a free log subscription operation binding the contract event 0xedcad282d2fc969322aa68caeb2b9bf9de762094ec288a9965535b8968dca17f.
//
// Solidity: event CreateDID(string indexed id, string didString)
func (_IoTeXDID *IoTeXDIDFilterer) WatchCreateDID(opts *bind.WatchOpts, sink chan<- *IoTeXDIDCreateDID, id []string) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IoTeXDID.contract.WatchLogs(opts, "CreateDID", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoTeXDIDCreateDID)
				if err := _IoTeXDID.contract.UnpackLog(event, "CreateDID", log); err != nil {
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

// ParseCreateDID is a log parse operation binding the contract event 0xedcad282d2fc969322aa68caeb2b9bf9de762094ec288a9965535b8968dca17f.
//
// Solidity: event CreateDID(string indexed id, string didString)
func (_IoTeXDID *IoTeXDIDFilterer) ParseCreateDID(log types.Log) (*IoTeXDIDCreateDID, error) {
	event := new(IoTeXDIDCreateDID)
	if err := _IoTeXDID.contract.UnpackLog(event, "CreateDID", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IoTeXDIDDeleteDIDIterator is returned from FilterDeleteDID and is used to iterate over the raw logs and unpacked data for DeleteDID events raised by the IoTeXDID contract.
type IoTeXDIDDeleteDIDIterator struct {
	Event *IoTeXDIDDeleteDID // Event containing the contract specifics and raw log

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
func (it *IoTeXDIDDeleteDIDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoTeXDIDDeleteDID)
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
		it.Event = new(IoTeXDIDDeleteDID)
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
func (it *IoTeXDIDDeleteDIDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoTeXDIDDeleteDIDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoTeXDIDDeleteDID represents a DeleteDID event raised by the IoTeXDID contract.
type IoTeXDIDDeleteDID struct {
	DidString common.Hash
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeleteDID is a free log retrieval operation binding the contract event 0xf8910781c5d425c8a634b0a60a6567ef96cae8b2aa97b0db44389fff2bbcc64f.
//
// Solidity: event DeleteDID(string indexed didString)
func (_IoTeXDID *IoTeXDIDFilterer) FilterDeleteDID(opts *bind.FilterOpts, didString []string) (*IoTeXDIDDeleteDIDIterator, error) {

	var didStringRule []interface{}
	for _, didStringItem := range didString {
		didStringRule = append(didStringRule, didStringItem)
	}

	logs, sub, err := _IoTeXDID.contract.FilterLogs(opts, "DeleteDID", didStringRule)
	if err != nil {
		return nil, err
	}
	return &IoTeXDIDDeleteDIDIterator{contract: _IoTeXDID.contract, event: "DeleteDID", logs: logs, sub: sub}, nil
}

// WatchDeleteDID is a free log subscription operation binding the contract event 0xf8910781c5d425c8a634b0a60a6567ef96cae8b2aa97b0db44389fff2bbcc64f.
//
// Solidity: event DeleteDID(string indexed didString)
func (_IoTeXDID *IoTeXDIDFilterer) WatchDeleteDID(opts *bind.WatchOpts, sink chan<- *IoTeXDIDDeleteDID, didString []string) (event.Subscription, error) {

	var didStringRule []interface{}
	for _, didStringItem := range didString {
		didStringRule = append(didStringRule, didStringItem)
	}

	logs, sub, err := _IoTeXDID.contract.WatchLogs(opts, "DeleteDID", didStringRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoTeXDIDDeleteDID)
				if err := _IoTeXDID.contract.UnpackLog(event, "DeleteDID", log); err != nil {
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

// ParseDeleteDID is a log parse operation binding the contract event 0xf8910781c5d425c8a634b0a60a6567ef96cae8b2aa97b0db44389fff2bbcc64f.
//
// Solidity: event DeleteDID(string indexed didString)
func (_IoTeXDID *IoTeXDIDFilterer) ParseDeleteDID(log types.Log) (*IoTeXDIDDeleteDID, error) {
	event := new(IoTeXDIDDeleteDID)
	if err := _IoTeXDID.contract.UnpackLog(event, "DeleteDID", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IoTeXDIDUpdateHashIterator is returned from FilterUpdateHash and is used to iterate over the raw logs and unpacked data for UpdateHash events raised by the IoTeXDID contract.
type IoTeXDIDUpdateHashIterator struct {
	Event *IoTeXDIDUpdateHash // Event containing the contract specifics and raw log

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
func (it *IoTeXDIDUpdateHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoTeXDIDUpdateHash)
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
		it.Event = new(IoTeXDIDUpdateHash)
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
func (it *IoTeXDIDUpdateHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoTeXDIDUpdateHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoTeXDIDUpdateHash represents a UpdateHash event raised by the IoTeXDID contract.
type IoTeXDIDUpdateHash struct {
	DidString common.Hash
	Hash      [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateHash is a free log retrieval operation binding the contract event 0xad2d38b24c03f1f0db04b9345ca745dcf8904f27f8977599ec4769743d5e898d.
//
// Solidity: event UpdateHash(string indexed didString, bytes32 hash)
func (_IoTeXDID *IoTeXDIDFilterer) FilterUpdateHash(opts *bind.FilterOpts, didString []string) (*IoTeXDIDUpdateHashIterator, error) {

	var didStringRule []interface{}
	for _, didStringItem := range didString {
		didStringRule = append(didStringRule, didStringItem)
	}

	logs, sub, err := _IoTeXDID.contract.FilterLogs(opts, "UpdateHash", didStringRule)
	if err != nil {
		return nil, err
	}
	return &IoTeXDIDUpdateHashIterator{contract: _IoTeXDID.contract, event: "UpdateHash", logs: logs, sub: sub}, nil
}

// WatchUpdateHash is a free log subscription operation binding the contract event 0xad2d38b24c03f1f0db04b9345ca745dcf8904f27f8977599ec4769743d5e898d.
//
// Solidity: event UpdateHash(string indexed didString, bytes32 hash)
func (_IoTeXDID *IoTeXDIDFilterer) WatchUpdateHash(opts *bind.WatchOpts, sink chan<- *IoTeXDIDUpdateHash, didString []string) (event.Subscription, error) {

	var didStringRule []interface{}
	for _, didStringItem := range didString {
		didStringRule = append(didStringRule, didStringItem)
	}

	logs, sub, err := _IoTeXDID.contract.WatchLogs(opts, "UpdateHash", didStringRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoTeXDIDUpdateHash)
				if err := _IoTeXDID.contract.UnpackLog(event, "UpdateHash", log); err != nil {
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

// ParseUpdateHash is a log parse operation binding the contract event 0xad2d38b24c03f1f0db04b9345ca745dcf8904f27f8977599ec4769743d5e898d.
//
// Solidity: event UpdateHash(string indexed didString, bytes32 hash)
func (_IoTeXDID *IoTeXDIDFilterer) ParseUpdateHash(log types.Log) (*IoTeXDIDUpdateHash, error) {
	event := new(IoTeXDIDUpdateHash)
	if err := _IoTeXDID.contract.UnpackLog(event, "UpdateHash", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IoTeXDIDUpdateURIIterator is returned from FilterUpdateURI and is used to iterate over the raw logs and unpacked data for UpdateURI events raised by the IoTeXDID contract.
type IoTeXDIDUpdateURIIterator struct {
	Event *IoTeXDIDUpdateURI // Event containing the contract specifics and raw log

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
func (it *IoTeXDIDUpdateURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoTeXDIDUpdateURI)
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
		it.Event = new(IoTeXDIDUpdateURI)
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
func (it *IoTeXDIDUpdateURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoTeXDIDUpdateURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoTeXDIDUpdateURI represents a UpdateURI event raised by the IoTeXDID contract.
type IoTeXDIDUpdateURI struct {
	DidString common.Hash
	Uri       string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateURI is a free log retrieval operation binding the contract event 0xa55ea16d9a68fded3c5b007d1fbba6ed1524d9354d44c66bcc54c721108586c4.
//
// Solidity: event UpdateURI(string indexed didString, string uri)
func (_IoTeXDID *IoTeXDIDFilterer) FilterUpdateURI(opts *bind.FilterOpts, didString []string) (*IoTeXDIDUpdateURIIterator, error) {

	var didStringRule []interface{}
	for _, didStringItem := range didString {
		didStringRule = append(didStringRule, didStringItem)
	}

	logs, sub, err := _IoTeXDID.contract.FilterLogs(opts, "UpdateURI", didStringRule)
	if err != nil {
		return nil, err
	}
	return &IoTeXDIDUpdateURIIterator{contract: _IoTeXDID.contract, event: "UpdateURI", logs: logs, sub: sub}, nil
}

// WatchUpdateURI is a free log subscription operation binding the contract event 0xa55ea16d9a68fded3c5b007d1fbba6ed1524d9354d44c66bcc54c721108586c4.
//
// Solidity: event UpdateURI(string indexed didString, string uri)
func (_IoTeXDID *IoTeXDIDFilterer) WatchUpdateURI(opts *bind.WatchOpts, sink chan<- *IoTeXDIDUpdateURI, didString []string) (event.Subscription, error) {

	var didStringRule []interface{}
	for _, didStringItem := range didString {
		didStringRule = append(didStringRule, didStringItem)
	}

	logs, sub, err := _IoTeXDID.contract.WatchLogs(opts, "UpdateURI", didStringRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoTeXDIDUpdateURI)
				if err := _IoTeXDID.contract.UnpackLog(event, "UpdateURI", log); err != nil {
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

// ParseUpdateURI is a log parse operation binding the contract event 0xa55ea16d9a68fded3c5b007d1fbba6ed1524d9354d44c66bcc54c721108586c4.
//
// Solidity: event UpdateURI(string indexed didString, string uri)
func (_IoTeXDID *IoTeXDIDFilterer) ParseUpdateURI(log types.Log) (*IoTeXDIDUpdateURI, error) {
	event := new(IoTeXDIDUpdateURI)
	if err := _IoTeXDID.contract.UnpackLog(event, "UpdateURI", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IoTeXDIDStorageABI is the input ABI used to generate the binding from.
const IoTeXDIDStorageABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"dids\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IoTeXDIDStorageFuncSigs maps the 4-byte function signature to its string representation.
var IoTeXDIDStorageFuncSigs = map[string]string{
	"f44ab516": "dids(string)",
}

// IoTeXDIDStorageBin is the compiled bytecode used for deploying new contracts.
var IoTeXDIDStorageBin = "0x608060405234801561001057600080fd5b50610253806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063f44ab51614610030575b600080fd5b6100d66004803603602081101561004657600080fd5b81019060208101813564010000000081111561006157600080fd5b82018360208201111561007357600080fd5b8035906020019184600183028401116401000000008311171561009557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610160945050505050565b604051808415151515815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561012357818101518382015260200161010b565b50505050905090810190601f1680156101505780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b80516020818301810180516000825292820193820193909320919092528054600180830154600280850180546040805161010096831615969096026000190190911692909204601f810188900488028501880190925281845260ff90941695919493918301828280156102145780601f106101e957610100808354040283529160200191610214565b820191906000526020600020905b8154815290600101906020018083116101f757829003601f168201915b505050505090508356fea265627a7a72315820d24d51bb6f1d61fa23ba567809965c228785afd80342420826b4654562ebf74264736f6c63430005100032"

// DeployIoTeXDIDStorage deploys a new Ethereum contract, binding an instance of IoTeXDIDStorage to it.
func DeployIoTeXDIDStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IoTeXDIDStorage, error) {
	parsed, err := abi.JSON(strings.NewReader(IoTeXDIDStorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IoTeXDIDStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IoTeXDIDStorage{IoTeXDIDStorageCaller: IoTeXDIDStorageCaller{contract: contract}, IoTeXDIDStorageTransactor: IoTeXDIDStorageTransactor{contract: contract}, IoTeXDIDStorageFilterer: IoTeXDIDStorageFilterer{contract: contract}}, nil
}

// IoTeXDIDStorage is an auto generated Go binding around an Ethereum contract.
type IoTeXDIDStorage struct {
	IoTeXDIDStorageCaller     // Read-only binding to the contract
	IoTeXDIDStorageTransactor // Write-only binding to the contract
	IoTeXDIDStorageFilterer   // Log filterer for contract events
}

// IoTeXDIDStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type IoTeXDIDStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IoTeXDIDStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IoTeXDIDStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IoTeXDIDStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IoTeXDIDStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IoTeXDIDStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IoTeXDIDStorageSession struct {
	Contract     *IoTeXDIDStorage  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IoTeXDIDStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IoTeXDIDStorageCallerSession struct {
	Contract *IoTeXDIDStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IoTeXDIDStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IoTeXDIDStorageTransactorSession struct {
	Contract     *IoTeXDIDStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IoTeXDIDStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type IoTeXDIDStorageRaw struct {
	Contract *IoTeXDIDStorage // Generic contract binding to access the raw methods on
}

// IoTeXDIDStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IoTeXDIDStorageCallerRaw struct {
	Contract *IoTeXDIDStorageCaller // Generic read-only contract binding to access the raw methods on
}

// IoTeXDIDStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IoTeXDIDStorageTransactorRaw struct {
	Contract *IoTeXDIDStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIoTeXDIDStorage creates a new instance of IoTeXDIDStorage, bound to a specific deployed contract.
func NewIoTeXDIDStorage(address common.Address, backend bind.ContractBackend) (*IoTeXDIDStorage, error) {
	contract, err := bindIoTeXDIDStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IoTeXDIDStorage{IoTeXDIDStorageCaller: IoTeXDIDStorageCaller{contract: contract}, IoTeXDIDStorageTransactor: IoTeXDIDStorageTransactor{contract: contract}, IoTeXDIDStorageFilterer: IoTeXDIDStorageFilterer{contract: contract}}, nil
}

// NewIoTeXDIDStorageCaller creates a new read-only instance of IoTeXDIDStorage, bound to a specific deployed contract.
func NewIoTeXDIDStorageCaller(address common.Address, caller bind.ContractCaller) (*IoTeXDIDStorageCaller, error) {
	contract, err := bindIoTeXDIDStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IoTeXDIDStorageCaller{contract: contract}, nil
}

// NewIoTeXDIDStorageTransactor creates a new write-only instance of IoTeXDIDStorage, bound to a specific deployed contract.
func NewIoTeXDIDStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*IoTeXDIDStorageTransactor, error) {
	contract, err := bindIoTeXDIDStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IoTeXDIDStorageTransactor{contract: contract}, nil
}

// NewIoTeXDIDStorageFilterer creates a new log filterer instance of IoTeXDIDStorage, bound to a specific deployed contract.
func NewIoTeXDIDStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*IoTeXDIDStorageFilterer, error) {
	contract, err := bindIoTeXDIDStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IoTeXDIDStorageFilterer{contract: contract}, nil
}

// bindIoTeXDIDStorage binds a generic wrapper to an already deployed contract.
func bindIoTeXDIDStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IoTeXDIDStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IoTeXDIDStorage *IoTeXDIDStorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IoTeXDIDStorage.Contract.IoTeXDIDStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IoTeXDIDStorage *IoTeXDIDStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IoTeXDIDStorage.Contract.IoTeXDIDStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IoTeXDIDStorage *IoTeXDIDStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IoTeXDIDStorage.Contract.IoTeXDIDStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IoTeXDIDStorage *IoTeXDIDStorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IoTeXDIDStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IoTeXDIDStorage *IoTeXDIDStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IoTeXDIDStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IoTeXDIDStorage *IoTeXDIDStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IoTeXDIDStorage.Contract.contract.Transact(opts, method, params...)
}

// Dids is a free data retrieval call binding the contract method 0xf44ab516.
//
// Solidity: function dids(string ) view returns(bool exist, bytes32 hash, string uri)
func (_IoTeXDIDStorage *IoTeXDIDStorageCaller) Dids(opts *bind.CallOpts, arg0 string) (struct {
	Exist bool
	Hash  [32]byte
	Uri   string
}, error) {
	ret := new(struct {
		Exist bool
		Hash  [32]byte
		Uri   string
	})
	out := ret
	err := _IoTeXDIDStorage.contract.Call(opts, out, "dids", arg0)
	return *ret, err
}

// Dids is a free data retrieval call binding the contract method 0xf44ab516.
//
// Solidity: function dids(string ) view returns(bool exist, bytes32 hash, string uri)
func (_IoTeXDIDStorage *IoTeXDIDStorageSession) Dids(arg0 string) (struct {
	Exist bool
	Hash  [32]byte
	Uri   string
}, error) {
	return _IoTeXDIDStorage.Contract.Dids(&_IoTeXDIDStorage.CallOpts, arg0)
}

// Dids is a free data retrieval call binding the contract method 0xf44ab516.
//
// Solidity: function dids(string ) view returns(bool exist, bytes32 hash, string uri)
func (_IoTeXDIDStorage *IoTeXDIDStorageCallerSession) Dids(arg0 string) (struct {
	Exist bool
	Hash  [32]byte
	Uri   string
}, error) {
	return _IoTeXDIDStorage.Contract.Dids(&_IoTeXDIDStorage.CallOpts, arg0)
}
