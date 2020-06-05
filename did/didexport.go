package main

import (
	"C"
	"fmt"
	"math/big"
)

//export CeateDID
func CeateDID(endpoint, privateKey, contract, abiString, gasPrice string, gasLimit uint64, id, didHash, url string) *C.char {
	gp, ok := new(big.Int).SetString(gasPrice, 10)
	if !ok {
		return C.CString("")
	}
	d, err := NewDID(endpoint, privateKey, contract, abiString, gp, gasLimit)
	if err != nil {
		return C.CString("")
	}
	h, err := d.CreateDID(id, didHash, url)
	if err != nil {
		return C.CString("")
	}
	return C.CString(h)
}

//export DeleteDID
func DeleteDID(endpoint, privateKey, contract, abiString, gasPrice string, gasLimit uint64, did string) *C.char {
	gp, ok := new(big.Int).SetString(gasPrice, 10)
	if !ok {
		return C.CString("")

	}
	d, err := NewDID(endpoint, privateKey, contract, abiString, gp, gasLimit)
	if err != nil {
		return C.CString("")
	}
	h, err := d.DeleteDID(did)
	if err != nil {
		return C.CString("")
	}
	return C.CString(h)
}

//export UpdateHash
func UpdateHash(endpoint, privateKey, contract, abiString, gasPrice string, gasLimit uint64, did, didHash string) *C.char {
	gp, ok := new(big.Int).SetString(gasPrice, 10)
	if !ok {
		return C.CString("")
	}
	d, err := NewDID(endpoint, privateKey, contract, abiString, gp, gasLimit)
	if err != nil {
		C.CString("")
	}
	h, err := d.UpdateHash(did, didHash)
	if err != nil {
		return C.CString("")
	}
	return C.CString(h)
}

//export UpdateUri
func UpdateUri(endpoint, privateKey, contract, abiString, gasPrice string, gasLimit uint64, did, uri string) *C.char {
	gp, ok := new(big.Int).SetString(gasPrice, 10)
	if !ok {
		return C.CString("")
	}
	d, err := NewDID(endpoint, privateKey, contract, abiString, gp, gasLimit)
	if err != nil {
		return C.CString("")
	}
	h, err := d.UpdateUri(did, uri)
	if err != nil {
		return C.CString("")
	}
	return C.CString(h)
}

//export GetHash
func GetHash(endpoint, contract, abiString, did string) *C.char {
	fmt.Println(endpoint, contract, did, abiString)
	d, err := NewDID(endpoint, "", contract, abiString, nil, 0)
	if err != nil {
		return C.CString("")
	}
	h, err := d.GetHash(did)
	if err != nil {
		return C.CString("")
	}
	return C.CString(h)
}

//export GetUri
func GetUri(endpoint, contract, abiString, did string) *C.char {
	d, err := NewDID(endpoint, "", contract, abiString, nil, 0)
	if err != nil {
		return C.CString("")
	}
	uri, err := d.GetUri(did)
	if err != nil {
		return C.CString("")
	}
	return C.CString(uri)
}
func main() {}
