package main

import (
	"C"
	"math/big"
)

//export CeateDID
func CeateDID(endpoint, privateKey, contract, abiString string, gasPrice *big.Int, gasLimit uint64, id, didHash, url string) (hash string, err error) {
	d, err := NewDID(endpoint, privateKey, contract, abiString, gasPrice, gasLimit)
	if err != nil {
		return
	}
	return d.CreateDID(id, didHash, url)
}

//export DeleteDID
func DeleteDID(endpoint, privateKey, contract, abiString string, gasPrice *big.Int, gasLimit uint64, did string) (hash string, err error) {
	d, err := NewDID(endpoint, privateKey, contract, abiString, gasPrice, gasLimit)
	if err != nil {
		return
	}
	return d.DeleteDID(did)
}

//export UpdateHash
func UpdateHash(endpoint, privateKey, contract, abiString string, gasPrice *big.Int, gasLimit uint64, did, didHash string) (hash string, err error) {
	d, err := NewDID(endpoint, privateKey, contract, abiString, gasPrice, gasLimit)
	if err != nil {
		return
	}
	return d.UpdateHash(did, didHash)
}

//export UpdateUri
func UpdateUri(endpoint, privateKey, contract, abiString string, gasPrice *big.Int, gasLimit uint64, did, uri string) (hash string, err error) {
	d, err := NewDID(endpoint, privateKey, contract, abiString, gasPrice, gasLimit)
	if err != nil {
		return
	}
	return d.DeleteDID(did)
}

//export GetHash
func GetHash(endpoint, privateKey, contract, abiString string, gasPrice *big.Int, gasLimit uint64, did string) (hash string, err error) {
	d, err := NewDID(endpoint, privateKey, contract, abiString, gasPrice, gasLimit)
	if err != nil {
		return
	}
	return d.GetHash(did)
}

//export GetUri
func GetUri(endpoint, privateKey, contract, abiString string, gasPrice *big.Int, gasLimit uint64, did string) (uri string, err error) {
	d, err := NewDID(endpoint, privateKey, contract, abiString, gasPrice, gasLimit)
	if err != nil {
		return
	}
	return d.GetUri(did)
}
