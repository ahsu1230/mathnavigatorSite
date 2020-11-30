package domains

import (
	"database/sql"
	"fmt"
	"regexp"
	"time"
)

var TABLE_USERS = "users"

type User struct {
	Id             uint         `json:"id"`
	CreatedAt      time.Time    `json:"createdAt" db:"created_at"`
	UpdatedAt      time.Time    `json:"updatedAt" db:"updated_at"`
	DeletedAt      sql.NullTime `json:"-" db:"deleted_at"`
	AccountId      uint         `json:"accountId" db:"account_id"`
	FirstName      string       `json:"firstName" db:"first_name"`
	MiddleName     NullString   `json:"middleName" db:"middle_name"`
	LastName       string       `json:"lastName" db:"last_name"`
	Email          string       `json:"email"`
	Phone          NullString   `json:"phone"`
	IsAdminCreated bool         `json:"isAdminCreated" db:"is_admin_created"`
	IsGuardian     bool         `json:"isGuardian" db:"is_guardian"`
	School         NullString   `json:"school" db:"school"`
	GraduationYear NullUint     `json:"graduationYear" db:"graduation_year"`
	Notes          NullString   `json:"notes"`
}

// Class Methods

func (user *User) Validate() error {
	messageFmt := "Invalid User: %s"

	// Retrieves the inputted values
	firstName := user.FirstName
	lastName := user.LastName
	email := user.Email
	school := user.School.String
	year := user.GraduationYear.Uint

	// First name validation
	if firstName == "" {
		return fmt.Errorf(messageFmt, "Please provide a first name")
	}

	// Last name validation
	if lastName == "" {
		return fmt.Errorf(messageFmt, "Please provide a last name")
	}

	// Email validation
	if matches, _ := regexp.MatchString(REGEX_EMAIL, email); !matches {
		return fmt.Errorf(messageFmt, "Invalid email format")
	}

	// Phone validation
	if user.Phone.Valid {
		if matches, _ := regexp.MatchString(REGEX_PHONE, user.Phone.String); !matches {
			return fmt.Errorf(messageFmt, "Invalid phone format")
		}
	}

	// School validation
	if matches, _ := regexp.MatchString(REGEX_WORDS, school); matches && user.School.Valid {
		return fmt.Errorf(messageFmt, "School contains non alphabetic characters")
	}

	// Year validation
	if year < 2000 && user.GraduationYear.Valid {
		return fmt.Errorf(messageFmt, "Invalid graduation year. Must be > 2000")
	}
	return nil
}
