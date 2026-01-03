package model

type PersonaTemplate struct {
	ID           string              `json:"id"`
	Title        string              `json:"title"`
	Description  string              `json:"description"`
	AgeRange     string              `json:"age_range"`
	PersonaCount int                 `json:"persona_count"`
	Constraints  map[string][]string `json:"constraints"`
}
