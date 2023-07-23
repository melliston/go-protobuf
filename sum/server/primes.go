package main

import (
	"log"

	pb "github.com/melliston/go-protobuf/sum/proto"
)

func (s *Server) Primes(in *pb.PrimeRequest, stream pb.SumService_PrimesServer) error {
	log.Printf("invoked Primes() with%+v\n", in)

	n := in.Number
	k := int64(2) // Starting point.

	for {
		if n > 1 {
			if n%k == 0 {
				stream.Send(&pb.PrimeResponse{
					Prime: k,
				})
				n = n / k
			} else {
				k = k + 1
			}
		} else {
			break
		}
	}

	return nil
}
