package handler

import (
	"fmt"
	"net/http"
)

func (h *handler) get(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body, r.Header, r.Method, r.URL, r.Header["Age"])

	w.Write([]byte("GET request received"))
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("POST request received"))
}

func put(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PUT request received"))
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DELETE request received"))
}
