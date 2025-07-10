package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"student-api/internal/models"
	"student-api/internal/services"
	"testing"
)

func TestOllamaHandler_GenerateSummary(t *testing.T) {
	setupTest()

	// Create a test student
	_ = services.CreateStudent(models.Student{
		Name:  "Alice Johnson",
		Age:   23,
		Email: "alice@example.com",
	})

	// Mock Ollama service for testing
	mockOllamaService := &MockOllamaService{}
	handler := &OllamaHandler{
		OllamaService: mockOllamaService,
	}

	tests := []struct {
		name           string
		url            string
		expectedStatus int
		mockError      bool
	}{
		{
			name:           "Valid summary generation",
			url:            "/students/1/summary",
			expectedStatus: http.StatusOK,
			mockError:      false,
		},
		{
			name:           "Non-existent student",
			url:            "/students/999/summary",
			expectedStatus: http.StatusNotFound,
			mockError:      false,
		},
		{
			name:           "Invalid ID format",
			url:            "/students/abc/summary",
			expectedStatus: http.StatusBadRequest,
			mockError:      false,
		},
		{
			name:           "Ollama service error",
			url:            "/students/1/summary",
			expectedStatus: http.StatusInternalServerError,
			mockError:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockOllamaService.ShouldError = tt.mockError

			req := httptest.NewRequest("GET", tt.url, nil)
			rr := httptest.NewRecorder()
			handler.GenerateSummary(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedStatus == http.StatusOK {
				var response map[string]interface{}
				json.Unmarshal(rr.Body.Bytes(), &response)
				if summary, exists := response["summary"]; !exists || summary == "" {
					t.Error("Expected summary in response")
				}
			}
		})
	}
}

type MockOllamaService struct {
	ShouldError bool
}

func (m *MockOllamaService) GenerateSummary(student models.Student) (string, error) {
	if m.ShouldError {
		return "", &mockError{message: "mock ollama error"}
	}
	return "Mock summary for " + student.Name, nil
}

type mockError struct {
	message string
}

func (e *mockError) Error() string {
	return e.message
}
