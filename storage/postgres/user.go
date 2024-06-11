package postgres

import (
	"database/sql"
	"errors"
	"log"

	"github.com/Mubinabd/auth_service/token"
	pb "github.com/Mubinabd/auth_service/genproto"
)

type UserManager struct {
	db *sql.DB
}

func NewUserManager(db *sql.DB) *UserManager {
	return &UserManager{db: db}
}

func (user *UserManager) RegisterUser(req *pb.UserCreate) (*pb.User, error) {

	query := `
			INSERT INTO users
			(id, username, password, email) 
			VALUES ($1, $2, $3, $4) RETURNING username, password, email `

	var resp pb.User
	err := user.db.QueryRow(query, req.Id, req.Username, req.Password, req.Email).
		Scan(&resp.Username, &resp.Password, &resp.Email)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (user *UserManager) LoginUser(login *pb.LoginReq) (*pb.Token, error) {
	var usernameDB, passwordDB, userID string
	query := `SELECT id, username, password FROM users WHERE username = $1`
	err := user.db.QueryRow(query, login.Username).Scan(&userID, &usernameDB, &passwordDB)
	if err != nil {
		return nil, err
	}
	qualify := true
	if passwordDB != login.Password || usernameDB != login.Username {
		qualify = false
	}
	if !qualify {
		return nil, errors.New("username or password incorrect")
	}
	token, err := token.GenereteJWTToken(userID, login.GetUsername())
	if err != nil {
		return nil, err
	}
	return token, nil
}


func (user *UserManager) GetUser(get *pb.ByUsername) (*pb.User, error) {
	query := "SELECT username,email,password FROM users WHERE username = $1"

	var us pb.User

	row := user.db.QueryRow(query, get.Username)
	if err := row.Scan(&us.Username, &us.Password, &us.Email); err != nil {
		log.Println("error while getting user:", err)
		return nil, err
	}

	return &us, nil
}
