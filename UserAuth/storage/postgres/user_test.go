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

	reg := &pb.UserCreate{
		Username: "Mubina",
		Password: "mubina0804",
		Email:    "Mubina@gmail.com",
	}

	mock.ExpectExec("INSERT INTO users").
		WithArgs(sqlmock.AnyArg(), reg.Username, reg.Password, reg.Email).
		WillReturnResult(sqlmock.NewResult(1, 1))

	user, err := userManager.RegisterUser(reg)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, reg.Username, user.Username)
	assert.Equal(t, reg.Password, user.Password)
	assert.Equal(t, reg.Email, user.Email)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestLoginUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	userManager := NewUserManager(db)

	loginReq := &pb.LoginReq{
		Username: "Mubina",
		Password: "mubina0804",
	}

	rows := sqlmock.NewRows([]string{"id", "username", "password", "email"}).
		AddRow("1", "Mubina", "mubina0804", "Mubina@gmail.com")

	mock.ExpectQuery("SELECT \\* FROM users WHERE username = \\$1 AND password = \\$2").
		WithArgs(loginReq.Username, loginReq.Password).
		WillReturnRows(rows)

	token, err := userManager.LoginUser(loginReq)
	assert.NoError(t, err)
	assert.NotNil(t, token)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	userManager := NewUserManager(db)

	getReq := &pb.ByUsername{
		Username: "Mubina",
	}

	rows := sqlmock.NewRows([]string{"id", "username", "password", "email"}).
		AddRow("1", "Mubina", "mubina0804", "mubina@gmail.com")

	mock.ExpectQuery("SELECT \\* FROM users WHERE username = \\$1").
		WithArgs(getReq.Username).
		WillReturnRows(rows)

	user, err := userManager.GetUser(getReq)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "Mubina", user.Username)
	assert.Equal(t, "mubina0804", user.Password)
	assert.Equal(t, "mubina@gmail.com", user.Email)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
