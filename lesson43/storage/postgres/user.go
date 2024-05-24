package postgres

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/husanmusa/NT_Golang_10/lesson43/model"
)

type UserStorage struct {
	db *sql.DB
}

func (u *UserStorage) Create(db *sql.DB, user *model.User) error {
	id := uuid.NewString()
	query := `
		INSERT INTO User (id, name, balance,email,phone)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := db.Exec(query, id, user.Name, user.Balance, user.Email, user.Phone)
	return err
}

func (u *UserStorage) Get(db *sql.DB, id string) (*model.User, error) {
	query := `
		SELECT id, name, balance, email, phone
		FROM User
		WHERE id = $1
	`
	row := db.QueryRow(query, id)

	var user model.User
	err := row.Scan(&user.ID, &user.Name, &user.Balance, &user.Email, &user.Phone)
	return &user, err
}
func (u *UserStorage) Update(db *sql.DB, user *model.User) error {
	query := `
		UPDATE User
		SET name = $2, balance = $3, email = $4, phone = $5, updated_at = now()
		WHERE id = $1
	`
	_, err := db.Exec(query, user.ID, user.Name, user.Balance, user.Email, user.Phone)
	return err
}

func (u *UserStorage) Delete(db *sql.DB, id string) error {
	query := `
		delete from User where id = $1
	`
	_, err := db.Exec(query, id)
	return err
}
