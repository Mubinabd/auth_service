package models

type Car struct {
	Id                       string
	Model, Num, Color, Owner string
	Year                     int
	CreatedAt, UpdatedAt     string
}
