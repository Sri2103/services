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
	client *ent.CartClient
}

func setUpEnt(db *database.DB) *ent.CartClient {
	drv := entsql.OpenDB(dialect.Postgres, db.Conn)
	client := ent.NewClient(ent.Driver(drv))
	return client.Debug().Cart
}

// AddItem implements Repository.
func (r *repo) AddItem(ctx context.Context, item *ent.CartItem) (*ent.Cart, error) {
	panic("unimplemented")
}

// DeleteCart implements Repository.
func (r *repo) DeleteCart(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

// RemoveItem implements Repository.
func (r *repo) RemoveItem(ctx context.Context, item *ent.CartItem) (*ent.Cart, error) {
	panic("unimplemented")
}

func New(db *database.DB) Repository {
	return &repo{
		DB:     db,
		client: setUpEnt(db),
	}
}
