package handler

import (
	"encoding/json"
	"github.com/husanmusa/NT_Golang_10/lesson35/group/models"
	"io"
	"net/http"
	"strings"
)

// StudentCourse Handlers

func (h *HandlerStruct) CreateStudentCourse(w http.ResponseWriter, r *http.Request) {
	var studentCourse models.StudentCourse
	if err := json.NewDecoder(r.Body).Decode(&studentCourse); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.StudentCourseDB.Create(&studentCourse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(studentCourse)
}

func (h *HandlerStruct) GetAllStudentCourses(w http.ResponseWriter, r *http.Request) {
	studentCourses, err := h.StudentCourseDB.ReadAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(studentCourses)
}

func (h *HandlerStruct) GetStudentCourseByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/student_course/")
	studentCourse, err := h.StudentCourseDB.Read(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(studentCourse)
}

func (h *HandlerStruct) GetStudentCourseByStudentID(w http.ResponseWriter, r *http.Request) {
	studentID := strings.TrimPrefix(r.URL.Path, "/student_course/st/")
	studentCourses, err := h.StudentCourseDB.ReadAllByStudentID(studentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(studentCourses)
}

func (h *HandlerStruct) GetStudentCourseByCourseID(w http.ResponseWriter, r *http.Request) {
	courseID := strings.TrimPrefix(r.URL.Path, "/student_course/cr/")
	studentCourses, err := h.StudentCourseDB.ReadAllByCourseID(&courseID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(studentCourses)
}

func (h *HandlerStruct) UpdateStudentCourse(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/student_course/")
	studentCourse, err := h.StudentCourseDB.Read(id)
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

	err = json.Unmarshal(body, studentCourse)
	if err != nil {
		http.Error(w, "Object could not be unmarshalled", http.StatusBadRequest)
		return
	}

	if err := h.StudentCourseDB.Update(studentCourse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("StudentCourse updated successfully"))
}

func (h *HandlerStruct) DeleteStudentCourse(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/student_course/")
	if err := h.StudentCourseDB.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("StudentCourse deleted successfully"))
}
