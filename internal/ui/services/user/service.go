package user_service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

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
	GetUser(ctx context.Context, userId string) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type service struct {
	AllClients *client.ClientAggregator
}

// EditUser implements UserService.
func (s *service) EditUser(ctx context.Context, req *User) (*User, error) {
	// create edit User request
	editedUserReq := s.AllClients.UserClient.NewRequest().SetBody(req)
	// send the request and get response
	res, err := editedUserReq.Put("")
	if err != nil {
		return nil, err
	}
	statusCodeInt, _ := strconv.Atoi(res.Status())
	switch statusCodeInt / 100 {
	case 2:
		return req, nil
	case 4:
		return nil, fmt.Errorf("user with id %s not found", req.Id)
	}

	return nil, fmt.Errorf("unhandled error")
}

// GetUser implements UserService.
func (s *service) GetUser(ctx context.Context, userId string) (*User, error) {
	getUserReq := s.AllClients.UserClient.NewRequest()
	res, err := getUserReq.Get("/" + userId)
	if err != nil {
		return nil, err
	}

	var u User
	err = json.Unmarshal(res.Body(), &u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// LoginUser implements UserService.
func (s *service) LoginUser(ctx context.Context, req *User) (*User, error) {
   r := s.AllClients.UserClient.NewRequest().SetBody(req)
   res, err := r.Post("/login")
   if err != nil {
	  return nil, fmt.Errorf("could not LoginUser user: %w", err)
   }
   var u User
   err = json.Unmarshal(res.Body(), &u)
   if err != nil {
	  return nil, err
   }
   if u.Id == "" {
	  return nil, ErrInvalidCred
   }
   return &u, nil
	
}
// SignUpUser implements UserService.
func (s *service) SignUpUser(ctx context.Context, req *User) (*User, error) {
	r := s.AllClients.UserClient.NewRequest().SetBody(req)
	res, err := r.Post("/")
	if err != nil {
		return nil, fmt.Errorf("could not add user: %w", err)
	}
	var u User
	err = json.Unmarshal(res.Body(), &u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// GetUserByEmail implements UserService.
func (s *service) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	panic("unimplemented")
}

func New(cfg *config.AppConfig) UserService {
	switch cfg.DevConfig.UseApi {
	case true:
		return NewMockService()

	case false:
		return &service{
			AllClients: client.AllClients(client.New(cfg)),
		}
	default:
		return nil
	}
}
