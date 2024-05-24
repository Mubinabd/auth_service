package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/husanmusa/NT_Golang_10/lesson43/model"
)

func (h *HTTPHandler) CreateProduct(c *gin.Context) {
	pr := model.CreateProduct{}
	err := c.BindJSON(&pr)
	if err != nil {
		panic(err)
	}

	err = h.service.PrService.CreateProduct(pr)
	if err != nil {
		c.Error(err)
	}

	c.String(201, "Success")
}
