package domains

import (
	"time"
)

var TABLE_ASKFORHELP = "askforhelp"

type AskForHelp struct {
	Id         uint      `json:"id"`
	CreatedAt  time.Time `json:"-" db:"created_at"`
	UpdatedAt  time.Time `json:"-" db:"updated_at"`
	DeletedAt  NullTime  `json:"-" db:"deleted_at"`
	Title      string    `json:"title"`
	Date       string    `json:"date"`
	TimeString string    `json:"timeString"`
	Subject    string    `json:"subject"`
	LocationId string    `json:"locationId" db:"location_id"`
}

func (askForHelp *AskForHelp) Validate() error {
	return nil
}
