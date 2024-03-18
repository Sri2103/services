package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Sri2103/services/internal/users/repository"
	"github.com/Sri2103/services/internal/users/service"
	"github.com/Sri2103/services/pkg/database"
	user_pb "github.com/Sri2103/services/pkg/rpc/user"
	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":8091") // listen on port 8081
	if err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
	dsn := "host=postgres  port=5432 dbname=services user=postgres password=harsha  sslmode=disable"
	DB, err := database.ConnectSQL(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()

	user_pb.RegisterUserServiceServer(server, service.New(repository.NewDB(DB)))
	fmt.Println("Starting grpc server for user")
	if err := server.Serve(listen); err != nil {
		log.Fatalf("Failed to start the grpc server")
	}
}
