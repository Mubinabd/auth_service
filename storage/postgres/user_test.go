package postgres

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	pb "github.com/Mubinabd/auth_service/genproto"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	userManager := NewUserManager(db)

	req := &pb.UserCreate{
		Id:       "1",
		Username: "dior",
		Password: "dior",
		Email:    "dior",
	}

	mock.ExpectQuery(`INSERT INTO users`).
		WithArgs(req.Id, req.Username, req.Password, req.Email).
		WillReturnRows(sqlmock.NewRows([]string{"username", "dior", "email"}).
			AddRow(req.Username, req.Password, req.Email))

	resp, err := userManager.RegisterUser(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, req.Username, resp.Username)
	assert.Equal(t, req.Password, resp.Password)
	assert.Equal(t, req.Email, resp.Email)
}

func TestLoginUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	userManager := NewUserManager(db)

	loginReq := &pb.LoginReq{
		Username: "dior",
		Password: "dior",
	}

	mock.ExpectQuery(`SELECT id, username, password FROM users WHERE username = \$1`).
		WithArgs(loginReq.Username).
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "dior"}).
			AddRow("1", loginReq.Username, loginReq.Password))

	mock.ExpectExec(`UPDATE users`).
		WillReturnResult(sqlmock.NewResult(1, 1))

	token, err := userManager.LoginUser(loginReq)
	assert.NoError(t, err)
	assert.NotNil(t, token)
}

func TestGetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	userManager := NewUserManager(db)

	getReq := &pb.ByUsername{
		Username: "dior",
	}

	mock.ExpectQuery(`SELECT username,email,password FROM users WHERE username = \$1`).
		WithArgs(getReq.Username).
		WillReturnRows(sqlmock.NewRows([]string{"username", "email", "dior"}).
			AddRow("dior", "dior", "dior"))

	user, err := userManager.GetUser(getReq)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "dior", user.Username)
	assert.Equal(t, "dior", user.Email)
	assert.Equal(t, "dior", user.Password)
}
