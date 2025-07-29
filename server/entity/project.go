package entity

import (
	"time"
)

type Project struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	GithubURL   string    `json:"github_url"`
	DemoURL     string    `json:"demo_url"`
	IsPinned    bool      `json:"is_pinned"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type InsertProjectParams struct {
	Name      string `json:"name"`
	GithubURL string `json:"github_url"`
	DemoURL   string `json:"demo_url"`
	IsPinned  bool   `json:"is_pinned"`
}

type UpdateProjectParams struct {
	ID        string
	Name      string
	GithubURL string
	DemoURL   string
	IsPinned  bool
}
