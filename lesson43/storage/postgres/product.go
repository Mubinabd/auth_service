package postgres

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/husanmusa/NT_Golang_10/lesson43/model"
)

type ProductStorage struct {
	db *sql.DB
}

func NewProductStorage(db *sql.DB) *ProductStorage {
	return &ProductStorage{db: db}
}

func (p *ProductStorage) CreateProduct(product *model.CreateProduct) error {
	id := uuid.NewString()
	query := `
		INSERT INTO product (id, name, price, category_id, expired_at)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := p.db.Exec(query, id, product.Name, product.Price, product.CategoryID, product.ExpiredAt)
	return err
}

func (p *ProductStorage) GetProduct(id string) (*model.Product, error) {
	// TODO there is not categories. change the query
	query := `
		SELECT id, name, price, category_id, expired_at, created_at, updated_at
		FROM product
		WHERE id = $1 and deleted_at=0
	`
	row := p.db.QueryRow(query, id)

	var product model.Product
	err := row.Scan(&product.ID,
		&product.Name,
		&product.Price,
		&product.Category,
		&product.ExpiredAt,
		&product.CreatedAt,
		&product.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductStorage) UpdateProduct(product *model.Product) error {
	query := `
		UPDATE product
		SET name = $2, price = $3, category_id = $4, expired_at = $5, updated_at = now()
		WHERE id = $1 and deleted_at=0
	`
	_, err := p.db.Exec(query, product.ID, product.Name, product.Price, product.Category, product.ExpiredAt)
	return err
}

func (p *ProductStorage) DeleteProduct(id string) error {
	query := `
		update product set deleted_at = Extract(epoch from now()) WHERE id = $1 and deleted_at=0
	`
	_, err := p.db.Exec(query, id)
	return err
}
