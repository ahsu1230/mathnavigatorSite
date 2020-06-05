package domains

import (
	"database/sql"
	"errors"
	"regexp"
	"time"
)

var TABLE_PEOPLE = "people"

//rowid,first,middle,last names, phone, emails, isguardian, guardianid, familyid
type People struct {
	Id         uint         `json:"id"`
	CreatedAt  time.Time    `json:"-" db:"created_at"`
	UpdatedAt  time.Time    `json:"-" db:"updated_at"`
	DeletedAt  sql.NullTime `json:"-" db:"deleted_at"`
	FirstName  string       `json:"firstName" db:"first_name"`
	LastName   string       `json:"lastName" db:"last_name"`
	MiddleName NullString   `json:"middleName" db:"middle_name"`
	Email      string       `json:"email"`
	Phone      string       `json:"phone"`
	IsGuardian bool         `json:"isGuardian" db:"is_guardian"`
	GuardianId NullUint     `json:"guardianId" db:"guardian_id"`
	FamilyId   uint			`json:"familyId" db:"family_id"`
}

// Class Methods

func (people *People) Validate() error {
	// Retrieves the inputted values
	id := people.Id
	firstName := people.FirstName
	lastName := people.LastName
	email := people.Email
	phone := people.Phone
	isGuardian := people.IsGuardian
	guardianId := people.GuardianId
	// familyId := people.FamilyId

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
		if guardianId.Valid {
			return errors.New("guardian cannot have a guardian id")
		}
	} else {
		if guardianId.Uint == id || !guardianId.Valid {
			return errors.New("invalid guardian id")
		}
	}

	// // Family validation
	// if familyId.Valid {
	// 		return errors.New("guardian cannot have a guardian id")
	// }
	
	// if familyId.Uint == id || !familyId.Valid {
	// 		return errors.New("invalid guardian id")
	// }

	return nil
}
