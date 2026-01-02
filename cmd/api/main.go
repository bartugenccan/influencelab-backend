package main

import (
	"log"

	"influencelab-backend/internal/config"
	"influencelab-backend/internal/router"
)

func main() {
	// Load environment variables
	config.Load()

	r := router.SetupRouter()

	log.Println("Starting server at :" + config.AppConfig.Port)
	r.Run(":" + config.AppConfig.Port)
}
