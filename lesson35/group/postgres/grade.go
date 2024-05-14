package postgres

import (
	"database/sql"
	"github.com/husanmusa/NT_Golang_10/lesson35/group/models"
	"time"

	"github.com/google/uuid"
)

type GradeDB struct {
	Db *sql.DB
}

func NewGradeDB(db *sql.DB) *GradeDB {
	return &GradeDB{Db: db}
}

func (gdb *GradeDB) Create(grade *models.Grade) error {
	err := gdb.Db.Ping()
	if err != nil {
		return err
	}

	_, err = gdb.Db.Exec(
		"INSERT INTO grade (grade_id, student_course_id, grade, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)",
		uuid.New().String(),
		grade.StudentCourseID,
		grade.Grade,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (gdb *GradeDB) Update(grade *models.Grade) error {
	err := gdb.Db.Ping()
	if err != nil {
		return err
	}

	query := `
		UPDATE grade
		SET student_course_id = $1,
		    grade = $2,
		    updated_at = $3
		WHERE grade_id = $4
		AND deleted_at IS NULL
	`
	_, err = gdb.Db.Exec(query,
		grade.StudentCourseID,
		grade.Grade,
		time.Now(),
		grade.GradeID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (gdb *GradeDB) Delete(id string) error {
	err := gdb.Db.Ping()
	if err != nil {
		return err
	}

	query := `
		UPDATE grade
		SET deleted_at = $1,
		    updated_at = $2
		WHERE grade_id = $3
		AND deleted_at IS NULL
	`
	_, err = gdb.Db.Exec(query,
		time.Now().Unix(),
		time.Now(),
		id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (gdb *GradeDB) Read(id *string) (*models.Grade, error) {
	err := gdb.Db.Ping()
	if err != nil {
		return nil, err
	}

	var grade models.Grade

	query := `
		SELECT grade_id, student_course_id, grade, created_at, updated_at, deleted_at
		FROM grade
		WHERE grade_id = $1
		AND deleted_at = 0
	`

	err = gdb.Db.QueryRow(query, id).Scan(
		&grade.GradeID,
		&grade.StudentCourseID,
		&grade.Grade,
		&grade.CreatedAt,
		&grade.UpdatedAt,
		&grade.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return &grade, nil
}

func (gdb *GradeDB) ReadAll() ([]*models.Grade, error) {
	err := gdb.Db.Ping()
	if err != nil {
		return nil, err
	}

	query := `SELECT grade_id, student_course_id, grade, created_at, updated_at, deleted_at FROM grade WHERE deleted_at = 0`

	rows, err := gdb.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var grades []*models.Grade

	for rows.Next() {
		grade := &models.Grade{}
		err = rows.Scan(
			&grade.GradeID,
			&grade.StudentCourseID,
			&grade.Grade,
			&grade.CreatedAt,
			&grade.UpdatedAt,
			&grade.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		grades = append(grades, grade)
	}

	return grades, nil
}

func (gdb *GradeDB) ReadAllByStudentCourseID(studentCourseID string) ([]*models.Grade, error) {
	err := gdb.Db.Ping()
	if err != nil {
		return nil, err
	}

	query := `SELECT grade_id, student_course_id, grade, created_at, updated_at, deleted_at 
	          FROM grade 
	          WHERE student_course_id = $1 
	          AND deleted_at = 0`

	rows, err := gdb.Db.Query(query, studentCourseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var grades []*models.Grade

	for rows.Next() {
		grade := &models.Grade{}
		err = rows.Scan(
			&grade.GradeID,
			&grade.StudentCourseID,
			&grade.Grade,
			&grade.CreatedAt,
			&grade.UpdatedAt,
			&grade.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		grades = append(grades, grade)
	}

	return grades, nil
}
