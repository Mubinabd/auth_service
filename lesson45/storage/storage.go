package storage

import (
	pb "github.com/husanmusa/NT_Golang_10/lesson45/genproto/coffee"
	pbcr "github.com/husanmusa/NT_Golang_10/lesson45/genproto/courier"
)

type StorageI interface {
	Coffee() CoffeeI
	Courier() CourierI
}

type CoffeeI interface {
	SelectCoffee(coffee *pb.BuyCoffee) (*pb.PreparedCoffee, error)
}

type CourierI interface {
	Deliver(order *pbcr.TakeOrder) (*pbcr.DeliverOrder, error)
}
