package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	generalRoutes(router)
	userRoutes(router)
	eventRoutes(router)
}

