package domains

import (
	"time"
)

var TABLE_ASKFORHELP = "askforhelp"

type AskForHelp struct {
	Id         uint      `json:"id"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
	DeletedAt  time.Time `json:"-"`
	Title      string    `json:"message"`
	Date       string    `json:"date"`
	TimeString string    `json:"timeString"`
	Subject    string    `json:"subject"`
	LocationId string    `json:"locationId"`
}

func (askForHelp *AskForHelp) Validate() error {
	return nil
}
