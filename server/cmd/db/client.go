package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool" // essa pool é um connection manager thread safe, faz o gerenciamento e por ser thread safe,
	// se houver várias conexões em paralelo com go routines, ele n vai quebrar
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
	

	return conn, nil
}
