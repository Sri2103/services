package repository

import (
	"context"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/Sri2103/services/pkg/database"
	"github.com/Sri2103/services/pkg/ent"
	"github.com/google/uuid"
)

type repo struct {
	DB     *database.DB
	client *ent.OrderClient
}

func setUpEnt(db *database.DB) *ent.OrderClient {
	drv := entsql.OpenDB(dialect.Postgres, db.Conn)
	client := ent.NewClient(ent.Driver(drv))
	return client.Debug().Order
}

// CreateOrder implements Repository.
func (r *repo) CreateOrder(ctx context.Context, order *ent.Order) (*ent.Order, error) {
	panic("unimplemented")
}

// DeleteOrder implements Repository.
func (r *repo) DeleteOrder(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

// GetOrder implements Repository.
func (r *repo) GetOrder(ctx context.Context, id uuid.UUID) (*ent.Order, error) {
	panic("unimplemented")
}

// ListOrders implements Repository.
func (r *repo) ListOrders(ctx context.Context) ([]*ent.Order, error) {
	panic("unimplemented")
}

// UpdateOrder implements Repository.
func (r *repo) UpdateOrder(ctx context.Context, order *ent.Order) (*ent.Order, error) {
	panic("unimplemented")
}

func New(db *database.DB) Repository {
	return &repo{
		DB:     db,
		client: setUpEnt(db),
	}
}
