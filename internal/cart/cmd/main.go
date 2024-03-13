package main

import (
	"log"
	"net"

	"github.com/Sri2103/services/internal/cart/service"
	cart_pb "github.com/Sri2103/services/pkg/rpc/cart"
	"google.golang.org/grpc"
)

func main() {
	// create a listener for the proto listener
	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	server := grpc.NewServer()

	cart_pb.RegisterCartServiceServer(server, service.New())
	server.Serve(listener) // start serving
}
