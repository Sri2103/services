package repo

import (
	"context"

	"github.com/Sri2103/services/pkg/ent"
	"github.com/google/uuid"
)

type Repo interface {
	GetProduct(ctx context.Context, id uuid.UUID) (*ent.Product, error)
	GetProductList(ctx context.Context) ([]*ent.Product, error)
	CreateProduct(ctx context.Context, p *ent.Product, c *ent.Category) (*ent.Product, error)
	UpdateProduct(ctx context.Context, p *ent.Product) (*ent.Product, error)
	CreateCategory(ctx context.Context, c *ent.Category) (*ent.Category, error)
	GetCategories(ctx context.Context) ([]*ent.Category, error)
	GetCategory(ctx context.Context, id uuid.UUID) (*ent.Category, error)
	DeleteCategory(ctx context.Context, id uuid.UUID) error
	GetProductsByCategory(ctx context.Context, id uuid.UUID, pageNumber int, pageSize int, sort string) ([]*ent.Product, int, error)
	GetCategorizedProducts(ctx context.Context, id uuid.UUID, pageNumber int, pageSize int, sort string) ([]*ent.Product, int, error)
}
