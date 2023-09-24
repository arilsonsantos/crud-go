package service

import (
	logger "github.com/arilsonsantos/crud-go.git/src/configuration/logger"
	"github.com/arilsonsantos/crud-go.git/src/errors"
	"github.com/arilsonsantos/crud-go.git/src/model/domain"
	"go.uber.org/zap"
)

func (ud *userDomainService) loginUserService(userDomain domain.UserDomainInterface) (
	domain.UserDomainInterface, string, *errors.ErrorDto,
) {
	logger.Info("Init login user domain/service.", zap.String("service", "LoginUserService"))

	userDomain.EncryptPassword()
	user, err := ud.findByEmailAndPassword(userDomain.GetEmail(), userDomain.GetPassword())
	if user == nil {
		return nil, "", err
	}

	token, err := user.GenerateToken()
	if err != nil {
		return nil, "", err
	}

	logger.Info(
		"User login service executed successfully.",
		zap.String("userId", user.GetID()),
		zap.String("journey", "LoginUserService"))

	return user, token, nil
}
