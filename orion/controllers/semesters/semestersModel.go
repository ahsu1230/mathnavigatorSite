package semesters

import (
	"database/sql"
)

type Semester struct {
	Id         uint
	CreatedAt  uint          `db:"created_at"`
	UpdatedAt  uint          `db:"updated_at"`
	DeletedAt  sql.NullInt64 `db:"deleted_at"`
	SemesterId string        `db:"semester_id" json:"semesterId"`
	Title      string        `json:"string"`
}
