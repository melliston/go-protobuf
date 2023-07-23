package main

import (
	"context"
	"log"

	pb "github.com/melliston/go-protobuf/sum/proto"
)

func doSum(c pb.SumServiceClient) {
	log.Println("doSum was invoked")
	req := &pb.SumRequest{
		Number1: 6,
		Number2: 10,
	}
	res, err := c.Sum(context.Background(), req)

	if err != nil {
		log.Fatalf("err unable to process request to server: %v\n", err)
	}

	log.Printf("sum %d + %d = %d\n", req.Number1, req.Number2, res.Sum)
}
