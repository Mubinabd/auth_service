package models

import (
	"time"
)

type Grade struct {
	GradeID         string     `json:"grade_id" db:"grade_id"`
	StudentCourseID string     `json:"student_course_id" db:"student_course_id"`
	Grade           float64    `json:"grade" db:"grade"`
	CreatedAt       *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt       int64      `json:"deleted_at" db:"deleted_at"`
}
