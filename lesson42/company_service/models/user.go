package models

import "time"

type User struct {
	UserID      string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	BirthDate   string `json:"birthday"`
	Gender      string `json:"gender"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt int       `json:"deleted_at"`
}

type UserCreated struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	BirthDate   string `json:"birthday"`
	Gender      string `json:"gender"`
}

type UserUpdated struct {
	UserID      string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	BirthDate   string `json:"birthday"`
	Gender      string `json:"gender"`
}

type Users struct {
	Users []User
	Count int
}
