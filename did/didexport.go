package main

import (
	"C"
	"fmt"
	"math/big"
)

const (
	Failure = iota
	Success
)

//export CeateDID
func CeateDID(endpoint, privateKey, contract, abiString, gasPrice string, gasLimit uint64, id, didHash, url string) (*C.char, uint64, *C.char) {
	gp, ok := new(big.Int).SetString(gasPrice, 10)
	if !ok {
		return C.CString(""), Failure, C.CString("gas price convert error")
	}
	d, err := NewDID(endpoint, privateKey, contract, abiString, gp, gasLimit)
	if err != nil {
		return C.CString(""), Failure, C.CString(err.Error())
	}
	h, err := d.CreateDID(id, didHash, url)
	if err != nil {
		return C.CString(""), Failure, C.CString(err.Error())
	}
	fmt.Println(h)
	return C.CString(h), Success, C.CString("")
}

//export DeleteDID
func DeleteDID(endpoint, privateKey, contract, abiString, gasPrice string, gasLimit uint64, did string) (*C.char, uint64, *C.char) {
	gp, ok := new(big.Int).SetString(gasPrice, 10)
	if !ok {
		return C.CString(""), Failure, C.CString("gas price convert error")

	}
	d, err := NewDID(endpoint, privateKey, contract, abiString, gp, gasLimit)
	if err != nil {
		return C.CString(""), Failure, C.CString(err.Error())
	}
	h, err := d.DeleteDID(did)
	if err != nil {
		return C.CString(""), Failure, C.CString(err.Error())
	}
	fmt.Println(h)
	return C.CString(h), Success, C.CString("")
}

//export UpdateHash
func UpdateHash(endpoint, privateKey, contract, abiString, gasPrice string, gasLimit uint64, did, didHash string) (*C.char, uint64, *C.char) {
	gp, ok := new(big.Int).SetString(gasPrice, 10)
	if !ok {
		return C.CString(""), Failure, C.CString("gas price convert error")
	}
	d, err := NewDID(endpoint, privateKey, contract, abiString, gp, gasLimit)
	if err != nil {
		return C.CString(""), Failure, C.CString(err.Error())
	}
	h, err := d.UpdateHash(did, didHash)
	if err != nil {
		return C.CString(""), Failure, C.CString(err.Error())
	}
	fmt.Println(h)
	return C.CString(h), Success, C.CString("")
}

//export UpdateUri
func UpdateUri(endpoint, privateKey, contract, abiString, gasPrice string, gasLimit uint64, did, uri string) (*C.char, uint64, *C.char) {
	gp, ok := new(big.Int).SetString(gasPrice, 10)
	if !ok {
		return C.CString(""), Failure, C.CString("gas price convert error")
	}
	d, err := NewDID(endpoint, privateKey, contract, abiString, gp, gasLimit)
	if err != nil {
		return C.CString(""), Failure, C.CString(err.Error())
	}
	h, err := d.UpdateUri(did, uri)
	if err != nil {
		return C.CString(""), Failure, C.CString(err.Error())
	}
	fmt.Println(h)
	return C.CString(h), Success, C.CString("")
}

//export GetHash
func GetHash(endpoint, contract, abiString, did string) (*C.char, uint64, *C.char) {
	fmt.Println(endpoint, contract, did, abiString)
	d, err := NewDID(endpoint, "", contract, abiString, nil, 0)
	if err != nil {
		return C.CString(""), Failure, C.CString(err.Error())
	}
	h, err := d.GetHash(did)
	if err != nil {
		return C.CString(""), Failure, C.CString(err.Error())
	}
	fmt.Println(h)
	return C.CString(h), Success, C.CString("")
}

//export GetUri
func GetUri(endpoint, contract, abiString, did string) (*C.char, uint64, *C.char) {
	d, err := NewDID(endpoint, "", contract, abiString, nil, 0)
	if err != nil {
		return C.CString(""), Failure, C.CString(err.Error())
	}
	uri, err := d.GetUri(did)
	if err != nil {
		return C.CString(""), Failure, C.CString(err.Error())
	}
	fmt.Println(uri)
	return C.CString(uri), Success, C.CString("")
}
func main() {}
