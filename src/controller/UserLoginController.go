package controller

import (
	"fmt"
	"github.com/arilsonsantos/crud-go.git/src/configuration/logger"
	"github.com/arilsonsantos/crud-go.git/src/controller/dto"
	"github.com/arilsonsantos/crud-go.git/src/errors/validation"
	"github.com/arilsonsantos/crud-go.git/src/model/domain"
	"github.com/arilsonsantos/crud-go.git/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func (uc *userControllerInterface) FindUserLogin(c *gin.Context) {
	logger.Info("Init user login controller", zap.String("UserLoginController", "FindUserLogin"))
	var userLogin dto.UserLogin

	if err := c.ShouldBindJSON(&userLogin); err != nil {
		logger.Error("Error trying to validate user login", err, zap.String("UserLoginController", "FindUserLogin"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userLogin)

	userLoginDomain := domain.NewUserLoginDomain(
		userLogin.Email,
		userLogin.Password,
	)

	domainResult, err := uc.userService.LoginUserService(userLoginDomain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User login found with success", zap.String("UserLoginController", "FindUserLogin"))

	c.JSON(http.StatusOK, view.ConvertUserDomainToUserDto(domainResult))
}
