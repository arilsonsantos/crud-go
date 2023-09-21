package service

import (
	"github.com/arilsonsantos/crud-go.git/src/errors"
	"github.com/arilsonsantos/crud-go.git/src/model/domain"
	"github.com/arilsonsantos/crud-go.git/src/model/repository"
)

func NewUserDomainService(userRepository repository.UserRepositoryInterface) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepositoryInterface
}

type UserDomainService interface {
	Create(domainInterface domain.UserDomainInterface) (domain.UserDomainInterface, *errors.ErrorDto)
	Find(string) (*domain.UserDomainInterface, *errors.ErrorDto) // Update the return type here
	Update(string, domain.UserDomainInterface) *errors.ErrorDto
	Delete(string) *errors.ErrorDto
}
