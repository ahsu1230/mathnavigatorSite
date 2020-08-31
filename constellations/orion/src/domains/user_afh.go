package domains

import (
	"database/sql"
	"time"
)

var TABLE_USER_AFH = "user_afh"

type UserAfh struct {
	Id        uint         `json:"id"`
	CreatedAt time.Time    `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time    `json:"updatedAt" db:"updated_at"`
	DeletedAt sql.NullTime `json:"-" db:"deleted_at"`
	UserId    uint         `json:"userId" db:"user_id"`
	AfhId     uint         `json:"afhId" db:"afh_id"`
	AccountId uint         `json:"accountId" db:"account_id"`
}
