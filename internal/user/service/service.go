package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Sri2103/services/internal/users/repository"
	"github.com/Sri2103/services/pkg/ent"
	user_pb "github.com/Sri2103/services/pkg/rpc/user"
)

var (
	ErrWrongPassword = errors.New("wrong password")
)

type userImpl struct {
	repo repository.Repo
	user_pb.UnimplementedUserServiceServer
}

func New(r repository.Repo) user_pb.UserServiceServer {
	return &userImpl{
		repo: r,
	}
}

func (u *userImpl) handleUser(_ context.Context, user *ent.User, err error) (*user_pb.CreateUserResp, error) {
	if err != nil {
		return nil, err
	}

	return &user_pb.CreateUserResp{
		Code: 200,
		Msg:  "success",
		User: &user_pb.User{
			Id:       user.ID.String(),
			Name:     user.Name,
			Email:    user.Email,
			UserName: user.Username,
			Password: user.Password,
			Role:     user.Edges.Role.Role,
		},
	}, nil
}

func (u *userImpl) handleLogin(_ context.Context, user *ent.User, err error) (*user_pb.LoginResp, error) {
	if err != nil {
		return nil, err
	}

	currentTime := time.Now().Format("2006-01-02 15:04:05")

	role := user.Edges.Role.Role

	fmt.Println(role, "Role of the user found ", user)

	return &user_pb.LoginResp{
		Code:    200,
		Msg:     "success",
		LoginAt: currentTime,
		UserInfo: &user_pb.User{
			Id:       user.ID.String(),
			Name:     user.Name,
			Email:    user.Email,
			UserName: user.Username,
			Password: user.Password,
			Role:     user.Edges.Role.Role,
		},
	}, nil
}

func (u *userImpl) Login(ctx context.Context, r *user_pb.LoginReq) (*user_pb.LoginResp, error) {
	getUser, err := u.repo.GetUserByEmail(ctx, r.GetEmail())
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	isValidPassword := getUser.Password == r.Password
	if !isValidPassword {
		return nil, ErrWrongPassword
	}

	return u.handleLogin(ctx, getUser, nil)
}

func (u *userImpl) CreateUser(ctx context.Context, r *user_pb.CreateUserReq) (*user_pb.CreateUserResp, error) {
	user := r.User

	var entUser ent.User
	entUser.Name = user.GetName()
	entUser.Username = user.GetUserName()
	entUser.Email = user.GetEmail()
	entUser.Password = user.GetPassword()

	// assert the role here
	role := &ent.Role{Role: user.GetRole()}

	createdUser, err := u.repo.CreateUserTransaction(ctx, &entUser, role)

	return u.handleUser(ctx, createdUser, err)
}

func (u *userImpl) GetUserById(ctx context.Context, r *user_pb.GetUserByIdReq) (*user_pb.CreateUserResp, error) {
	user, err := u.repo.GetUserById(ctx, r.UserId)
	return u.handleUser(ctx, user, err)
}

func (u *userImpl) EditUser(ctx context.Context, r *user_pb.UpdateUserReq) (*user_pb.UpdateUserResp, error) {
	// update user
	updatedUser, err := u.repo.UpdateUser(ctx, r.User.Id, &ent.User{
		Name:     r.User.Name,
		Email:    r.User.Email,
		Username: r.User.UserName,
		Password: r.User.Password,
	})
	if err != nil {
		return nil, err
	}

	return &user_pb.UpdateUserResp{
		Code: 200,
		Msg:  "success",
		User: &user_pb.User{
			Id:       updatedUser.ID.String(),
			Name:     updatedUser.Name,
			Email:    updatedUser.Email,
			UserName: updatedUser.Username,
			Password: updatedUser.Password,
		},
	}, nil
}
