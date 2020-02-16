package announce

import (
	"database/sql"
)

type Announce struct {
	Id				uint
	CreatedAt		uint
	UpdatedAt		uint
	DeletedAt		sql.NullInt64
	AnnounceId		string
	Title			string
	Message			string
}