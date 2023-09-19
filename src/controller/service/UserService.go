package service

import (
	"fmt"
	"github.com/arilsonsantos/crud-go.git/src/controller/model/request"
	"github.com/arilsonsantos/crud-go.git/src/errors/validation"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)
}

func Update(c *gin.Context) {}

func Delete(c *gin.Context) {}

func FindBydId(c *gin.Context) {}

func FindBydEmail(c *gin.Context) {}
