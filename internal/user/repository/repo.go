package repository

import (
	"context"

	"github.com/Sri2103/services/pkg/ent"
)

type Repo interface {
	GetUserById(ctx context.Context, userId string) (*ent.User, error)
	CreateUser(ctx context.Context, user *ent.User, role *ent.Role) (*ent.User, error)
	UpdateUser(ctx context.Context, id string, user *ent.User) (*ent.User, error)
	GetUserByEmail(ctx context.Context, email string) (*ent.User, error)
	CreateUserTransaction(ctx context.Context, user *ent.User, rl *ent.Role) (*ent.User, error)
}
