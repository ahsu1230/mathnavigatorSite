package domains

import (
	"github.com/pkg/errors"
)

type AppError struct {
    Code     int    `json:"code"`
	Message  string `json:"message"`
	Error	 string  `json:"error"`
}

func ErrSQL() string {
	return "Error processing SQL query"
}

func ErrRepo(domain string) string {
	return fmt.Sprintf("Error from %s Repo", domain)
}

func ErrParse(fieldName string) string {
	return fmt.Sprintf("Error from parsing `%s`", fieldName)
}

func ErrDomain(domain string) string {
	return fmt.Sprintf("Error validating domain %s", domain string)
}

func ErrBindJSON(structName string) string {
	return fmt.Sprintf("Error binding incoming JSON request to %s", structName)
}