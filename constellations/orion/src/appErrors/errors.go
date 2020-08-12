package appErrors

import (
	"database/sql"
	"io/ioutil"
	"net/http"
	"github.com/pkg/errors"
)

type ResponseError struct {
    Code     int    `json:"code"`
	Message  string `json:"message"`
	Error	 error  `json:"error"`
}

// var ERR_PARSE = errors.New("PARSING_ERROR")
// var ERR_DOMAIN = errors.New("INVALID_DOMAIN_ERROR")
// var ERR_BIND_JSON = errors.New("BIND_JSON_ERROR")
// var ERR_SQL = errors.New("SQL_ERROR")
// var ERR_DB_PREPARE = errors.New("DB_PREPARE_ERROR")
// var ERR_DB_EXEC = errors.New("DB_EXEC_ERROR")

var ERR_INVALID_EMAIL = errors.New("INVALID_EMAIL_ERROR")
var ERR_INVALID_PASSWORD = errors.New("INVALID_PASSWORD_ERROR")
var ERR_DB_EXEC_MISMATCH = errors.New("DB_EXEC_MISMATCH_ERROR")
var ERR_DB = errors.New("DB_ERROR")

func WrapInvalidDomain(e error, reason string) error {
	return errors.Wrapf(e, "Invalid Domain (%s)", reason)
}

func WrapBindJSON(e error, request *http.Request) error {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return errors.Wrap(err, "Error binding and reaindg request JSON")
	}
	return errors.Wrapf(e, "Error binding request JSON (%v)", body)
}

func WrapRepo(e error) error {
	return e	// should we wrap?
}

func WrapParse(e error, v interface{}) error {
	return errors.Wrapf(e, "Error parsing parameter (%v)", v)
}

func WrapDbPrepare(e error, statement string) error {
	return errors.Wrapf(e, "Error preparing SQL statement (%s)", statement)
}

func WrapDbQuery(e error, statement string, v ...interface{}) error {
	return errors.Wrapf(e, "Error querying SQL statement (%s) with args (%v)", statement, v)
}

func WrapDbExec(e error, statement string, v ...interface{}) error {
	return errors.Wrapf(e, "Error executing SQL statement (%s) with args (%v)", statement, v)
}

func HandleDbResult(result sql.Result, expected int64, message string) error {
	numAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "Error calling result.RowsAffected()")
	}

	if numAffected != expected {
		return errors.Wrap(ERR_DB_EXEC_MISMATCH, message)
	}
	return nil
}