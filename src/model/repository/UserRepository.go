package repository

import (
	"context"
	"fmt"
	"github.com/arilsonsantos/crud-go.git/src/configuration/logger"
	"github.com/arilsonsantos/crud-go.git/src/errors"
	"github.com/arilsonsantos/crud-go.git/src/model/domain"
	"github.com/arilsonsantos/crud-go.git/src/model/repository/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
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

func (ur *userRepositoryInterface) FindByEmail(email string) (
	domain.UserDomainInterface, *errors.ErrorDto) {

	logger.Info("Init findByEmail user repository", zap.String("userRepository", "findByEmail"))

	collectionName := os.Getenv(MongodbUserCollection)
	collection := ur.databaseConnection.Collection(collectionName)

	userEntity := &entity.UserEntity{}
	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(
		context.Background(),
		filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this email: %s", email)
			logger.Error(errorMessage, err, zap.String("userRepository", "findByEmail"))
			return nil, errors.NotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by email"
		logger.Error(errorMessage, err, zap.String("userRepository", "findByEmail"))
		return nil, errors.InternalServerError(errorMessage)
	}

	logger.Info("User found successfully",
		zap.String("userRepository", "findByEmail"),
		zap.String("userId", userEntity.ID.Hex()),
		zap.String("Email", userEntity.Email))
	return entity.UserEntityToDomain(*userEntity), nil
}

func (ur *userRepositoryInterface) FindById(ID string) (
	domain.UserDomainInterface, *errors.ErrorDto) {

	logger.Info("Init findByEmail user repository", zap.String("userRepository", "findByEmail"))

	collectionName := os.Getenv(MongodbUserCollection)
	collection := ur.databaseConnection.Collection(collectionName)

	userEntity := &entity.UserEntity{}
	//filter := bson.D{{Key: "ID", Value: ID}}
	objectId, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.D{{Key: "_id", Value: objectId}}
	err := collection.FindOne(
		context.Background(),
		filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this ID: %s", ID)
			logger.Error(errorMessage, err, zap.String("userRepository", "findById"))
			return nil, errors.NotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by ID"
		logger.Error(errorMessage, err, zap.String("userRepository", "findById"))
		return nil, errors.InternalServerError(errorMessage)
	}

	logger.Info("User found successfully",
		zap.String("userRepository", "findById"),
		zap.String("userId", userEntity.ID.Hex()))
	return entity.UserEntityToDomain(*userEntity), nil
}
