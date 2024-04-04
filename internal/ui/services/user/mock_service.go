package user_service

import (
	"context"
	"fmt"
)

type mockUserService struct {
	returnUsers []*User
	returnErr   error
}

var userDummy = &User{
	Id:       "1234",
	Name:     " user-1",
	Email:    "duser.one@example.com",
	Password: "$2a$10",
	UserName: "user-1",
}

var user_2 = &User{
	Id:       "1235",
	Name:     " user-2",
	Email:    "duser.two@example.com",
	Password: "$2a$10",
	UserName: "user-2",
}

// EditUser implements UserService.
func (m *mockUserService) EditUser(ctx context.Context, user *User) (*User, error) {

	return user, m.returnErr
}

// GetUser implements UserService.
func (m *mockUserService) GetUser(ctx context.Context, id string) (*User, error) {
	for _, user := range m.returnUsers {
		if user.Id == id {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user with id %s not found", id)
}

// LoginUser implements UserService.
func (m *mockUserService) LoginUser(ctx context.Context, u *User) (*User, error) {
	for _, user := range m.returnUsers {
		if user.Email == u.Email && user.Password == u.Password {
			return user, nil
		}
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

// GetUserByEmail implements UserService.
func (m *mockUserService) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	for _, user := range m.returnUsers {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user with email %s not found", email)
}


func NewMockService() UserService {

	return &mockUserService{
		returnUsers: []*User{userDummy, user_2},
		returnErr:   nil,
	}
}
