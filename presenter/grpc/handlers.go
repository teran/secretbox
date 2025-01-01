package grpc

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/teran/secretbox/presenter/grpc/proto/v1"
	"github.com/teran/secretbox/service"
)

type Handlers interface {
	proto.SecretBoxServiceServer

	Register(*grpc.Server)
}

type handlers struct {
	proto.UnimplementedSecretBoxServiceServer
	svc service.Service
}

func New(svc service.Service) Handlers {
	return &handlers{
		svc: svc,
	}
}

func (h *handlers) GetSecret(ctx context.Context, in *proto.GetSecretRequest) (*proto.GetSecretResponse, error) {
	secret, err := h.svc.GetSecret(ctx, in.GetName(), in.GetToken())
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &proto.GetSecretResponse{
		Secret: secret,
	}, nil
}

func (h *handlers) Register(srv *grpc.Server) {
	proto.RegisterSecretBoxServiceServer(srv, h)
}
