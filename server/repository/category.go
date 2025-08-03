package repository

import (
	"context"
	"time"

	"github.com/albqvictor1508/portfolio/entity"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type CategoryRepository struct {
	Conn *pgxpool.Pool
}

func NewCategory(conn *pgxpool.Pool) CategoryRepository {
	return CategoryRepository{
		Conn: conn,
	}
}

func (cr *CategoryRepository) Insert(category *entity.Category) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		INSERT INTO categories (name)
		VALUES ($1)
		RETURNING id
	`

	var id int
	err := cr.Conn.QueryRow(
		ctx,
		query,
		category.Name,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (cr *CategoryRepository) FindByID(id int) (entity.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		SELECT
			id,
			name
		FROM
			categories
		WHERE
			id = $1
	`

	var category entity.Category
	err := cr.Conn.QueryRow(
		ctx,
		query,
		id,
	).Scan(&category.ID, &category.Name)

	if err == pgx.ErrNoRows {
		return entity.Category{}, nil
	}

	if err != nil {
		return entity.Category{}, err
	}

	return category, nil
}

func (cr *CategoryRepository) FindByName(name string) (entity.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		SELECT
			id,
			name
		FROM
			categories
		WHERE
			name = $1
	`

	var category entity.Category
	err := cr.Conn.QueryRow(
		ctx,
		query,
		name,
	).Scan(&category.ID, &category.Name)

	if err == pgx.ErrNoRows {
		return entity.Category{}, nil
	}

	if err != nil {
		return entity.Category{}, err
	}

	return category, nil
}

func (cr *CategoryRepository) GetCategories() ([]entity.Category, error) {
	var categoryList []entity.Category

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		SELECT
			id,
			name
		FROM
			categories
	`

	rows, err := cr.Conn.Query(
		ctx,
		query,
	)
	if err != nil {
		return []entity.Category{}, err
	}

	for rows.Next() {
		var categoryObj entity.Category
		err := rows.Scan(
			&categoryObj.ID,
			&categoryObj.Name,
		)
		if err != nil {
			return []entity.Category{}, err
		}
		categoryList = append(categoryList, categoryObj)
	}

	return categoryList, nil
}

func (cr *CategoryRepository) DeleteCategoryByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := `
		DELETE FROM
			categories
		WHERE
			id = $1
	`

	_, err := cr.Conn.Exec(
		ctx,
		query,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}
