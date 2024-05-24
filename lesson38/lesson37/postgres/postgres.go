package postgres

import (
	"context"
	"log"

	_ "github.com/lib/pq"
)

// const (
// 	host     = "localhost"
// 	user     = "sayyidmuhammad"
// 	dbname   = "todo"
// 	password = "root"
// 	port     = 5432
// )

func DBConn(config config.Config) (*pgx.Conn, error) {
	var (
		db  *pgx.Conn
		err error
	)
	db, err = pgx.Connect(context.Background(), "postgresql://sayyidmuhammad:root@localhost:5432/todo")
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
	// dbCon := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
	// 	host, user, dbname, password, port)
	// db, err = sql.Open("postgres", dbCon)
	// if err != nil {
	// 	return nil, err
	// }

	return db, err
}

// func Ping(db *sql.DB) error {
// 	err := db.Ping()
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
