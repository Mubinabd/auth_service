package handler

import "github.com/husanmusa/NT_Golang_10/lesson49/grpc/client"

type Handler struct {
	srvs *client.GrpcClients
}

func NewHandler(srvs *client.GrpcClients) *Handler {
	return &Handler{srvs: srvs}
}
