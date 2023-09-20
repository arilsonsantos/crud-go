package service

import (
	"github.com/arilsonsantos/crud-go.git/src/errors"
	"github.com/arilsonsantos/crud-go.git/src/model"
	"github.com/arilsonsantos/crud-go.git/src/model/repository"
)

func NewUserDomainService(userRepository repository.UserRepositoryInterface) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepositoryInterface
}

type UserDomainService interface {
	Create(domainInterface model.UserDomainInterface) (model.UserDomainInterface, *errors.ErrorDto)
	Find(string) (*model.UserDomainInterface, *errors.ErrorDto) // Update the return type here
	Update(string, model.UserDomainInterface) *errors.ErrorDto
	Delete(string) *errors.ErrorDto
}
