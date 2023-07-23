package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/melliston/go-protobuf/sum/proto"
)

var addr string = "0.0.0.0:50051"

func main() {
	// Create a new connection with no SSL - For use when testing
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewSumServiceClient(conn) // Should really have named the service something more generic than summing.
	// doSum(c)
	doPrimes(c)

}
