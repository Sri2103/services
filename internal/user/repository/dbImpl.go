package repository

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect"
	"github.com/Sri2103/services/pkg/database"
	"github.com/Sri2103/services/pkg/ent/user"
	"github.com/google/uuid"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/Sri2103/services/pkg/ent"
)

var (
	ErrNoSuchUser = fmt.Errorf("no such user")
)

type dbImpl struct {
	DB     *database.DB
	client *ent.UserClient
}

func setUpEnt(db *database.DB) *ent.UserClient {
	drv := entsql.OpenDB(dialect.Postgres, db.Conn)
	client := ent.NewClient(ent.Driver(drv))
	return client.Debug().User
}

func NewDB(db *database.DB) Repo {
	return &dbImpl{
		DB:     db,
		client: setUpEnt(db),
	}
}

// CreateUser implements Repo.
func (d *dbImpl) CreateUser(ctx context.Context, user *ent.User) (*ent.User, error) {
	user, err := d.client.Create().
		SetName(user.Name).
		SetEmail(user.Email).
		SetUsername(user.Username).
		SetID(user.ID).
		SetPassword(user.Password).
		Save(ctx)

	return user, err
}

// GetUserById implements Repo.
func (d *dbImpl) GetUserById(ctx context.Context, userId string) (*ent.User, error) {
	userGet, err := d.client.Get(ctx, uuid.MustParse(userId))
	return userGet, err
}

// UpdateUser implements Repo.
func (d *dbImpl) UpdateUser(ctx context.Context, id string, user *ent.User) (*ent.User, error) {
	// get and then update
	u, err := d.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	updater := d.client.UpdateOneID(u.ID)
	if user.Name != "" {
		updater = updater.SetName(user.Name)
	}
	if user.Email != "" {
		updater = updater.SetEmail(user.Email)
	}
	return updater.Save(ctx)
}

// GetUserByEmail implements Repo
func (d *dbImpl) GetUserByEmail(ctx context.Context, email string) (*ent.User, error) {
	userGet, err := d.client.Query().Where(user.Email(email)).Only(ctx)
	if ent.IsNotFound(err) {
		return nil, ErrNoSuchUser
	} else if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return userGet, err
}
