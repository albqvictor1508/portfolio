package entity

import (
	"time"
)

type Project struct {
	ID           int          `json:"id"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	GithubURL    string       `json:"github_url"`
	DemoURL      string       `json:"demo_url"`
	PhotoURL     string       `json:"photo_url"`
	IsPinned     bool         `json:"is_pinned"`
	Category     *Category    `json:"category,omitempty"`
	Technologies []Technology `json:"technologies"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
}
