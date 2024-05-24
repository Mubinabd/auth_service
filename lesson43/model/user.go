package model

type User struct {
	ID        string `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Balance   int    `json:"balance" db:"balance"`
	Email     string `json:"email" db:"email"`
	Phone     string `json:"phone" db:"phone"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
	DeletedAt int64  `json:"deleted_at" db:"deleted_at"`
}
