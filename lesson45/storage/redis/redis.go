package redis

import (
	"database/sql"
	pb "github.com/husanmusa/NT_Golang_10/lesson45/genproto/coffee"
	pbr "github.com/husanmusa/NT_Golang_10/lesson45/genproto/courier"
	"github.com/husanmusa/NT_Golang_10/lesson45/storage"
)

type Storage struct {
	Db *sql.DB
}

func NewRedisStorage() (storage.StorageI, error) {
	return &Storage{}, nil
}

type son int

func (son) SelectCoffee(*pb.BuyCoffee) (*pb.PreparedCoffee, error) { return nil, nil }

func (s *Storage) Coffee() storage.CoffeeI {

	return son(1)
}

type san int

func (san) Deliver(*pbr.TakeOrder) (*pbr.DeliverOrder, error) { return nil, nil }

func (s *Storage) Courier() storage.CourierI {

	return san(1)
}
