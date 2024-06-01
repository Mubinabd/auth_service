package service

import (
	"context"
	"fmt"
	pb "github.com/husanmusa/NT_Golang_10/lesson45/genproto/coffee"
	"github.com/husanmusa/NT_Golang_10/lesson45/storage/postgres"
	"github.com/husanmusa/NT_Golang_10/lesson45/storage/redis"
)

type coffeeService struct {
	//stg storage.StorageI
	stg  postgres.Storage
	stgR redis.Storage
	pb.UnimplementedCoffeeServiceServer
}

func NewCoffeeService(stg *postgres.Storage, stgR redis.Storage) *coffeeService {
	return &coffeeService{stg: *stg, stgR: stgR}
}

func (c *coffeeService) BuyingCoffee(ctx context.Context, coffee *pb.BuyCoffee) (*pb.PreparedCoffee, error) {
	if !coffee.IsPaid {
		return nil, fmt.Errorf("coffee is not paid")
	}

	res, err := c.stg.CoffeeS.SelectCoffee(coffee)
	if err != nil {
		return nil, err
	}

	return res, nil
}
