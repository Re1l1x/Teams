package repository

import (
	"context"

	"github.com/Re1l1x/Teams/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PersonRepository struct {
	db *pgxpool.Pool
}

func NewPersonRepository(db *pgxpool.Pool) *PersonRepository {
	return &PersonRepository{db: db}
}

func (r *PersonRepository) Create(ctx context.Context, person *models.Person) error {
	query := `
	INSERT INTO people(first_name, last_name, birth_year, group_id)
	VALUES ($1, $2, $3, $4)
	RETURNING id;
	`

	return r.db.QueryRow(
		ctx,
		query,
		person.FirstName,
		person.LastName,
		person.BirthYear,
		person.GroupID,
	).Scan(&person.ID)
}

func (r *PersonRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Person, error) {
	query := `
	SELECT id, first_name, last_name, birth_year, group_id
	FROM people
	WHERE id = $1;
	`

	var p models.Person

	err := r.db.QueryRow(ctx, query, id).Scan(
		&p.ID,
		&p.FirstName,
		&p.LastName,
		&p.BirthYear,
		&p.GroupID,
	)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *PersonRepository) GetAll(ctx context.Context) ([]models.Person, error) {
	query := `
	SELECT id, first_name, last_name, birth_year, group_id
	FROM people
	ORDER BY last_name, first_name;
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []models.Person

	for rows.Next() {
		var p models.Person

		if err := rows.Scan(
			&p.ID,
			&p.FirstName,
			&p.LastName,
			&p.BirthYear,
			&p.GroupID,
		); err != nil {
			return nil, err
		}

		res = append(res, p)
	}

	return res, rows.Err()
}

func (r *PersonRepository) Update(ctx context.Context, person *models.Person) error {
	query := `
	UPDATE people
	SET first_name = $1,
		last_name = $2,
		birth_year = $3,
		group_id = $4
	WHERE id = $5;
	`

	_, err := r.db.Exec(
		ctx,
		query,
		person.FirstName,
		person.LastName,
		person.BirthYear,
		person.GroupID,
		person.ID,
	)

	return err
}

func (r *PersonRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `
	DELETE FROM people
	WHERE id = $1;
	`

	_, err := r.db.Exec(ctx, query, id)
	return err
}

func (r *PersonRepository) GetByGroup(ctx context.Context, groupID uuid.UUID) ([]models.Person, error) {
	query := `
	SELECT id, first_name, last_name, birth_year, group_id
	FROM people
	WHERE group_id = $1
	ORDER BY last_name, first_name;
	`

	rows, err := r.db.Query(ctx, query, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []models.Person

	for rows.Next() {
		var p models.Person

		if err := rows.Scan(
			&p.ID,
			&p.FirstName,
			&p.LastName,
			&p.BirthYear,
			&p.GroupID,
		); err != nil {
			return nil, err
		}

		res = append(res, p)
	}

	return res, rows.Err()
}

func (r *PersonRepository) GetByGroups(ctx context.Context, groupIDs []uuid.UUID) ([]models.Person, error) {
	query := `
	SELECT id, first_name, last_name, birth_year, group_id
	FROM people
	WHERE group_id = ANY($1)
	ORDER BY last_name, first_name;
	`

	rows, err := r.db.Query(ctx, query, groupIDs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []models.Person

	for rows.Next() {
		var p models.Person

		if err := rows.Scan(
			&p.ID,
			&p.FirstName,
			&p.LastName,
			&p.BirthYear,
			&p.GroupID,
		); err != nil {
			return nil, err
		}

		res = append(res, p)
	}

	return res, rows.Err()
}
