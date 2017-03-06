package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "pb"
)

const (
	address     = "172.17.0.1:8080"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTimeServiceClient(conn)

	r, err := c.Now(context.Background(), &pb.NowRequest{Msg: "Hi Simon"})
	if err != nil {
		log.Fatalf("could not get Now: %v", err)
	}
	log.Printf("Now: %s", r.Now)
}
