package repository

import (
	"context"

	"github.com/Sri2103/services/pkg/ent"
	"github.com/google/uuid"
)

type Repository interface {
	AddItem(ctx context.Context, item *ent.CartItem) (*ent.Cart, error)
	RemoveItem(ctx context.Context, item *ent.CartItem) (*ent.Cart, error)
	DeleteCart(ctx context.Context, id uuid.UUID) error
}
