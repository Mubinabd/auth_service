package postgres

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/husanmusa/NT_Golang_10/lesson43/model"
)

type CategoryStorage struct {
	db *sql.DB
}

func (c *CategoryStorage) Create(db *sql.DB, cat *model.Category) error {
	id := uuid.NewString()
	query := `
		INSERT INTO category (id, name)
		VALUES ($1, $2)
	`
	_, err := db.Exec(query, id, cat.Name)
	return err
}

func (c *CategoryStorage) Get(db *sql.DB, id string) (*model.Category, error) {
	query := `
		SELECT id, name
		FROM Category
		WHERE id = $1
	`
	row := db.QueryRow(query, id)

	var cat model.Category
	err := row.Scan(&cat.ID,
		&cat.Name)
	if err != nil {
		return nil, err
	}

	return &cat, nil
}

func (c *CategoryStorage) Update(db *sql.DB, category *model.Category) error {
	query := `
		UPDATE category
		SET name = $2, updated_at = now()
		WHERE id = $1 
	`
	_, err := db.Exec(query, category.ID, category.Name)
	return err
}

func (c *CategoryStorage) Delete(db *sql.DB, id string) error {
	query := `
		delete from category where id = $1 
	`
	_, err := db.Exec(query, id)
	return err
}
