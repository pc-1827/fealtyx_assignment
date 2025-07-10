package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"student-api/internal/models"
)

type OllamaServiceInterface interface {
	GenerateSummary(student models.Student) (string, error)
}

type OllamaService struct {
	BaseURL string
}

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type OllamaResponse struct {
	Response string `json:"response"`
}

func NewOllamaService() *OllamaService {
	return &OllamaService{
		BaseURL: "http://localhost:11434",
	}
}

func (s *OllamaService) GenerateSummary(student models.Student) (string, error) {
	prompt := fmt.Sprintf(`Generate a professional summary for this student profile:
Name: %s
Age: %d
Email: %s
ID: %d

Please provide a brief, professional summary of this student in 2-3 sentences. Return only the summary without any prefixes or headers.`,
		student.Name, student.Age, student.Email, student.ID)

	reqBody := OllamaRequest{
		Model:  "llama3",
		Prompt: prompt,
		Stream: false,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(s.BaseURL+"/api/generate", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var ollamaResp OllamaResponse
	if err := json.NewDecoder(resp.Body).Decode(&ollamaResp); err != nil {
		return "", err
	}

	cleanedResponse := cleanSummaryResponse(ollamaResp.Response)

	return cleanedResponse, nil
}

func cleanSummaryResponse(response string) string {
	cleaned := strings.ReplaceAll(response, "\\n", " ")
	cleaned = strings.ReplaceAll(cleaned, "\\t", " ")
	cleaned = strings.ReplaceAll(cleaned, "\\r", " ")
	cleaned = strings.ReplaceAll(cleaned, "\n", " ")
	cleaned = strings.ReplaceAll(cleaned, "\t", " ")
	cleaned = strings.ReplaceAll(cleaned, "\r", " ")

	prefixes := []string{
		"Here is a professional summary for the student profile:",
		"Here's a professional summary for the student profile:",
		"Professional summary:",
		"Summary:",
		"Here is a brief summary:",
		"Here's a brief summary:",
	}

	for _, prefix := range prefixes {
		if strings.HasPrefix(strings.TrimSpace(cleaned), prefix) {
			cleaned = strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(cleaned), prefix))
			break
		}
	}

	cleaned = strings.Join(strings.Fields(cleaned), " ")
	cleaned = strings.TrimSpace(cleaned)

	return cleaned
}
