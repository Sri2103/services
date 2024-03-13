package repo

import (
	"context"
	"errors"

	"github.com/Sri2103/services/internal/products/model"
	"github.com/Sri2103/services/pkg/ent"
)

var products []*ent.Product = []*ent.Product{
	&ent.Product{Id: 1, Name: "Laptop", Price: 1500},
	&ent.Product{Id: 2, Name: "Smartphone", Price: 800},
}

type data struct {
	Products []*model.Product
}

func NewData() Repo {
	d := &data{Products: products}
	return d
}

func (d *data) GetProduct(ctx context.Context, id int) (*ent.Product, error) {
	for _, v := range d.Products {
		if v.Id == int32(id) {
			return v, nil
		}
	}
	return nil, errors.New("product not found")
}

func (d *data) GetProductList(ctx context.Context) ([]*ent.Product, error) {
	return d.Products, nil
}
