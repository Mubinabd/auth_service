package models

import (
	"time"
)

type Student struct {
	StudentID string     `json:"student_id"`
	Name      string     `json:"name"`
	LastName  string     `json:"lastname"`
	Phone     string     `json:"phone"`
	Age       int        `json:"age"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt int64      `json:"deleted_at"`
}
