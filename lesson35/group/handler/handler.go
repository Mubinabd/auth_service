package handler

import (
	muxes "github.com/gorilla/mux"
	"github.com/husanmusa/NT_Golang_10/lesson35/group/postgres"
	"net/http"
)

type HandlerStruct struct {
	Student         postgres.StudentDB
	Course          postgres.CourseDB
	GradeDB         postgres.GradeDB
	StudentCourseDB postgres.StudentCourseDB
}

func NewHandler(student *postgres.StudentDB) *HandlerStruct {
	return &HandlerStruct{Student: *student}
}

func Handler(student *postgres.StudentDB, course *postgres.CourseDB, student_course *postgres.StudentCourseDB, grade *postgres.GradeDB) *http.Server {
	mux := http.NewServeMux()
	handler := HandlerStruct{Student: *student, Course: *course, StudentCourseDB: *student_course, GradeDB: *grade}

	m := muxes.NewRouter()

	mux.HandleFunc("POST /student", handler.Create)
	mux.HandleFunc("GET /student", handler.GetAll)
	mux.HandleFunc("GET /student/", handler.GetById)
	mux.HandleFunc("PUT /student/", handler.Update)
	mux.HandleFunc("DELETE /student/", handler.Delete)

	mux.HandleFunc("POST /course", handler.CreateCourse)
	mux.HandleFunc("GET /course", handler.GetAllCourses)
	mux.HandleFunc("GET /course/", handler.GetCourseByID)
	mux.HandleFunc("PUT /course/", handler.UpdateCourse)
	mux.HandleFunc("DELETE /course/", handler.DeleteCourse)
	m.HandleFunc("/course-students/", handler.GetStudentsCourse).Methods("GET")
	mux.Handle("/course-students/", m)

	mux.HandleFunc("POST /student_course", handler.CreateStudentCourse)
	mux.HandleFunc("GET /student_course", handler.GetAllStudentCourses)
	mux.HandleFunc("GET /student_course/", handler.GetStudentCourseByID)
	mux.HandleFunc("GET /student_course/st/", handler.GetStudentCourseByStudentID)
	mux.HandleFunc("GET /student_course/cr/", handler.GetStudentCourseByCourseID)
	mux.HandleFunc("PUT /student_course/", handler.UpdateStudentCourse)
	mux.HandleFunc("DELETE /student_course/", handler.DeleteStudentCourse)

	mux.HandleFunc("POST /grade", handler.CreateGrade)
	mux.HandleFunc("GET /grade", handler.GetAllGrades)
	mux.HandleFunc("GET /grade/", handler.GetGradeByID)
	mux.HandleFunc("GET /grade/stc/", handler.GetGradeByStudentCourseID)
	mux.HandleFunc("PUT /grade/", handler.UpdateGrade)
	mux.HandleFunc("DELETE /grade/", handler.DeleteGrade)

	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}
