package models

import (
	"errors"
	"regexp"
)

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func (s *Student) Validate() error {
	if s.Name == "" {
		return errors.New("name is required")
	}
	if s.Age <= 0 {
		return errors.New("age must be a positive integer")
	}
	if !isValidEmail(s.Email) {
		return errors.New("invalid email format")
	}
	return nil
}

func isValidEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
