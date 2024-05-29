package service

import (
	"context"

	pbu "server/genproto/user"
	grpcclient "server/service/grpc-client"
	"server/storage"
	s "server/storage"

	"github.com/jmoiron/sqlx"
)

type UserServer struct {
	Storage s.IStorage
}

func NewUserServer(db *sqlx.DB, grpcClinet grpcclient.Clients) *UserServer {
	return &UserServer{
		Storage: storage.NewStoragePg(db),
	}
}

func (u *UserServer) Register(ctx context.Context, req *pbu.RegisterReq) (*pbu.UserRes, error) {
	res, err := u.Storage.User().CreateUser(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *UserServer) GetUserById(context.Context, *pbu.UserIdReq) (*pbu.UserRes, error) {

	return nil, nil
}

func (u *UserServer) DeleteUser(ctx context.Context, req *pbu.UserIdReq) (*pbu.DeleteRes, error) {
	return nil, nil
}

func (u *UserServer) UpdateUser(context.Context, *pbu.UpdateReq) (*pbu.UserRes, error) {

	return nil, nil
}
