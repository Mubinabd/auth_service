package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	username := c.GetHeader("username")
	email := c.GetHeader("email")
	password := c.GetHeader("password")
	if username != "username" || email != "email" || password != "password"  {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	timestamp := time.Now()
	path := c.Request.URL.Path
	log.Printf("Authenticated request at %s for path %s\n", timestamp.Format(time.RFC3339), path)
	c.Next()
}

