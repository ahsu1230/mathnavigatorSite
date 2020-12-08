package domains

import (
	"database/sql"
	"time"
)

var TABLE_USER_CLASSES = "user_classes"

const (
	USER_CLASS_PENDING   string = "pending"
	USER_CLASS_ENROLLED  string = "enrolled"
	USER_CLASS_TRIAL     string = "trial"
	USER_CLASS_DISMISSED string = "dismissed"
)

var ALL_USER_CLASS_STATES = []string{
	USER_CLASS_PENDING,
	USER_CLASS_ENROLLED,
	USER_CLASS_TRIAL,
	USER_CLASS_DISMISSED,
}

type UserClass struct {
	Id        uint         `json:"id"`
	CreatedAt time.Time    `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time    `json:"updatedAt" db:"updated_at"`
	DeletedAt sql.NullTime `json:"-" db:"deleted_at"`
	ClassId   string       `json:"classId" db:"class_id"`
	UserId    uint         `json:"userId" db:"user_id"`
	AccountId uint         `json:"accountId" db:"account_id"`
	State     string       `json:"state" db:"state"`
}

func (userClass *UserClass) Validate() error {
	return nil
}
