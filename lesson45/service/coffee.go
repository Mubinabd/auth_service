package service

import (
	"context"
	"fmt"
	pb "github.com/husanmusa/NT_Golang_10/lesson45/genproto/coffee"
	"github.com/husanmusa/NT_Golang_10/lesson45/storage/postgres"
)

var coffees = map[string]struct {
	Price  int32
	Volume int32
}{
	"latte":           {17000, 40},
	"cappucino":       {17000, 50},
	"double expresso": {13000, 80},
}

type coffeeService struct {
	//stg storage.StorageI
	stg postgres.Storage
	pb.UnimplementedCoffeeServiceServer
}

func NewCoffeeService(stg *postgres.Storage) *coffeeService {
	return &coffeeService{stg: *stg}
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
