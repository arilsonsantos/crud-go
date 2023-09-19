package main

import (
	"github.com/arilsonsantos/crud-go.git/src/configuration/logger"
	"github.com/arilsonsantos/crud-go.git/src/configuration/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	logger.Info("Starting application...")
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatal("Deu pau!")
	//}
	//fmt.Println(os.Getenv("TESTE"))

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
