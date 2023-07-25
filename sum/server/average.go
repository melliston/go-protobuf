package main

import (
	"io"
	"log"

	pb "github.com/melliston/go-protobuf/sum/proto"
)

func (s *Server) Average(stream pb.SumService_AverageServer) error {
	log.Printf("invoked Average()")
	var total int64
	var count int64
	for {
		count++
		msg, err := stream.Recv()

		if err == io.EOF {
			// Return the average
			average := total / count
			stream.SendMsg(&pb.AverageResponse{
				Average: float64(average),
			})
			break
		}
		if err != nil {
			log.Fatalf("err receiving stream: %+v", err)
		}
		total += msg.Number

	}
	return nil
}
