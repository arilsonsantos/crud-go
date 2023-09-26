package service

import (
	logger "github.com/arilsonsantos/crud-go.git/src/configuration/logger"
	"github.com/arilsonsantos/crud-go.git/src/errors"
	"github.com/arilsonsantos/crud-go.git/src/model/domain"
	"go.uber.org/zap"
)

func (ud *userDomainService) Create(userDomain domain.UserDomainInterface) (
	domain.UserDomainInterface, *errors.ErrorDto,
) {
	logger.Info("Init create user domain/service.", zap.String("journey", "createUser"))
	userDomain.EncryptPassword()

	user, _ := ud.FindByEmail(userDomain.GetEmail())
	if user != nil {
		return nil, errors.BadRequestError("Email is already registered in another account.")
	}

	userRepositoryInterface, err := ud.userRepository.Create(userDomain)
	if err != nil {
		return nil, err
	}
	return userRepositoryInterface, nil
}
