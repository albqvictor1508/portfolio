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

func (r *Repository) Insert(project internal.Project) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.Conn.Exec(
		ctx,
		"INSERT INTO projects () VALUES()",
	)

	return err
}
