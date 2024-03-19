package main

import (
	"fmt"
	"log"
	"net"

	repository "github.com/Sri2103/services/internal/cart/repo"
	"github.com/Sri2103/services/internal/cart/service"
	"github.com/Sri2103/services/pkg/database"
	cart_pb "github.com/Sri2103/services/pkg/rpc/cart"
	"google.golang.org/grpc"
)

func main() {
	// create a listener for the proto listener
	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	dsn := "host=postgres  port=5432 dbname=services user=postgres password=harsha  sslmode=disable"
	DB, err := database.ConnectSQL(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Conn.Close()
	server := grpc.NewServer()

	cart_pb.RegisterCartServiceServer(server, service.New(repository.New(DB)))
	fmt.Println("Starting grpc server for cart")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to start the grpc server on cart")
	}

}
