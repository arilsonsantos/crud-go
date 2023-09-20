package controller

import (
	"fmt"
	"net/http"

	"github.com/arilsonsantos/crud-go.git/src/controller/dto"
	"github.com/arilsonsantos/crud-go.git/src/model"
	"github.com/arilsonsantos/crud-go.git/src/view"

	"github.com/arilsonsantos/crud-go.git/src/configuration/logger"
	"github.com/arilsonsantos/crud-go.git/src/errors/validation"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) Create(c *gin.Context) {
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

	domainResult, err := uc.service.Create(userDomain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User added with success", zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, view.ConvertUserDomainToUserDto(domainResult))
}

func (uc *userControllerInterface) Update(c *gin.Context) {}

func (uc *userControllerInterface) Delete(c *gin.Context) {}

func (uc *userControllerInterface) FindById(c *gin.Context) {}

func (uc *userControllerInterface) FindByEmail(c *gin.Context) {}
