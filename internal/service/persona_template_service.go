package service

import (
	"influencelab-backend/internal/data"
	"influencelab-backend/internal/model"
)

type PersonaTemplateService struct{}

func NewPersonaTemplateService() *PersonaTemplateService {
	return &PersonaTemplateService{}
}

func (s *PersonaTemplateService) GetFreeTemplates() []model.PersonaTemplate {
	return data.FreePersonaTemplates
}
