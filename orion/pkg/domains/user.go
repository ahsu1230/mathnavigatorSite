package domains

import (
	"database/sql"
	"errors"
	"regexp"
	"time"
)

var TABLE_USERS = "users"

type User struct {
	Id         uint
	CreatedAt  time.Time    `db:"created_at"`
	UpdatedAt  time.Time    `db:"updated_at"`
	DeletedAt  sql.NullTime `db:"deleted_at"`
	FirstName  string       `json:"firstName"`
	LastName   string       `json:"lastName"`
	MiddleName string       `json:"middleName"`
	Email      string       `json:"email"`
	Phone      string       `json:"phone"`
	IsGuardian bool         `json:"isGuardian"`
	GuardianId uint         `db:"guardian_id" json:"guardianId"`
}

// Class Methods

func (user *User) Validate() error {
	// Retrieves the inputted values
	id := user.Id
	firstName := user.FirstName
	lastName := user.LastName
	email := user.Email
	phone := user.Phone
	isGuardian := user.IsGuardian
	guardianId := user.GuardianId

	// First name validation
	if firstName == "" || len(firstName) > 32 {
		return errors.New("invalid first name")
	}

	// Last name validation
	if lastName == "" || len(lastName) > 32 {
		return errors.New("invalid last name")
	}

	// Email validation
	if matches, _ := regexp.MatchString(REGEX_EMAIL, email); !matches || len(email) > 64 {
		return errors.New("invalid email")
	}

	// Phone validation
	if matches, _ := regexp.MatchString(REGEX_PHONE, phone); !matches || len(phone) > 24 {
		return errors.New("invalid phone")
	}

	// Guardian validation
	if isGuardian {
		if guardianId != 0 {
			return errors.New("guardian cannot have a guardian id")
		}
	} else {
		if guardianId == id {
			return errors.New("invalid guardian id")
		}
	}

	return nil
}
