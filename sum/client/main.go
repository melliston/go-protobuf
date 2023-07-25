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

	// Receive the cmd line args
	// if len(os.Args) < 2 {
	// 	log.Fatalf("usage: executable.file <number>\n")
	// }
	// arg := os.Args[1]
	// number, err := strconv.Atoi(arg)
	// if err != nil {
	// 	log.Fatalf("unable to convert to int64 %s\n", arg)
	// }
	// doPrimes(c, int64(number))

	numbers := []int64{1, 5, 3, 22}
	doAverage(c, numbers)
}
