#!/bin/bash
apt-get update
sudo apt-get install openssl -y
sudo apt-get install libssl-dev -y
sudo apt-get install autoconf -y
cd vendor/github.com/ethereum/go-ethereum/crypto/secp256k1/libsecp256k1
./autogen.sh
./configure
make
sudo make install