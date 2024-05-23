package handlers

import (
	"net/http"
	"project/models"

	"github.com/gin-gonic/gin"
)

func (h *HTTPHandler) GetInterview(c *gin.Context) {
	id := c.Param("id")
	interview, err := h.IM.GetInterviewByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error getting interview": err.Error()})
		return
	}
	c.JSON(http.StatusOK, interview)
}

func (h *HTTPHandler) CreateInterview(c *gin.Context) {
	var interview models.InterviewCreated
	c.BindJSON(&interview)
	err := h.IM.CreateInterview(&interview)
	if err != nil {
		c.JSON(500, gin.H{"Error creating interview": err.Error()})
		h.Logger.ERROR.Println("Interview not created: ", err.Error())
		return
	}
	c.String(200, "Interview created")
	h.Logger.INFO.Println("Interview created")
}

func (h *HTTPHandler) UpdateInterview(c *gin.Context) {
	var interview models.InterviewUpdated
	err := c.BindJSON(&interview)
	if err != nil{
		c.JSON(500, gin.H{"Error updating interview": err.Error()})
		h.Logger.ERROR.Println("Interview not updated: ", err.Error())
		return
	}
	err = h.IM.UpdateInterview(&interview)
	if err != nil {
		c.JSON(500, gin.H{"Error updating interview": err.Error()})
		h.Logger.ERROR.Println("Interview not updated: ", err.Error())
		return
	}
	c.String(200, "Interview updated")
	h.Logger.INFO.Println("Interview updated")
}

func (h *HTTPHandler) DeleteInterview(c *gin.Context) {
	id := c.Param("id")
	err := h.IM.DeleteInterview(id)
	if err != nil {
		c.JSON(500, gin.H{"Error deleting interview": err.Error()})
		h.Logger.ERROR.Println("Interview not deleted: ", err.Error())
		return
	}
	c.String(200, "Interview deleted")
	h.Logger.INFO.Println("Interview deleted")
}

func (h *HTTPHandler) GetAllInterviews(c *gin.Context) {
	companyID := c.Query("companyID")
	interviews, err := h.IM.GetAllInterviews(companyID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error getting all interviews": err.Error()})
		return
	}
	c.JSON(http.StatusOK, interviews)
}