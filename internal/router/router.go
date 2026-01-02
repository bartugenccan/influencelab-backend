package router

import (
	"influencelab-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/health", handler.HealthCheck)
		api.POST("/analyze", handler.Analyze)
	}

	return r
}
