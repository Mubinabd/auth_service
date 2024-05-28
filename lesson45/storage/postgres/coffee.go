package postgres

import (
	"database/sql"
	pb "github.com/husanmusa/NT_Golang_10/lesson45/genproto/coffee"
)

type CoffeeRepo struct {
	db *sql.DB
}

func NewCoffeeRepo(db *sql.DB) *CoffeeRepo {
	return &CoffeeRepo{db: db}
}

func (c *CoffeeRepo) SelectCoffee(coffee *pb.BuyCoffee) (*pb.PreparedCoffee, error) {
	res := pb.PreparedCoffee{}
	err := c.db.QueryRow("select name, volume, price from coffee where name = $1", coffee.GetName()).
		Scan(&res.Name, &res.Volume, &res.Price)

	return &res, err
}
