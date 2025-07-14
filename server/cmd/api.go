package api

import (
	"time"
)

type Project struct {
	ID          string    `json:"id"`
	CategoryID  int64     `json:"categoryId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	GithubURL   string    `json:"githubUrl"`
	DemoURL     string    `json:"demoUrl"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Technology struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ProjectTechnologies struct {
	ID           int64     `json:"id"`
	ProjectID    int64     `json:"projectId"`
	TechnologyID int64     `json:"technologyId"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

