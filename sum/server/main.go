package main

import (
	"log"
	"net"

	pb "github.com/melliston/go-protobuf/sum/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.SumServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("unable to listen on %s: %v\n", addr, err)
	}

	log.Printf("listening on %s", addr)

	s := grpc.NewServer()
	pb.RegisterSumServiceServer(s, &Server{})
	log.Println("service registered")
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v\n", err.Error())
	}
}
