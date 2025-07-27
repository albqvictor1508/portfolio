package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

var Connect pgxpool.Conn

func NewConnection(connectionString string) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)

	defer cancel()

	conn, err := pgxpool.Connect(ctx, connectionString)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	return conn, nil
}
