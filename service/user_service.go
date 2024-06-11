package service

import (
	pb "github.com/Mubinabd/auth_service/genproto"
	"github.com/Mubinabd/auth_service/storage"
	"github.com/google/uuid"
)

type UserService struct {
	storage storage.StorageI
}

func NewUserService(storage storage.StorageI) *UserService {
	return &UserService{
		storage: storage,
	}
}

func (s *UserService) RegisterUser(user *pb.UserCreate) (*pb.User, error) {
	id := uuid.NewString()
	user.Id = id
	return s.storage.User().RegisterUser(user)
}
func (s *UserService) GetUser(username *pb.ByUsername) (*pb.User, error) {
    return s.storage.User().GetUser(username)
}
func (s *UserService) LoginUser(logreq *pb.LoginReq) (*pb.Token, error) {
    return s.storage.User().LoginUser(logreq)
}
