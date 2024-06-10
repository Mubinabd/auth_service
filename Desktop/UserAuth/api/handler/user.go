package handler

import (
	"fmt"
	"net/http"
	pb "github.com/Mubinabd/auth_service/genproto"

	"github.com/gin-gonic/gin"
)
func (h *HandlerStruct) RegisterUser(c *gin.Context) {
	var (
		user pb.UserCreate
		err      error
	)

	if err = c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error when binding JSON: " + err.Error()})
		return
	}

	h.User.RegisterUser(c.Request.Context(), &user)
	fmt.Println(user.Id)

	c.String(http.StatusOK, "User registered successfully")
}

func (h *HandlerStruct)GetUser(c *gin.Context) {
	username := c.Query("username")
	user, err := h.User.GetUser(c.Request.Context(), &pb.ByUsername{Username: username})
	if err != nil {
		c.JSON(400, gin.H{"Error when logging user": err.Error()})
		return
	}
	c.JSON(200, user)
}


func (h *HandlerStruct)LoginUser(c *gin.Context) {
	var(
		login pb.LoginReq
		err error
	)
	_, err = h.User.LoginUser(c.Request.Context(), &login)
	if err != nil {
		c.JSON(400, gin.H{"Error when logging user": err.Error()})
		return
	}
	c.JSON(200, "User logging successfully")
}
