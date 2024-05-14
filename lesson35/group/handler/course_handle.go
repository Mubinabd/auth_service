package handler

import (
	"encoding/json"
	"github.com/husanmusa/NT_Golang_10/lesson35/group/models"
	"io"
	"log"
	"net/http"
	"strings"
)

func (h *HandlerStruct) GetAllCourses(w http.ResponseWriter, r *http.Request) {
	courses, err := h.Course.ReadAll()
	if err != nil {
		log.Printf("error while reading from course: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	data, err := json.MarshalIndent(&courses, "", "  ")
	if err != nil {
		log.Printf("error while marshaling courses: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (h *HandlerStruct) GetCourseByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/course/")
	course, err := h.Course.Read(&id)
	if err != nil {
		log.Printf("error while reading course by ID: %v", err)
		http.Error(w, "Course Not Found", http.StatusNotFound)
		return
	}
	data, err := json.MarshalIndent(&course, "", "  ")
	if err != nil {
		log.Printf("error while marshaling course: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (h *HandlerStruct) CreateCourse(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "Body could not be read", http.StatusBadRequest)
		return
	}

	var newCourse models.Course
	err = json.Unmarshal(body, &newCourse)
	if err != nil {
		http.Error(w, "Object could not be unmarshalled", http.StatusBadRequest)
		return
	}

	err = h.Course.Create(&newCourse)
	if err != nil {
		http.Error(w, "Error while creating course"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Course created successfully"))
}

func (h *HandlerStruct) UpdateCourse(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/course/")
	course, err := h.Course.Read(&id)
	if err != nil {
		log.Printf("error while reading course by ID: %v", err)
		http.Error(w, "Course Not Found", http.StatusNotFound)
		return
	}

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "Body could not be read", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, course)
	if err != nil {
		http.Error(w, "Object could not be unmarshalled", http.StatusBadRequest)
		return
	}

	err = h.Course.Update(course)
	if err != nil {
		http.Error(w, "Error while updating course", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Course updated successfully"))
}

func (h *HandlerStruct) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/course/")
	err := h.Course.Delete(&id)
	if err != nil {
		log.Printf("error while deleting course: %v", err)
		http.Error(w, "Error while deleting course", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Course deleted successfully"))
}

func (h *HandlerStruct) GetStudentsCourse(w http.ResponseWriter, r *http.Request) {
	courses, err := h.Course.GetBestStudents()
	if err != nil {
		log.Printf("error while reading from course: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	data, err := json.MarshalIndent(&courses, "", "  ")
	if err != nil {
		log.Printf("error while marshaling courses: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
