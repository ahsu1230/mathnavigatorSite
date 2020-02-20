package achieve

import (
	"database/sql"
)

type Achieve struct {
	Id          uint
	CreatedAt   uint          `db:"created_at"`
	UpdatedAt   uint          `db:"updated_at"`
	DeletedAt   sql.NullInt64 `db:"deleted_at"`
	Year        uint          `json:"programId"`
	Message     string        `json:"message"`
}
