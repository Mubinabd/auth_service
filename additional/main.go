package main

import (
	"log"

	"gitgub.com/husanmusa/NT_Golang_10/additional/handler"
	"gitgub.com/husanmusa/NT_Golang_10/additional/postgres"
)

func main() {
	db, err := postgres.DBConn()
	if err != nil {
		panic(err)
	}

	user := postgres.NewUser(db)

	server := handler.Handler(db, user)

	log.Println("Server started on localhost:8080")
	log.Fatal(server.ListenAndServe())
}
