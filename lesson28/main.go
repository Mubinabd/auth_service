package main

import (
	"fmt"

	"github.com/husanmusa/NT_Golang_10/lesson28/models"
	"github.com/husanmusa/NT_Golang_10/lesson28/postgres"
)

func main() {
	db, err := postgres.ConnectDb()
	if err != nil {
		panic(err)
	}

	card := postgres.NewCardRepo(db)
	// user := postgres.NewUserRepo(db)

	c := models.Card{
		Name:      "axadjonov sardorbek",
		CardNum:   "8600098309872345",
		ExpiredAt: "2029-12-01",
		Password:  1111,
		CardType:  "uzcard",
		UserID:    "85ce1f38-db6f-420a-ac58-6b370cc2ff5e",
	}

	// u := models.User{
	// 	Name: "sardorbek",
	// 	Username: "sardorbek",
	// 	Password: "1234",
	// 	Phone: "998200070424",
	// }

	tr := postgres.NewTransRepo(db)

	// user.Create(&u)
	err = card.Create(&c)

	card1, err := card.GetById("34534")
	if err != nil {
		panic(err)
	}

	card1.Amount, err = tr.GetBalance(card1.ID)
	if err != nil {
		panic(err)
	}

	fmt.Println(card1)

	fmt.Println(err)
}
