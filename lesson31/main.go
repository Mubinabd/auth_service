package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	faker "github.com/go-faker/faker/v3"

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
	defer db.Close()

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 350000; i++ {
		model := faker.Name()
		year := faker.YearString()
		num := faker.SetRandomNumberBoundaries(1000, 2000)
		color := faker.Word()
		owner := faker.Name()
		_, err = db.Exec("insert into cars (model, year, num, color, owner) values ($1,$2,$3,$4,$5)",
			model, year, num, color, owner)
		if err != nil {
			fmt.Println(err)
		}
		if i%1000 == 0 {
			fmt.Println(i)
		}
	}
}
