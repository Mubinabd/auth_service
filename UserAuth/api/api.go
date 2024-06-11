package api

import (
	"github.com/Mubinabd/auth_service/api/handler"
	"github.com/Mubinabd/auth_service/api/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Title Online Voting System Swagger UI
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name role
func NewGin(h *handler.HandlerStruct) *gin.Engine {
	r := gin.Default()
	
	r.Use(middleware.AuthMiddleware)

	r.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	user := r.Group("/auth")
	{
		user.POST("/register", h.RegisterUser)
		user.POST("/login", h.LoginUser)
		user.GET("/profile", h.GetUser)
	}
	return r
}
