package announce

import (
	"database/sql"
)

type Announce struct {
	Id        uint
	CreatedAt uint          `db:"created_at"`
	UpdatedAt uint          `db:"updated_at"`
	DeletedAt sql.NullInt64 `db:"deleted_at"`
	PostedAt  uint          `db:"posted_at"`
	Author    string        `json:"author"`
	Message   string        `json:"message"`
}
