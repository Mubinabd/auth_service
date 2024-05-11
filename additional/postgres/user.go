package postgres

import (
	"database/sql"

	"gitgub.com/husanmusa/NT_Golang_10/lesson28/models"
)

type UserSt struct {
	Db *sql.DB
}

func NewUser(db *sql.DB) *UserSt {
	return &UserSt{db}
}

func (u *UserSt) GetAll() *[]models.User {
	u.Db.Query()
}
