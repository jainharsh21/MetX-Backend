package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jainharsh21/MetX-Backend/controllers"
)

func userRoutes(router *gin.Engine) {
	router.GET("/users", controllers.GetUsers)
	router.POST("/user", controllers.CreateUser)
	router.GET("/user/:userId", controllers.GetUser)
	router.PUT("/user/:userId", controllers.UpdateUser)
	router.DELETE("/user/:userId", controllers.DeleteUser)
}