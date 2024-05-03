package postgres

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/husanmusa/NT_Golang_10/lesson28/models"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *userRepo {
	return &userRepo{db}
}

func (u *userRepo) Create(req *models.User) error {
	id := uuid.NewString()

	_, err := u.db.Exec("insert into users(id, name, username, password, phone) values($1, $2, $3, $4, $5)",
		id, req.Name, req.Username, req.Password, req.Phone)

	return err
}

func (u *userRepo) GetById(id string) (*models.User, error) {
	user := &models.User{ID: id}
	err := u.db.QueryRow("select name, username, phone, created_at from users where id=$1 and deleted_at=0", id).
		Scan(user.Name, user.Username, user.Phone, user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepo) GetAll() (*models.GetAllUsersResp, error) {
	users := []models.User{}

	rows, err := u.db.Query("select name, username, phone, created_at from users where deleted_at = 0")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := models.User{}

		err = rows.Scan(&user.Name, &user.Username, &user.Phone, &user.CreatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	var count int
	err = u.db.QueryRow("select count(1) from users").Scan(&count)

	return &models.GetAllUsersResp{&users, count}, nil
}

func (u *userRepo) Update(user *models.User) error {
	_, err := u.db.Exec("update users set name = $1, username=$2, phone=$3, updated_at=now() where id=$4 and deleted_at=0",
		user.Name, user.Username, user.Phone, user.ID)

	return err
}

func (u *userRepo) Delete(id string) error {
	_, err := u.db.Exec("update users set deleted_at=date_part('epoch', current_timestamp)::INT where id=$1 and deleted_at=0", id)

	return err
}
