package service

import (
	"github.com/arilsonsantos/crud-go.git/src/errors"
	"github.com/arilsonsantos/crud-go.git/src/model/domain"
	"github.com/arilsonsantos/crud-go.git/src/model/repository"
)

func NewUserDomainService(userRepository repository.UserRepositoryInterface) UserDomainServiceInterface {
	return &userDomainService{
		userRepository,
	}
}

type userDomainService struct {
	userRepository repository.UserRepositoryInterface
}

type UserDomainServiceInterface interface {
	Create(domainInterface domain.UserDomainInterface) (domain.UserDomainInterface, *errors.ErrorDto)
	FindByEmail(email string) (domain.UserDomainInterface, *errors.ErrorDto)
	FindById(id string) (domain.UserDomainInterface, *errors.ErrorDto)
	Update(string, domain.UserDomainInterface) *errors.ErrorDto
	Delete(string) *errors.ErrorDto
	LoginUserService(domainInterface domain.UserDomainInterface) (domain.UserDomainInterface, string, *errors.ErrorDto)
}
