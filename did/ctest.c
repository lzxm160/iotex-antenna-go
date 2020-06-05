#include <stdio.h>
#include <string.h>
#include "didexport.h"

int main ()
{
    char privateKey[] = "414efa99dfac6f4095d6954713fb0085268d400d6a05a8ae8a69b5b1c10b4bed"
	char endpoint[] = "api.testnet.iotex.one:443"
	char contract[] = "io1zgs5gqjl679qlj4gqqpa9t329r8f5gr8xc9lr0"
	char did[] = "did:io:0x0ddfC506136fb7c050Cc2E9511eccD81b15e7426"
	char abi[] = ""
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
    GetHash(Endpoint,PrivateKey,Contract,Abi,);
    return 0;
}