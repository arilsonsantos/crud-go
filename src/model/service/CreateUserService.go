package service

import (
	logger "github.com/arilsonsantos/crud-go.git/src/configuration/logger"
	"github.com/arilsonsantos/crud-go.git/src/errors"
	"github.com/arilsonsantos/crud-go.git/src/model/domain"
	"go.uber.org/zap"
)

func (us *userDomainService) Create(userDomain domain.UserDomainInterface) (
	domain.UserDomainInterface, *errors.ErrorDto,
) {
	logger.Info("Init create user domain/service.", zap.String("journey", "createUser"))
	userDomain.EncryptPassword()
	userRepositoryInterface, err := us.userRepository.Create(userDomain)
	if err != nil {
		return nil, err
	}
	return userRepositoryInterface, nil
}
