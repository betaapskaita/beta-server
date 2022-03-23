package handler

import (
	"context"

	"github.com/betaapskaita/beta-server/libs/rpcs"
	"github.com/betaapskaita/beta-server/proto"
	"github.com/betaapskaita/beta-server/repositories"
)

type Greeter struct {
	repository repositories.AccountRepository
}

func NewGreeterServer(s *rpcs.RpcServer, repository repositories.AccountRepository) *Greeter {
	server := &Greeter{repository}
	if s != nil {
		proto.RegisterAccountServiceServer(s.Grpc, server)
	}
	return server
}

func (g *Greeter) AuthenticateByEmailAndPassword(ctx context.Context, req *proto.User) (*proto.Account, error) {
	return &proto.Account{
		Token: "token is got",
	}, nil
}
