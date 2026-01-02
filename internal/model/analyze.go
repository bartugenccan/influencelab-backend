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

// CoachAnalysis represents the structured feedback from AI coach
type CoachAnalysis struct {
	OverallScore   int      `json:"overall_score"`
	VisualScore    int      `json:"visual_score"`
	CaptionScore   int      `json:"caption_score"`
	AlignmentScore int      `json:"alignment_score"`
	Strengths      []string `json:"strengths"`
	Improvements   []string `json:"improvements"`
	RevisedCaption string   `json:"revised_caption"`
	QuickWins      []string `json:"quick_wins"`
}

type AnalyzeResponse struct {
	Mode            string         `json:"mode"`
	CoachAnalysis   *CoachAnalysis `json:"coach_analysis,omitempty"`
	PersonaReaction string         `json:"persona_reaction,omitempty"`
	Error           string         `json:"error,omitempty"`
}
