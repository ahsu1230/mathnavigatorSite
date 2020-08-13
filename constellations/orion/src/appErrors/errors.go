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

// Domain errors
// These errors indicate that there could be a problem with a
// invalid field value in a domain object
var ERR_INVALID_EMAIL = errors.New("INVALID_EMAIL_ERROR")
var ERR_INVALID_PASSWORD = errors.New("INVALID_PASSWORD_ERROR")

// JSON errors
// These errors indicate that there could be a problem with the incoming JSON request body
var ERR_JSON_NULL_BODY = errors.New("JSON_NULL_BODY_ERROR")
var ERR_JSON_READ_BODY = errors.New("JSON_READ_BODY_ERROR")
var ERR_JSON_BIND_BODY = errors.New("JSON_BIND_BODY_ERROR")

// Repo errors
// These errors indicate that there could be a problem with the Go code in Repo
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
var ERR_PARSE = errors.New("PARSE_ERROR")

func WrapInvalidDomain(e error, reason string) error {
	return errors.Wrapf(e, "Invalid Domain (%s)", reason)
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

func WrapDbPrepare(e error, statement string) error {
	return errors.Wrapf(
		ERR_REPO_PREPARE,
		"Error (%v) preparing SQL statement (%s)",
		e,
		statement,
	)
}

func WrapDbQuery(e error, statement string, v ...interface{}) error {
	if errors.Is(e, sql.ErrNoRows) {
		return errors.Wrapf(
			ERR_SQL_NO_ROWS,
			"No rows found (%v) from querying SQL statement (%s) with args (%v)",
			e,
			statement,
			v,
		)
	}

	return errors.Wrapf(
		ERR_REPO_QUERY,
		"Error (%v) querying SQL statement (%s) with args (%v)",
		e,
		statement,
		v,
	)
}

func WrapDbExec(e error, statement string, v ...interface{}) error {
	return errors.Wrapf(
		ERR_REPO_EXEC,
		"Error (%v) executing SQL statement (%s) with args (%v)",
		e,
		statement,
		v,
	)
}

func HandleDbResult(result sql.Result, expected int64, message string) error {
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
	// Errors are usually caught from Prepare, Query, Exec
	return err

	// // if errors.Is(err, sql.ErrNoRows) {
	// if sql.ErrNoRows == err {
	// 	return errors.Wrapf(
	// 		ERR_SQL_NO_ROWS,
	// 		"SQL Query resulted in no rows (%v)",
	// 		err,
	// 	)
	// } else if errors.Is(err, sql.ErrConnDone) {
	// 	return errors.Wrapf(
	// 		ERR_SQL_CONN_DONE,
	// 		"Attempted to execute SQL but connection has already closed (%v)",
	// 		err,
	// 	)
	// } else if errors.Is(err, sql.ErrTxDone) {
	// 	return errors.Wrapf(
	// 		ERR_SQL_TX_DONE,
	// 		"Attempted to execute SQL operation within a transaction that has already been committed or rolled back (%v)",
	// 		err,
	// 	)
	// }

	// // Attempt to parse as MySQL Error
	// me, ok := err.(*mysql.MySQLError)
	// if !ok {
	// 	// If not a MySQL Error, return as a generic repo error
	// 	return errors.Wrapf(ERR_REPO, "Error running Repo function (%v)", err)
	// }

	// var mysqlErr error
	// switch me.Number {
	// 	case 1062:
	// 		mysqlErr = errors.Wrapf(
	// 			ERR_MYSQL_DUPLICATE_ENTRY,
	// 			"MySQL Duplicate Entry error %d (%v)",
	// 			me.Number,
	// 			err,
	// 		)
	// 	default:
	// 		mysqlErr = errors.Wrapf(
	// 			ERR_MYSQL_UNKNOWN,
	// 			"Unknown MySQL error %d (%v)",
	// 			me.Number,
	// 			err,
	// 		)
	// }
	// return mysqlErr
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
