package postgres

import (
	"database/sql"
	"fmt"
	"github.com/husanmusa/NT_Golang_10/lesson45/storage"
	_ "github.com/lib/pq"
)

type Storage struct {
	Db       *sql.DB
	CoffeeS  storage.CoffeeI
	CourierS storage.CourierI
}

func NewPostgresStorage() (*Storage, error) {
	con := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", "postgres", "pass", "localhost", 5432, "go10")
	db, err := sql.Open("postgres", con)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	//cf := NewCoffeeRepo(db)
	//cr := NewCourierRepo(db)
	return &Storage{Db: db}, err
}

func (s *Storage) Coffee() storage.CoffeeI {
	if s.CoffeeS == nil {
		s.CoffeeS = NewCoffeeRepo(s.Db)
	}

	return s.CoffeeS
}

func (s *Storage) Courier() storage.CourierI {
	if s.CourierS == nil {
		s.CourierS = NewCourierRepo(s.Db)
	}

	return s.CourierS
}
