package main

import (
	"database/sql"
	"fmt"
	"github.com/husanmusa/NT_Golang_10/lesson38/config"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load()

	dbCon := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDatabase)
	fmt.Println(dbCon)
	db, err := sql.Open("postgres", dbCon)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}
