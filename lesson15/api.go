package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/time", vaqt)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}

func hello (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<b>Hello World</b>")
	fmt.Fprintf(w, "Hello Worldddd")

}

func vaqt(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	t:= time.Now()
	fmt.Fprint(w, t)
}