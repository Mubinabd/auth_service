package handlers

import (
	"net/http"
	"project/models"

	"github.com/gin-gonic/gin"
)

func (h *HTTPHandler) GetResume(c *gin.Context) {
	id := c.Param("id")
	user, err := h.RM.GetResumeByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error getting resume": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *HTTPHandler) GetAllResumes(c *gin.Context) {
	position := c.Query("position")
	min_exp := c.Query("min_exp")

	users, err := h.RM.GetAllResumes("", position, min_exp)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error getting all users": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *HTTPHandler) CreateResume(c *gin.Context) {
	var resume models.ResumeCreated
	c.BindJSON(&resume)
	err := h.RM.CreateResume(&resume)
	if err != nil {
		c.JSON(500, gin.H{"Error creating resume": err.Error()})
		h.Logger.ERROR.Println("Resume not created: ", err.Error())
		return
	}
	c.String(200, "Resume created")
	h.Logger.INFO.Println("Resume created")
}

func (h *HTTPHandler) UpdateResume(c *gin.Context) {
	var resume models.ResumeUpdated
	err := c.BindJSON(&resume)
	if err != nil {
		c.JSON(500, gin.H{"Error updating resume": err.Error()})
		h.Logger.ERROR.Println("Resume not updated: ", err.Error())
		return
	}
	err = h.RM.UpdateResume(&resume)
	if err != nil {
		c.JSON(500, gin.H{"Error updating resume": err.Error()})
		h.Logger.ERROR.Panicln("Resume not updated", err.Error())
		return
	}
	c.String(200, "Resume updated")
	h.Logger.INFO.Println("Resume updated")
}

func (h *HTTPHandler) DeleteResume(c *gin.Context) {
	id := c.Param("id")
	err := h.RM.DeleteResume(id)
	if err != nil {
		c.JSON(500, gin.H{"Error deleting resume": err.Error()})
		h.Logger.ERROR.Println("Resume not deleted", err.Error())
		return
	}
	c.String(200, "Resume deleted")
	h.Logger.INFO.Println("Resume deleted")
}