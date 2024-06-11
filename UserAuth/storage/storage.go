package storage

import (
	pb "github.com/Mubinabd/auth_service/genproto"
)
type StorageI interface {
	User() UserI
}

type UserI interface {
	RegisterUser(user *pb.UserCreate) (*pb.User,error)
	LoginUser(login *pb.LoginReq) (*pb.Token, error)
	GetUser(get *pb.ByUsername) (*pb.User, error)
}
