package data

import "influencelab-backend/internal/model"

var FreePersonaTemplates = []model.PersonaTemplate{
	{
		ID:           "fitness_wellness",
		Title:        "Fitness & Wellness",
		Description:  "People focused on fitness, health and routines",
		AgeRange:     "20-40",
		PersonaCount: 10,
		Constraints: map[string][]string{
			"experience_level": {"beginner", "intermediate", "advanced"},
			"lifestyle":        {"gym", "home workout", "outdoor"},
		},
	},
	{
		ID:           "stay_at_home_parents",
		Title:        "Stay-at-home Parents",
		Description:  "Parents managing home, children and daily routines",
		AgeRange:     "25-45",
		PersonaCount: 10,
		Constraints: map[string][]string{
			"focus": {"children", "time-saving", "budget"},
		},
	},
}
