package service

import (
	"context"

	"github.com/Sri2103/services/internal/users/repository"
	"github.com/Sri2103/services/pkg/ent"
	user_pb "github.com/Sri2103/services/pkg/rpc/user"
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

func (u *userImpl) Login(context.Context, *user_pb.LoginReq) (*user_pb.LoginResp, error) {
	// TODO: implement login logic here
	return nil, nil
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

func (u *userImpl) EditUser(context.Context, *user_pb.UpdateUserReq) (*user_pb.UpdateUserResp, error) {
	return nil, nil
}
