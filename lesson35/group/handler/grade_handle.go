package handler

import (
	"encoding/json"
	"github.com/husanmusa/NT_Golang_10/lesson35/group/models"
	"io"
	"net/http"
	"strings"
)

// Grade Handlers

func (h *HandlerStruct) CreateGrade(w http.ResponseWriter, r *http.Request) {
	var grade models.Grade
	if err := json.NewDecoder(r.Body).Decode(&grade); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.GradeDB.Create(&grade); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(grade)
}

func (h *HandlerStruct) GetAllGrades(w http.ResponseWriter, r *http.Request) {
	grades, err := h.GradeDB.ReadAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(grades)
}

func (h *HandlerStruct) GetGradeByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/grade/")
	grade, err := h.GradeDB.Read(&id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(grade)
}

func (h *HandlerStruct) GetGradeByStudentCourseID(w http.ResponseWriter, r *http.Request) {
	studentCourseID := strings.TrimPrefix(r.URL.Path, "/grade/stc/")
	grades, err := h.GradeDB.ReadAllByStudentCourseID(studentCourseID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(grades)
}

func (h *HandlerStruct) UpdateGrade(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/grade/")
	grade, err := h.GradeDB.Read(&id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "Body could not be read", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, grade)
	if err != nil {
		http.Error(w, "Object could not be unmarshalled", http.StatusBadRequest)
		return
	}

	if err := h.GradeDB.Update(grade); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Grade updated successfully"))
}

func (h *HandlerStruct) DeleteGrade(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/grade/")
	if err := h.GradeDB.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Grade deleted successfully"))
}
