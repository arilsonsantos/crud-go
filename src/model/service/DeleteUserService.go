package service

import (
	"github.com/arilsonsantos/crud-go.git/src/configuration/logger"
	"github.com/arilsonsantos/crud-go.git/src/errors"
	"go.uber.org/zap"
)

func (ud *userDomainService) Delete(userId string) *errors.ErrorDto {
	logger.Info("Init delete user domain/service.", zap.String("DeleteUserService", "delete"))

	err := ud.userRepository.Delete(userId)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("DeleteUserService", "delete"))

		return err
	}

	logger.Info("User deleted successfully",
		zap.String("userId", userId),
		zap.String("UpdateUserService", "delete"))

	return nil
}
