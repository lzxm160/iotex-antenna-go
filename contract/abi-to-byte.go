// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package contract

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto"
)

//
//export function getAbiFunctions(abi: Array<any>): AbiByFunc {
//  const abiFunctions = ({} as any) as AbiByFunc;
//  abi.forEach(f => {
//    if (f.type === "function") {
//      abiFunctions[f.name] = f;
//    }
//  });
//
//  return abiFunctions;
//}
//
//function getArgTypes(fnAbi: {
//  inputs: Array<{ name: string; type: string }>;
//}): Array<EthAbiDecodeParametersType> {
//  const args = [] as Array<EthAbiDecodeParametersType>;
//  fnAbi.inputs.forEach(field => {
//    args.push({ name: field.name, type: field.type });
//  });
//  return args;
//}
//
//export function getHeaderHash(
//  fnAbi: any,
//  args: Array<EthAbiDecodeParametersType>
//): string {
//  const inputs = args.map(i => {
//    return i.type;
//  });
//  const signature = `${fnAbi.name}(${inputs.join(",")})`;
//  const keccak256 = hash256b(signature).toString("hex");
//  return keccak256.slice(0, 8);
//}
//
func GetFuncHash(fun string) string {
	return hex.EncodeToString(crypto.Keccak256([]byte(fun))[:4])
}

//export function encodeArguments(
//  args: Array<EthAbiDecodeParametersType>,
//  userInput: UserInput
//): string {
//  const types = [] as Array<any>;
//  const values = [] as Array<any>;
//
//  (args || []).forEach(arg => {
//    if (arg.type === "bool") {
//      types.push("uint256");
//    } else {
//      types.push(arg.type);
//    }
//    if (userInput.hasOwnProperty(arg.name)) {
//      let value = userInput[arg.name];
//      if (arg.type === "address") {
//        value = address.fromString(value).stringEth();
//      }
//      values.push(value);
//    } else {
//      values.push("");
//    }
//  });
//  try {
//    const encoded = ethereumjs.rawEncode(types, values);
//    return encoded.toString("hex");
//  } catch (e) {
//    throw new Error(`failed to rawEncode: ${e.stack}`);
//  }
//}
//
//type UserInput = {
//  [key: string]: any;
//};
//
//export function encodeInputData(
//  abiByFunc: AbiByFunc,
//  fnName: string,
//  userInput: UserInput
//): string {
//  const fnAbi = abiByFunc[fnName];
//  const args = getArgTypes(fnAbi);
//  const header = getHeaderHash(fnAbi, args);
//  const encodedArgs = encodeArguments(args, userInput);
//  return `${header}${encodedArgs}`;
//}
