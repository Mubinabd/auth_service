package handler

import (
	"database/sql"
	"net/http"

	"gitgub.com/husanmusa/NT_Golang_10/additional/postgres"
)

type handler struct {
	db   *sql.DB
	User *postgres.UserSt
}

func NewHandler(db *sql.DB, user *postgres.UserSt) *handler {
	return &handler{
		db:   db,
		User: user,
	}
}

func Handler(db *sql.DB, user *postgres.UserSt) *http.Server {
	mux := http.NewServeMux()
	user.GetAll()
	handler := handler{db, user}

	mux.HandleFunc("GET /qwer", handler.get)
	mux.HandleFunc("GET /id/", getById)
	mux.HandleFunc("GET /user", handler.getAllUsers)
	mux.HandleFunc("POST /user", createUser)

	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}
