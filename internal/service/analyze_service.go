package service

import (
	"errors"
	"fmt"
	"mime/multipart"

	"influencelab-backend/internal/model"
)

type AnalyzeInput struct {
	Caption      string
	AnalysisType string
	Persona      string
	MediaFile    *multipart.FileHeader
}

func AnalyzeContent(input AnalyzeInput) (model.AnalyzeResponse, error) {
	if input.AnalysisType != model.AnalysisTypeCoach && input.AnalysisType != model.AnalysisTypePersona {
		return model.AnalyzeResponse{}, errors.New("invalid analysis_type")
	}

	mediaInfo := fmt.Sprintf(
		"Received media: %s (%d bytes)",
		input.MediaFile.Filename,
		input.MediaFile.Size,
	)

	// coach mode
	if input.AnalysisType == model.AnalysisTypeCoach {
		return model.AnalyzeResponse{
			Mode: model.ModeAICoach,
			CoachFeedback: []string{
				"The visual does not immediately communicate value",
				"The caption hook is weak in the first line",
				"Consider adding emotional context",
			},
			MediaSummary: mediaInfo,
		}, nil
	}

	// persona mode
	return model.AnalyzeResponse{
		Mode:            model.ModePersonaSimulation,
		PersonaReaction: "This content feels neutral and not relatable for this audience",
		MediaSummary:    mediaInfo,
	}, nil
}
