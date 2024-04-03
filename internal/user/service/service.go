package service

import (
	"context"
	"errors"
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

func (u *userImpl) Login(ctx context.Context, r *user_pb.LoginReq) (*user_pb.LoginResp, error) {
	getUser, err := u.repo.GetUserById(ctx, r.GetEmail())
	if err != nil {
		return nil, err
	}
	isValidPassword := getUser.Password == r.Password
	if !isValidPassword {
		return nil, ErrWrongPassword
	}
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	// render loginResp
	return &user_pb.LoginResp{
		Code:    200,
		Msg:     "success",
		LoginAt: currentTime,
		UserInfo: &user_pb.User{
			Id:       getUser.ID.String(),
			Name:     getUser.Name,
			Email:    getUser.Email,
			UserName: getUser.Username,
			Password: getUser.Password,
		},
	}, nil
}
func (u *userImpl) CreateUser(ctx context.Context, r *user_pb.CreateUserReq) (*user_pb.CreateUserResp, error) {
	user := r.User
	createdUser, err := u.repo.CreateUser(ctx, &ent.User{
		Name:     user.GetName(),
		Password: user.GetPassword(),
		Username: user.GetUserName(),
		Email:    user.GetEmail(),
	})

	if err != nil {
		return nil, err
	}
	return &user_pb.CreateUserResp{
		Code: 200,
		Msg:  "success",
		User: &user_pb.User{
			Id:       createdUser.ID.String(),
			Name:     createdUser.Name,
			Email:    createdUser.Email,
			UserName: createdUser.Username,
			Password: createdUser.Password,
		},
	}, nil
}

func (u *userImpl) GetUserById(ctx context.Context, r *user_pb.GetUserByIdReq) (*user_pb.CreateUserResp, error) {
	user, err := u.repo.GetUserById(ctx, r.UserId)
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
			Password: string(user.Password),
		},
	}, nil
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
