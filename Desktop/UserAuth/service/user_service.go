package service

import (
	"context"
	pb "github.com/Mubinabd/auth_service/genproto"
	"github.com/Mubinabd/auth_service/storage"
)

type UserService struct {
	stg storage.StorageI
	pb.UnimplementedUserServiceServer
}

func NewUserService(stg storage.StorageI) *UserService {
	return &UserService{stg: stg}
}

func (user *UserService)RegisterUser(ctx context.Context, req *pb.UserCreate) (*pb.User, error) {
	_,err := user.stg.User().RegisterUser(req)
	if err != nil {
		return nil, err
	}
	return &pb.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}, nil
}

func (user *UserService)LoginUser(ctx context.Context, req *pb.LoginReq) (*pb.Token, error) {
	token, err := user.stg.User().LoginUser(req)
	if err != nil {
		return nil, err
	}
	return token, nil
}
func (user *UserService) GetUser(ctx context.Context, req *pb.ByUsername) (*pb.User, error) {
	res, err := user.stg.User().GetUser(req)
	if err != nil {
		return nil, err
	}
	return &pb.User{
		Username: res.Username,
		Password: res.Password,
		Email:    res.Email,
	}, nil
}
