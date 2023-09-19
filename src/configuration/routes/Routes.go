package routes

import (
	userservice "github.com/arilsonsantos/crud-go.git/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("users/:userId", userservice.FindBydId)
	r.GET("users/byEmail/:userEmail", userservice.FindBydEmail)
	r.POST("users", userservice.Create)
	r.PUT("users/:userId", userservice.Update)
	r.DELETE("users/:userId", userservice.Delete)

}
