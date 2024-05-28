package storage

import pb "github.com/husanmusa/NT_Golang_10/lesson45/genproto/coffee"

type StorageI interface {
	Coffee() CoffeeI
}

type CoffeeI interface {
	SelectCoffee(coffee *pb.BuyCoffee) (*pb.PreparedCoffee, error)
}
