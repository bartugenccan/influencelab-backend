package service

import (
	"errors"

	"influencelab-backend/internal/ai"
	"influencelab-backend/internal/model"
)

func AnalyzeContent(input model.AnalyzeRequest) (model.AnalyzeResponse, error) {
	if input.AnalysisType != model.AnalysisTypeCoach && input.AnalysisType != model.AnalysisTypePersona {
		return model.AnalyzeResponse{}, errors.New("invalid analysis_type")
	}

	// coach mode
	if input.AnalysisType == model.AnalysisTypeCoach {
		analysis, err := ai.AnalyzeWithGemini(
			input.Caption,
			input.Media,
		)

		if err != nil {
			return model.AnalyzeResponse{
				Mode:  model.ModeAICoach,
				Error: err.Error(),
			}, nil
		}

		return model.AnalyzeResponse{
			Mode:          model.ModeAICoach,
			CoachAnalysis: analysis,
		}, nil
	}

	// persona mode (TODO: implement persona simulation)
	return model.AnalyzeResponse{
		Mode:            model.ModePersonaSimulation,
		PersonaReaction: "This content feels neutral and not relatable for this audience",
	}, nil
}
