package models

import (
	"time"
)

type RoleType string

const (
	// Admin represents a "todo" task type
	Admin RoleType = "admin"
	// Users represents a "done" task type
	Users RoleType = "user"
)

// User represents the "users" table
type User struct {
	ID        string    `json:"id" db:"id" uri:"id"`
	Name      string    `json:"name" db:"name"`
	Lastname  string    `json:"lastname" db:"lastname"`
	Phone     string    `json:"phone" db:"phone"`
	Email     string    `json:"email" db:"email"`
	Age       int       `json:"age" db:"age"`
	Role      RoleType  `json:"role" db:"role"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt int64     `json:"deleted_at" db:"deleted_at"`
}

type AllUsers struct {
	Count uint
	Users []User
}

type UserId struct {
	ID string `json:"id" db:"id" uri:"id"`
}
