package controller

import (
	"github.com/arilsonsantos/crud-go.git/src/model/service"
	"github.com/gin-gonic/gin"
)

func NewUserControllerInterface(serviceInterface service.UserDomainServiceInterface) UserControllerInterface {
	return &userControllerInterface{
		userService: serviceInterface,
	}
}

type UserControllerInterface interface {
	Create(c *gin.Context)
	FindById(c *gin.Context)
	FindByEmail(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	FindUserLogin(c *gin.Context)
}

type userControllerInterface struct {
	userService service.UserDomainServiceInterface
}
