package service

import (
	"context"

	user_pb "github.com/Sri2103/services/pkg/rpc/user"
)

type userImpl struct {
	user_pb.UnimplementedUserServiceServer
}

func New() user_pb.UserServiceServer {
	return &userImpl{}
}

func (u *userImpl) Login(context.Context, *user_pb.LoginReq) (*user_pb.LoginResp, error) {
	return nil, nil
}
func (u *userImpl) CreateUser(context.Context, *user_pb.CreateUserReq) (*user_pb.CreateUserResp, error) {
	return nil, nil
}

func (u *userImpl) GetUserById(context.Context, *user_pb.CreateUserReq) (*user_pb.CreateUserResp, error) {
	return nil, nil
}

func (u *userImpl) EditUser(context.Context, *user_pb.UpdateUserReq) (*user_pb.UpdateUserResp, error) {
	return nil, nil
}
