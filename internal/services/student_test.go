package services

import (
	"student-api/internal/models"
	"sync"
	"testing"
)

func TestCreateStudent(t *testing.T) {
	ResetStudents()

	student := models.Student{
		Name:  "John Doe",
		Age:   20,
		Email: "john@example.com",
	}

	result := CreateStudent(student)

	if result.ID != 1 {
		t.Errorf("Expected ID 1, got %d", result.ID)
	}
	if result.Name != student.Name {
		t.Errorf("Expected name %s, got %s", student.Name, result.Name)
	}
}

func TestGetAllStudents(t *testing.T) {
	ResetStudents()

	students := GetAllStudents()
	if len(students) != 0 {
		t.Errorf("Expected empty list, got %d students", len(students))
	}

	CreateStudent(models.Student{Name: "Student1", Age: 20, Email: "s1@example.com"})
	CreateStudent(models.Student{Name: "Student2", Age: 21, Email: "s2@example.com"})

	students = GetAllStudents()
	if len(students) != 2 {
		t.Errorf("Expected 2 students, got %d", len(students))
	}
}

func TestGetStudentByID(t *testing.T) {
	ResetStudents()

	student := CreateStudent(models.Student{
		Name:  "Jane Doe",
		Age:   22,
		Email: "jane@example.com",
	})

	result, err := GetStudentByID(student.ID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.Name != student.Name {
		t.Errorf("Expected name %s, got %s", student.Name, result.Name)
	}

	_, err = GetStudentByID(999)
	if err == nil {
		t.Error("Expected error for non-existing student")
	}
}

func TestUpdateStudent(t *testing.T) {
	ResetStudents()

	student := CreateStudent(models.Student{
		Name:  "Original",
		Age:   20,
		Email: "original@example.com",
	})

	updatedData := models.Student{
		Name:  "Updated",
		Age:   25,
		Email: "updated@example.com",
	}

	result, err := UpdateStudent(student.ID, updatedData)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result.Name != updatedData.Name {
		t.Errorf("Expected name %s, got %s", updatedData.Name, result.Name)
	}
	if result.ID != student.ID {
		t.Errorf("Expected ID %d, got %d", student.ID, result.ID)
	}

	_, err = UpdateStudent(999, updatedData)
	if err == nil {
		t.Error("Expected error for non-existing student")
	}
}

func TestDeleteStudent(t *testing.T) {
	ResetStudents()

	student := CreateStudent(models.Student{
		Name:  "To Delete",
		Age:   20,
		Email: "delete@example.com",
	})

	err := DeleteStudent(student.ID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	_, err = GetStudentByID(student.ID)
	if err == nil {
		t.Error("Expected error after deletion")
	}

	err = DeleteStudent(999)
	if err == nil {
		t.Error("Expected error for non-existing student")
	}
}

func TestConcurrency(t *testing.T) {
	ResetStudents()

	var wg sync.WaitGroup
	numGoroutines := 10

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(index int) {
			defer wg.Done()
			CreateStudent(models.Student{
				Name:  "Student" + string(rune(index)),
				Age:   20 + index,
				Email: "student" + string(rune(index)) + "@example.com",
			})
		}(i)
	}
	wg.Wait()

	students := GetAllStudents()
	if len(students) != numGoroutines {
		t.Errorf("Expected %d students, got %d", numGoroutines, len(students))
	}

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			GetAllStudents()
		}()
	}
	wg.Wait()
}
