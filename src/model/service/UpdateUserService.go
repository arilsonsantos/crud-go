package service

import (
	logger "github.com/arilsonsantos/crud-go.git/src/configuration/logger"
	"github.com/arilsonsantos/crud-go.git/src/errors"
	"github.com/arilsonsantos/crud-go.git/src/model/domain"
	"go.uber.org/zap"
)

func (ud *userDomainService) Update(userId string, userDomain domain.UserDomainInterface) *errors.ErrorDto {
	logger.Info("Init update user domain/service.", zap.String("UpdateUserService", "update"))

	err := ud.userRepository.Update(userId, userDomain)
	if err != nil {
		logger.Error("Error traying to call repository",
			err,
			zap.String("UpdateUserService", "update"))

		return err
	}

	logger.Info("User updated successfully",
		zap.String("userId", userId),
		zap.String("UpdateUserService", "update"))

	return nil
}
