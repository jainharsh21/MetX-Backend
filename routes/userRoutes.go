package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jainharsh21/MetX-Backend/controllers"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)
	router.GET("/users", controllers.GetUsers)
	router.POST("/user", controllers.CreateUser)
	router.GET("/user/:userId", controllers.GetUser)
	router.PUT("/user/:userId", controllers.UpdateUser)
	router.DELETE("/user/:userId", controllers.DeleteUser)
	router.NoRoute(notFound)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
	return
}