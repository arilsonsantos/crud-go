package routes

import (
	"github.com/arilsonsantos/crud-go.git/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("usersById/:userId", controller.FindUserBydId)
	r.GET("usersByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("users", controller.CreateUser)
	r.PUT("users/:userId", controller.UpdateUser)
	r.DELETE("users/:userId", controller.DeleteUser)

}
