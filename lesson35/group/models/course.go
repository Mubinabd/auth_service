package models

import "time"

type Course struct {
	CourseID   string     `json:"course_id"`
	CourseName string     `json:"course_name"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"update_at"`
	DeletedAt  int64      `json:"delete_at"`
}

type BestStudentsByGroup struct {
	CourseName, StudentName string
	Grade                   float32
}
