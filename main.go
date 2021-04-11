package main

import (
	"github.com/jainharsh21/MetX-Backend/config"
	"github.com/jainharsh21/MetX-Backend/env"
)

func main()  {
	env.SetEnvVar()
	// Database
	config.Connect()
}