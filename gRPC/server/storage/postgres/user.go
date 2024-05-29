package postgres

import (
	pbu "server/genproto/user"
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) CreateUser(user *pbu.RegisterReq) (*pbu.UserRes, error) {

	return nil, nil

}
