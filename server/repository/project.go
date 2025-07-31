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

	tx, err := r.Conn.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	var id int
	err = tx.QueryRow(ctx,
		"INSERT INTO projects (name, description, github_url, demo_url, is_pinned, category_id) VALUES($1, $2, $3, $4 ,$5, $6) RETURNING id",
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

	if len(project.Technologies) > 0 {
		for _, tech := range project.Technologies {
			_, err := tx.Exec(ctx, "INSERT INTO project_technologies (project_id, technology_id) VALUES ($1, $2)", id, tech.ID)
			if err != nil {
				return 0, err
			}
		}
	}

	return id, tx.Commit(ctx)
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

	tx, err := pr.Conn.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx,
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

	_, err = tx.Exec(ctx, "DELETE FROM project_technologies WHERE project_id = $1", project.ID)
	if err != nil {
		return 0, err
	}

	if len(project.Technologies) > 0 {
		for _, tech := range project.Technologies {
			_, err := tx.Exec(ctx, "INSERT INTO project_technologies (project_id, technology_id) VALUES ($1, $2)", project.ID, tech.ID)
			if err != nil {
				return 0, err
			}
		}
	}

	return project.ID, tx.Commit(ctx)
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

	rows, err := pr.Conn.Query(ctx, "SELECT t.id, t.name, t.photo_url FROM technologies t JOIN project_technologies pt ON t.id = pt.technology_id WHERE pt.project_id = $1", id)
	if err != nil {
		return entity.Project{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var tech entity.Technology
		if err := rows.Scan(&tech.ID, &tech.Name, &tech.PhotoURL); err != nil {
			return entity.Project{}, err
		}
		project.Technologies = append(project.Technologies, tech)
	}

	return project, nil
}

func (pr *ProjectRepository) GetProjects() ([]entity.Project, error) {
	var projectList []entity.Project

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rows, err := pr.Conn.Query(
		ctx,
		"SELECT p.id, p.name, p.description, p.github_url, p.demo_url, p.is_pinned, p.category_id, p.created_at, p.updated_at"+
			"FROM projects p"+
			"INNER JOIN project_technologies pt ON pt.project_id = p.id",
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

		techRows, err := pr.Conn.Query(ctx, "SELECT t.id, t.name, t.photo_url FROM technologies t JOIN project_technologies pt ON t.id = pt.technology_id WHERE pt.project_id = $1", projectObj.ID)
		if err != nil {
			return []entity.Project{}, err
		}
		defer techRows.Close()

		for techRows.Next() {
			var tech entity.Technology
			if err := techRows.Scan(&tech.ID, &tech.Name, &tech.PhotoURL); err != nil {
				return []entity.Project{}, err
			}
			projectObj.Technologies = append(projectObj.Technologies, tech)
		}

		projectList = append(projectList, projectObj)
	}

	return projectList, nil
}
