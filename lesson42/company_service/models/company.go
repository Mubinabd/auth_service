package models

import "time"

type Company struct {
	CompanyID string `json:"id"`
	Name      string `json:"name"`
	Location  string `json:"location"`
	Workers   int    `json:"workers"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt int       `json:"deleted_at"`
}

type CompanyCreated struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Workers  int    `json:"workers"`
}

type CompanyUpdated struct {
	CompanyID string `json:"id"`
	Name      string `json:"name"`
	Location  string `json:"location"`
	Workers   int    `json:"workers"`
}

type Companies struct {
	Companies []Company
	Count     int
}
