package main

import "fmt"

type Product struct {
	id          int
	name        string
	price       int
	expiredDate string
	isNew       bool
	barCode     int
	madeIn      string
}

type Order struct {
	id   int
	date string
	prod Product
}

type User struct {
	id     int
	name   string
	orders []*Order
}

func main() {

	products := []*Product{
		{
			1,
			"Water",
			10000,
			"2022-01-01",
			true,
			534645,
			"Uzbekistan",
		},
		{
			2,
			"Juice",
			10000,
			"2022-01-01",
			false,
			534645,
			"Uzbekistan",
		},
		{
			3,
			"Water",
			10000,
			"2022-01-01",
			true,
			534645,
			"Uzbekistan",
		},
	}
	//user := &User{
	//	id:   1,
	//	name: "John",
	//}
	var basket []Product
	for {
		fmt.Println("1. mahsulot tanlash\n2. savatni ko'rish\n3. rasmiylashtirish\n4. buyurtmalarim")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			getAllProducts(products)
			fmt.Println("Mahsulot idsini kiriting:")
			fmt.Scan(&choice)

			prod := getProduct(products, choice)
			basket = append(basket, prod)
		}
	}
}

func getAllProducts(prs []*Product) {

	for _, v := range prs {
		fmt.Printf("%+v", v)
	}
}

func getProduct(pr []*Product, id int) Product {
	for _, v := range pr {
		if v.id == id {
			return *v
		}
	}
	return Product{}
}
