package model

type Transaction struct {
	ID         string `json:"id" db:"id"`
	Basket     Basket `json:"basket_id" db:"basket_id"`
	TotalPrice int    `json:"total_price" db:"total_price"`
}
