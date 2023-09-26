package routes

import (
	controller "github.com/arilsonsantos/crud-go.git/src/controller"
	"github.com/arilsonsantos/crud-go.git/src/model/domain"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	const pathUserId string = "/users/:userId"

	r.GET(pathUserId, domain.VerifyTokenMiddleware, userController.FindById)
	r.GET("/users/byEmail/:userEmail", domain.VerifyTokenMiddleware, userController.FindByEmail)
	r.POST("/users", userController.Create)
	r.PUT(pathUserId, userController.Update)
	r.DELETE(pathUserId, userController.Delete)
	r.POST("/login", userController.FindUserLogin)

}
