package repo

import (
	"context"
	"errors"

	"github.com/Sri2103/services/pkg/ent"
	"github.com/google/uuid"
)

var products []*ent.Product = []*ent.Product{
	{Name: "Laptop", Price: 1500, ID: uuid.Must(uuid.Parse("07e4b48c-1686-4bd6-84ac-313e228473f9"))},
	{Name: "Smartphone", Price: 800, ID: uuid.Must(uuid.Parse("31a60615-4257-4bc2-9230-79a02c801286"))},
}

type data struct {
	Products []*ent.Product
}

// GetCategorizedProducts implements Repo.
func (d *data) GetCategorizedProducts(ctx context.Context, id uuid.UUID, pageNumber int, pageSize int, sort string) ([]*ent.Product, int, error) {
	panic("unimplemented")
}

// GetProductsByCategory implements Repo.
func (d *data) GetProductsByCategory(ctx context.Context, id uuid.UUID, pageNumber int, pageSize int, sort string) ([]*ent.Product, int, error) {
	panic("unimplemented")
}

// CreateCategory implements Repo.
func (d *data) CreateCategory(ctx context.Context, c *ent.Category) (*ent.Category, error) {
	panic("unimplemented")
}

// DeleteCategory implements Repo.
func (d *data) DeleteCategory(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

// GetCategories implements Repo.
func (d *data) GetCategories(ctx context.Context) ([]*ent.Category, error) {
	panic("unimplemented")
}

// GetCategory implements Repo.
func (d *data) GetCategory(ctx context.Context, id uuid.UUID) (*ent.Category, error) {
	panic("unimplemented")
}

func NewData() Repo {
	d := &data{Products: products}
	return d
}

func (d *data) GetProduct(ctx context.Context, id uuid.UUID) (*ent.Product, error) {
	for _, v := range d.Products {
		if v.ID == id {
			return v, nil
		}
	}
	return nil, errors.New("product not found")
}

func (d *data) GetProductList(ctx context.Context) ([]*ent.Product, error) {
	return d.Products, nil
}

func (d *data) CreateProduct(ctx context.Context, p *ent.Product, _ *ent.Category) (*ent.Product, error) {
	panic("not implemented")
}

func (d *data) UpdateProduct(ctx context.Context, p *ent.Product) (*ent.Product, error) {
	panic("not implemented")
}
