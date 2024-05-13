package handler

import (
	"net/http"

	"github.com/husanmusa/NT_Golang_10/lesson34/postgres"
)

type handler struct {
	car *postgres.CarRepo
}

func NewHandler(car *postgres.CarRepo) *handler {
	return &handler{car}
}

func Handler(car *postgres.CarRepo) *http.Server {
	mux := http.NewServeMux()

	handler := NewHandler(car)

	mux.HandleFunc("POST /car", handler.Create)

	return &http.Server{
		Addr:    ":8090",
		Handler: mux,
	}
}
