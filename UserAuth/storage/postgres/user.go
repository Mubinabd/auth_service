package postgres

import (
	"database/sql"
	"errors"
	"log"

	pb "github.com/Mubinabd/auth_service/genproto"

	"github.com/google/uuid"
)

type UserManager struct {
	db *sql.DB
}

func NewUserManager(db *sql.DB) *UserManager {
	return &UserManager{db: db}
}

func (user *UserManager) RegisterUser(reg *pb.UserCreate) (*pb.User, error) {
	id := uuid.NewString()
	query := "INSERT INTO users(id, username, password, email) VALUES ($1, $2, $3, $4)"

	_, err := user.db.Exec(query, id, reg.Username, reg.Password, reg.Email)
	if err != nil {
		log.Println("error while creating user:", err)
		return nil, err
	}

	return &pb.User{
		Username: reg.Username,
		Password: reg.Password,
		Email:    reg.Email,
	}, nil
}

func (user *UserManager) LoginUser(login *pb.LoginReq) (*pb.Token, error) {
	query := "SELECT * FROM users WHERE username = $1 AND password = $2"

	var id, username, password, email string

	row := user.db.QueryRow(query, login.Username, login.Password)
	if err := row.Scan(&id, &username, &password, &email); err != nil {
		log.Println("error while logging in user:", err)
		return nil, err
	}

	return &pb.Token{}, nil
}

func (user *UserManager) GetUser(get *pb.ByUsername) (*pb.User, error) {
	if user.db == nil {
		return nil, errors.New("database connection is nil")
	}

	query := "SELECT * FROM users WHERE username = $1"
	var id, username, password, email string

	row := user.db.QueryRow(query, get.Username)
	if err := row.Scan(&id, &username, &password, &email); err != nil {
		log.Println("error while getting user:", err)
		return nil, err
	}

	return &pb.User{
		Username: username,
		Password: password,
		Email:    email,
	}, nil
}
