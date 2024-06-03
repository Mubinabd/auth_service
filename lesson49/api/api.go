package api

import (
	"github.com/gin-gonic/gin"
	_ "github.com/husanmusa/NT_Golang_10/lesson49/api/docs"
	"github.com/husanmusa/NT_Golang_10/lesson49/api/handler"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewGin(handler *handler.Handler) *gin.Engine {

	r := gin.Default()
	r.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/coffee", handler.BuyCoffee)
	r.POST("/courier", handler.Delivering)

	return r
}
