package main

import (
	fmt "fmt"
	"github.com/arilsonsantos/crud-go.git/src/configuration/logger"
	"github.com/arilsonsantos/crud-go.git/src/configuration/routes"
	"github.com/arilsonsantos/crud-go.git/src/controller"
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

	//Init dependencies
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
