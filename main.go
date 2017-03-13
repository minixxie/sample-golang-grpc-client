package main

import "log"
import "os"
import "golang.org/x/net/context"
import "google.golang.org/grpc"
import pb "pb"

const (
	ldevAddress    = "ldev-sample-golang-grpc-microservice:80"
	devAddress    = "dev-sample-golang-grpc-microservice:80"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
    address := ldevAddress
    if os.Getenv("ENV") == "dev" {
        address = devAddress
    }

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
