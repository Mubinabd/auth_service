package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Storage struct {
	Db      *sql.DB
	CoffeeS *CoffeeRepo
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
	cf := NewCoffeeRepo(db)
	return &Storage{Db: db, CoffeeS: cf}, err
}

//func (s *Storage) Coffee() storage.CoffeeI {
//	if s.CoffeeS == nil {
//		s.CoffeeS = &CoffeeRepo{s.Db}
//	}
//
//	return s.CoffeeS
//}
