package repository

import (
	"context"

	"github.com/Re1l1x/Teams/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type GroupRepository struct {
	db *pgxpool.Pool
}

func NewGroupRepository(db *pgxpool.Pool) *GroupRepository {
	return &GroupRepository{db: db}
}

func (r *GroupRepository) Create(ctx context.Context, group *models.Group) error {
	query := `
	INSERT INTO groups(name, parent_id)
	VALUES ($1, $2)
	RETURNING id;
	`

	return r.db.QueryRow(ctx, query, group.Name, group.ParentID).Scan(&group.ID)
}

func (r *GroupRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Group, error) {
	query := `
	SELECT id, name, parent_id
	FROM groups
	WHERE id = $1;
	`

	var g models.Group

	err := r.db.QueryRow(ctx, query, id).Scan(
		&g.ID,
		&g.Name,
		&g.ParentID,
	)

	if err != nil {
		return nil, err
	}

	return &g, nil
}

func (r *GroupRepository) GetAll(ctx context.Context) ([]models.Group, error) {
	query := `
	SELECT id, name, parent_id
	FROM groups
	ORDER BY name;
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []models.Group

	for rows.Next() {
		var g models.Group

		if err := rows.Scan(&g.ID, &g.Name, &g.ParentID); err != nil {
			return nil, err
		}

		res = append(res, g)
	}

	return res, rows.Err()
}

func (r *GroupRepository) Update(ctx context.Context, group *models.Group) error {
	query := `
	UPDATE groups
	SET name = $1, parent_id = $2
	WHERE id = $3;
	`

	_, err := r.db.Exec(ctx, query, group.Name, group.ParentID, group.ID)
	return err
}

func (r *GroupRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `
	DELETE FROM groups
	WHERE id = $1;
	`

	_, err := r.db.Exec(ctx, query, id)
	return err
}

func (r *GroupRepository) GetPeopleByGroup(ctx context.Context, groupID uuid.UUID) ([]models.Person, error) {
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

func (r *GroupRepository) GetPeopleRecursive(ctx context.Context, groupID uuid.UUID) ([]models.Person, error) {
	query := `
	WITH RECURSIVE group_tree AS (
		SELECT id
		FROM groups
		WHERE id = $1

		UNION ALL

		SELECT g.id
		FROM groups g
		INNER JOIN group_tree gt ON g.parent_id = gt.id
	)

	SELECT p.id, p.first_name, p.last_name, p.birth_year, p.group_id
	FROM people p
	INNER JOIN group_tree gt ON p.group_id = gt.id
	ORDER BY p.last_name, p.first_name;
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

func (r *GroupRepository) CountPeople(ctx context.Context, groupID uuid.UUID) (int, error) {
	query := `
	SELECT COUNT(*)
	FROM people
	WHERE group_id = $1;
	`

	var cnt int
	err := r.db.QueryRow(ctx, query, groupID).Scan(&cnt)

	return cnt, err
}

func (r *GroupRepository) CountPeopleRecursive(ctx context.Context, groupID uuid.UUID) (int, error) {
	query := `
	WITH RECURSIVE group_tree AS (
		SELECT id
		FROM groups
		WHERE id = $1

		UNION ALL

		SELECT g.id
		FROM groups g
		INNER JOIN group_tree gt ON g.parent_id = gt.id
	)

	SELECT COUNT(*)
	FROM people p
	WHERE p.group_id IN (SELECT id FROM group_tree);
	`

	var cnt int
	err := r.db.QueryRow(ctx, query, groupID).Scan(&cnt)

	return cnt, err
}
