package main

import (
	"context"

	"github.com/iotexproject/iotex-proto/golang/iotexapi"

	"github.com/iotexproject/iotex-antenna-go/v2/examples/service"
)

type GetInfoExample interface {
	GetChainMeta(ctx context.Context, in *iotexapi.GetChainMetaRequest) (*iotexapi.GetChainMetaResponse, error)
	GetBlockMetas(ctx context.Context, in *iotexapi.GetBlockMetasRequest) (*iotexapi.GetBlockMetasResponse, error)
	GetActions(ctx context.Context, in *iotexapi.GetActionsRequest) (*iotexapi.GetActionsResponse, error)
}

type iotexService struct {
	service.IotexService
}

func NewIotexService(accountPrivate, endpoint string, secure bool) GetInfoExample {
	return &iotexService{
		service.NewIotexService(accountPrivate, endpoint, secure),
	}
}

func (s *iotexService) GetChainMeta(ctx context.Context, in *iotexapi.GetChainMetaRequest) (*iotexapi.GetChainMetaResponse, error) {
	err := s.Connect()
	if err != nil {
		return nil, err
	}
	return s.ReadOnlyClient().API().GetChainMeta(ctx, in)
}

func (s *iotexService) GetBlockMetas(ctx context.Context, in *iotexapi.GetBlockMetasRequest) (*iotexapi.GetBlockMetasResponse, error) {
	err := s.Connect()
	if err != nil {
		return nil, err
	}
	return s.ReadOnlyClient().API().GetBlockMetas(ctx, in)
}

func (s *iotexService) GetActions(ctx context.Context, in *iotexapi.GetActionsRequest) (*iotexapi.GetActionsResponse, error) {
	err := s.Connect()
	if err != nil {
		return nil, err
	}
	return s.ReadOnlyClient().API().GetActions(ctx, in)
}
