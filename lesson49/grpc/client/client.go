package client

import (
	"github.com/husanmusa/NT_Golang_10/lesson49/config"
	"github.com/husanmusa/NT_Golang_10/lesson49/genproto/coffee"
	"github.com/husanmusa/NT_Golang_10/lesson49/genproto/courier"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClients struct {
	CoffeeService  coffee.CoffeeServiceClient
	CourierService courier.CourierServiceClient
}

func NewGrpcClients(cfg config.Config) (*GrpcClients, error) {
	conn, err := grpc.NewClient(cfg.COFFEE_SERVICE_HOST+cfg.COFFEE_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &GrpcClients{
		CoffeeService:  coffee.NewCoffeeServiceClient(conn),
		CourierService: courier.NewCourierServiceClient(conn),
	}, nil
}
