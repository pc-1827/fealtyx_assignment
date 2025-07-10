package models

import (
	"testing"
)

func TestStudentValidation(t *testing.T) {
	tests := []struct {
		name        string
		student     Student
		expectError bool
		errorMsg    string
	}{
		{
			name: "Valid student",
			student: Student{
				Name:  "John Doe",
				Age:   20,
				Email: "john@example.com",
			},
			expectError: false,
		},
		{
			name: "Empty name",
			student: Student{
				Name:  "",
				Age:   20,
				Email: "john@example.com",
			},
			expectError: true,
			errorMsg:    "name is required",
		},
		{
			name: "Invalid age - zero",
			student: Student{
				Name:  "John Doe",
				Age:   0,
				Email: "john@example.com",
			},
			expectError: true,
			errorMsg:    "age must be a positive integer",
		},
		{
			name: "Invalid age - negative",
			student: Student{
				Name:  "John Doe",
				Age:   -5,
				Email: "john@example.com",
			},
			expectError: true,
			errorMsg:    "age must be a positive integer",
		},
		{
			name: "Invalid email format",
			student: Student{
				Name:  "John Doe",
				Age:   20,
				Email: "invalid-email",
			},
			expectError: true,
			errorMsg:    "invalid email format",
		},
		{
			name: "Empty email",
			student: Student{
				Name:  "John Doe",
				Age:   20,
				Email: "",
			},
			expectError: true,
			errorMsg:    "invalid email format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.student.Validate()

			if tt.expectError {
				if err == nil {
					t.Error("Expected validation error, got none")
				} else if err.Error() != tt.errorMsg {
					t.Errorf("Expected error message '%s', got '%s'", tt.errorMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Expected no validation error, got %v", err)
				}
			}
		})
	}
}

func TestEmailValidation(t *testing.T) {
	tests := []struct {
		email string
		valid bool
	}{
		{"test@example.com", true},
		{"user.name@domain.co.uk", true},
		{"user+tag@example.org", true},
		{"invalid-email", false},
		{"@example.com", false},
		{"user@", false},
		{"user@.com", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.email, func(t *testing.T) {
			result := isValidEmail(tt.email)
			if result != tt.valid {
				t.Errorf("Expected %v for email %s, got %v", tt.valid, tt.email, result)
			}
		})
	}
}
