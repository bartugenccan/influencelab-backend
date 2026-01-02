package main

import (
	"log"

	"influencelab-backend/internal/router"
)

func main() {
	r := router.SetupRouter()

	log.Println("Starting server at :8080")
	r.Run(":8080")
}
