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
	CategoryID  *int      `json:"category_id,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
