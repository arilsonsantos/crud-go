package main

import (
	"context"
	"github.com/arilsonsantos/crud-go.git/src/configuration/database"
	"github.com/arilsonsantos/crud-go.git/src/configuration/logger"
	"github.com/arilsonsantos/crud-go.git/src/configuration/routes"
	"github.com/arilsonsantos/crud-go.git/src/controller"
	"github.com/arilsonsantos/crud-go.git/src/model/repository"
	"github.com/arilsonsantos/crud-go.git/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func main() {
	logger.Info("Starting application...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Deu pau!")
	}

	ctx := context.Background()
	db, err := database.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatalf("Error trying to connect to database, error=%s", err.Error())
		return
	}

	userController := initDependencies(db)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	userRepository := repository.NewUserRepositoryInterface(database)
	userDomainService := service.NewUserDomainService(userRepository)
	return controller.NewUserControllerInterface(userDomainService)
}
