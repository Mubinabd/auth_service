package main

import (
	"fmt"
	"lesson12/models"
)

func main() {
	atm := models.NewATM()
	cards := models.NewCard()
	person := models.NewPerson(cards)

	fmt.Println("Bankomatdasiz")
	input := 0
	fmt.Scan(&input)

	card := person.SelectCard(input)
	fmt.Println("Parolni kiriting: ")
	fmt.Scan(&input)

	if !atm.CheckPass(input, card) {
		panic("Password Xato!")
	}

	for {
		fmt.Println("1. Balans\n2.Naq pul olish\n3.Parol o'zgartirish\n4. Exit")
		fmt.Scan(&input)
		switch input {
		case 1:
			fmt.Printf("Sizda %d so'm pul bor", atm.CheckBalance(card))
		case 2:
			fmt.Println("yechmoqchi bo'lgan summani kiriting")
			fmt.Scan(&input)
			if !atm.WithDraw(input, card) {
				fmt.Println("Balansda yetarli mablag' mavjudmas")
			} else {
				fmt.Println("Muvaffaqiyatli")
			}
		case 3:
			pass := 0
			fmt.Println("Yangi Parol kiriting: ")
			fmt.Scan(&pass)
			fmt.Println("Yangi parolni takror kiriting")
			fmt.Scan(&input)
			if atm.ChangePass(pass, input, &card) {
				fmt.Println("Muvaffaqiyatli")
			} else {
				fmt.Println("Muvaffaqiyatsiz")
			}
		default:
			return
		}
	}
}

// ATM
// structs:
//     ATM
//     Person
//     Card

// ------------------––––––----------------------––––––----
// type summer interface {
// 	sum() string
// 	sub() string
// }

// type float struct {
// 	x float64
// 	y float64
// }

// func (f float) sum() string {
// 	return fmt.Sprint(f.x + f.y)
// }

// func (f float) sub() string {
// 	return fmt.Sprint(f.x - f.y)
// }

// type integr struct {
// 	x, y int
// }

// func (i integr) sum() string {
// 	return fmt.Sprint(i.x + i.y)
// }

// func (i integr) sub() string {
// 	return fmt.Sprint(i.x - i.y)
// }

// type str string

// func (s str) sum() string {
// 	return string(s + s)
// }

// func main() {
// 	// INTERFACES
// 	f := float{4, 5}
// 	i := integr{65, 234}
// 	var fp summer = f
// 	var ip summer = i
// 	printer(fp)

// 	printer(ip)

// 	var j interface{}

// 	fmt.Println(j)
// 	j = 34
// 	fmt.Println(j)
// 	j = "Hello"
// 	fmt.Println(j)
// 	j = f
// 	fmt.Println(j)
// }

// func printer(s summer) {
// 	fmt.Println(s.sum())
// 	fmt.Println(s.sub())
// }
