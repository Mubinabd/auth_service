package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"gitgub.com/husanmusa/NT_Golang_10/additional/models"
)

func (h *handler) getById(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	fmt.Println(r.URL.Query()["id"])
	id := strings.TrimPrefix(r.URL.Path, "/id/")
	fmt.Println(id)
	w.Write([]byte("GetById request received"))
}

func (h *handler) getAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query(), r.URL.Query()["age"], r.URL.Query()["g"], len(r.URL.Query()["g"]))

	h.User.GetAll(r.URL.Query()["age"])

	w.Write([]byte(""))
}

func createUser(w http.ResponseWriter, r *http.Request) {
	body := r.Body

	buffer := make([]byte, 1024)
	n, _ := body.Read(buffer)
	fmt.Println(string(buffer))
	u := models.User{}
	err := json.Unmarshal(buffer[:n], &u)
	if err != nil {
		panic(err)
	}
	w.Write([]byte(u.Name))
}
