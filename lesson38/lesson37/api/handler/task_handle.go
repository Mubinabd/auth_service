package handler

import (
	"context"
	"github.com/husanmusa/NT_Golang_10/lesson38/lesson37/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *HandlerStruct) CreateTaskHandler(c *gin.Context) {
	var task models.Task
	err := c.BindJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error binding data: " + err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	err = h.Task.Create(ctx, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task created successfully"})
}

func (h *HandlerStruct) UpdateTaskHandler(c *gin.Context) {
	var task models.Task
	err := c.BindJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error binding data: " + err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	err = h.Task.Update(ctx, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

func (h *HandlerStruct) DeleteTaskHandler(c *gin.Context) {
	taskID := c.Param("id")

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	err := h.Task.Delete(ctx, &taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
func (h *HandlerStruct) GetTaskByIdHandler(c *gin.Context) {
	// Get task ID from URL parameter
	taskID := c.Param("id")

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	// Fetch the task from the database
	task, err := h.Task.GetById(ctx, &taskID)
	if err != nil {
		if err.Error() == "no rows found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get task: " + err.Error()})
		}
		return
	}

	// Return the task data
	c.JSON(http.StatusOK, gin.H{"task": task})
}

func (h *HandlerStruct) GetAllTasksHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	tasks, err := h.Task.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get tasks: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func (h *HandlerStruct) GetTaskByDate(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	userId := c.Query("user_id")
	from := c.Query("from")
	to := c.Query("to")
	fromDate, err := time.Parse("02.01.2006", from)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'from' date format. Use 'dd.mm.yyyy'."})
		return
	}

	toDate, err := time.Parse("02.01.2006", to)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'to' date format. Use 'dd.mm.yyyy'."})
		return
	}

	// Dummy response
	tasks, err := h.Task.GetByUserIdAndDateRange(ctx, userId, fromDate, toDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'from' date format. Use 'dd.mm.yyyy'."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func (h *HandlerStruct) GetFinishedTasks(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	userId := c.Query("user_id")

	tasks, err := h.Task.GetFinishedTasks(ctx, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "sql error" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func (h *HandlerStruct) GetTasksByType(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	userId := c.Query("user_id")
	types := c.Query("type")

	tasks, err := h.Task.GetTasksByType(ctx, userId, models.TaskType(types))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "sql error" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}
