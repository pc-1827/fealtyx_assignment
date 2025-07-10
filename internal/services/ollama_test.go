package services

import (
	"fmt"
	"strings"
	"student-api/internal/models"
	"testing"
)

func TestCleanSummaryResponse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Response with newlines",
			input:    "Here is a professional summary for the student profile:\n\nJohn is a great student.\nHe works hard.",
			expected: "John is a great student. He works hard.",
		},
		{
			name:     "Response with escape characters",
			input:    "Here is a professional summary for the student profile:\\n\\nJohn is motivated.\\tHe is 25 years old.",
			expected: "John is motivated. He is 25 years old.",
		},
		{
			name:     "Response with multiple spaces",
			input:    "John    is   a    motivated     student.",
			expected: "John is a motivated student.",
		},
		{
			name:     "Response with prefix and formatting",
			input:    "Here's a professional summary for the student profile:\n\n\tJohn Updated is ambitious and focused.",
			expected: "John Updated is ambitious and focused.",
		},
		{
			name:     "Clean response without issues",
			input:    "John is an excellent student with great potential.",
			expected: "John is an excellent student with great potential.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := cleanSummaryResponse(tt.input)
			if result != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

func TestOllamaServiceGenerateSummary(t *testing.T) {
	mockService := &MockOllamaService{}
	student := models.Student{
		ID:    1,
		Name:  "Test Student",
		Age:   20,
		Email: "test@example.com",
	}

	summary, err := mockService.GenerateSummary(student)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if strings.Contains(summary, "\\n") || strings.Contains(summary, "\n") {
		t.Errorf("Summary contains escape characters: %s", summary)
	}

	if strings.Contains(summary, "Here is") || strings.Contains(summary, "Here's") {
		t.Errorf("Summary contains unwanted prefix: %s", summary)
	}
}

type MockOllamaService struct{}

func (m *MockOllamaService) GenerateSummary(student models.Student) (string, error) {
	rawResponse := "Here is a professional summary for the student profile:\n\n" + student.Name + " is an excellent student with great potential. At " + fmt.Sprintf("%d", student.Age) + " years old, they demonstrate strong academic abilities."

	return cleanSummaryResponse(rawResponse), nil
}
