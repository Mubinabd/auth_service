package main

import (
	"net/http"
)

func main() {
	
	http.HandleFunc("/", greetingFunc)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func greetingFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HELLOOO"))
}
