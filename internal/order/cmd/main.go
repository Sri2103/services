package main

import (
	"log"
	"net"

	"github.com/Sri2103/services/internal/order/service"
	order_pb "github.com/Sri2103/services/pkg/rpc/order"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	order_pb.RegisterOrderServiceServer(server, service.New())

	log.Println("Starting server on port 8083...")
	if err = server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
