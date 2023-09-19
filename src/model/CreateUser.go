package model

import (
	"fmt"
	logger "github.com/arilsonsantos/crud-go.git/src/configuration/logger"
	"github.com/arilsonsantos/crud-go.git/src/errors"
	"go.uber.org/zap"
)

func (u *UserDomain) Create() *errors.ErrorDto {
	logger.Info("Init create user domain/service.", zap.String("journey", "createUser"))
	u.EncryptPassword()
	fmt.Println(u)
	return nil
}
