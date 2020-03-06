package users

import (
	"database/sql"
)

type User struct {
	Id         uint
	CreatedAt  uint          `db:"created_at"`
	UpdatedAt  uint          `db:"updated_at"`
	DeletedAt  sql.NullInt64 `db:"deleted_at"`
	FirstName  string        `json:"firstName"`
	LastName   string        `json:"lastName"`
	MiddleName string        `json:"middleName"`
	Email      string        `json:"email"`
	Phone      string        `json:"phone"`
	IsGuardian bool          `json:"isGuardian"`
	GuardianId uint          `db:"program_id" json:"guardianId"`
}
