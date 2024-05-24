package model

type Product struct {
	ID        string   `json:"id" db:"id"`
	Name      string   `json:"name" db:"name"`
	Price     int      `json:"price" db:"price"`
	Category  Category `json:"category_id" db:"category_id"`
	ExpiredAt string   `json:"expired_at" db:"expired_at"`
	CreatedAt string   `json:"created_at" db:"created_at"`
	UpdatedAt string   `json:"updated_at" db:"updated_at"`
	DeletedAt int64    `json:"deleted_at" db:"deleted_at"`
}

type CreateProduct struct {
	ID         string `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	Price      int    `json:"price" db:"price"`
	CategoryID string `json:"category_id" db:"category_id"`
	ExpiredAt  string `json:"expired_at" db:"expired_at"`
	CreatedAt  string `json:"created_at" db:"created_at"`
	UpdatedAt  string `json:"updated_at" db:"updated_at"`
	DeletedAt  int64  `json:"deleted_at" db:"deleted_at"`
}
