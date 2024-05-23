package models

import "time"

type Vacancy struct {
	VacancyID   string  `json:"id"`
	Name        string  `json:"name"`
	Position    string  `json:"position"`
	Experience  int     `json:"min_exp"`
	Company     Company `json:"company"`
	Description string  `json:"description"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt int       `json:"deleted_at"`
}

type VacancyCreated struct {
	Name        string `json:"name"`
	Position    string `json:"position"`
	Experience  int    `json:"min_exp"`
	CompanyID   string `json:"company_id"`
	Description string `json:"description"`
}

type VacancyUpdated struct {
	VacancyID   string `json:"id"`
	Name        string `json:"name"`
	Position    string `json:"position"`
	Experience  int    `json:"min_exp"`
	Description string `json:"description"`
}

type Vacancies struct {
	Vacancies []Vacancy
	Count     int
}
