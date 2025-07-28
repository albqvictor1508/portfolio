package internal

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID        uuid.UUID `json:"id"`
	GithubURL string    `json:"github_url"`
	DemoURL   string    `json:"demo_url"`
	IsPinned  bool      `json:"is_pinned"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateProjectParams struct {
	ID        string
	GithubURL string
	DemoURL   string
	IsPinned  bool
}
