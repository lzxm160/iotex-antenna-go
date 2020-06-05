#include <stdio.h>
#include <string.h>
#include "../didlib.h"

int main ()
{
    char privateKey[] = "414efa99dfac6f4095d6954713fb0085268d400d6a05a8ae8a69b5b1c10b4bed";
	char endpoint[] = "api.testnet.iotex.one:443";
	char contract[] = "io1zgs5gqjl679qlj4gqqpa9t329r8f5gr8xc9lr0";
	char did[] = "did:io:0x0ddfC506136fb7c050Cc2E9511eccD81b15e7426";
	char abi[] = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"didString\",\"type\":\"string\"}],\"name\":\"CreateDID\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"didString\",\"type\":\"string\"}],\"name\":\"DeleteDID\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"didString\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"UpdateHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"didString\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"UpdateURI\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"createDID\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"deleteDID\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"dids\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"getHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"getURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"updateHash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"updateURI\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]";
    char gasPrice[] = "1000000000000"
    char uri[] = "urixxx"
    char id[] = ""
    GoString PrivateKey;
    PrivateKey.p = privateKey;
    PrivateKey.n = strlen(privateKey);
    GoString Endpoint;
    Endpoint.p = endpoint;
    Endpoint.n = strlen(endpoint);
    GoString Contract;
    Contract.p = contract;
    Contract.n = strlen(contract);
    GoString Abi;
    Abi.p = abi;
    Abi.n = strlen(abi);
    GoString Did;
    Did.p = did;
    Did.n = strlen(did);
    GoString GasPrice;
    GasPrice.p = gasPrice;
    GasPrice.n = strlen(gasPrice);
    GoString Uri;
    Uri.p = uri;
    Uri.n = strlen(uri);
    GoString id;
    Id.p = id;
    Id.n = strlen(id);
    //CeateDID
    char* ret=CeateDID(Endpoint,PrivateKey,Contract,Abi,GasPrice,1000000,Id,PrivateKey,Uri);
    printf("CeateDID %s\n",ret);
    //GetHash
    ret=GetHash(Endpoint,Contract,Abi,Did);
    printf("GetHash %s\n",ret);

    return 0;
}