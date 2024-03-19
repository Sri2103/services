package main

import (
	"log"
	"net"

	"github.com/Sri2103/services/internal/order/repository"
	"github.com/Sri2103/services/internal/order/service"
	"github.com/Sri2103/services/pkg/database"
	order_pb "github.com/Sri2103/services/pkg/rpc/order"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatal(err)
	}
	dsn := "host=postgres  port=5432 dbname=services user=postgres password=harsha  sslmode=disable"
	DB, err := database.ConnectSQL(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Conn.Close()

	server := grpc.NewServer()

	order_pb.RegisterOrderServiceServer(server, service.New(repository.New(DB)))

	log.Println("Starting order server  grpc on port 8083...")
	if err = server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve order service: %v", err)
	}
}
