package model

import "mime/multipart"

// Analysis type constants
const (
	AnalysisTypeCoach   = "coach"
	AnalysisTypePersona = "persona"
)

// Response mode constants
const (
	ModeAICoach           = "AI_COACH"
	ModePersonaSimulation = "PERSONA_SIMULATION"
)

type AnalyzeRequest struct {
	Caption      string                `form:"caption"`
	AnalysisType string                `form:"analysis_type" binding:"required"`
	Persona      string                `form:"persona"`
	Media        *multipart.FileHeader `form:"media" binding:"required"`
}

type AnalyzeResponse struct {
	Mode            string   `json:"mode"`
	CoachFeedback   []string `json:"coach_feedback,omitempty"`
	PersonaReaction string   `json:"persona_reaction,omitempty"`
	MediaSummary    string   `json:"media_summary"`
}
