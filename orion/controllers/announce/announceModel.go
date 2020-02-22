package announce

import (
	"database/sql"
)

type Announce struct {
	Id				string
	CreatedAt		uint
	UpdatedAt		uint
	DeletedAt		sql.NullInt64
	PostedAt		uint
	Author			string
	Message			string
}