package repository

import (
	"database/sql"

	"github.com/albqvictor1508/server/cmd"
)

type ProjectRepository interface {
	CreateProject(project *api.Project) (string, error)
	GetProjectByID(id string) (*api.Project, error)
	GetProjects() ([]*api.Project, error)
	UpdateProject(project *api.Project) error
	DeleteProject(id string) error
}

type projectRepository struct {
	db *sql.DB
}

func NewProjectRepository(db *sql.DB) ProjectRepository {
	return &projectRepository{db: db}
}

func (r *projectRepository) CreateProject(p *api.Project) (string, error) {
	query := `
		INSERT INTO projects (category_id, name, description, github_url, demo_url)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	var projectID string
	err := r.db.QueryRow(query, p.CategoryID, p.Name, p.Description, p.GithubURL, p.DemoURL).Scan(&projectID)

	if err != nil {
		return "", err
	}
	return projectID, nil
}

func (r *projectRepository) GetProjectByID(id string) (*api.Project, error) {
	query := `
		SELECT id, category_id, name, description, github_url, demo_url, created_at, updated_at
		FROM projects
		WHERE id = $1
	`
	var p api.Project
	err := r.db.QueryRow(query, id).Scan(&p.ID, &p.CategoryID, &p.Name, &p.Description, &p.GithubURL, &p.DemoURL, &p.CreatedAt, &p.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

func (r *projectRepository) GetProjects() ([]*api.Project, error) {
	query := `
		SELECT id, category_id, name, description, github_url, demo_url, created_at, updated_at
		FROM projects
		ORDER BY created_at DESC
	`
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var projects []*api.Project
	for rows.Next() {
		var p api.Project
		if err := rows.Scan(&p.ID, &p.CategoryID, &p.Name, &p.Description, &p.GithubURL, &p.DemoURL, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		projects = append(projects, &p)
	}

	return projects, nil
}

func (r *projectRepository) UpdateProject(p *api.Project) error {
	query := `
		UPDATE projects
		SET category_id = $1, name = $2, description = $3, github_url = $4, demo_url = $5, updated_at = NOW()
		WHERE id = $6
	`
	result, err := r.db.Exec(query, p.CategoryID, p.Name, p.Description, p.GithubURL, p.DemoURL, p.ID)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *projectRepository) DeleteProject(id string) error {
	query := "DELETE FROM projects WHERE id = $1"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
