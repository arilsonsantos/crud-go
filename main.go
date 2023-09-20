package main

import (
	"context"
	fmt "fmt"
	"github.com/arilsonsantos/crud-go.git/src/configuration/database"
	"github.com/arilsonsantos/crud-go.git/src/configuration/logger"
	"github.com/arilsonsantos/crud-go.git/src/configuration/routes"
	"github.com/arilsonsantos/crud-go.git/src/controller"
	"github.com/arilsonsantos/crud-go.git/src/model/repository"
	"github.com/arilsonsantos/crud-go.git/src/model/service"
	godotenv "github.com/joho/godotenv"
	"log"
	os "os"

	"github.com/gin-gonic/gin"
)

func main() {
	logger.Info("Starting application...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Deu pau!")
	}
	fmt.Println(os.Getenv("TESTE"))

	ctx := context.Background()
	db, err := database.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatalf("Error trying to connect to database, error=%s", err.Error())
		return
	}

	//Init dependencies

	userRepository := repository.NewUserRepositoryInterface(db)
	userDomainService := service.NewUserDomainService(userRepository)
	userController := controller.NewUserControllerInterface(userDomainService)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
