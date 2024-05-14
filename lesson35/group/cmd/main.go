package main

import (
	"github.com/husanmusa/NT_Golang_10/lesson35/group/handler"
	"github.com/husanmusa/NT_Golang_10/lesson35/group/postgres"
	"log"
)

func main() {
	db, err := postgres.DBConn()
	if err != nil {
		panic(err)
	}
	student := postgres.NewStudent(db)
	course := postgres.NewCourse(db)
	student_course := postgres.NewStudentCourse(db)
	grade := postgres.NewGradeDB(db)
	server := handler.Handler(student, course, student_course, grade)
	log.Println("Server started on localhost:8080")
	log.Fatal(server.ListenAndServe())
}
