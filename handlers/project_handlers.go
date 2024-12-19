package handlers

import (
	"encoding/json"
	"net/http"
	"project-api/database"
	"project-api/models"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

// GetProjects retrieves all projects from the database
func GetProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := database.DB.Query("SELECT id, name, description, technologies, github FROM projects")
	if err != nil {
		http.Error(w, `{"error":"Failed to fetch projects"}`, http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	projects := []models.Project{}
	for rows.Next() {
		var project models.Project
		// Note the use of pq.Array(&project.TechStack)
		if err := rows.Scan(&project.ID, &project.Name, &project.Description, pq.Array(&project.TechStack), &project.GitHubURL); err != nil {
			http.Error(w, `{"error":"Failed to scan project"}`, http.StatusInternalServerError)
			return
		}
		projects = append(projects, project)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, `{"error":"Failed to read projects"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(projects)
}

// GetProject retrieves a single project by ID from the database
func GetProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, `{"error":"Project ID is required"}`, http.StatusBadRequest)
		return
	}

	var project models.Project
	query := "SELECT id, name, description, technologies, github FROM projects WHERE id = $1"
	err := database.DB.QueryRow(query, id).Scan(&project.ID, &project.Name, &project.Description, pq.Array(&project.TechStack), &project.GitHubURL)
	if err != nil {
		http.Error(w, `{"error":"Project not found"}`, http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(project)
}

// CreateProject adds a new project to the database
func CreateProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var project models.Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		http.Error(w, `{"error":"Invalid input"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	project.ID = uuid.New().String()
	query := `INSERT INTO projects (id, name, description, technologies, github) VALUES ($1, $2, $3, $4, $5)`
	_, err := database.DB.Exec(query,
		project.ID,
		project.Name,
		project.Description,
		pq.Array(project.TechStack),
		project.GitHubURL,
	)
	if err != nil {
		http.Error(w, `{"error":"Failed to create project"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(project)
}

// UpdateProject updates an existing project in the database
func UpdateProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var project models.Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Project ID is required", http.StatusBadRequest)
		return
	}

	query := `UPDATE projects SET name = $1, description = $2, technologies = $3, github = $4 WHERE id = $5`
	_, err := database.DB.Exec(query, project.Name, project.Description, pq.Array(project.TechStack), project.GitHubURL, id)
	if err != nil {
		http.Error(w, "Failed to update project", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Project updated successfully"})
}

// DeleteProject removes a project from the database
func DeleteProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, `{"error":"Project ID is required"}`, http.StatusBadRequest)
		return
	}

	query := `DELETE FROM projects WHERE id = $1`
	_, err := database.DB.Exec(query, id)
	if err != nil {
		http.Error(w, `{"error":"Failed to delete project"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Project deleted successfully"})
}
