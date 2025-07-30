package repository

import (
	"context"
	"errors"
	"time"

	"github.com/albqvictor1508/portfolio/entity"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ExperienceRepository struct {
	Conn *pgxpool.Pool
}

func (er *ExperienceRepository) NewExperience(conn *pgxpool.Pool) ExperienceRepository {
	return ExperienceRepository{
		Conn: conn,
	}
}

func (er *ExperienceRepository) Insert(project *entity.Project) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var id int
	err := er.Conn.QueryRow(ctx,
		"INSERT INTO experiences (name, description, github_url, demo_url, is_pinned, category_id) VALUES($1, $2, $3, $4 ,$5, $6) RETURNING id",
		project.Name,
		project.Description,
		project.GithubURL,
		project.DemoURL,
		project.IsPinned,
		project.CategoryID,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (pr *ProjectRepository) FindByName(name string) (entity.Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	project := entity.Project{}
	err := pr.Conn.QueryRow(
		ctx,
		"SELECT id, name, description, github_url, demo_url, is_pinned, category_id, created_at, updated_at FROM projects WHERE name = $1",
		name,
	).Scan(
		&project.ID,
		&project.Name,
		&project.Description,
		&project.GithubURL,
		&project.DemoURL,
		&project.IsPinned,
		&project.CategoryID,
		&project.CreatedAt,
		&project.UpdatedAt,
	)
	if err == pgx.ErrNoRows {
		return entity.Project{}, nil
	}

	if err != nil {
		return entity.Project{}, err
	}

	return project, nil
}

func (pr *ProjectRepository) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query, err := pr.Conn.Exec(
		ctx,
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

func (pr *ProjectRepository) Update(project *entity.Project) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := pr.Conn.Exec(ctx,
		"UPDATE projects SET name = $1, description = $2, github_url = $3, demo_url = $4, is_pinned = $5, category_id = $6, updated_at = NOW() WHERE id = $7",
		project.Name,
		project.Description,
		project.GithubURL,
		project.DemoURL,
		project.IsPinned,
		project.CategoryID,
		project.ID,
	)
	if err != nil {
		return 0, err
	}

	return project.ID, nil
}

func (pr *ProjectRepository) FindByID(id int) (entity.Project, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	project := entity.Project{}
	err := pr.Conn.QueryRow(
		ctx,
		"SELECT id, name, description, github_url, demo_url, is_pinned, category_id, created_at, updated_at FROM projects WHERE id = $1",
		id,
	).Scan(
		&project.ID,
		&project.Name,
		&project.Description,
		&project.GithubURL,
		&project.DemoURL,
		&project.IsPinned,
		&project.CategoryID,
		&project.CreatedAt,
		&project.UpdatedAt,
	)
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rows, err := pr.Conn.Query(
		ctx,
		"SELECT id, name, description, github_url, demo_url, is_pinned, category_id, created_at, updated_at FROM projects",
	)
	if err != nil {
		return []entity.Project{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var projectObj entity.Project
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

	return projectList, nil
}
