package main

import (
	"C"
	"math/big"
)

const (
	Failure = 0
	Success = 1
)

//export CeateDID returns transaction hash,transaction if success,error message
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
	return C.CString(h), Success, C.CString("")
}

//export DeleteDID returns transaction hash,transaction if success,error message
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
	return C.CString(h), Success, C.CString("")
}

//export UpdateHash returns transaction hash,transaction if success,error message
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
	return C.CString(h), Success, C.CString("")
}

//export UpdateUri returns transaction hash,transaction if success,error message
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
	return C.CString(h), Success, C.CString("")
}

//export GetHash returns did hash,transaction if success,error message
func GetHash(endpoint, contract, abiString, did string) (*C.char, uint64, *C.char) {
	d, err := NewDID(endpoint, "", contract, abiString, nil, 0)
	if err != nil {
		return C.CString(""), Failure, C.CString(err.Error())
	}
	h, err := d.GetHash(did)
	if err != nil {
		return C.CString(""), Failure, C.CString(err.Error())
	}
	return C.CString(h), Success, C.CString("")
}

//export GetUri returns did uri,transaction if success,error message
func GetUri(endpoint, contract, abiString, did string) (*C.char, uint64, *C.char) {
	d, err := NewDID(endpoint, "", contract, abiString, nil, 0)
	if err != nil {
		return C.CString(""), Failure, C.CString(err.Error())
	}
	uri, err := d.GetUri(did)
	if err != nil {
		return C.CString(""), Failure, C.CString(err.Error())
	}
	return C.CString(uri), Success, C.CString("")
}
func main() {}
