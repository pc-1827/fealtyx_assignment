package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"student-api/internal/services"
)

type OllamaHandler struct {
	OllamaService services.OllamaServiceInterface
}

func (h *OllamaHandler) GenerateSummary(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/students/")
	idStr := strings.TrimSuffix(path, "/summary")

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

	summary, err := h.OllamaService.GenerateSummary(student)
	if err != nil {
		http.Error(w, "Failed to generate summary", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"summary": summary,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
