package domains

import (
	"database/sql"
	"errors"
	"regexp"
	"time"
)

var TABLE_USERS = "users"

type User struct {
	Id             uint         `json:"id"`
	CreatedAt      time.Time    `json:"-" db:"created_at"`
	UpdatedAt      time.Time    `json:"-" db:"updated_at"`
	DeletedAt      sql.NullTime `json:"-" db:"deleted_at"`
	FirstName      string       `json:"firstName" db:"first_name"`
	LastName       string       `json:"lastName" db:"last_name"`
	MiddleName     NullString   `json:"middleName" db:"middle_name"`
	Email          string       `json:"email"`
	Phone          string       `json:"phone"`
	IsGuardian     bool         `json:"isGuardian" db:"is_guardian"`
	AccountId      uint         `json:"accountId" db:"account_id"`
	Notes          NullString   `json:"notes"`
	School         NullString   `json:"school" db:"school"`
	GraduationYear NullUint     `json:"graduationYear" db:"graduation_year"`
}

// Class Methods

func (user *User) Validate() error {
	// Retrieves the inputted values
	firstName := user.FirstName
	lastName := user.LastName
	email := user.Email
	phone := user.Phone

	// First name validation
	if firstName == "" {
		return errors.New("invalid first name")
	}

	// Last name validation
	if lastName == "" {
		return errors.New("invalid last name")
	}

	// Email validation
	if matches, _ := regexp.MatchString(REGEX_EMAIL, email); !matches {
		return errors.New("invalid email")
	}

	// Phone validation
	if matches, _ := regexp.MatchString(REGEX_PHONE, phone); !matches {
		return errors.New("invalid phone")
	}

	return nil
}
