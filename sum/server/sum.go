package main

import (
	"context"
	"log"

	pb "github.com/melliston/go-protobuf/sum/proto"
)

func (*Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("greet function was invoked")
	return &pb.SumResponse{
		Sum: in.Number1 + in.Number2,
	}, nil
}
