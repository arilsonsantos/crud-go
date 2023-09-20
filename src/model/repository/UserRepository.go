package repository

import (
	"context"
	"github.com/arilsonsantos/crud-go.git/src/configuration/logger"
	"github.com/arilsonsantos/crud-go.git/src/errors"
	"github.com/arilsonsantos/crud-go.git/src/model"
	"os"
)

var (
	MongodbUserCollection = "MONGODB_USER_COLLECTION"
)

func (ur *userRepositoryInterface) Create(
	userDomainInterface model.UserDomainInterface,
) (model.UserDomainInterface, *errors.ErrorDto) {
	logger.Info("Init create user repository")

	collectionName := os.Getenv(MongodbUserCollection)
	collection := ur.databaseConnection.Collection(collectionName)

	value, err := userDomainInterface.GetJsonValue()
	if err != nil {
		return nil, errors.InternalServerError(err.Error())
	}

	result, err := collection.InsertOne(context.Background(), value)

	if err != nil {
		return nil, errors.InternalServerError(err.Error())
	}

	userDomainInterface.SetID(result.InsertedID.(string))

	return userDomainInterface, nil

}
