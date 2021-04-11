package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jainharsh21/MetX-Backend/config"
	"github.com/jainharsh21/MetX-Backend/routes"
	"github.com/jainharsh21/MetX-Backend/env"
)

func main() {
	env.SetEnvVar()
	// Database
	config.Connect()
	// Init Router
	router := gin.Default()

	// Route Handlers / Endpoints
	routes.Routes(router)

	log.Fatal(router.Run(":3000"))
}
