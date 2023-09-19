package service

import (
	"fmt"
	logger "github.com/arilsonsantos/crud-go.git/src/configuration/logger"
	"github.com/arilsonsantos/crud-go.git/src/errors"
	"github.com/arilsonsantos/crud-go.git/src/model"
	"go.uber.org/zap"
)

func (uc *userDomainService) Create(userDomain model.UserDomainInterface) *errors.ErrorDto {
	logger.Info("Init create user domain/service.", zap.String("journey", "createUser"))
	userDomain.EncryptPassword()
	fmt.Println(userDomain.GetPassword())
	return nil
}
