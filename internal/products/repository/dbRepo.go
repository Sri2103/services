package repo

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/Sri2103/services/pkg/database"
	"github.com/Sri2103/services/pkg/ent"
	"github.com/google/uuid"
)

type dbImpl struct {
	DB     *database.DB
	client *ent.Client
}

// GetProductsByCategory implements Repo.
func (d *dbImpl) GetProductsByCategory(ctx context.Context, id uuid.UUID) ([]*ent.Product, error) {
	panic("unimplemented")
}

// CreateCategory implements Repo.
func (d *dbImpl) CreateCategory(ctx context.Context, c *ent.Category) (*ent.Category, error) {
	var created *ent.Category
	tx, err := d.client.Tx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			// Log rollback error instead of ignoring
			fmt.Println("Rollback error", err.Error())
		}
	}()

	created, err = tx.Category.Create().
		SetName(c.Name).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return created, nil

}

// DeleteCategory implements Repo.
func (d *dbImpl) DeleteCategory(ctx context.Context, id uuid.UUID) error {
	// delete the category
	if err := d.client.Category.DeleteOneID(id).Exec(ctx); err != nil {
		return err
	}
	return nil

}

// GetCategories implements Repo.
func (d *dbImpl) GetCategories(ctx context.Context) ([]*ent.Category, error) {
	categories, err := d.client.Category.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

// GetCategory implements Repo.
func (d *dbImpl) GetCategory(ctx context.Context, id uuid.UUID) (*ent.Category, error) {

	c, err := d.client.Category.Get(ctx, id)
	if err != nil {
		return nil, errors.New("category not found")
	}
	return c, nil
}

func setUpEnt(db *database.DB) *ent.Client {
	drv := entsql.OpenDB(dialect.Postgres, db.Conn)
	client := ent.NewClient(ent.Driver(drv))
	return client.Debug()
}

func NewDB(db *database.DB) Repo {
	return &dbImpl{
		DB:     db,
		client: setUpEnt(db),
	}
}

func (d *dbImpl) GetProduct(ctx context.Context, id uuid.UUID) (*ent.Product, error) {

	p, err := d.client.Product.Get(ctx, id)
	if err != nil {
		return nil, errors.New("product not found")
	}
	return p, nil

}

func (d *dbImpl) GetProductList(ctx context.Context) ([]*ent.Product, error) {
	products, err := d.client.Product.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (d *dbImpl) CreateProduct(ctx context.Context, p *ent.Product, _ *ent.Category) (*ent.Product, error) {
	created, err := d.client.Product.Create().
		SetName(p.Name).
		SetPrice(p.Price).
		SetDescription(p.Description).
		SetColor(p.Color).
		SetImages(p.Images).
		Save(ctx)
	if err != nil {

		return nil, err
	}
	return created, nil
}

func (d *dbImpl) UpdateProduct(ctx context.Context, p *ent.Product) (*ent.Product, error) {

	updated, err := d.client.Product.UpdateOneID(p.ID).
		SetName(p.Name).
		SetPrice(p.Price).
		SetDescription(p.Description).
		SetColor(p.Color).
		SetImages(p.Images).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return updated, nil
}
