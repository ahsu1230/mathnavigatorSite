package utils

import (
	"database/sql"
	"errors"
)

func AppendError(errorString string, id string, err error) string {
	if err == nil {
		return errorString
	}
	if len(id) > 0 {
		errorString += "id: " + id + ", error: " + err.Error() + "\n"
	} else {
		errorString += err.Error() + "\n"
	}
	return errorString
}

func HandleSqlExecResult(result sql.Result, expected int64, message string) error {
	numAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if numAffected != expected {
		return errors.New(message)
	}
	return nil
}
