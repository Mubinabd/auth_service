package api

import (
	// Adjust this import path to match your project structure

	"github.com/husanmusa/NT_Golang_10/lesson38/lesson37/api/handler"
	"github.com/husanmusa/NT_Golang_10/lesson38/lesson37/config"
	"github.com/jackc/pgx/v5"

	"github.com/gin-gonic/gin"
)

func NewGin(db *pgx.Conn, config config.Config) *gin.Engine {
	handler := handler.NewHandler(db, config)
	router := gin.Default()

	router.POST("/user", handler.CreateUserHandler)       // POST /user
	router.GET("/user/:id", handler.GetUserByIdHandler)   // GET /user/:id
	router.GET("/user", handler.GetAllUsersHandler)       // GET /user
	router.PUT("/user", handler.UpdateUserHandler)        // PUT /user/
	router.DELETE("/user/:id", handler.DeleteUserHandler) // DELETE /user/

	router.POST("/task", handler.CreateTaskHandler)       // POST /task
	router.GET("/task/:id", handler.GetTaskByIdHandler)   // GET /task/:id
	router.GET("/task", handler.GetAllTasksHandler)       // GET /tasks
	router.PUT("/task", handler.UpdateTaskHandler)        // PUT /task
	router.DELETE("/task/:id", handler.DeleteTaskHandler) // DELETE /task

	// get user's tasks through userid
	router.GET("/user_task/:user_id", handler.GetTaskByUserIdHandler)

	// get tasks by date
	router.GET("/tasks", handler.GetTaskByDate)

	// get tasks after-deadline?user_id=e041b094-ddcc-46fe-be6e-80797e0732b0
	router.GET("/tasks/after-deadline", handler.GetFinishedTasks)

	// getting tasks by type tasks?user_id&type=done
	router.GET("/tasks/type", handler.GetTasksByType)

	// get users by role /users?role=admin
	router.GET("/user/role", handler.GetUserByRoleHandler) // GET /user/:id

	return router
}
