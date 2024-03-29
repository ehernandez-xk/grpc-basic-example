package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/ehernandez-xk/grpc-basic-example/myservice"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (a *server) GetGreeter(ctx context.Context, in *pb.MyGreeter) (*pb.TheReplay, error) {
	return &pb.TheReplay{Reply: "Hi, this is a message from the server"}, nil
}

func main() {
	fmt.Println("server on port: ", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}
