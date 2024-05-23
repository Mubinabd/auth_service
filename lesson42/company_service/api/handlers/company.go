package handlers

import (
	"net/http"
	"project/models"

	"github.com/gin-gonic/gin"
)

func (h *HTTPHandler) GetCompany(c *gin.Context) {
	id := c.Param("id")
	user, err := h.CM.GetCompanyByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error getting company": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *HTTPHandler) GetAllCompanies(c *gin.Context) {
	location := c.Query("location")
	companies, err := h.CM.GetAllCompanies(location)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error getting all companies": err.Error()})
		return
	}
	c.JSON(http.StatusOK, companies)
}

func (h *HTTPHandler) CreateCompany(c *gin.Context) {
	var company models.CompanyCreated
	c.BindJSON(&company)
	err := h.CM.CreateCompany(&company)
	if err != nil {
		c.JSON(500, gin.H{"Error creating company": err.Error()})
		h.Logger.ERROR.Println("Company not created: ", err.Error())
		return
	}
	c.String(200, "Company created")
	h.Logger.INFO.Println("Company created")
}

func (h *HTTPHandler) UpdateCompany(c *gin.Context) {
	var company models.CompanyUpdated
	err := c.BindJSON(&company)
	if err != nil {
		c.JSON(500, gin.H{"Error updating company": err.Error()})
		h.Logger.ERROR.Println("Company not updated: ", err.Error())
		return
	}
	err = h.CM.UpdateCompany(&company)
	if err != nil {
		c.JSON(500, gin.H{"Error updating company": err.Error()})
		h.Logger.ERROR.Println("Company not updated: ", err.Error())
	}
	c.String(200, "Company updated")
	h.Logger.INFO.Println("Company updated")
}

func (h *HTTPHandler) DeleteCompany(c *gin.Context) {
	id := c.Param("id")
	err := h.CM.DeleteCompany(id)
	if err != nil {
		c.JSON(500, gin.H{"Error deleting company": err.Error()})
		h.Logger.ERROR.Println("Company not deleted: ", err.Error())
		return
	}
	c.String(200, "Company deleted")
	h.Logger.INFO.Println("Company deleted")
}