package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"student-api/internal/models"
	"student-api/internal/services"
	"testing"
)

func setupTest() {
	services.ResetStudents()
}

func TestCreateStudent(t *testing.T) {
	setupTest()

	tests := []struct {
		name           string
		requestBody    interface{}
		expectedStatus int
		expectedName   string
	}{
		{
			name: "Valid student creation",
			requestBody: models.Student{
				Name:  "John Doe",
				Age:   20,
				Email: "john@example.com",
			},
			expectedStatus: http.StatusCreated,
			expectedName:   "John Doe",
		},
		{
			name: "Invalid student - missing name",
			requestBody: models.Student{
				Age:   20,
				Email: "john@example.com",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Invalid student - missing email",
			requestBody: models.Student{
				Name: "John Doe",
				Age:  20,
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Invalid student - invalid age",
			requestBody: models.Student{
				Name:  "John Doe",
				Age:   -5,
				Email: "john@example.com",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Invalid JSON",
			requestBody:    "invalid json",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonData, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest("POST", "/students", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			CreateStudent(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedStatus == http.StatusCreated {
				var student models.Student
				json.Unmarshal(rr.Body.Bytes(), &student)
				if student.Name != tt.expectedName {
					t.Errorf("Expected name %s, got %s", tt.expectedName, student.Name)
				}
				if student.ID != 1 {
					t.Errorf("Expected ID 1, got %d", student.ID)
				}
			}
		})
	}
}

func TestGetAllStudents(t *testing.T) {
	setupTest()

	req := httptest.NewRequest("GET", "/students", nil)
	rr := httptest.NewRecorder()
	GetAllStudents(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}

	var students []models.Student
	json.Unmarshal(rr.Body.Bytes(), &students)
	if len(students) != 0 {
		t.Errorf("Expected empty list, got %d students", len(students))
	}

	services.CreateStudent(models.Student{Name: "Test", Age: 25, Email: "test@example.com"})

	req = httptest.NewRequest("GET", "/students", nil)
	rr = httptest.NewRecorder()
	GetAllStudents(rr, req)

	json.Unmarshal(rr.Body.Bytes(), &students)
	if len(students) != 1 {
		t.Errorf("Expected 1 student, got %d", len(students))
	}
}

func TestGetStudentByID(t *testing.T) {
	setupTest()

	_ = services.CreateStudent(models.Student{
		Name:  "Jane Doe",
		Age:   22,
		Email: "jane@example.com",
	})

	tests := []struct {
		name           string
		url            string
		expectedStatus int
		expectedID     int
	}{
		{
			name:           "Valid student ID",
			url:            "/students/1",
			expectedStatus: http.StatusOK,
			expectedID:     1,
		},
		{
			name:           "Non-existent student ID",
			url:            "/students/999",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "Invalid student ID format",
			url:            "/students/abc",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tt.url, nil)
			rr := httptest.NewRecorder()
			GetStudentByID(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedStatus == http.StatusOK {
				var returnedStudent models.Student
				json.Unmarshal(rr.Body.Bytes(), &returnedStudent)
				if returnedStudent.ID != tt.expectedID {
					t.Errorf("Expected ID %d, got %d", tt.expectedID, returnedStudent.ID)
				}
			}
		})
	}
}

func TestUpdateStudent(t *testing.T) {
	setupTest()

	services.CreateStudent(models.Student{
		Name:  "Original Name",
		Age:   20,
		Email: "original@example.com",
	})

	tests := []struct {
		name           string
		url            string
		requestBody    interface{}
		expectedStatus int
		expectedName   string
	}{
		{
			name: "Valid update",
			url:  "/students/1",
			requestBody: models.Student{
				Name:  "Updated Name",
				Age:   25,
				Email: "updated@example.com",
			},
			expectedStatus: http.StatusOK,
			expectedName:   "Updated Name",
		},
		{
			name: "Non-existent student",
			url:  "/students/999",
			requestBody: models.Student{
				Name:  "Test",
				Age:   25,
				Email: "test@example.com",
			},
			expectedStatus: http.StatusNotFound,
		},
		{
			name: "Invalid student data",
			url:  "/students/1",
			requestBody: models.Student{
				Name:  "",
				Age:   25,
				Email: "test@example.com",
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Invalid ID format",
			url:            "/students/abc",
			requestBody:    models.Student{Name: "Test", Age: 25, Email: "test@example.com"},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonData, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest("PUT", tt.url, bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			UpdateStudent(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			if tt.expectedStatus == http.StatusOK {
				var student models.Student
				json.Unmarshal(rr.Body.Bytes(), &student)
				if student.Name != tt.expectedName {
					t.Errorf("Expected name %s, got %s", tt.expectedName, student.Name)
				}
			}
		})
	}
}

func TestDeleteStudent(t *testing.T) {
	setupTest()

	services.CreateStudent(models.Student{
		Name:  "To Be Deleted",
		Age:   20,
		Email: "delete@example.com",
	})

	tests := []struct {
		name           string
		url            string
		expectedStatus int
	}{
		{
			name:           "Valid deletion",
			url:            "/students/1",
			expectedStatus: http.StatusNoContent,
		},
		{
			name:           "Non-existent student",
			url:            "/students/999",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "Invalid ID format",
			url:            "/students/abc",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("DELETE", tt.url, nil)
			rr := httptest.NewRecorder()
			DeleteStudent(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}
		})
	}
}
