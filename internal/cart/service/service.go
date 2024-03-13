package service

import (
	"context"
	"fmt"

	cart_pb "github.com/Sri2103/services/pkg/rpc/cart"
)

type service struct {
	cart_pb.UnimplementedCartServiceServer
}

func New() cart_pb.CartServiceServer {
	return &service{}
}

func (s *service) AddItem(ctx context.Context, req *cart_pb.AddItemRequest) (*cart_pb.AddItemResponse, error) {
	// TODO: Implement your logic here
	fmt.Println("Got a request to add an item to the cart")
	return &cart_pb.AddItemResponse{}, nil
}

func (s *service) GetCart(ctx context.Context, req *cart_pb.GetCartRequest) (*cart_pb.GetCartResponse, error) {
	return &cart_pb.GetCartResponse{}, nil
}

func (s *service) RemoveItem(context.Context, *cart_pb.RemoveItemRequest) (*cart_pb.RemoveItemResponse, error) {
	return &cart_pb.RemoveItemResponse{}, nil
}


