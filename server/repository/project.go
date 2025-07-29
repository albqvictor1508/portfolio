package repository

import (
	"context"
	"errors"
	"time"

	"github.com/albqvictor1508/portfolio/entity"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ProjectRepository struct {
	Conn *pgxpool.Pool
}

func NewProjectRepo(connection *pgxpool.Pool) ProjectRepository {
	return ProjectRepository{
		Conn: connection,
	}
}

func (r *ProjectRepository) Insert(project *entity.Project) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var id int
	err := r.Conn.QueryRow(ctx,
		"INSERT INTO projects (name, description, github_url, demo_url, is_pinned) VALUES($1, $2, $3, $4 ,$5) RETURNING id",
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

func (pr *ProjectRepository) FindByName(name string) (entity.Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	project := entity.Project{Name: name}
	err := pr.Conn.QueryRow(
		ctx,
		"SELECT p.id, p.name, p.description, p.github_url, p.demo_url, p.is_pinned FROM projects p WHERE p.name = $1",
		name,
	).Scan(&project.ID, &project.Name, &project.Description, &project.GithubURL, &project.DemoURL, &project.IsPinned)
	if err == pgx.ErrNoRows {
		return entity.Project{}, nil
	}

	if err != nil {
		return entity.Project{}, err
	}

	return project, nil
}

func (pr *ProjectRepository) Delete(ctx context.Context, id int) error {
	query, err := pr.Conn.Exec(ctx,
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

func (pr *ProjectRepository) Update(ctx context.Context, project entity.Project) error {
	_, err := pr.Conn.Exec(ctx,
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

func (pr *ProjectRepository) FindByID(ctx context.Context, id int) (entity.Project, error) {
	project := entity.Project{ID: id}
	err := pr.Conn.QueryRow(
		ctx,
		"SELECT p.github_url, p.demo_url, p.is_pinned FROM projects p WHERE p.id = $1",
		id,
	).Scan(&project.GithubURL, &project.DemoURL, &project.IsPinned)
	if err == pgx.ErrNoRows {
		return entity.Project{}, nil
	}

	if err != nil {
		return entity.Project{}, err
	}

	return project, nil
}

func (pr *ProjectRepository) GetProjects() ([]entity.Project, error) {
	var projectList []entity.Project
	var projectObj entity.Project

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rows, err := pr.Conn.Query(
		ctx,
		"SELECT * FROM projects p",
	)
	if err != nil {
		return []entity.Project{}, err
	}

	for rows.Next() {
		err := rows.Scan(
			&projectObj.ID,
			&projectObj.Name,
			&projectObj.Description,
			&projectObj.GithubURL,
			&projectObj.DemoURL,
			&projectObj.IsPinned,
			&projectObj.CategoryID,
			&projectObj.CreatedAt,
			&projectObj.UpdatedAt,
		)
		if err != nil {
			return []entity.Project{}, err
		}

		projectList = append(projectList, projectObj)
	}

	rows.Close()
	return projectList, nil
}
