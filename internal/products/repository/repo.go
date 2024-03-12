package repo

import (
	"context"

	"github.com/Sri2103/services/internal/products/model"
)

type Repo interface {
	GetProduct(ctx context.Context, id int) (*model.Product, error)
	GetProductList(ctx context.Context) ([]*model.Product, error)
}
