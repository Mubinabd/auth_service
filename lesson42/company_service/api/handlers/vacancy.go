package handlers

import (
	"net/http"
	"project/models"

	"github.com/gin-gonic/gin"
)

func (h *HTTPHandler) GetVacancy(c *gin.Context) {
	id := c.Param("id")
	vacancy, err := h.VM.GetVacancyByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error getting vacancy": err.Error()})
		return
	}
	c.JSON(http.StatusOK, vacancy)
}

func (h *HTTPHandler) GetAllVacancies(c *gin.Context) {
	position := c.Query("position")
	min_exp := c.Query("min_exp")
	companyID := c.Query("company_id")

	vacancies, err := h.VM.GetAllVacancies(position, min_exp, companyID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error getting all vacancies": err.Error()})
		return
	}
	c.JSON(http.StatusOK, vacancies)
}

func (h *HTTPHandler) CreateVacancy(c *gin.Context) {
	var vacancy models.VacancyCreated
	c.BindJSON(&vacancy)
	err := h.VM.CreateVacancy(&vacancy)
	if err != nil {
		c.JSON(500, gin.H{"Error creating vacancy": err.Error()})
		h.Logger.ERROR.Println("Vacancy not created: ", err.Error())
		return
	}
	c.String(200, "Vacancy created")
	h.Logger.INFO.Println("Vacancy created")
}

func (h *HTTPHandler) UpdateVacancy(c *gin.Context) {
	var vacancy models.VacancyUpdated
	err := c.BindJSON(&vacancy)
	if err != nil {
		c.JSON(500, gin.H{"Error updating vacancy": err.Error()})
		h.Logger.ERROR.Println("Vacancy not updated: ", err.Error())
		return
	}
	err = h.VM.UpdateVacancy(&vacancy)
	if err != nil {
		c.JSON(500, gin.H{"Error updating vacancy": err.Error()})
		h.Logger.ERROR.Println("Vacancy not updated: ", err.Error())
		return
	}
	c.String(200, "Vacancy updated")
	h.Logger.INFO.Println("Vacancy updated")
}

func (h *HTTPHandler) DeleteVacancy(c *gin.Context) {
	id := c.Param("id")
	err := h.VM.DeleteVacancy(id)
	if err != nil {
		c.JSON(500, gin.H{"Error deleting vacancy": err.Error()})
		h.Logger.ERROR.Println("Vacancy not deleted: ", err.Error())
		return
	}
	c.String(200, "Vacancy deleted")
	h.Logger.INFO.Println("Vacancy deleted")
}
