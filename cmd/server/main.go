package main

import (
	"log"
	"net/http"
	"strings"
	"student-api/internal/handlers"
	"student-api/internal/services"
)

func main() {
	ollamaService := services.NewOllamaService()
	ollamaHandler := &handlers.OllamaHandler{
		OllamaService: ollamaService,
	}

	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			handlers.GetAllStudents(w, r)
		case "POST":
			handlers.CreateStudent(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/students/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/summary") {
			if r.Method == "GET" {
				ollamaHandler.GenerateSummary(w, r)
			} else {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
			return
		}

		switch r.Method {
		case "GET":
			handlers.GetStudentByID(w, r)
		case "PUT":
			handlers.UpdateStudent(w, r)
		case "DELETE":
			handlers.DeleteStudent(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
