package domains

import (
	"database/sql"
	"regexp"
	"time"

	"fmt"
)

var TABLE_ACCOUNTS = "accounts"

type Account struct {
	Id           uint         `json:"id"`
	CreatedAt    time.Time    `json:"-" db:"created_at"`
	UpdatedAt    time.Time    `json:"-" db:"updated_at"`
	DeletedAt    sql.NullTime `json:"-" db:"deleted_at"`
	PrimaryEmail string       `json:"primaryEmail" db:"primary_email"`
	Password     string       `json:"password" db:"password"`
}

// Class Methods

func (account *Account) Validate() error {
	messageFmt := "Invalid Account: %s"
	// Retrieves the inputted values
	primaryEmail := account.PrimaryEmail
	password := account.Password

	if matches, _ := regexp.MatchString(REGEX_EMAIL, primaryEmail); !matches {
		return fmt.Errorf(messageFmt, "Invalid Email")
	}
	if len(password) < 8 {
		return fmt.Errorf(messageFmt, "Password must be at least 8 characters long")
	}
	return nil
}
