package repository

import (
	"context"
	"time"

	"github.com/albqvictor1508/portfolio/entity"
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

	var id int
	err := cr.Conn.QueryRow(
		ctx,
		"INSERT INTO categories c VALUES ($1)",
		category.Name,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
