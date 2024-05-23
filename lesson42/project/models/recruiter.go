package models

import "time"

type Recruiter struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	PhoneNumber string  `json:"phone_number"`
	BirthDate   string  `json:"birthday"`
	Gender      string  `json:"gender"`
	Company     Company `json:"company"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt int       `json:"deleted_at"`
}

type RecruiterCreated struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	BirthDate   string `json:"birthday"`
	Gender      string `json:"gender"`
	CompanyID   string `json:"company_id"`
}

type RecruiterUpdated struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	BirthDate   string `json:"birthday"`
	Gender      string `json:"gender"`
}

type RecruiterUpdatedCompany struct {
	ID        string `json:"id"`
	CompanyID string `json:"company_id"`
}

type Recruiters struct {
	Recruiters []Recruiter
	Count      int
}
