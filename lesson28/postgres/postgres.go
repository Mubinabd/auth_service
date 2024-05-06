package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	user     = "sardorbek"
	dbname   = "postgres"
	password = "1111"
	port     = 5432
)

func ConnectDb() (*sql.DB, error) {
	dbCon := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		host, user, dbname, password, port)
	db, err := sql.Open("postgres", dbCon)
	if err != nil {
		return nil, err
	}

	return db, err
}

func Ping(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	}

	return nil
}
