package main

import (
	"context"

	"github.com/iotexproject/iotex-antenna-go/v2/examples/service"
)

const (
	protocolID          = "staking"
	readBucketsLimit    = 30000
	readCandidatesLimit = 20000
)

type xrc20Example interface {
	Transfer(ctx context.Context, to string) (string, error)
}

type iotexService struct {
	service.IotexService
}

func NewIotexService(accountPrivate, endpoint string, secure bool) xrc20Example {
	return &iotexService{
		service.NewIotexService(accountPrivate, endpoint, secure),
	}
}
func (s *iotexService) Transfer(ctx context.Context, to string) (string, error) {

}
