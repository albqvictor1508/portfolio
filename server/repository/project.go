package repository

import (
	"context"
	"errors"
	"time"

	"github.com/albqvictor1508/portfolio/entity"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	Conn *pgxpool.Pool
}

func (r *Repository) Insert(project *entity.Project) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var id int
	err := r.Conn.QueryRow(ctx,
		"INSERT INTO projects (name, description, github_url, demo_url, is_pinned) VALUES($1, $2, $3, $4) RETURNING id",
		project.Name,
		project.Description,
		project.GithubURL,
		project.DemoURL,
		project.IsPinned,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repository) FindByName(name string) (entity.Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	project := entity.Project{Name: name}
	err := r.Conn.QueryRow(
		ctx,
		"SELECT p.id, p.name, p.description, p.github_url, p.demo_url, p.is_pinned FROM projects p WHERE p.name = $1",
	).Scan(&project.ID, &project.Name, &project.Description, &project.GithubURL, &project.DemoURL, &project.IsPinned)
	if err != nil {
		return entity.Project{}, err
	}

	return project, nil
}

func (r *Repository) Delete(ctx context.Context, id uuid.UUID) error {
	query, err := r.Conn.Exec(ctx,
		"DELETE FROM projects WHERE id = $1",
		id,
	)

	if query.RowsAffected() == 0 {
		return errors.New("THIS PROJECT NOT EXISTS")
	}
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Update(ctx context.Context, project entity.Project) error {
	_, err := r.Conn.Exec(ctx,
		"UPDATE projects SET github_url = $1, demo_url = $2, is_pinned = $3, updated_at = NOW() WHERE id = $4",
		project.GithubURL,
		project.DemoURL,
		project.IsPinned,
		project.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) FindByID(ctx context.Context, id uuid.UUID) (entity.Project, error) {
	project := entity.Project{ID: id}
	err := r.Conn.QueryRow(
		ctx,
		"SELECT p.github_url, p.demo_url, p.is_pinned FROM projects p WHERE p.id = $1",
		id,
	).Scan(&project.GithubURL, &project.DemoURL, &project.IsPinned)
	if err != nil {
		return entity.Project{}, nil
	}

	return project, nil
}
