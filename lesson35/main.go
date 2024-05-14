package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v, Name: %v\n", vars["id"], vars["name"])
}
func first(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("FIRST"))
}

func second(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("SECOND"))
}

type son int

func (son) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("SON"))
}

func main() {
	r := mux.NewRouter()

	raqam := son(1)
	s := r.PathPrefix("/products").Handler(raqam).Subrouter()
	// "/products/"
	s.HandleFunc("/first", first)

	//http.Handle("/", raqam)

	//s.PathPrefix("/second").
	//l.HandleFunc("/three", first)
	//s.HandleFunc("/second", second)
	//r.HandleFunc("/articles", ArticlesHandler)
	//http.Handle("/", r)

	panic(http.ListenAndServe(":8070", r))
}
