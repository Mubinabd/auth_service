package postgres

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/husanmusa/NT_Golang_10/lesson43/model"
)

type ItemStorage struct {
	db *sql.DB
}

func (i *ItemStorage) Create(db *sql.DB, item *model.Item) error {
	id := uuid.NewString()
	query := `
		INSERT INTO item (id, basket_id, product_id, quantity)
		VALUES ($1, $2, $3, $4)
	`
	_, err := db.Exec(query, id, item.BasketID, item.ProductID, item.Quantity)
	return err
}
