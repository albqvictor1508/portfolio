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

func NewExperience(conn *pgxpool.Pool) ExperienceRepository {
	return ExperienceRepository{
		Conn: conn,
	}
}

func (er *ExperienceRepository) Insert(experience entity.Experience) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		INSERT INTO experiences
			(company_name, description, photo_url, role, start_date, end_date, category_id)
		VALUES
			($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`

	var id int
	err := er.Conn.QueryRow(ctx,
		query,
		experience.CompanyName,
		experience.Description,
		experience.PhotoURL,
		experience.Role,
		experience.StartDate,
		experience.EndDate,
		experience.CategoryID,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (er *ExperienceRepository) FindByName(name string) (entity.Experience, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		SELECT
			id, company_name, description, photo_url, role, start_date, end_date, category_id
		FROM
			experiences
		WHERE
			company_name = $1
	`

	experience := entity.Experience{}
	err := er.Conn.QueryRow(
		ctx,
		query,
		name,
	).Scan(
		&experience.ID,
		&experience.CompanyName,
		&experience.Description,
		&experience.PhotoURL,
		&experience.Role,
		&experience.StartDate,
		&experience.EndDate,
		&experience.CategoryID,
	)
	if err == pgx.ErrNoRows {
		return entity.Experience{}, nil
	}

	if err != nil {
		return entity.Experience{}, err
	}

	return experience, nil
}

func (er *ExperienceRepository) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		DELETE FROM
			experiences
		WHERE
			id = $1
	`

	cmd, err := er.Conn.Exec(
		ctx,
		query,
		id,
	)

	if cmd.RowsAffected() == 0 {
		return errors.New("THIS EXPERIENCE NOT EXISTS")
	}
	if err != nil {
		return err
	}

	return nil
}

func (er *ExperienceRepository) Update(experience *entity.Experience) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		UPDATE
			experiences
		SET
			company_name = $1,
			description = $2,
			photo_url = $3,
			role = $4,
			start_date = $5,
			end_date = $6,
			category_id = $7
		WHERE
			id = $8
	`

	_, err := er.Conn.Exec(ctx,
		query,
		experience.CompanyName,
		experience.Description,
		experience.PhotoURL,
		experience.Role,
		experience.StartDate,
		experience.EndDate,
		experience.CategoryID,
		experience.ID,
	)
	if err != nil {
		return 0, err
	}

	return experience.ID, nil
}

func (er *ExperienceRepository) FindByID(id int) (entity.Experience, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		SELECT
			id, company_name, description, photo_url, role, start_date, end_date, category_id
		FROM
			experiences
		WHERE
			id = $1
	`

	experience := entity.Experience{}
	err := er.Conn.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&experience.ID,
		&experience.CompanyName,
		&experience.Description,
		&experience.PhotoURL,
		&experience.Role,
		&experience.StartDate,
		&experience.EndDate,
		&experience.CategoryID,
	)
	if err == pgx.ErrNoRows {
		return entity.Experience{}, nil
	}

	if err != nil {
		return entity.Experience{}, err
	}

	return experience, nil
}

func (pr *ExperienceRepository) GetExperiences() ([]entity.Experience, error) {
	var experienceList []entity.Experience

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		SELECT
			id, company_name, description, photo_url, role, category_id, start_date, end_date
		FROM
			experiences
	`

	rows, err := pr.Conn.Query(
		ctx,
		query,
	)
	if err != nil {
		return []entity.Experience{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var experienceObj entity.Experience
		err := rows.Scan(
			&experienceObj.ID,
			&experienceObj.CompanyName,
			&experienceObj.Description,
			&experienceObj.PhotoURL,
			&experienceObj.Role,
			&experienceObj.CategoryID,
			&experienceObj.StartDate,
			&experienceObj.EndDate,
		)
		if err != nil {
			return []entity.Experience{}, err
		}

		experienceList = append(experienceList, experienceObj)
	}

	return experienceList, nil
}
