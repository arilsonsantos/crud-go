package main

import (
	"fmt"
	"log"
	"os"

	"github.com/arilsonsantos/crud-go.git/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Deu pau!")
	}
	fmt.Println(os.Getenv("TESTE"))

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
