package domains

import (
	"database/sql"
	"errors"
	"regexp"
	"time"
)

var TABLE_FAMILY = "family"

//ID,password and primary contact email
//possibly add family name?
type Family struct {
	Id         uint         `json:"id"`
	CreatedAt  time.Time    `json:"-" db:"created_at"`
	UpdatedAt  time.Time    `json:"-" db:"updated_at"`
	DeletedAt  sql.NullTime `json:"-" db:"deleted_at"`
	PrimaryEmail string     `json:"primaryEmail db:"primary_email"` 
	Password   string       `json:"password" db:"password"`
}

// Class Methods

func (family *Family) Validate() error {
	// Retrieves the inputted values
	password := family.Password

	if password == "" {
		return errors.New("invalid password")
	}
	return nil
}
