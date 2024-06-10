package handler

import (
	pb "github.com/Mubinabd/auth_service/genproto"

	"google.golang.org/grpc"
)

type HandlerStruct struct {
	User  pb.UserServiceClient
}

func NewHandlerStruct(connClient *grpc.ClientConn) *HandlerStruct {
	return &HandlerStruct{
		User:   pb.NewUserServiceClient(connClient),
		
	}
}
