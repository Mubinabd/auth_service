package main

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

func main() {
	dbCon := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		host, user, dbname, password, port)
	db, err := sql.Open("postgres", dbCon)
	if err != nil {
		panic(err)
	}

	tr, err := db.Begin()

	defer func() {
		tr.Commit()
	}()
	defer func() {
		if err != nil {
			err = tr.Rollback()
			if err != nil {
				panic(err)
			}
		}
	}()

	res, err := tr.Exec("update books set is_sold = false where id = 4")
	if err != nil {
		panic(err)
	}
	re, _ := res.RowsAffected()
	if re == 0 {
		panic(fmt.Errorf("nothing has changed"))
	}

	res, err = tr.Exec("update books set is_sold = true where id = 5")
	if err != nil {
		panic(err)
	}
	re, _ = res.RowsAffected()
	if re == 0 {
		err = fmt.Errorf("nothing has changed")
		panic(err)
	}
}
