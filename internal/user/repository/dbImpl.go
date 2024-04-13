package repository

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect"
	"github.com/Sri2103/services/pkg/database"
	"github.com/Sri2103/services/pkg/ent/role"
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
	client *ent.Client
}

func setUpEnt(db *database.DB) *ent.Client {
	drv := entsql.OpenDB(dialect.Postgres, db.Conn)
	client := ent.NewClient(ent.Driver(drv))
	return client
}

func NewDB(db *database.DB) Repo {
	return &dbImpl{
		DB:     db,
		client: setUpEnt(db),
	}
}

// CreateUser implements Repo.
func (d *dbImpl) CreateUser(ctx context.Context, user *ent.User, role *ent.Role) (*ent.User, error) {
	u, err := d.client.User.Create().
		SetName(user.Name).
		SetEmail(user.Email).
		SetUsername(user.Username).
		SetPassword(user.Password).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	_, err = d.client.Role.Create().AddUserIDs(u.ID).SetRole(role.Role).Save(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println(u.Edges.Role, "Role to be saved ")
	return u, nil
}

// GetUserById implements Repo.
func (d *dbImpl) GetUserById(ctx context.Context, userId string) (*ent.User, error) {
	userGet, err := d.client.User.Get(ctx, uuid.MustParse(userId))
	return userGet, err
}

// UpdateUser implements Repo.
func (d *dbImpl) UpdateUser(ctx context.Context, id string, user *ent.User) (*ent.User, error) {
	// get and then update
	u, err := d.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	updater := d.client.User.UpdateOneID(u.ID)
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
	userGet, err := d.client.User.Query().
		Where(user.Email(email)).
		WithRole().
		Only(ctx)
	if ent.IsNotFound(err) {
		return nil, ErrNoSuchUser
	} else if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	fmt.Println(userGet.Edges.Role, "Role of the user found ", userGet)
	return userGet, err
}

// create a chain of events with transactions
func (d *dbImpl) CreateUserTransaction(ctx context.Context, user *ent.User, rl *ent.Role) (*ent.User, error) {
	tx, err := d.client.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}
	u, err := tx.User.Create().
		SetName(user.Name).
		SetPassword(user.Password).
		SetEmail(user.Email).
		SetUsername(user.Username).
		Save(ctx)
	if err != nil {
		return nil, d.rollBackTxn(ctx, tx)
	}
	r, err := tx.Role.Query().Where(role.Role(rl.Role)).Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, d.rollBackTxn(ctx, tx)
		}
		// create role and set to user
		r, err = tx.Role.Create().SetRole(rl.Role).Save(ctx)
		if err != nil {
			return nil, d.rollBackTxn(ctx, tx)
		}
		u, err = tx.User.UpdateOne(u).SetRole(r).Save(ctx)
		if err != nil {
			return nil, d.rollBackTxn(ctx, tx)
		}
	} else {
		u, err = tx.User.UpdateOne(u).SetRole(r).Save(ctx)
		if err != nil {
			return nil, d.rollBackTxn(ctx, tx)
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, d.rollBackTxn(ctx, tx)
	}
	u.Edges.Role = r

	fmt.Println(u.Edges.Role, "Role to be saved ", u)

	return u, nil
}

func (d *dbImpl) rollBackTxn(_ context.Context, tx *ent.Tx) error {
	err := tx.Rollback()
	if err != nil {
		return fmt.Errorf("failed to rollback transaction: %w", err)
	}
	return nil
}
