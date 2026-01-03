package handler

import (
	"influencelab-backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PersonaTemplateHandler struct {
	service *service.PersonaTemplateService // Handler service'e ihtiyaç duyuyor dependency injection
}

func NewPersonaTemplateHandler(s *service.PersonaTemplateService) *PersonaTemplateHandler {
	return &PersonaTemplateHandler{service: s}
}

func (h *PersonaTemplateHandler) GetFreeTemplates(c *gin.Context) {
	templates := h.service.GetFreeTemplates()
	c.JSON(http.StatusOK, gin.H{
		"templates": templates,
	})
}
