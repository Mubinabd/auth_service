package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/husanmusa/NT_Golang_10/lesson37/api/handler"
)

func NewGin(db *sql.DB) *gin.Engine {
	r := gin.Default()

	handler := handler.Handler{db}

	r.POST("/test", handler.Test)
	r.POST("/car", handler.CarCreate)

	return r
}
