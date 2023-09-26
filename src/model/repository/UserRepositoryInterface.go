package repository

import (
	"github.com/arilsonsantos/crud-go.git/src/errors"
	"github.com/arilsonsantos/crud-go.git/src/model/domain"
	mongo "go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepositoryInterface(database *mongo.Database) UserRepositoryInterface {
	return &userRepositoryInterface{database}
}

type userRepositoryInterface struct {
	databaseConnection *mongo.Database
}

type UserRepositoryInterface interface {
	Create(domainInterface domain.UserDomainInterface) (domain.UserDomainInterface, *errors.ErrorDto)
	FindByEmail(email string) (domain.UserDomainInterface, *errors.ErrorDto)
	FindById(email string) (domain.UserDomainInterface, *errors.ErrorDto)
	Update(userId string, userDomainInterface domain.UserDomainInterface) *errors.ErrorDto
	Delete(userId string) *errors.ErrorDto
	FindByEmailAndPassword(email string, password string) (domain.UserDomainInterface, *errors.ErrorDto)
}
