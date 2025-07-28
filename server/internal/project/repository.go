package project

import (
	"context"
	"time"

	"github.com/albqvictor1508/portfolio/internal"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	Conn *pgxpool.Conn
}

func (r *Repository) Insert(project internal.Project) (internal.Project, error) {
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
