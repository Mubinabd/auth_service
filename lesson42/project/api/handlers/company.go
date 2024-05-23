package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"project/models"
)

func (h *HTTPHandler) GetCompany(c *gin.Context) {
	id := c.Param("id")
	h.Logger.INFO.Println(id)
	url := fmt.Sprintf("http://localhost:8090/company/%s", id)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusBadGateway, gin.H{"Error getting company": err.Error()})
		return
	}
	defer resp.Body.Close()

	var comp = models.Company{}
	err = json.NewDecoder(resp.Body).Decode(&comp)
	h.Logger.INFO.Println(comp)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"Error decoding": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comp)
}

//	func (h *HTTPHandler) GetAllCompanies(c *gin.Context) {
//		location := c.Query("location")
//		companies, err := h.CM.GetAllCompanies(location)
//		if err != nil {
//			c.JSON(http.StatusNotFound, gin.H{"Error getting all companies": err.Error()})
//			return
//		}
//		c.JSON(http.StatusOK, companies)
//	}
//
//	func (h *HTTPHandler) CreateCompany(c *gin.Context) {
//		var company models.CompanyCreated
//		c.BindJSON(&company)
//		err := h.CM.CreateCompany(&company)
//		if err != nil {
//			c.JSON(500, gin.H{"Error creating company": err.Error()})
//			h.Logger.ERROR.Println("Company not created: ", err.Error())
//			return
//		}
//		c.String(200, "Company created")
//		h.Logger.INFO.Println("Company created")
//	}
func (h *HTTPHandler) UpdateCompany(c *gin.Context) {
	var data map[string]interface{}
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"Error decoding": err.Error()})
		return
	}
	h.Logger.INFO.Println(data)
	buf, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"Error decoding": err.Error()})
		return
	}
	h.Logger.INFO.Println(string(buf))

	// send the PUT request to the API
	url := "http://localhost:8090/company/"
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(buf))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"Error decoding": err.Error()})
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusBadGateway, gin.H{"Error decoding": err.Error()})
		return
	}

	defer resp.Body.Close()
	c.String(200, "Company updated")
	h.Logger.INFO.Println("Company updated")
}

//
//func (h *HTTPHandler) DeleteCompany(c *gin.Context) {
//	id := c.Param("id")
//	err := h.CM.DeleteCompany(id)
//	if err != nil {
//		c.JSON(500, gin.H{"Error deleting company": err.Error()})
//		h.Logger.ERROR.Println("Company not deleted: ", err.Error())
//		return
//	}
//	c.String(200, "Company deleted")
//	h.Logger.INFO.Println("Company deleted")
//}
