package api

import (
	"time"
)

// Project representa a estrutura de um projeto no banco de dados.
// As tags `json:"..."` definem como cada campo será representado no JSON da API.
type Project struct {
	ID          int64     `json:"id"`
	CategoryID  int64     `json:"categoryId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	GithubURL   string    `json:"githubUrl"`
	DemoURL     string    `json:"demoUrl"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// Technology representa uma tecnologia que pode ser associada a um projeto.
type Technology struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// ProjectTechnology é a tabela de associação entre um Projeto e uma Tecnologia.
type ProjectTechnology struct {
	ID           int64     `json:"id"`
	ProjectID    int64     `json:"projectId"`
	TechnologyID int64     `json:"technologyId"` // Adicionado para relacionar com a tecnologia
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// Category representa a categoria de um projeto.
type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// Error representa uma estrutura padrão para erros da API.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

