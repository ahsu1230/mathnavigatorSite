package announce

import (
	"database/sql"
)

type Announce struct {
	Id				string			`db:"id"`
	CreatedAt		uint			`db:"created_at"`
	UpdatedAt		uint			`db:"updated_at"`
	DeletedAt		sql.NullInt64	`db:"deleted_at"`
	PostedAt		uint			`db:"posted_at"`
	Author			string			`db:"author"`
	Message			string			`db:"message"`
}