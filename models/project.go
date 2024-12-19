package models

type Project struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	TechStack   []string `json:"tech_stack"`
	GitHubURL   string   `json:"github_url"`
}
