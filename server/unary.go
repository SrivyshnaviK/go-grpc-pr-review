package main

import (
	// "log"
	"context"

	pb "vyshnavi/grpc-go/proto"
)


func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParams) (*pb.HelloResponse, error) {

	return &pb.HelloResponse {
		Message: "Hello",
	}, nil
}
