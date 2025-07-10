package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"student-api/internal/models"
	"student-api/internal/services"
)

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if student.Name == "" || student.Email == "" || student.Age <= 0 {
		http.Error(w, "Invalid student data", http.StatusBadRequest)
		return
	}

	createdStudent := services.CreateStudent(student)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdStudent)
}

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	students := services.GetAllStudents()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func GetStudentByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/students/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	student, err := services.GetStudentByID(id)
	if err != nil {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/students/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	var student models.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if student.Name == "" || student.Email == "" || student.Age <= 0 {
		http.Error(w, "Invalid student data", http.StatusBadRequest)
		return
	}

	updatedStudent, err := services.UpdateStudent(id, student)
	if err != nil {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedStudent)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/students/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	if err := services.DeleteStudent(id); err != nil {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
