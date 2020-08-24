package main

import (
	"context"
	"crypto/tls"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials"

	"github.com/iotexproject/iotex-proto/golang/iotexapi"

	"github.com/iotexproject/iotex-antenna-go/v2/account"
	"github.com/iotexproject/iotex-antenna-go/v2/iotex"
)

type IotexService interface {
	GetChainMeta(ctx context.Context, in *iotexapi.GetChainMetaRequest, opts ...grpc.CallOption) (*iotexapi.GetChainMetaResponse, error)
}

type iotexService struct {
	sync.RWMutex
	endpoint       string
	secure         bool
	accountPrivate string

	grpcConn       *grpc.ClientConn
	authedClient   iotex.AuthedClient
	readOnlyClient iotex.ReadOnlyClient
}

func NewIotexService(accountPrivate, endpoint string, secure bool) IotexService {
	return &iotexService{
		endpoint:       endpoint,
		secure:         secure,
		accountPrivate: accountPrivate,
	}
}

func (s *iotexService) GetChainMeta(ctx context.Context, in *iotexapi.GetChainMetaRequest, opts ...grpc.CallOption) (*iotexapi.GetChainMetaResponse, error) {
	err := s.connect()
	if err != nil {
		return nil, err
	}
	if len(opts) != 0 {
		return s.readOnlyClient.API().GetChainMeta(ctx, in, opts...)
	}
	return s.readOnlyClient.API().GetChainMeta(ctx, in)
}

func (s *iotexService) connect() (err error) {
	s.Lock()
	defer s.Unlock()
	// Check if the existing connection is good.
	if s.grpcConn != nil && s.grpcConn.GetState() != connectivity.Shutdown {
		return
	}
	opts := []grpc.DialOption{}
	if s.secure {
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	s.grpcConn, err = grpc.Dial(s.endpoint, opts...)
	if err != nil {
		return
	}
	if s.accountPrivate != "" {
		creator, err := account.HexStringToAccount(s.accountPrivate)
		if err != nil {
			return err
		}
		s.authedClient = iotex.NewAuthedClient(iotexapi.NewAPIServiceClient(s.grpcConn), creator)
	} else {
		s.readOnlyClient = iotex.NewReadOnlyClient(iotexapi.NewAPIServiceClient(s.grpcConn))
	}
	return
}
