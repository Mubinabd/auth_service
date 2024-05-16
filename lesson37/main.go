package main

import (
	"database/sql"
	"github.com/husanmusa/NT_Golang_10/lesson37/api"
	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres",
		`postgres://postgres:pass@localhost:5432/postgres?sslmode=disable`)
	if err != nil {
		panic(err)
	}
	r := api.NewGin(db)
	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
