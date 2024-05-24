package api

import (
	"github.com/gin-gonic/gin"
	"github.com/husanmusa/NT_Golang_10/lesson43/api/handlers"
)

func NewGin(h *handlers.HTTPHandler) *gin.Engine {
	r := gin.Default()

	user := r.Group("/product")
	user.POST("", h.CreateProduct)

	return r
}
