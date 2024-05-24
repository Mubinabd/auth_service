package model

import "github.com/google/uuid"

type Item struct {
	ID        uuid.UUID `json:"id" db:"id"`
	BasketID  uuid.UUID `json:"basket_id" db:"basket_id"`
	ProductID uuid.UUID `json:"product_id" db:"product_id"`
	Quantity  int       `json:"quantity" db:"quantity"`
}

type CreateItem struct {
	BasketID  uuid.UUID `json:"basket_id" db:"basket_id"`
	ProductID uuid.UUID `json:"product_id" db:"product_id"`
	Quantity  int       `json:"quantity" db:"quantity"`
}
