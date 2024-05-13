package postgres

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/husanmusa/NT_Golang_10/lesson34/models"
)

type CarRepo struct {
	DB *sql.DB
}

func NewCarRepo(db *sql.DB) *CarRepo {
	return &CarRepo{db}
}

func (u *CarRepo) Create(car models.Car) error {
	query := "insert into cars (id, model, year, num, color, owner) values ($1, $2, $3,$4,$5,$6)"
	_, err := u.DB.Exec(query, uuid.NewString(), car.Model, car.Year, car.Num, car.Color, car.Owner)

	return err
}
func (u *CarRepo) Get(id string) (*models.Car, error) {
	car := models.Car{}

	err := u.DB.QueryRow("select id, model, year, num, color, owner from cars where id = $1 and deleted_at = 0").
		Scan(&car.Id, &car.Model, &car.Year, &car.Num, &car.Color, &car.Owner)
	if err != nil {
		return nil, err
	}

	return &car, nil
}
