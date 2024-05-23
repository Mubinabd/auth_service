package models

import "time"

type Interview struct {
	InterviewID string    `json:"id"`
	User        User      `json:"user"`
	Vacancy     Vacancy   `json:"vacancy"`
	Recruiter   Recruiter `json:"recruiter"`
	Date        string    `json:"date"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt int       `json:"deleted_at"`
}

type InterviewCreated struct {
	UserID      string `json:"user_id"`
	VacancyID   string `json:"vacancy_id"`
	RecruiterID string `json:"recruiter_id"`
	Date        string `json:"date"`
}

type InterviewUpdated struct {
	InterviewID string `json:"id"`
	UserID      string `json:"user_id"`
	VacancyID   string `json:"vacancy_id"`
	RecruiterID string `json:"recruiter_id"`
	Date        string `json:"date"`
}

type Interviews struct {
	Interviews []Interview
	Count      int
}
