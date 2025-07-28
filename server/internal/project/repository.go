package project

import (
	"context"
	"time"

	"github.com/albqvictor1508/portfolio/internal"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type RepositoryPg struct {
	Conn *pgxpool.Conn
}

func (r *RepositoryPg) Insert(project internal.Project) (internal.Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := r.Conn.QueryRow(
		ctx,
		"INSERT INTO projects (github_url, demo_url, is_pinned) VALUES($1, $2, $3) RETURNING id, github_url, demo_url, is_pinned",
		project.GithubURL,
		project.DemoURL,
		project.IsPinned,
	).Scan(&project.ID, &project.GithubURL, &project.DemoURL, &project.IsPinned)
	if err != nil {
		return internal.Project{}, err
	}

	return project, nil
}

func (r *RepositoryPg) Delete(ctx context.Context, id uuid.UUID) error {
	tag, err := r.Conn.Exec(
		ctx,
		"DELETE FROM projects p WHERE p.id = $1",
		id,
	)

	if tag.RowsAffected() == 0 {
		return nil
	}

	return err
}

func (r *RepositoryPg) Update(ctx context.Context, project internal.Project) error {
	_, err := r.Conn.Exec(
		ctx,
		"UPDATE projects p SET github_url = $1, demo_url = $2, is_pinned = $3 WHERE p.id = $4",
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

func (r *RepositoryPg) FindByID(ctx context.Context, id uuid.UUID) (internal.Project, error) {
	project := internal.Project{ID: id}
	err := r.Conn.QueryRow(
		ctx,
		"SELECT p.github_url, p.demo_url, p.is_pinned FROM projects p WHERE p.id = $1",
		id,
	).Scan(&project.GithubURL, &project.DemoURL, &project.IsPinned)

	if err == pgx.ErrNoRows {
		return internal.Project{}, nil
	}

	if err != nil {
		return internal.Project{}, nil
	}

	return project, nil
}
