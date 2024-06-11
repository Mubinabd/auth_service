package postgres

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	pb "github.com/Mubinabd/auth_service/genproto"
	// token "github.com/Mubinabd/auth_service/token"
	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/require"
)

func TestRegisterUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	userStorage := NewUserStorage(db)

	user := &pb.UserCreate{
		Id:       "1",
		Username: "testuser",
		Password: "testpassword",
		Email:    "test@example.com",
	}

	mock.ExpectQuery("insert into users").
		WithArgs(user.Id, user.Username, user.Password, user.Email).
		WillReturnRows(sqlmock.NewRows([]string{"username", "password", "email"}).AddRow(user.Username, user.Password, user.Email))

	registeredUser, err := userStorage.RegisterUser(user)
	assert.NoError(t, err)
	assert.Equal(t, user.Username, registeredUser.Username)
	assert.Equal(t, user.Password, registeredUser.Password)
	assert.Equal(t, user.Email, registeredUser.Email)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUserInfo(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	userStorage := NewUserStorage(db)

	username := "testuser"
	mock.ExpectQuery("select username, password, email from users where username = \\$1").
		WithArgs(username).
		WillReturnRows(sqlmock.NewRows([]string{"username", "password", "email"}).AddRow("testuser", "testpassword", "test@example.com"))

	userInfo, err := userStorage.GetUserInfo(&pb.ByUsername{Username: username})
	assert.NoError(t, err)
	assert.Equal(t, "testuser", userInfo.Username)
	assert.Equal(t, "testpassword", userInfo.Password)
	assert.Equal(t, "test@example.com", userInfo.Email)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// type MockTokenGenerator struct{}

// func (m *MockTokenGenerator) GenerateJWTToken(userID, username string) (*pb.Token, error) {
//     return &pb.Token{
//         AccessToken:  "testAccessToken",
//         RefreshToken: "testRefreshToken",
//     }, nil
// }

// func TestLoginUser(t *testing.T) {
//     db, mock, err := sqlmock.New()
//     require.NoError(t, err)
//     defer db.Close()

//     mockTokenGenerator := &MockTokenGenerator{}
//     userStorage := NewUserStorage(db, mockTokenGenerator)

//     loginReq := &pb.LoginReq{
//         Username: "testuser",
//         Password: "testpassword",
//     }

//     mock.ExpectQuery("select id, username, password from users where username = \\$1").
//         WithArgs(loginReq.Username).
//         WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password"}).AddRow("1", "testuser", "testpassword"))

//     result, err := userStorage.Loginuser(loginReq)
//     require.NoError(t, err)
//     assert.Equal(t, "testAccessToken", result.AccessToken)
//     assert.Equal(t, "testRefreshToken", result.RefreshToken)

//     err = mock.ExpectationsWereMet()
//     require.NoError(t, err)
// }
