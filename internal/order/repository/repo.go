package repository

import (
	"context"

	"github.com/Sri2103/services/pkg/ent"
	"github.com/google/uuid"
)

type Repository interface {
	CreateOrder(ctx context.Context, order *ent.Order) (*ent.Order, error)
	GetOrder(ctx context.Context, id uuid.UUID) (*ent.Order, error)
	ListOrders(ctx context.Context) ([]*ent.Order, error)
	UpdateOrder(ctx context.Context, order *ent.Order) (*ent.Order, error)
	DeleteOrder(ctx context.Context, id uuid.UUID) error
}
