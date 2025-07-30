package repository

import (
	"context"
	"time"

	"github.com/albqvictor1508/portfolio/entity"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type TechnologyRepository struct {
	Conn *pgxpool.Pool
}

func NewTechnology(conn *pgxpool.Pool) TechnologyRepository {
	return TechnologyRepository{
		Conn: conn,
	}
}

func (cr *TechnologyRepository) Insert(technology *entity.Technology) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var id int
	err := cr.Conn.QueryRow(
		ctx,
		"INSERT INTO technologies (name) VALUES ($1) RETURNING id",
		technology.Name,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (cr *TechnologyRepository) FindByID(id int) (entity.Technology, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var technology entity.Technology
	err := cr.Conn.QueryRow(
		ctx,
		"SELECT t.id, t.name FROM technologies t WHERE t.id = $1",
		id,
	).Scan(&technology.ID, &technology.Name)

	if err == pgx.ErrNoRows {
		return entity.Technology{}, nil
	}

	if err != nil {
		return entity.Technology{}, err
	}

	return technology, nil
}

func (cr *TechnologyRepository) FindByName(name string) (entity.Technology, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var technology entity.Technology
	err := cr.Conn.QueryRow(
		ctx,
		"SELECT c.id, c.name FROM categories c WHERE c.name = $1",
		name,
	).Scan(&technology.ID, &technology.Name)

	if err == pgx.ErrNoRows {
		return entity.Technology{}, nil
	}

	if err != nil {
		return entity.Technology{}, err
	}

	return technology, nil
}

func (cr *TechnologyRepository) GetTechnologies() ([]entity.Technology, error) {
	var technologyList []entity.Technology
	var technologyObj entity.Technology

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rows, err := cr.Conn.Query(
		ctx,
		"SELECT * FROM technologies t",
	)
	if err != nil {
		return []entity.Technology{}, err
	}

	for rows.Next() {
		err := rows.Scan(
			&technologyObj.ID,
			&technologyObj.Name,
		)
		if err != nil {
			return []entity.Technology{}, err
		}
		technologyList = append(technologyList, technologyObj)
	}

	return technologyList, nil
}

func (cr *CategoryRepository) DeleteTechnologyByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := cr.Conn.Exec(
		ctx,
		"DELETE FROM technology t WHERE t.id = $1",
		id,
	)
	if err != nil {
		return err
	}
	return nil
}
