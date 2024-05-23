package postgresql

import (
	"database/sql"
	"fmt"
	"log"
	cf "project/config"

	_ "github.com/lib/pq"
)

func ConnectDB(config cf.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable", config.DB_HOST, config.DB_USER, config.DB_NAME, config.DB_PASSWORD, config.DB_PORT)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	log.Println("Successfully connected to the db pgsql!")
	return db, nil
}
