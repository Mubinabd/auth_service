package handler

import (
	pb "github.com/Mubinabd/auth_service/genproto"

	"github.com/gin-gonic/gin"
)
// @Router 				/auth/register [POST]
// @Summary 			REGISTER USER
// @Description		 	This api registers user
// @Tags 				USER
// @Accept 				json
// @Produce 			json
// @Param data 			body pb.UserCreate true "UserCreate"
// @Success 201 		{object} pb.User
// @Failure 400 		string Error
func (h *HandlerStruct) RegisterUser(c *gin.Context) {
	var req pb.UserCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user, err := h.UserService.RegisterUser(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}

// @Router 				/auth/profile/{username} [GET]
// @Summary 			GET USER
// @Description		 	This api GETS user by username
// @Tags 				USER
// @Accept 				json
// @Produce 			json
// @Security    		BearerAuth
// @Param 			    username path string true "USERNAME"
// @Success 200			{object} pb.User
// @Failure 400 		string Error
// @Failure 404 		string Error
func (h *HandlerStruct) GetUser(c *gin.Context) {
	var req pb.ByUsername
	username := c.Param("username")
	req.Username = username
	
	user, err := h.UserService.GetUser(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}

// @Router 				/auth/login [POST]
// @Summary 			Login USER
// @Description		 	This api logs  user in
// @Tags 				USER
// @Accept 				json
// @Produce 			json
// @Param data 			body pb.LoginReq true "LoginReq"
// @Success 201 		{object} pb.Token
// @Failure 400 		string Error
func (h *HandlerStruct) LoginUser(c *gin.Context) {
	var req pb.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	token, err := h.UserService.LoginUser(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, token)
}
