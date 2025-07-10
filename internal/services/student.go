package services

import (
	"errors"
	"student-api/internal/models"
	"sync"
)

var (
	students = make(map[int]models.Student)
	nextID   = 1
	mutex    = sync.RWMutex{}
)

func CreateStudent(student models.Student) models.Student {
	mutex.Lock()
	defer mutex.Unlock()

	student.ID = nextID
	nextID++
	students[student.ID] = student
	return student
}

func GetAllStudents() []models.Student {
	mutex.RLock()
	defer mutex.RUnlock()

	result := make([]models.Student, 0, len(students))
	for _, student := range students {
		result = append(result, student)
	}
	return result
}

func GetStudentByID(id int) (models.Student, error) {
	mutex.RLock()
	defer mutex.RUnlock()

	student, exists := students[id]
	if !exists {
		return models.Student{}, errors.New("student not found")
	}
	return student, nil
}

func UpdateStudent(id int, student models.Student) (models.Student, error) {
	mutex.Lock()
	defer mutex.Unlock()

	if _, exists := students[id]; !exists {
		return models.Student{}, errors.New("student not found")
	}

	student.ID = id
	students[id] = student
	return student, nil
}

func DeleteStudent(id int) error {
	mutex.Lock()
	defer mutex.Unlock()

	if _, exists := students[id]; !exists {
		return errors.New("student not found")
	}

	delete(students, id)
	return nil
}

func ResetStudents() {
	mutex.Lock()
	defer mutex.Unlock()
	students = make(map[int]models.Student)
	nextID = 1
}
