package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/husanmusa/NT_Golang_10/lesson34/models"
)

func (h handler) Create(w http.ResponseWriter, r *http.Request) {
	// car := postgres.NewCarRepo(h.db)

	// buffer := make([]byte, 1024)
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.Write([]byte("Body could not be read. err:" + err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newCar := models.Car{}
	err = json.Unmarshal(body, &newCar)
	if err != nil {
		w.Write([]byte("Object could not be Unmarshalled. err:" + err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.car.Create(newCar)
	if err != nil {
		w.Write([]byte("Error while creating a car. err:" + err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write([]byte("Successed"))
	w.WriteHeader(http.StatusCreated)
}
