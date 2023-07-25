package main

import (
	"context"
	"log"
	"time"

	pb "github.com/melliston/go-protobuf/sum/proto"
)

func doAverage(c pb.SumServiceClient, numbers []int64) {
	log.Printf("invoked doAverage() with %+v", numbers)

	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("unable to create stream: %v\n", err)
	}

	for _, num := range numbers {
		log.Printf("sending request: %d", num)
		stream.Send(&pb.AverageRequest{
			Number: num,
		})
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error receiving response ")
	}

	log.Printf("average received: %f", res.Average)
}
