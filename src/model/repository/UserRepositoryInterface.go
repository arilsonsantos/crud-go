package repository

import (
	"github.com/arilsonsantos/crud-go.git/src/errors"
	"github.com/arilsonsantos/crud-go.git/src/model"
	mongo "go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepositoryInterface(
	database *mongo.Database,
) UserRepositoryInterface {
	return &userRepositoryInterface{
		database,
	}
}

type userRepositoryInterface struct {
	databaseConnection *mongo.Database
}

type UserRepositoryInterface interface {
	Create(domainInterface model.UserDomainInterface) (model.UserDomainInterface, *errors.ErrorDto)
}
