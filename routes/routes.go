package routes

import (
	"net/http"
	"project-api/handlers"
)

// Initialize sets up the HTTP routes for the application
func Initialize() *http.ServeMux {
	router := http.NewServeMux()

	// Define routes
	router.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetProjects(w, r)
		case http.MethodPost:
			handlers.CreateProject(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	router.HandleFunc("/project", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetProject(w, r)
		case http.MethodPut:
			handlers.UpdateProject(w, r)
		case http.MethodDelete:
			handlers.DeleteProject(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return router
}
