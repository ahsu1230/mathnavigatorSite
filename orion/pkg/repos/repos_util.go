package repos

import (
	"database/sql"
	"errors"
)

func SetupRepos(db *sql.DB) {
	ProgramRepo.Initialize(db)
	ClassRepo.Initialize(db)
	LocationRepo.Initialize(db)
	AnnounceRepo.Initialize(db)
	AchieveRepo.Initialize(db)
	SemesterRepo.Initialize(db)
	SessionRepo.Initialize(db)
	UserRepo.Initialize(db)
	// AccountRepo.Initialize(db)
}

func handleSqlExecResult(result sql.Result, expected int64, errorMessage string) error {
	numAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if numAffected != expected {
		return errors.New(errorMessage)
	}
	return nil
}

func getPublishError(errorString string) error {
	switch len(errorString) {
	case 0:
		return nil
	case 1:
		return errors.New("one domain failed to publish: " + errorString)
	default:
		return errors.New("one or more domains failed to publish:" + errorString)
	}
}

func appendError(errorString string, id string, err error) string {
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
