package handlers

import (
	"github.com/husanmusa/NT_Golang_10/lesson43/service"
)

type HTTPHandler struct {
	service *service.Service
}

func NewHTTPHandler(service *service.Service) *HTTPHandler {
	return &HTTPHandler{service: service}
}
