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

	projectQuery := `
		INSERT INTO projects
			(name, description, github_url, demo_url, is_pinned, category_id)
		VALUES
			($1, $2, $3, $4 ,$5, $6)
		RETURNING id
	`

	var id int
	err = tx.QueryRow(ctx,
		projectQuery,
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
		techQuery := `
			INSERT INTO project_technologies
				(project_id, technology_id)
			VALUES
				($1, $2)
		`
		for _, tech := range project.Technologies {
			_, err := tx.Exec(ctx, techQuery, id, tech.ID)
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

	query := `
		SELECT
			id, name, description, github_url, demo_url, is_pinned, category_id, created_at, updated_at
		FROM
			projects
		WHERE
			name = $1
	`

	project := entity.Project{}
	err := pr.Conn.QueryRow(
		ctx,
		query,
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

	query := `
		DELETE FROM
			projects
		WHERE
			id = $1
	`

	cmd, err := pr.Conn.Exec(
		ctx,
		query,
		id,
	)

	if cmd.RowsAffected() == 0 {
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

	updateQuery := `
		UPDATE
			projects
		SET
			name = $1, description = $2, github_url = $3, demo_url = $4, is_pinned = $5, category_id = $6, updated_at = NOW()
		WHERE
			id = $7
	`

	_, err = tx.Exec(ctx,
		updateQuery,
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

	deleteTechQuery := `DELETE FROM project_technologies WHERE project_id = $1`
	_, err = tx.Exec(ctx, deleteTechQuery, project.ID)
	if err != nil {
		return 0, err
	}

	if len(project.Technologies) > 0 {
		insertTechQuery := `
			INSERT INTO project_technologies
				(project_id, technology_id)
			VALUES
				($1, $2)
		`
		for _, tech := range project.Technologies {
			_, err := tx.Exec(ctx, insertTechQuery, project.ID, tech.ID)
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

	projectQuery := `
		SELECT
			id, name, description, github_url, demo_url, is_pinned, category_id, created_at, updated_at
		FROM
			projects
		WHERE
			id = $1
	`

	project := entity.Project{}
	err := pr.Conn.QueryRow(
		ctx,
		projectQuery,
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

	techQuery := `
		SELECT
			t.id, t.name, t.photo_url
		FROM
			technologies t
		JOIN
			project_technologies pt ON t.id = pt.technology_id
		WHERE
			pt.project_id = $1
	`
	rows, err := pr.Conn.Query(ctx, techQuery, id)
	if err != nil {
		return entity.Project{}, err
	}
	defer rows.Close()

	project.Technologies = []entity.Technology{}
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

	projectQuery := `
		SELECT
			id, name, description, github_url, demo_url, is_pinned, category_id, created_at, updated_at
		FROM
			projects
	`

	rows, err := pr.Conn.Query(
		ctx,
		projectQuery,
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

		techQuery := `
			SELECT
				t.id, t.name, t.photo_url
			FROM
				technologies t
			JOIN
				project_technologies pt ON t.id = pt.technology_id
			WHERE
				pt.project_id = $1
		`
		techRows, err := pr.Conn.Query(ctx, techQuery, projectObj.ID)
		if err != nil {
			return []entity.Project{}, err
		}

		projectObj.Technologies = []entity.Technology{}
		for techRows.Next() {
			var tech entity.Technology
			if err := techRows.Scan(&tech.ID, &tech.Name, &tech.PhotoURL); err != nil {
				return []entity.Project{}, err
			}
			projectObj.Technologies = append(projectObj.Technologies, tech)
		}
		techRows.Close()

		projectList = append(projectList, projectObj)
	}

	return projectList, nil
}
