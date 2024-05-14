package handler

import (
	"encoding/json"
	"github.com/husanmusa/NT_Golang_10/lesson35/group/models"
	"io"
	"log"
	"net/http"
	"strings"
)

func (student HandlerStruct) GetAll(w http.ResponseWriter, r *http.Request) {
	students, err := student.Student.ReadAll()
	if err != nil {
		log.Fatal("error while reading from student", err)
	}
	data, err := json.MarshalIndent(&students, "\t", " ")
	if err != nil {
		log.Fatal("error while marshaling student", err)
	}
	w.Write(data)
}
func (student HandlerStruct) GetById(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/student/")
	students, err := student.Student.Read(&id)
	if err != nil {
		log.Fatal("error while reading from student", err)
	}

	data, err := json.MarshalIndent(&students, "\t", " ")
	if err != nil {
		log.Fatal(err)
	}
	w.Write(data)

}
func (student HandlerStruct) Create(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.Write([]byte("Body could not be read. err:" + err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	NewStudent := models.Student{}
	err = json.Unmarshal(body, &NewStudent)
	if err != nil {
		w.Write([]byte("Object could not be Unmarshalled. err:" + err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = student.Student.Create(&NewStudent)
	if err != nil {
		w.Write([]byte("Error while creating a car. err:" + err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write([]byte("Successed"))
	w.WriteHeader(http.StatusCreated)
}

func (student HandlerStruct) Update(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/student/")
	students, err := student.Student.Read(&id)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.Write([]byte("Body could not be read. err:" + err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &students)
	if err != nil {
		w.Write([]byte("Object could not be Unmarshalled. err:" + err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = student.Student.Update(students)
	if err != nil {
		w.Write([]byte("Error while updating a car. err:" + err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write([]byte("Successed"))
	w.WriteHeader(http.StatusCreated)
}

func (student HandlerStruct) Delete(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/student/")
	err := student.Student.Delete(&id)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte("Succesfully deleted!"))
}
