package main

import (
	"context"
	"fmt"

	"github.com/iotexproject/iotex-proto/golang/iotexapi"
)

func main() {
	s := NewIotexService("73c7b4a62bf165dccf8ebdea8278db811efd5b8638e2ed9683d2d94889450426", "api.testnet.iotex.one:443", true)
	r, err := s.GetChainMeta(context.Background(), &iotexapi.GetChainMetaRequest{})
	fmt.Println("chain info", r, err)
}
