package service

import (
	"github.com/husanmusa/NT_Golang_10/lesson43/model"
	"github.com/husanmusa/NT_Golang_10/lesson43/storage/postgres"
	"log/slog"
)

type ProductService struct {
	product *postgres.ProductStorage
}

func NewProductService(pr *postgres.ProductStorage) *ProductService {
	return &ProductService{product: pr}
}

func (p *ProductService) CreateProduct(pr model.CreateProduct) error {
	slog.Info("CreateProduct Service", "product", pr)
	err := p.product.CreateProduct(&pr)
	if err != nil {
		slog.Error("error while CreateProduct Service", "err", err)
		return err
	}

	return nil
}

func (p *ProductService) GetProduct(id string) (*model.Product, error) {
	slog.Info("GetProduct Service", "productId req", id)
	prod, err := p.GetProduct(id)
	if err != nil {
		slog.Error("error while GetProduct Service", "err", err)
		return nil, err
	}

	slog.Info("GetProduct Service", "product resp", prod)

	return prod, err
}
