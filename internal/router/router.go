package router

import (
	"influencelab-backend/internal/handler"
	"influencelab-backend/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	personaService := service.NewPersonaTemplateService()
	personaHandler := handler.NewPersonaTemplateHandler(personaService)

	api := r.Group("/api")
	{
		api.GET("/health", handler.HealthCheck)
		api.POST("/analyze", handler.Analyze)
		api.GET("/persona-templates/free", personaHandler.GetFreeTemplates)
	}

	return r
}
