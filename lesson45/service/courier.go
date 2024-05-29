package service

import (
	"context"
	"fmt"
	pb "github.com/husanmusa/NT_Golang_10/lesson45/genproto/courier"
	"github.com/husanmusa/NT_Golang_10/lesson45/storage/postgres"
)

type courierService struct {
	//stg storage.StorageI
	stg postgres.Storage
	pb.UnimplementedCourierServiceServer
}

func NewCourierService(stg *postgres.Storage) *courierService {
	return &courierService{stg: *stg}
}

func (c *courierService) Deliver(ctx context.Context, o *pb.TakeOrder) (*pb.DeliverOrder, error) {
	if !o.IsPaid {
		return nil, fmt.Errorf("order is not paid for yet")
	}

	res, err := c.stg.Courier.Deliver(o)
	if err != nil {
		return nil, err
	}

	return res, nil
}
