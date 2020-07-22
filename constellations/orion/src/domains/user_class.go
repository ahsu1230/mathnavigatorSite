package domains

import (
	"database/sql"
	"time"
)

var TABLE_USERCLASS = "userclass"

type UserClass struct {
	Id        uint         `json:"id"`
	CreatedAt time.Time    `json:"-" db:"created_at"`
	UpdatedAt time.Time    `json:"-" db:"updated_at"`
	DeletedAt sql.NullTime `json:"-" db:"deleted_at"`
	UserId    uint         `json:"userId" db:"user_id"`
	ClassId   string       `json:"classId" db:"class_id"`
	AccountId uint         `json:"accountId" db:"account_id"`
	State     uint         `json:"state" db:"state"`
}

func (userClass *UserClass) Validate() error {
	return nil
}
