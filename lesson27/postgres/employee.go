package postgres

import (
	"database/sql"

	"github.com/husanmusa/NT_Golang_10/lesson27/models"

	"github.com/google/uuid"
)

type EmployeeSt struct {
	Db *sql.DB
}

func NewEmployee(db *sql.DB) *EmployeeSt {
	return &EmployeeSt{db}
}

func (e *EmployeeSt) Create() error {
	dId := "5939ed2c-4b4b-43ae-a0ab-695175c3a731"
	uuid := uuid.New()
	_, err := e.Db.Exec(
		"insert into employee(id, name,department_id, salary) values($1, $2, $3, $4)",
		uuid.String(),
		"Falonchi Pistonchi",
		dId,
		234534,
	)
	if err != nil {
		return err
	}

	return nil
}

func (employee *EmployeeSt) getById(id string) (*models.Employee, error) {
	e := models.Employee{}
	err := employee.Db.QueryRow(`select e.id, e.name, d.name as d_name, e.salary from employee e join department d on
	d.id=e.department_id`).Scan(&e.Id, &e.Name, &e.DepartmentName, &e.Salary)
	if err != nil {
		return nil, err
	}

	return &e, nil
}
