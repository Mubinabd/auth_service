package main

import (
	"github.com/husanmusa/NT_Golang_10/lesson28/postgres"
)

func main() {
	db, err := postgres.ConnectDb()
	if err != nil {
		panic(err)
	}

	card := postgres.NewCardRepo(db)

	card.Create()
}
