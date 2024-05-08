package main

import (
	"database/sql"
	"fmt"
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

	var a int
	for i := 0; i < 5000; i++ {
		// model := faker.Name()
		go func() {
			year := faker.YearString()
			// num := faker.SetRandomNumberBoundaries(1000, 2000)
			// color := faker.Word()
			// owner := faker.Name()
			t := time.Now()
			err = db.QueryRow("select count(1) from cars where year = $1",
				year).Scan(&a)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(i, a, time.Now().Sub(t))
		}()
		// if i%1000 == 0 {
		// 	fmt.Println(i)
		// }
	}
	time.Sleep(5*time.Second)
}
