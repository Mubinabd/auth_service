package postgres

import (
	"database/sql"
	"github.com/husanmusa/NT_Golang_10/lesson35/group/models"
	"time"

	"github.com/google/uuid"
)

type StudentCourseDB struct {
	Db *sql.DB
}

func NewStudentCourse(db *sql.DB) *StudentCourseDB {
	return &StudentCourseDB{Db: db}
}

func (scdb *StudentCourseDB) Create(studentCourse *models.StudentCourse) error {
	err := scdb.Db.Ping()
	if err != nil {
		return err
	}

	_, err = scdb.Db.Exec(
		"INSERT INTO student_course (student_course_id, student_id, course_id, created_at, updated_at ) VALUES ($1, $2, $3, $4, $5)",
		uuid.New().String(),
		studentCourse.StudentID,
		studentCourse.CourseID,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (scdb *StudentCourseDB) Update(studentCourse *models.StudentCourse) error {
	err := scdb.Db.Ping()
	if err != nil {
		return err
	}

	query := `
		UPDATE student_course
		SET student_id = $1,
		    course_id = $2,
		    updated_at = $3
		WHERE student_course_id = $4
		AND deleted_at = 0
	`
	_, err = scdb.Db.Exec(query,
		studentCourse.StudentID,
		studentCourse.CourseID,
		time.Now(),
		studentCourse.StudentCourseID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (scdb *StudentCourseDB) Delete(id string) error {
	err := scdb.Db.Ping()
	if err != nil {
		return err
	}

	query := `
		UPDATE student_course
		SET deleted_at = $1,
		    updated_at = $2
		WHERE student_course_id = $3
		AND deleted_at = 0
	`
	_, err = scdb.Db.Exec(query,
		time.Now().Unix(),
		time.Now(),
		id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (scdb *StudentCourseDB) Read(id string) (*models.StudentCourse, error) {
	err := scdb.Db.Ping()
	if err != nil {
		return nil, err
	}

	var studentCourse models.StudentCourse

	query := `
		SELECT student_course_id, student_id, course_id, created_at, updated_at, deleted_at
		FROM student_course
		WHERE student_course_id = $1
		AND deleted_at = 0
	`

	err = scdb.Db.QueryRow(query, id).Scan(
		&studentCourse.StudentCourseID,
		&studentCourse.StudentID,
		&studentCourse.CourseID,
		&studentCourse.CreatedAt,
		&studentCourse.UpdatedAt,
		&studentCourse.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return &studentCourse, nil
}

func (scdb *StudentCourseDB) ReadAll() ([]*models.StudentCourse, error) {
	err := scdb.Db.Ping()
	if err != nil {
		return nil, err
	}

	query := `SELECT student_course_id, student_id, course_id, created_at, updated_at, deleted_at FROM student_course WHERE deleted_at = 0`

	rows, err := scdb.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var studentCourses []*models.StudentCourse

	for rows.Next() {
		studentCourse := &models.StudentCourse{}
		err = rows.Scan(
			&studentCourse.StudentCourseID,
			&studentCourse.StudentID,
			&studentCourse.CourseID,
			&studentCourse.CreatedAt,
			&studentCourse.UpdatedAt,
			&studentCourse.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		studentCourses = append(studentCourses, studentCourse)
	}

	return studentCourses, nil
}

func (scdb *StudentCourseDB) ReadAllByStudentID(studentID string) ([]*models.StudentCourse, error) {
	err := scdb.Db.Ping()
	if err != nil {
		return nil, err
	}

	query := `SELECT student_course_id, student_id, course_id, created_at, updated_at, deleted_at 
	          FROM student_course 
	          WHERE student_id = $1 
	          AND deleted_at = 0`

	rows, err := scdb.Db.Query(query, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var studentCourses []*models.StudentCourse

	for rows.Next() {
		studentCourse := &models.StudentCourse{}
		err = rows.Scan(
			&studentCourse.StudentCourseID,
			&studentCourse.StudentID,
			&studentCourse.CourseID,
			&studentCourse.CreatedAt,
			&studentCourse.UpdatedAt,
			&studentCourse.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		studentCourses = append(studentCourses, studentCourse)
	}

	return studentCourses, nil
}
func (scdb *StudentCourseDB) ReadAllByCourseID(courseID *string) ([]*models.StudentCourse, error) {
	err := scdb.Db.Ping()
	if err != nil {
		return nil, err
	}

	query := `SELECT student_course_id, student_id, course_id, created_at, updated_at, deleted_at 
	          FROM student_course 
	          WHERE course_id = $1 
	          AND deleted_at = 0`

	rows, err := scdb.Db.Query(query, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var studentCourses []*models.StudentCourse

	for rows.Next() {
		studentCourse := &models.StudentCourse{}
		err = rows.Scan(
			&studentCourse.StudentCourseID,
			&studentCourse.StudentID,
			&studentCourse.CourseID,
			&studentCourse.CreatedAt,
			&studentCourse.UpdatedAt,
			&studentCourse.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		studentCourses = append(studentCourses, studentCourse)
	}

	return studentCourses, nil
}
