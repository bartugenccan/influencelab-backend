package service

import (
	"errors"
	"fmt"

	"influencelab-backend/internal/ai"
	"influencelab-backend/internal/model"
)

func AnalyzeContent(input model.AnalyzeRequest) (model.AnalyzeResponse, error) {
	if input.AnalysisType != model.AnalysisTypeCoach && input.AnalysisType != model.AnalysisTypePersona {
		return model.AnalyzeResponse{}, errors.New("invalid analysis_type")
	}

	mediaInfo := fmt.Sprintf(
		"Received media: %s (%d bytes)",
		input.Media.Filename,
		input.Media.Size,
	)

	// coach mode
	if input.AnalysisType == model.AnalysisTypeCoach {

		feedback, err := ai.AnalyzeWithGemini(
			input.Caption,
			input.Media,
		)

		if err != nil {
			return model.AnalyzeResponse{}, err
		}

		return model.AnalyzeResponse{
			Mode:          model.ModeAICoach,
			CoachFeedback: []string{feedback},
		}, nil
	}

	// persona mode
	return model.AnalyzeResponse{
		Mode:            model.ModePersonaSimulation,
		PersonaReaction: "This content feels neutral and not relatable for this audience",
		MediaSummary:    mediaInfo,
	}, nil
}
