package service

import (
	"fmt"
	"github.com/arilsonsantos/crud-go.git/src/configuration/logger"
	"github.com/arilsonsantos/crud-go.git/src/controller/model/request"
	"github.com/arilsonsantos/crud-go.git/src/controller/model/response"
	"github.com/arilsonsantos/crud-go.git/src/errors/validation"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func Create(c *gin.Context) {
	logger.Info("Init create user")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user", err, zap.String("journey", "createUser"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)

	userResponse := response.UserResponse{
		Id:    "123",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	logger.Info("User added with success", zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, userResponse)
}

func Update(c *gin.Context) {}

func Delete(c *gin.Context) {}

func FindBydId(c *gin.Context) {}

func FindBydEmail(c *gin.Context) {}
