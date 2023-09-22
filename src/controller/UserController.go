package controller

import (
	"fmt"
	"github.com/arilsonsantos/crud-go.git/src/configuration/logger"
	"github.com/arilsonsantos/crud-go.git/src/controller/dto"
	"github.com/arilsonsantos/crud-go.git/src/errors"
	"github.com/arilsonsantos/crud-go.git/src/errors/validation"
	"github.com/arilsonsantos/crud-go.git/src/model/domain"
	"github.com/arilsonsantos/crud-go.git/src/model/repository/entity"
	"github.com/arilsonsantos/crud-go.git/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"net/http"
	"net/mail"
)

func (uc *userControllerInterface) Create(c *gin.Context) {
	logger.Info("Init create user controller", zap.String("UserController", "Create"))
	var userRequest dto.UserRequestDto

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user", err, zap.String("journey", "createUser"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)

	userDomain := domain.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	domainResult, err := uc.userService.Create(userDomain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User added with success", zap.String("journey", "createUser"))

	c.JSON(http.StatusCreated, view.ConvertUserDomainToUserDto(domainResult))
}

func (uc *userControllerInterface) Update(c *gin.Context) {
	logger.Info("Init update user controller", zap.String("UserController", "Update"))
	var UserUpdateRequestDto dto.UserUpdateRequestDto

	if err := c.ShouldBindJSON(&UserUpdateRequestDto); err != nil {
		logger.Error("Error trying to validate user", err, zap.String("journey", "createUser"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errorRest := errors.BadRequestError("Invalid userId, must be a Hex value")
		c.JSON(errorRest.Code, errorRest)
	}

	userDomain := domain.NewUserUpdateDomain(
		UserUpdateRequestDto.Name,
		UserUpdateRequestDto.Age,
	)

	err := uc.userService.Update(userId, userDomain)
	if err != nil {
		logger.Error("Error trying to call userService update",
			err,
			zap.String("userService", "update"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User updated with success",
		zap.String("userId", userId),
		zap.String("journey", "createUser"))

	c.Status(http.StatusOK)
}

func (uc *userControllerInterface) Delete(c *gin.Context) {
	logger.Info("Init delete user controller", zap.String("UserController", "Delete"))

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errorRest := errors.BadRequestError("Invalid userId, must be a Hex value")
		c.JSON(errorRest.Code, errorRest)
	}

	err := uc.userService.Delete(userId)
	if err != nil {
		logger.Error("Error trying to call userService delete",
			err,
			zap.String("userService", "delete"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User deleted with success",
		zap.String("userId", userId),
		zap.String("journey", "createUser"))

	c.Status(http.StatusOK)
}

func (uc *userControllerInterface) FindById(c *gin.Context) {
	logger.Info("Init findById user controller", zap.String("UserController", "FindById"))

	userId := c.Param("userId")

	//if _, err := uuid.Parse(userId); err != nil {
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errorMessage := errors.BadRequestError("UserId is not valid ID")
		logger.Error(errorMessage.Message, err, zap.String("UserController", "FindById"))
		c.JSON(errorMessage.Code, errorMessage)
	}

	userDomain, err := uc.userService.FindById(userId)
	if err != nil {
		c.JSON(err.Code, err)
		logger.Error("Error trying to call service", err, zap.String("FindUserService", "FindById"))
		return
	}

	logger.Info("User found successfully.", zap.String("UserController", "FindById"))

	c.JSON(http.StatusOK, view.ConvertUserDomainToUserDto(userDomain))
}

func (uc *userControllerInterface) FindByEmail(c *gin.Context) {
	logger.Info("Init findById user controller", zap.String("UserController", "FindByEmail"))

	userEmail := c.Param("userEmail")

	if _, err := mail.ParseAddress(userEmail); err != nil {
		errorMessage := errors.BadRequestError("Email is not valid.")
		logger.Error(errorMessage.Message, err, zap.String("UserController", "FindByEmail"))
		c.JSON(errorMessage.Code, errorMessage)
	}

	userDomain, err := uc.userService.FindByEmail(userEmail)
	if err != nil {
		c.JSON(err.Code, err)
		logger.Error("Error trying to call service", err, zap.String("FindUserService", "FindByEmail"))
		return
	}

	logger.Info("User found successfully.", zap.String("UserController", "FindByEmail"))

	c.JSON(http.StatusOK, entity.UserDomainToEntity(userDomain))
}
