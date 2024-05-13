package main

import (
	"github.com/husanmusa/NT_Golang_10/lesson34/handler"
	"github.com/husanmusa/NT_Golang_10/lesson34/postgres"
)

func main() {

	db, err := postgres.DBConn()
	if err != nil {
		panic(err)
	}

	car := postgres.NewCarRepo(db)

	server := handler.Handler(car)

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
