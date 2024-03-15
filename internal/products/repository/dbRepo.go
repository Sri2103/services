package repo

import (
	"context"
	"errors"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/Sri2103/services/pkg/database"
	"github.com/Sri2103/services/pkg/ent"
	"github.com/google/uuid"
)

type dbImpl struct {
	DB     *database.DB
	client *ent.ProductClient
}

func setUpEnt(db *database.DB) *ent.ProductClient {
	drv := entsql.OpenDB(dialect.Postgres, db.Conn)
	client := ent.NewClient(ent.Driver(drv)).Product
	return client
}

func New(db *database.DB) Repo {
	return &dbImpl{
		DB:     db,
		client: setUpEnt(db),
	}
}

func (d *dbImpl) GetProduct(ctx context.Context, id uuid.UUID) (*ent.Product, error) {

	p, err := d.client.Get(ctx, id)
	if err != nil {
		return nil, errors.New("product not found")
	}
	return p, nil

}

func (d *dbImpl) GetProductList(ctx context.Context) ([]*ent.Product, error) {
	products, err := d.client.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (d *dbImpl) CreateProduct(ctx context.Context, p *ent.Product) (*ent.Product, error) {
	created, err := d.client.Create().SetName(p.Name).SetPrice(p.Price).SetDescription(p.Description).Save(ctx)
	if err != nil {
		return nil, err
	}
	return created, nil
}

