package main

import (
	"github.com/arilsonsantos/crud-go.git/src/controller"
	"github.com/arilsonsantos/crud-go.git/src/model/repository"
	"github.com/arilsonsantos/crud-go.git/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	userRepository := repository.NewUserRepositoryInterface(database)
	userDomainService := service.NewUserDomainService(userRepository)
	return controller.NewUserControllerInterface(userDomainService)
}
