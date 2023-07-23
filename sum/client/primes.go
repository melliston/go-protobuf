package main

import (
	"context"
	"io"
	"log"

	pb "github.com/melliston/go-protobuf/sum/proto"
)

func doPrimes(c pb.SumServiceClient, number int64) {
	log.Printf("doPrimes() was invoked with %d\n", number)

	req := &pb.PrimeRequest{
		Number: number,
	}

	stream, err := c.Primes(context.Background(), req)
	if err != nil {
		log.Fatalf("error invoking Primes() with %+v\n", req)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error receiving message from Primes(), %+v\n", err)
		}
		log.Printf("%d\n", msg.Prime)
	}

}
