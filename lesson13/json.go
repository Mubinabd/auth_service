package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name, LastName string
	Age            int
	Info           []Info
}

type Info struct {
	Id     int    `json:"id"`
	UserId int    `json:"user_id"`
	Title  string `json:"taytil"`
	Body   string `json:"body"`
}

func main() {
	// text := Person{"Jhon", "Bek"}

	// d, err := json.Marshal(text)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(d))
	// fmt.Printf("%T, %+v", d, string(d))

	data := `{"Name":"Jhon","LastName":"Bek", "Info":[{
		"user_id": 1,
		"id": 1,
		"title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
		"body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit	 molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"
	  },
	  {
		"user_id": 1,
		"id": 2,
		"title": "qui est esse",
		"body": "est rerum tempore vitae\nsequi sint nihil reprehenderit dolor beatae ea dolores neque\nfugiat blanditiis voluptate porro vel nihil molestiae ut reiciendis\nqui aperiam non debitis possimus qui neque nisi nulla"
	  }]}`

	p := Person{}
	fmt.Println(json.Valid([]byte(data)))
	err := json.Unmarshal([]byte(data), &p)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", p)
}
