package handlers

import (
	"net/http"
	"project/models"

	"github.com/gin-gonic/gin"
)

// GetUser godoc
// @ID get_user
// @Router /user/{id} [GET]
// @Summary Get User
// @Description GEt User
// @Tags User
// @Accept json
// @Produce json
// @Param user path string true "id"
// @Success 200 {object} models.User "User data"
// @Response 400 {object} string "Bad Request"
// @Failure 500 {object} string "Server Error"
func (h *HTTPHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.UM.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error getting user": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *HTTPHandler) GetAllUsers(c *gin.Context) {
	from := c.Query("from")
	to := c.Query("to")
	gender := c.Query("gender")

	users, err := h.UM.GetAllUsers(from, to, gender)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error getting all users": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// CreateUser godoc
// @ID create_user
// @Router /user [POST]
// @Summary Create User
// @Description Create User
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.UserCreated true "UserRequest"
// @Success 201 {object} string "User data"
// @Response 400 {object} string "Bad Request"
// @Failure 500 {object} string "Server Error"
func (h *HTTPHandler) CreateUser(c *gin.Context) {
	var user models.UserCreated
	c.BindJSON(&user)
	err := h.UM.CreateUser(&user)
	if err != nil {
		c.JSON(500, gin.H{"Error creating user": err.Error()})
		h.Logger.ERROR.Println("User not created: ", err.Error())
		return
	}

	c.String(200, "User created")
	h.Logger.INFO.Println("User created")
}

func (h *HTTPHandler) UpdateUser(c *gin.Context) {
	var user models.UserUpdated
	c.BindJSON(&user)
	err := h.UM.UpdateUser(&user)
	if err != nil {
		c.JSON(500, gin.H{"Error updating user": err.Error()})
		h.Logger.ERROR.Panicln("User not updated", err.Error())
		return
	}
	c.String(200, "User updated")
	h.Logger.INFO.Println("User updated")
}

func (h *HTTPHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := h.UM.DeleteUser(id)
	if err != nil {
		c.JSON(500, gin.H{"Error deleting user": err.Error()})
		h.Logger.ERROR.Println("User not deleted", err.Error())
		return
	}
	c.String(200, "User deleted")
	h.Logger.INFO.Println("User deleted")
}

func (h *HTTPHandler) GetAllUserResumes(c *gin.Context) {
	id := c.Param("id")
	interviews, err := h.RM.GetAllResumes(id, "", "")

	if err != nil {
		c.JSON(500, gin.H{"Error getting all resumes": err.Error()})
		return
	}
	c.JSON(200, interviews)
}

func (h *HTTPHandler) GetAllUserInterviews(c *gin.Context) {
	//id := c.Param("id")
	////interviews, err := h.IM.GetInterviewsBy(id, true, false)
	//
	//if err != nil {
	//	c.JSON(500, gin.H{"Error getting all interviews": err.Error()})
	//	return
	//}
	//c.JSON(200, interviews)
}
