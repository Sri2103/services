package main

import (
	"fmt"
	"log"
	"net"

	repo "github.com/Sri2103/services/internal/products/repository"
	productImplementation "github.com/Sri2103/services/internal/products/service"
	product_pb "github.com/Sri2103/services/pkg/rpc/product"
	"google.golang.org/grpc"
)

func main() {
	// create a listener
	listen, err := net.Listen("tcp", ":8081") // listen on port 8081
	if err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
	server := grpc.NewServer()

	product_pb.RegisterProductServiceServer(server, productImplementation.New(repo.NewData()))
	fmt.Println("Starting grpc server for products")
	if err := server.Serve(listen); err != nil {
		log.Fatalf("Failed to start the grpc server")
	}

}
