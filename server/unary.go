package main

import (
	"context"

	pb "github.com/i101dev/template-gRPC/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {

	res := &pb.HelloResponse{
		Message: "@server.unary - SayHello - success",
	}

	return res, nil
}
