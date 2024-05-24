package service

import (
	"database/sql"
	"github.com/husanmusa/NT_Golang_10/lesson43/storage/postgres"
)

type Service struct {
	PrService *ProductService
}

func InitServices(db *sql.DB) *Service {
	product := NewProductService(postgres.NewProductStorage(db))

	return &Service{product}
}
