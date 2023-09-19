package controller

import (
	"fmt"
	"github.com/arilsonsantos/crud-go.git/src/controller/dto"
	"github.com/arilsonsantos/crud-go.git/src/model"
	service "github.com/arilsonsantos/crud-go.git/src/model/service"
	"net/http"

	"github.com/arilsonsantos/crud-go.git/src/configuration/logger"
	"github.com/arilsonsantos/crud-go.git/src/errors/validation"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func Create(c *gin.Context) {
	logger.Info("Init create user")
	var userRequest dto.UserRequestDto

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user", err, zap.String("journey", "createUser"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)

	userDomain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	service := service.NewUserDomainService()
	if err := service.Create(userDomain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User added with success", zap.String("journey", "createUser"))

	c.String(http.StatusOK, "")
}

func Update(c *gin.Context) {}

func Delete(c *gin.Context) {}

func FindBydId(c *gin.Context) {}

func FindBydEmail(c *gin.Context) {}
