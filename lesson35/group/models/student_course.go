package models

import (
	"time"

	"github.com/google/uuid"
)

type StudentCourse struct {
	StudentCourseID string     `json:"student_course_id" db:"student_course_id"`
	StudentID       string     `json:"student_id" db:"student_id"`
	CourseID        uuid.UUID  `json:"course_id" db:"course_id"`
	CreatedAt       *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt       int64      `json:"deleted_at" db:"deleted_at"`
}
