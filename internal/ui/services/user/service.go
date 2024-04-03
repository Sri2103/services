package user_service

import (
	"context"
	"errors"

	"github.com/Sri2103/services/internal/ui/config"
	"github.com/Sri2103/services/internal/ui/services/client"
)

var (
	ErrWrongPassword = errors.New("wrong password")
	ErrUserNotFound  = errors.New("user not found")
	ErrUserExists    = errors.New("user already exists")
	ErrWrongEmail    = errors.New("wrong email")
	ErrWrongUsername = errors.New("wrong username")
	ErrWrongId       = errors.New("wrong id")
	ErrWrongName     = errors.New("wrong name")
	ErrInvalidCred   = errors.New("invalid credentials")
)

type UserService interface {
	LoginUser(ctx context.Context, req *User) (*User, error)
	SignUpUser(ctx context.Context, req *User) (*User, error)
	EditUser(ctx context.Context, req *User) (*User, error)
	GetUser(ctx context.Context, req *User) (*User, error)
}

type service struct {
	AllClients *client.ClientAggregator
}

// EditUser implements UserService.
func (s *service) EditUser(ctx context.Context, req *User) (*User, error) {
	panic("unimplemented")
}

// GetUser implements UserService.
func (s *service) GetUser(ctx context.Context, req *User) (*User, error) {
	panic("unimplemented")
}

// LoginUser implements UserService.
func (s *service) LoginUser(ctx context.Context, req *User) (*User, error) {
	panic("unimplemented")
}

// SignUpUser implements UserService.
func (s *service) SignUpUser(ctx context.Context, req *User) (*User, error) {
	panic("unimplemented")
}

func New(cfg *config.AppConfig) UserService {
	switch cfg.DevConfig.UseApi {
	case false:
		return NewMockService()

	case true:
		return &service{
			AllClients: client.AllClients(client.New(cfg)),
		}
	default:
		return nil
	}
}
