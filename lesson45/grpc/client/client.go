package client

import (
	"github.com/husanmusa/NT_Golang_10/lesson45/genproto/coffee"
	"github.com/husanmusa/NT_Golang_10/lesson45/genproto/courier"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClients struct {
	CoffeeService  coffee.CoffeeServiceClient
	CourierService courier.CourierServiceClient
}

func NewGrpcClients() (*GrpcClients, error) {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &GrpcClients{
		CoffeeService:  coffee.NewCoffeeServiceClient(conn),
		CourierService: courier.NewCourierServiceClient(conn),
	}, nil
}
