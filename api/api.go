package api

import (
	"github.com/Mubinabd/auth_service/api/handler"
	"github.com/Mubinabd/auth_service/middleware"
	_ "github.com/Mubinabd/auth_service/docs"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Title Online Voting System Swagger UI
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewGin(h *handler.HandlerStruct) *gin.Engine {
	r := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost", "http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}
	r.Use(cors.New(corsConfig))

	r.Use(middleware.Middleware())

	user := r.Group("/auth")
	{
		user.POST("/register", h.RegisterUser)
		user.POST("/login", h.LoginUser)
		user.GET("/profile/:username", h.GetUser)
	}
	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return r
}
