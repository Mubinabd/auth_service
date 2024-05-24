package model

type Basket struct {
	ID   string `json:"id" db:"id"`
	User User   `json:"user_id" db:"user_id"`
}
