package grpcclient

import (
	"fmt"
	"server/config"
	pbp "server/genproto/post"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Clients interface {
	Post() pbp.PostServiceClient
}

type ServiceMeneger struct {
	postService pbp.PostServiceClient
}

func NewServiceMeneger(cfg config.Config) (*ServiceMeneger, error) {
	connPost, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.PostServiceHost, cfg.PostServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("comment service dial host:%s, port:%s", cfg.PostServiceHost, cfg.PostServicePort)
	}

	return &ServiceMeneger{
		postService: pbp.NewPostServiceClient(connPost),
	}, nil
}

func (s *ServiceMeneger) Post() pbp.PostServiceClient {
	return s.postService
}
