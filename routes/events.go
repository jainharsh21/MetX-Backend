package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jainharsh21/MetX-Backend/controllers"
)

func eventRoutes(router *gin.Engine) {
	router.GET("/events", controllers.GetEvents)
	router.PATCH("/events/:eventId/addAttendee/:attendeeId", controllers.AddAttendee)
	router.POST("/event", controllers.CreateEvent)
	router.GET("/event/:eventId", controllers.GetEvent)
	router.PUT("/event/:eventId", controllers.UpdateEvent)
	router.DELETE("/event/:eventId", controllers.DeleteEvent)
}
