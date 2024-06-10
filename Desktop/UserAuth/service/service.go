package service

type Service struct {
	UM  *UserService

}

func NewService(user *UserService) *Service {
	return &Service{
		UM: user,
	}
}