package postgres

import (
	"database/sql"
	"github.com/husanmusa/NT_Golang_10/lesson35/group/models"
	"time"

	"github.com/google/uuid"
)

type StudentDB struct {
	Db *sql.DB
}

func NewStudent(db *sql.DB) *StudentDB {
	return &StudentDB{db}
}

func (studentdb *StudentDB) Create(std *models.Student) error {
	err := studentdb.Db.Ping()
	if err != nil {
		return err
	}

	_, err = studentdb.Db.Exec(
		"INSERT INTO student(student_id, name,lastname, phone, age, created_at) VALUES($1, $2, $3, $4, $5, $6)",
		uuid.New().String(),
		std.Name,
		std.LastName,
		std.Phone,
		std.Age,
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (studentdb *StudentDB) Update(std *models.Student) error {
	err := studentdb.Db.Ping()
	if err != nil {
		return err
	}

	query := `
	UPDATE
		student
	SET
		name = $1,
		lastname = $2,
		phone = $3,
		age = $4,
		created_at = $5,
		update_at = $6
	WHERE
		student_id = $7
`
	_, err = studentdb.Db.Exec(query,
		std.Name,
		std.LastName,
		std.Phone,
		std.Age,
		std.CreatedAt,
		time.Now(),
		std.StudentID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (studentdb *StudentDB) Delete(id *string) error {
	err := studentdb.Db.Ping()
	if err != nil {
		return err
	}
	query := `
  UPDATE
    student
  SET
    delete_at = $1,
    update_at = $2
  WHERE
    student_id = $3 
  AND 
	delete_at = 0
`

	_, err = studentdb.Db.Exec(query,
		time.Now().Unix(),
		time.Now(),
		id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (studentdb *StudentDB) Read(id *string) (*models.Student, error) {
	err := studentdb.Db.Ping()
	if err != nil {
		return nil, err
	}
	query := `
	SELECT 
		student_id,
		name,
		lastname,
		phone,
		age,
		created_at,
		update_at,
		delete_at
	FROM 
		student
	WHERE
		student_id = $1
`

	row := studentdb.Db.QueryRow(query, id)

	student := models.Student{}
	err = row.Scan(
		&student.StudentID,
		&student.Name,
		&student.LastName,
		&student.Phone,
		&student.Age,
		&student.CreatedAt,
		&student.UpdatedAt,
		&student.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return &student, nil

}

func (studentdb *StudentDB) ReadAll() ([]*models.Student, error) {
	err := studentdb.Db.Ping()
	if err != nil {
		return nil, err
	}

	query := `SELECT * FROM student`

	rows, err := studentdb.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []*models.Student

	for rows.Next() {
		student := &models.Student{}
		err = rows.Scan(
			&student.StudentID,
			&student.Name,
			&student.LastName,
			&student.Phone,
			&student.Age,
			&student.CreatedAt,
			&student.UpdatedAt,
			&student.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return students, nil
}
