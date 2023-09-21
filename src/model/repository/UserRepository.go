package repository

import (
	"context"
	"github.com/arilsonsantos/crud-go.git/src/configuration/logger"
	"github.com/arilsonsantos/crud-go.git/src/errors"
	"github.com/arilsonsantos/crud-go.git/src/model/domain"
	"github.com/arilsonsantos/crud-go.git/src/model/repository/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
)

var (
	MongodbUserCollection = "MONGODB_USER_COLLECTION"
)

func (ur *userRepositoryInterface) Create(userDomainInterface domain.UserDomainInterface) (
	domain.UserDomainInterface, *errors.ErrorDto,
) {
	logger.Info("Init create user repository")

	collectionName := os.Getenv(MongodbUserCollection)
	collection := ur.databaseConnection.Collection(collectionName)

	value := entity.UserDomainToEntity(userDomainInterface)

	result, err := collection.InsertOne(context.Background(), value)

	if err != nil {
		return nil, errors.InternalServerError(err.Error())
	}

	value.ID = result.InsertedID.(primitive.ObjectID)

	return entity.UserEntityToDomain(*value), nil

}
