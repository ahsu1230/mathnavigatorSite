package domains

import (
	"errors"
	"fmt"
	"strings"
)

type UnpublishedDomains struct {
	Programs  []Program  `json:"programs"`
	Classes   []Class    `json:"classes"`
	Locations []Location `json:"locations"`
	Achieves  []Achieve  `json:"achieves"`
	Semesters []Semester `json:"semesters"`
	Sessions  []Session  `json:"sessions"`
}

type PublishErrorBody struct {
	RowId    uint   `json:"rowId,omitempty"`
	StringId string `json:"stringId,omitempty"`
	Error    error  `json:"error"`
}

func Concatenate(initMessage string, errorList []PublishErrorBody, isString bool) error {
	var errorStrings []string
	errorStrings = append(errorStrings, initMessage)
	for _, errorBody := range errorList {
		if isString {
			errorStrings = append(errorStrings, errorBody.StringId+": "+errorBody.Error.Error())
		} else {
			errorStrings = append(errorStrings, fmt.Sprint(errorBody.RowId)+": "+errorBody.Error.Error())
		}
	}
	return errors.New(strings.Join(errorStrings, "\n"))
}
