package domains

import (
	"database/sql"
	"time"
)

var TABLE_USER_CLASSES = "user_classes"

const (
	USER_CLASS_PENDING  uint = 0
	USER_CLASS_ACCEPTED uint = 1
	USER_CLASS_TRIAL    uint = 2
)

type UserClass struct {
	Id        uint         `json:"id"`
	CreatedAt time.Time    `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time    `json:"updatedAt" db:"updated_at"`
	DeletedAt sql.NullTime `json:"-" db:"deleted_at"`
	ClassId   string       `json:"classId" db:"class_id"`
	UserId    uint         `json:"userId" db:"user_id"`
	AccountId uint         `json:"accountId" db:"account_id"`
	State     uint         `json:"state" db:"state"`
}

func (userClass *UserClass) Validate() error {
	return nil
}
