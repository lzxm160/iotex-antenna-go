#!/bin/bash
apt-get update
apt-get install autoconf -y
cd vendor/github.com/ethereum/go-ethereum/crypto/secp256k1/libsecp256k1
./autogen.sh
./configure
make
sudo make install