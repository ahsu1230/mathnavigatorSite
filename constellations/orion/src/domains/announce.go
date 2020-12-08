package domains

import (
	"fmt"
	"regexp"
	"time"
)

var TABLE_ANNOUNCEMENTS = "announcements"

type Announce struct {
	Id         uint      `json:"id"`
	CreatedAt  time.Time `json:"-" db:"created_at"`
	UpdatedAt  time.Time `json:"-" db:"updated_at"`
	DeletedAt  NullTime  `json:"-" db:"deleted_at"`
	PostedAt   time.Time `json:"postedAt" db:"posted_at"`
	Author     string    `json:"author"`
	Message    string    `json:"message"`
	OnHomePage bool      `json:"onHomePage" db:"on_home_page"`
}

func (announce *Announce) Validate() error {
	messageFmt := "Invalid Announcement: %s"

	// Retrieves the inputted values
	author := announce.Author
	message := announce.Message

	// Author validation
	if matches, _ := regexp.MatchString(REGEX_LETTER, author); !matches {
		return fmt.Errorf(messageFmt, "Must be a valid author's name")
	}

	// Message validation
	if matches, _ := regexp.MatchString(REGEX_LETTER, message); !matches {
		return fmt.Errorf(messageFmt, "Message must contain at least one alphabetic letter")
	}

	return nil
}
