package service

import (
	"context"
	"fmt"
	pb "github.com/husanmusa/NT_Golang_10/lesson45/genproto/courier"
	"github.com/husanmusa/NT_Golang_10/lesson45/grpc/client"
	"github.com/husanmusa/NT_Golang_10/lesson45/storage"
)

type courierService struct {
	//stg storage.StorageI
	stg  storage.StorageI
	clnt client.GrpcClients
	pb.UnimplementedCourierServiceServer
}

func NewCourierService(stg storage.StorageI, clnt client.GrpcClients) *courierService {
	return &courierService{stg: stg, clnt: clnt}
}

func (c *courierService) Deliver(ctx context.Context, o *pb.TakeOrder) (*pb.DeliverOrder, error) {
	if !o.IsPaid {
		return nil, fmt.Errorf("order is not paid for yet")
	}

	//c.clnt.CoffeeService.BuyingCoffee()

	res, err := c.stg.Courier().Deliver(o)
	if err != nil {
		return nil, err
	}

	return res, nil
}
