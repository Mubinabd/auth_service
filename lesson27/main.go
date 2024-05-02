package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/google/uuid"
)

const (
	host     = "localhost"
	user     = "husanmusa"
	dbname   = "postgres"
	password = "pass"
	port     = 5432
)

type Employee struct {
	Id             string
	Name           string
	DepartmentName string
	Salary         int
}

func main() {
	dbCon := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		host, user, dbname, password, port)
	db, err := sql.Open("postgres", dbCon)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(fmt.Errorf("error while ping. err:%e", err))
	}
	dId := "5939ed2c-4b4b-43ae-a0ab-695175c3a731"
	uuid := uuid.New()
	_, err = db.Exec(
		"insert into employee(id, name,department_id, salary) values($1, $2, $3, $4)",
		uuid.String(),
		"John Doe",
		dId,
		5346436,
	)
	if err != nil {
		panic(fmt.Errorf("error while exec. err:%e", err))
	}

	// e := Employee{}
	// err = db.QueryRow(`select e.id, e.name, d.name as d_name, e.salary from employee e join department d on
	// d.id=e.department_id`).Scan(&e.Id, &e.Name, &e.DepartmentName, &e.Salary)
	// if err != nil {
	// 	panic(fmt.Errorf("error while get. err:%e", err))
	// }

	// fmt.Println(e)

	es := []Employee{}
	rows, err := db.Query(`select e.id, e.name, d.name as d_name, e.salary from employee e join department d on 
	d.id=e.department_id`)
	if err != nil {
		panic(fmt.Errorf("error while rows. err:%e", err))
	}
	for rows.Next() {
		emp := Employee{}
		rows.Scan(&emp.Id, &emp.Name, &emp.DepartmentName, &emp.Salary)
		if err != nil {
			panic(fmt.Errorf("error while scan for. err:%e", err))
		}
		es = append(es, emp)
	}

	fmt.Println(es)

	fmt.Println("successfull")
}
