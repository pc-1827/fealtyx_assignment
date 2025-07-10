package utils

import (
	"encoding/json"
	"net/http"
)

func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func ErrorResponse(w http.ResponseWriter, status int, message string) {
	http.Error(w, message, status)
}

func SuccessResponse(w http.ResponseWriter, data interface{}) {
	JSONResponse(w, http.StatusOK, data)
}
