package repository

import (
	"context"
	"time"

	"github.com/albqvictor1508/portfolio/entity"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct{}

type RepositoryPg struct {
	Conn *pgxpool.Pool
}

func (r *RepositoryPg) Insert(project entity.Project) (entity.Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := r.Conn.QueryRow(ctx,
		"INSERT INTO projects (github_url, demo_url, is_pinned, name) VALUES($1, $2, $3, $4) RETURNING id",
		project.GithubURL,
		project.DemoURL,
		project.IsPinned,
		project.Name,
	).Scan(&project.ID)
	if err != nil {
		return entity.Project{}, err
	}

	return project, nil
}

func (r *RepositoryPg) Delete(ctx context.Context, id uuid.UUID) error {
	tx, err := r.Conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx,
		"DELETE FROM projects WHERE id = $1",
		id,
	)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (r *RepositoryPg) Update(ctx context.Context, project entity.Project) error {
	tx, err := r.Conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx,
		"UPDATE projects SET github_url = $1, demo_url = $2, is_pinned = $3, updated_at = NOW() WHERE id = $4",
		project.GithubURL,
		project.DemoURL,
		project.IsPinned,
		project.ID,
	)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (r *RepositoryPg) FindByID(ctx context.Context, id uuid.UUID) (entity.Project, error) {
	project := entity.Project{ID: id}
	err := r.Conn.QueryRow(
		ctx,
		"SELECT p.github_url, p.demo_url, p.is_pinned FROM projects p WHERE p.id = $1",
		id,
	).Scan(&project.GithubURL, &project.DemoURL, &project.IsPinned)

	if err == pgx.ErrNoRows {
		return entity.Project{}, nil
	}

	if err != nil {
		return entity.Project{}, nil
	}

	return project, nil
}
