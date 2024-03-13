package service

import (
	"context"

	order_pb "github.com/Sri2103/services/pkg/rpc/order"
)

type orderImpl struct {
	order_pb.UnimplementedOrderServiceServer
}

func New() order_pb.OrderServiceServer {
	return &orderImpl{}
}

func (o *orderImpl) CreateOrder(context.Context, *order_pb.OrderRequest) (*order_pb.OrderResponse, error) {
	panic("Not implemented")
}

func (o *orderImpl) GetOrder(context.Context, *order_pb.GetOrderRequest) (*order_pb.OrderResponse, error) {
	panic("Not implemented")
}

func (o *orderImpl) ListOrders(context.Context, *order_pb.ListOrdersRequest) (*order_pb.ListOrdersResponse, error) {
	panic("Not implemented")
}

func (o *orderImpl) UpdateOrder(context.Context, *order_pb.OrderRequest) (*order_pb.OrderResponse, error) {
	panic("to implement")
}

// Delete an order
func (o *orderImpl) DeleteOrder(context.Context, *order_pb.DeleteOrderRequest) (*order_pb.DeleteOrderResponse, error) {
	panic("to implement")
}
