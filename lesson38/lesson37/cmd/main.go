package main

import (
	"github.com/husanmusa/NT_Golang_10/lesson38/lesson37/api"
	"github.com/husanmusa/NT_Golang_10/lesson38/lesson37/config"
	"github.com/husanmusa/NT_Golang_10/lesson38/lesson37/postgres"
)

func main() {
	cfg := config.Load()
	db, err := postgres.DBConn(cfg)
	if err != nil {
		panic(err)
	}
	server := api.NewGin(db, cfg)

	server.Run(":8080")
}
