package user_service

import (
	"context"
	"fmt"
)

type mockUserService struct {
	returnUser *User
	returnErr  error
}

var userDummy = &User{
	Id:       "1234",
	Name:     " user-1",
	Email:    "duser.one@example.com",
	Password: "$2a$10",
	UserName: "user-1",
}

// EditUser implements UserService.
func (m *mockUserService) EditUser(ctx context.Context, req *User) (*User, error) {

	if m.returnErr != nil {
		return nil, m.returnErr
	}

	return m.returnUser, nil
}

// GetUser implements UserService.
func (m *mockUserService) GetUser(ctx context.Context, req *User) (*User, error) {
	if m.returnUser != nil && m.returnUser.Id == req.Id {
		return m.returnUser, nil
	}

	return nil, fmt.Errorf("user with id %s not found", req.Id)
}

// LoginUser implements UserService.
func (m *mockUserService) LoginUser(ctx context.Context, req *User) (*User, error) {
	if m.returnUser != nil && m.returnUser.Email == req.Email && m.returnUser.Password == req.Password {
		return m.returnUser, nil
	}

	return nil, ErrInvalidCred
}

// SignUpUser implements UserService.
func (m *mockUserService) SignUpUser(ctx context.Context, req *User) (*User, error) {
	if m.returnErr != nil {
		return nil, m.returnErr
	}

	newUser := &User{
		Id:       req.Id,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		UserName: req.UserName,
	}

	return newUser, nil
}

func NewMockService() UserService {

	return &mockUserService{
		returnUser: userDummy,
		returnErr:  nil,
	}
}
