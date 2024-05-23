package models

type Resume struct {
	ResumeID    string `json:"id"`
	Position    string `json:"position"`
	Experience  int    `json:"experience"`
	Description string `json:"description"`
	User        User   `json:"user"`
}

type ResumeCreated struct {
	Position    string `json:"position"`
	Experience  int    `json:"experience"`
	Description string `json:"description"`
	User        string   `json:"user_id"`
}

type ResumeUpdated struct {
	ResumeID    string `json:"id"`
	Position    string `json:"position"`
	Experience  int    `json:"experience"`
	Description string `json:"description"`
}

type Resumes struct {
	Resumes []Resume
	Count   int
}
