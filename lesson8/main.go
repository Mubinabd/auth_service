//package main
//
//import "fmt"
//
//// struct - har xil tipdagi ma'lumotlarni bitta tip ostiga
//// (o'zgaruvchi) yig'ish
//type Product struct {
//	name        string
//	price       int
//	expiredDate string
//	isNew       bool
//	barCode     int
//	madeIn      string
//}
//
//type Order struct {
//	name string
//	date string
//	prod []Product
//}
//
////type Electronic struct {
////	year  int
////	brand string
////	name  string
////	price int
////}
////
////var Electronic struct {
////	year  int
////	brand string
////	name  string
////	price int
////}
//
////type Computer struct {
////	Electronic
////	RAM         int
////	monitorSize float64
////}
//
//func main() {
//	var Electronic = struct {
//		year  int
//		brand string
//		name  string
//		price int
//	}{
//		year:  2022,
//		brand: "Samsung",
//		name:  "S22",
//		price: 20000,
//	}
//
//	fmt.Println(Electronic)
//
//	//var p Product
//	//var p1 Product
//	//var products []Product
//
//	//var product = Product{
//	//	name:        "Milk",
//	//	price:       10000,
//	//	expiredDate: "2022-01-01",
//	//	madeIn:      "Uzbekistan",
//	//}
//	//
//	//product.isNew = true
//	//product.barCode = 534645
//	//
//	//var ordered Order
//	//ordered.name = "1 buyurtma"
//	//ordered.date = "today"
//	//ordered.prod = append(ordered.prod, product)
//	//
//	//var order2 = Order{
//	//	name: "2 buyurtma",
//	//	date: "tomorrow",
//	//	prod: []Product{
//	//		{
//	//			name:        "Water",
//	//			price:       10000,
//	//			expiredDate: "2022-01-01",
//	//			isNew:       true,
//	//			barCode:     534645,
//	//			madeIn:      "Uzbekistan",
//	//		},
//	//		{"Juice",
//	//			10000,
//	//			"2022-01-01",
//	//			false,
//	//			534645,
//	//			"Uzbekistan",
//	//		},
//	//	},
//	//}
//	//order2.prod[0].madeIn = "Qashqadaryo"
//	//fmt.Printf("name: %s\n sana: %s\n", order2.name, order2.date)
//	//
//	//for _, v := range order2.prod {
//	//	if v.expiredDate[:4] == "2022" {
//	//		fmt.Println(v)
//	//	}
//	//}
//
//	//cmp := []Computer{
//	//	Electronic{
//	//		year:  2022,
//	//		brand: "Apple",
//	//		name:  "MacBook",
//	//		price: 10000000,
//	//	},
//	//	16,
//	//	15.5,
//	//}
//	//
//	//cmp.year = 54
//	//
//	//if cmp.brand = "HP"
//	//
//	//fmt.Println(cmp.year)
//
//	var p = &Product{}
//
//	p.name = "hello"
//
//	fmt.Println(p)
//}
//
//// user
//// products
//// order
//
//// 3ta struct yasaladi.  User product sotib ola oladi.
//// Sotib olingan product ordersga
//// yoziladi
//func makeOrder(user User, products []Product) Order {
//
//	return order
//}
