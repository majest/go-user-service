package main

import (
	"github.com/go-kit/kit/transport/grpc"
	"github.com/majest/user-service/pb"
	"github.com/majest/user-service/server"
	"golang.org/x/net/context"
)

type grpcBinding struct {
	findOne grpc.Handler
	save    grpc.Handler
}

func newGRPCBinding(ctx context.Context, svc server.UserService) grpcBinding {
	return grpcBinding{
		findOne: grpc.NewServer(ctx, server.MakeFindOneEndpoint(svc), server.DecodeFindOneRequest, server.EncodeFindOneResponse),
	}
}

func (b grpcBinding) FindOne(ctx context.Context, req *pb.UserFindOneRequest) (*pb.UserResponse, error) {
	_, resp, err := b.findOne.ServeGRPC(ctx, req)
	return resp.(*pb.UserResponse), err
}

func (b grpcBinding) Save(ctx context.Context, req *pb.UserSaveRequest) (*pb.UserSaveResponse, error) {
	_, resp, err := b.save.ServeGRPC(ctx, req)
	return resp.(*pb.UserSaveResponse), err
}
