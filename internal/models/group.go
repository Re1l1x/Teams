package models

import "github.com/google/uuid"

type Group struct {
	ID       uuid.UUID  `db:"id" json:"id"`
	Name     string     `db:"name" json:"name"`
	ParentID *uuid.UUID `db:"parent_id" json:"parent_id,omitempty"`
}
