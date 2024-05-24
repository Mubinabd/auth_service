package handler

import (
	"context"
	"github.com/husanmusa/NT_Golang_10/lesson38/lesson37/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (User HandlerStruct) GetAllUsersHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	data, err := User.User.GetAll(ctx)
	if err != nil {
		if err.Error() == "no rows found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": data})
}
func (h *HandlerStruct) GetUserByIdHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	userID := c.Param("id")

	user, err := h.User.GetById(ctx, &userID)
	if err != nil {
		if err.Error() == "no rows found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
func (User HandlerStruct) CreateUserHandler(c *gin.Context) {
	var (
		user models.User
		err  error
	)

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	// Read body from JSON
	if err = c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error Binding data: " + err.Error()})
		return
	}

	if err := User.User.Create(ctx, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
func (User HandlerStruct) UpdateUserHandler(c *gin.Context) {
	var (
		user models.User
		err  error
	)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	// Read user data from JSON
	if err = c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error Binding data: " + err.Error()})
		return
	}

	// Perform the update operation
	if err := User.User.Update(ctx, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (user HandlerStruct) DeleteUserHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	userID := c.Param("id")

	err := user.User.Delete(ctx, &userID)
	if err != nil {
		if err.Error() == "no rows found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error" + err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted Success"})
}

func (h *HandlerStruct) GetTaskByUserIdHandler(c *gin.Context) {
	// Get task ID from URL parameter
	userID := c.Param("user_id")

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	// Fetch the task from the database
	task, err := h.Task.GetByUserId(ctx, &userID)
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

func (h *HandlerStruct) GetUserByRoleHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	role := c.Query("role")

	user, err := h.User.GetByRole(ctx, &role)
	if err != nil {
		if err.Error() == "no rows found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error" + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
