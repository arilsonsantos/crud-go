package service

import (
	"github.com/arilsonsantos/crud-go.git/src/errors"
	"github.com/arilsonsantos/crud-go.git/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	Create(domainInterface model.UserDomainInterface) *errors.ErrorDto
	Find(string) (*model.UserDomainInterface, *errors.ErrorDto) // Update the return type here
	Update(string, model.UserDomainInterface) *errors.ErrorDto
	Delete(string) *errors.ErrorDto
}
