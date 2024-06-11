package handler

import "github.com/Mubinabd/auth_service/service"

type HandlerStruct struct {
	UserService *service.UserService
}

func NewHandler(userService *service.UserService) *HandlerStruct {
	return &HandlerStruct{
		UserService: userService,
	}
}
