package models

import "github.com/google/uuid"

type Person struct {
	ID        uuid.UUID `db:"id" json:"id"`
	FirstName string    `db:"first_name" json:"first_name"`
	LastName  string    `db:"last_name" json:"last_name"`
	BirthYear int       `db:"birth_year" json:"birth_year"`
	GroupID   uuid.UUID `db:"group_id" json:"group_id"`
}
