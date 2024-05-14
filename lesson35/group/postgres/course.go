package postgres

import (
	"database/sql"
	"github.com/husanmusa/NT_Golang_10/lesson35/group/models"
	"time"

	"github.com/google/uuid"
)

type CourseDB struct {
	Db *sql.DB
}

func NewCourse(db *sql.DB) *CourseDB {
	return &CourseDB{db}
}

func (coursedb *CourseDB) Create(course *models.Course) error {
	err := coursedb.Db.Ping()
	if err != nil {
		return err
	}

	_, err = coursedb.Db.Exec(
		"INSERT INTO course(course_id, course_name, create_at) VALUES($1, $2, $3)",
		uuid.New().String(),
		course.CourseName,
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}
func (coursedb *CourseDB) Update(course *models.Course) error {
	err := coursedb.Db.Ping()
	if err != nil {
		return err
	}

	query := `
		UPDATE
			course
		SET
			course_name = $1,
			update_at = $2
		WHERE
			course_id = $3
	`
	_, err = coursedb.Db.Exec(query,
		course.CourseName,
		time.Now(), // Use time.Now() directly here
		course.CourseID,
	)
	if err != nil {
		return err
	}
	return nil
}
func (coursedb *CourseDB) Delete(id *string) error {
	err := coursedb.Db.Ping()
	if err != nil {
		return err
	}

	query := `
		UPDATE
			course
		SET
			delete_at = $1,
			update_at = $2
		WHERE
			course_id = $3 
		AND 
			delete_at = 0
	`

	stmt, err := coursedb.Db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		time.Now().Unix(),
		time.Now(),
		id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (coursedb *CourseDB) Read(courseID *string) (*models.Course, error) {
	err := coursedb.Db.Ping()
	if err != nil {
		return nil, err
	}

	var course models.Course

	query := `
		SELECT 
			course_id, course_name, create_at, update_at, delete_at
		FROM 
			course
		WHERE 
			course_id = $1 
		AND 
			delete_at = 0
	`

	err = coursedb.Db.QueryRow(query, courseID).Scan(
		&course.CourseID,
		&course.CourseName,
		&course.CreatedAt,
		&course.UpdatedAt,
		&course.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return &course, nil
}
func (coursedb *CourseDB) ReadAll() ([]*models.Course, error) {
	err := coursedb.Db.Ping()
	if err != nil {
		return nil, err
	}

	query := `SELECT * FROM course`

	rows, err := coursedb.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []*models.Course

	for rows.Next() {
		course := &models.Course{}
		err = rows.Scan(
			&course.CourseID,
			&course.CourseName,
			&course.CreatedAt,
			&course.UpdatedAt,
			&course.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	return courses, nil
}

func (c *CourseDB) GetBestStudents() ([]*models.BestStudentsByGroup, error) {
	query := `WITH top_students AS (
					SELECT
					c.course_name,
					s.name AS student_name,
					g.grade,
					RANK() OVER (PARTITION BY c.course_id ORDER BY g.grade DESC) AS rank
					FROM
					course c
					JOIN student_course sc ON c.course_id = sc.course_id
					JOIN grade g ON sc.student_course_id = g.student_course_id
					JOIN student s ON sc.student_id = s.student_id
					)
					
					SELECT
					course_name, student_name, grade
					FROM
					top_students
					WHERE
					rank = 1`

	rows, err := c.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []*models.BestStudentsByGroup

	for rows.Next() {
		student := &models.BestStudentsByGroup{}
		err = rows.Scan(
			&student.CourseName,
			&student.StudentName,
			&student.Grade,
		)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return students, nil
}
