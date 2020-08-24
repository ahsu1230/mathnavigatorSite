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
	CreatedAt    time.Time    `json:"createdAt" db:"created_at"`
	UpdatedAt    time.Time    `json:"updatedAt" db:"updated_at"`
	DeletedAt    sql.NullTime `json:"-" db:"deleted_at"`
	PrimaryEmail string       `json:"primaryEmail" db:"primary_email"`
	Password     string       `json:"password" db:"password"`
}

type AccountUser struct {
	Account Account `json:"account"`
	User    User    `json:"user"`
}

type AccountSum struct {
	Account Account `json:"account"`
	Balance int     `json:"balance"`
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
