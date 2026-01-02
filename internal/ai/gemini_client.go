package ai

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"regexp"

	"influencelab-backend/internal/config"
	"influencelab-backend/internal/model"
)

func AnalyzeWithGemini(
	caption string,
	image *multipart.FileHeader,
) (*model.CoachAnalysis, error) {

	apiKey := config.AppConfig.GeminiAPIKey
	if apiKey == "" {
		return nil, fmt.Errorf("GEMINI_API_KEY not set")
	}

	file, err := image.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open image: %w", err)
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(file)

	encodedImage := base64.StdEncoding.EncodeToString(buf.Bytes())

	// Determine mime type
	mimeType := image.Header.Get("Content-Type")
	if mimeType == "" {
		mimeType = "image/jpeg" // default fallback
	}

	prompt := fmt.Sprintf(`You are an expert Social Media Content Coach specializing in Instagram, TikTok, and YouTube content optimization.

## Your Task
Analyze the provided image and caption as a complete content piece. Provide strategic, actionable feedback that will help maximize engagement and reach.

## Content to Analyze
**Caption:** "%s"
**Visual:** [Attached Image]

## Analysis Framework
Evaluate the content across these dimensions:

### 1. Visual Impact (Score 1-10)
- First impression and scroll-stopping power
- Color harmony and composition
- Brand consistency and aesthetic appeal
- Image quality and clarity

### 2. Caption Effectiveness (Score 1-10)
- Hook strength (first line)
- Storytelling and emotional connection
- Call-to-action clarity
- Hashtag strategy (if applicable)

### 3. Content-Caption Alignment (Score 1-10)
- How well the visual and caption work together
- Message consistency
- Target audience clarity

## Required Output Format
Respond ONLY with valid JSON (no markdown, no code blocks, no extra text):
{
  "overall_score": <number 1-10>,
  "visual_score": <number 1-10>,
  "caption_score": <number 1-10>,
  "alignment_score": <number 1-10>,
  "strengths": ["<strength 1>", "<strength 2>"],
  "improvements": ["<specific improvement 1>", "<specific improvement 2>", "<specific improvement 3>"],
  "revised_caption": "<Your improved version of the caption>",
  "quick_wins": ["<easy fix 1>", "<easy fix 2>"]
}

Be direct, specific, and constructive. Focus on actionable improvements, not generic advice.`, caption)

	payload := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]interface{}{
					{
						"inlineData": map[string]string{
							"mimeType": mimeType,
							"data":     encodedImage,
						},
					},
					{
						"text": prompt,
					},
				},
			},
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Use gemini-2.0-flash (current model that supports vision)
	url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent?key=" + apiKey

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Read the full response body for debugging
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Check for HTTP errors
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Gemini API error (status %d): %s", resp.StatusCode, string(respBody))
	}

	var result model.GeminiResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	if len(result.Candidates) == 0 {
		return nil, fmt.Errorf("no response from Gemini. Raw response: %s", string(respBody))
	}

	if len(result.Candidates[0].Content.Parts) == 0 {
		return nil, fmt.Errorf("empty response parts from Gemini")
	}

	// Extract the text response
	textResponse := result.Candidates[0].Content.Parts[0].Text

	// Parse the JSON from the response (handle markdown code blocks if present)
	jsonStr := extractJSON(textResponse)

	var analysis model.CoachAnalysis
	if err := json.Unmarshal([]byte(jsonStr), &analysis); err != nil {
		return nil, fmt.Errorf("failed to parse coach analysis: %w. Raw: %s", err, textResponse)
	}

	return &analysis, nil
}

// extractJSON extracts JSON from a string, handling markdown code blocks
func extractJSON(s string) string {
	// Try to find JSON in markdown code block
	re := regexp.MustCompile("(?s)```(?:json)?\\s*({.*?})\\s*```")
	matches := re.FindStringSubmatch(s)
	if len(matches) > 1 {
		return matches[1]
	}

	// Try to find raw JSON object
	re = regexp.MustCompile(`(?s)\{.*\}`)
	match := re.FindString(s)
	if match != "" {
		return match
	}

	return s
}
