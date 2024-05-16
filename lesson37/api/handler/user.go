package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/husanmusa/NT_Golang_10/lesson37/models"
	"net/http"
)

func (h Handler) Test(c *gin.Context) {
	c.String(http.StatusOK, "Hello World from Test")
	var person models.Person
	err := c.BindJSON(&person)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	fmt.Println(person)
	c.String(200, "SUCCESS")
}
