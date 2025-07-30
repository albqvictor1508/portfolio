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
		"INSERT INTO technologies (name, photo_url) VALUES ($1, $2) RETURNING id",
		technology.Name,
		technology.PhotoURL,
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
		"SELECT t.id, t.name, t.photo_url FROM technologies t WHERE t.id = $1",
		id,
	).Scan(&technology.ID, &technology.Name, &technology.PhotoURL)

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
		"SELECT t.id, t.name, t.photo_url FROM technologies t WHERE t.name = $1",
		name,
	).Scan(&technology.ID, &technology.Name, &technology.PhotoURL)

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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rows, err := cr.Conn.Query(
		ctx,
		"SELECT id, name, photo_url FROM technologies",
	)
	if err != nil {
		return []entity.Technology{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var technologyObj entity.Technology
		err := rows.Scan(
			&technologyObj.ID,
			&technologyObj.Name,
			&technologyObj.PhotoURL,
		)
		if err != nil {
			return []entity.Technology{}, err
		}
		technologyList = append(technologyList, technologyObj)
	}

	return technologyList, nil
}

func (cr *TechnologyRepository) DeleteTechnologyByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := cr.Conn.Exec(
		ctx,
		"DELETE FROM technologies WHERE id = $1",
		id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (cr *TechnologyRepository) Update(technology *entity.Technology) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := cr.Conn.Exec(ctx,
		"UPDATE technologies SET name = $1, photo_url = $2 WHERE id = $3",
		technology.Name,
		technology.PhotoURL,
		technology.ID,
	)
	if err != nil {
		return 0, err
	}

	return technology.ID, nil
}
