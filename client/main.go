package main

import (
	"log"

	pb "github.com/ehernandez-xk/basic-grpc/myservice"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	r, err := c.GetGreeter(context.Background(), &pb.MyGreeter{Name: "eddy"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("greeting: %v", r.Reply)
}
