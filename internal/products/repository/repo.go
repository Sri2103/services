package repo

import (
	"context"

	"github.com/Sri2103/services/pkg/ent"
	"github.com/google/uuid"
)

type Repo interface {
	GetProduct(ctx context.Context, id uuid.UUID) (*ent.Product, error)
	GetProductList(ctx context.Context) ([]*ent.Product, error)
}
