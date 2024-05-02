package main

import (
	"github.com/husanmusa/NT_Golang_10/lesson27/postgres"
)



func main() {
	
	db, err := postgres.DBConn()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = postgres.Ping(db)
	if err != nil {
		panic(err)
	}

	// test := postgres.EmployeeSt{db}
	emp := postgres.NewEmployee(db)
	

	emp.Create()	
	
	
	// student, course ga CRUD. Bularni hammasini CLIda bajara olish lozim.



	// es := []Employee{}
	// rows, err := db.Query(`select e.id, e.name, d.name as d_name, e.salary from employee e join department d on 
	// d.id=e.department_id`)
	// if err != nil {
	// 	panic(fmt.Errorf("error while rows. err:%e", err))
	// }
	// for rows.Next() {
	// 	emp := Employee{}
	// 	rows.Scan(&emp.Id, &emp.Name, &emp.DepartmentName, &emp.Salary)
	// 	if err != nil {
	// 		panic(fmt.Errorf("error while scan for. err:%e", err))
	// 	}
	// 	es = append(es, emp)
	// }

	// fmt.Println(es)

	// fmt.Println("successfull")
}
