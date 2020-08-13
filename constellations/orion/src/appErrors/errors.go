package appErrors

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"

	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   error  `json:"error"`
}

// JSON errors
// These errors indicate that there could be a problem with the incoming JSON request body
var ERR_JSON_NULL_BODY = errors.New("JSON_NULL_BODY_ERROR")
var ERR_JSON_READ_BODY = errors.New("JSON_READ_BODY_ERROR")
var ERR_JSON_BIND_BODY = errors.New("JSON_BIND_BODY_ERROR")

// Repo errors
// These errors indicate that there could be a problem with the Go code in Repo
var ERR_REPO_TX_BEGIN = errors.New("REPO_TX_BEGIN_ERROR")
var ERR_REPO_TX_COMMIT = errors.New("REPO_TX_COMMIT_ERROR")
var ERR_REPO_TX_ROLLBACK = errors.New("REPO_TX_ROLLBACK_ERROR")

var ERR_REPO_PREPARE = errors.New("REPO_PREPARE_ERROR")
var ERR_REPO_QUERY = errors.New("REPO_QUERY_ERROR")
var ERR_REPO_EXEC = errors.New("REPO_EXEC_ERROR")
var ERR_REPO_BAD_RESULTS = errors.New("REPO_BAD_RESULTS_ERROR")
var ERR_REPO_EXEC_MISMATCH = errors.New("REPO_EXEC_MISMATCH_ERROR")
var ERR_REPO = errors.New("REPO_ERROR") // Generic Repo error

// SQL errors
// These errors originate from the "database/sql" package
var ERR_SQL_CONN_DONE = errors.New("SQL_CONN_DONE_ERROR")
var ERR_SQL_TX_DONE = errors.New("SQL_TX_DONE_ERROR")
var ERR_SQL_NO_ROWS = errors.New("SQL_NO_ROWS_ERROR")

// MYSQL errors
// These errors indicate that there was a conflict in the MySQL database state
// These errors originate from the "mysql" package
var ERR_MYSQL_DUPLICATE_ENTRY = errors.New("MYSQL_DUP_ENTRY_ERROR")
var ERR_MYSQL_UNKNOWN = errors.New("MYSQL_UNKNOWN_ERROR")

// Other errors
var ERR_CTRL = errors.New("CTRL_ERROR")
var ERR_INVALID_DOMAIN = errors.New("INVALID_DOMAIN_ERROR")
var ERR_PARSE = errors.New("PARSE_ERROR")

func WrapInvalidDomain(reason string) error {
	return errors.Wrapf(ERR_INVALID_DOMAIN, "Invalid Domain (%s)", reason)
}

func WrapCtrlf(message string, v ...interface{}) error {
	return errors.Wrapf(ERR_CTRL, message, v)
}

func WrapBindJSON(e error, request *http.Request) error {
	if request.Body == nil {
		return errors.Wrap(ERR_JSON_NULL_BODY, "Unexpected null JSON body")
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return errors.Wrapf(ERR_JSON_READ_BODY, "Error (%v) binding and reading request JSON", err)
	}
	return errors.Wrapf(ERR_JSON_BIND_BODY, "Error (%v) binding request JSON (%v)", e, body)
}

func WrapParse(e error, v interface{}) error {
	return errors.Wrapf(ERR_PARSE, "Error (%v) parsing parameter (%v)", e, v)
}

func WrapDbTxBegin(e error) error {
	return errors.Wrapf(ERR_REPO_TX_BEGIN, "Error (%v) starting a transaction", e)
}

func WrapDbTxCommit(e error) error {
	return errors.Wrapf(ERR_REPO_TX_COMMIT, "Error (%v) commiting a transaction", e)
}

func WrapDbTxRollback(e error) error {
	return errors.Wrapf(ERR_REPO_TX_ROLLBACK, "Error (%v) rolling back a transaction", e)
}

func WrapDbPrepare(e error, statement string) error {
	return errors.Wrapf(
		ERR_REPO_PREPARE,
		"Error (%v) preparing SQL statement (%s)",
		e,
		statement,
	)
}

func WrapDbQuery(e error, statement string, v ...interface{}) error {
	return wrapDbErrorHelper(ERR_REPO_QUERY, "querying SQL statement", e, statement, v)
}

func WrapDbExec(e error, statement string, v ...interface{}) error {
	return wrapDbErrorHelper(ERR_REPO_EXEC, "executing SQL statement", e, statement, v)
}

func wrapDbErrorHelper(
	parentError error,
	verb string,
	originalErr error,
	statement string,
	v ...interface{}) error {

	if err, found := checkSqlError(originalErr, statement, v); found {
		return err
	}

	return errors.Wrapf(
		parentError,
		"Error (%v) %s (%s) with args (%v)",
		originalErr,
		verb,
		statement,
		v,
	)
}

func checkSqlError(e error, statement string, v ...interface{}) (error, bool) {
	if errors.Is(e, sql.ErrNoRows) {
		return errors.Wrapf(
			ERR_SQL_NO_ROWS,
			"No rows found (%v) from querying SQL statement (%s) with args (%v)",
			e,
			statement,
			v,
		), true
	} else if errors.Is(e, sql.ErrTxDone) {
		return errors.Wrapf(
			ERR_SQL_TX_DONE,
			"Attempted to operate (%s) when SQL transaction has already been committed or rolled back. (%v)",
			statement,
			e,
		), true
	} else {
		return nil, false
	}
}

func ValidateDbResult(result sql.Result, expected int64, message string) error {
	numAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrapf(
			ERR_REPO_BAD_RESULTS,
			"Error (%v) calling result.RowsAffected(). %s",
			err,
			message,
		)
	}

	if numAffected != expected {
		return errors.Wrap(ERR_REPO_EXEC_MISMATCH, message)
	}
	return nil
}

func WrapRepo(err error) error {
	// TODO: do we want to wrap repo errors?
	// Errors are usually caught from Prepare, Query, Exec, and ValidateDbResult
	return err
}

// For testing purposes only

func TestMySQLDuplicateEntryError() error {
	return &mysql.MySQLError{
		1062,
		"Fake duplicate entry",
	}
}

func TestDbNoRowsError() error {
	return ERR_SQL_NO_ROWS
}
