package main

import (
	json "encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	file, _ := os.Open("j.json")
	defer file.Close()
	fileText, _ := io.ReadAll(file)

	product := Product{}
	temp := Temp{}

	json.Unmarshal(fileText, &product)

	json.Unmarshal(fileText, &temp)
	fmt.Println(temp)
	t, err := time.Parse("02/01 15:04:05 2006", temp.Time)
	if err != nil {
		panic(err)
	}
	product.CreatedAt = t
	fmt.Printf("%v %T", product, product.CreatedAt)
}

type Product struct {
	Id                 int
	Title              string
	Price              float64
	Quantity           int
	Total              int
	DiscountPercentage int
	DiscountedPrice    int
	Thumbnail          string
	CreatedAt          time.Time
}

type Temp struct {
	Time string `json:"created_at"`
}
