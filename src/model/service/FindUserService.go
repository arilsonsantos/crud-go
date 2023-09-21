package service

import (
	"github.com/arilsonsantos/crud-go.git/src/configuration/logger"
	"github.com/arilsonsantos/crud-go.git/src/errors"
	"github.com/arilsonsantos/crud-go.git/src/model/domain"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindByEmail(email string) (domain.UserDomainInterface, *errors.ErrorDto) {
	logger.Info("Init find user domain/service.", zap.String("FindUserService", "FindByEmail"))
	return ud.userRepository.FindByEmail(email)
}

func (ud *userDomainService) FindById(id string) (
	domain.UserDomainInterface, *errors.ErrorDto,
) {
	logger.Info("Init find user domain/service.", zap.String("FindUserService", "FindById"))
	return ud.userRepository.FindById(id)
}
