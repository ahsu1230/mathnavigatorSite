package repos

import (
	"database/sql"
	"errors"
)

func SetupRepos(db *sql.DB) {
	ProgramRepo.Initialize(db)
	// ClassRepo.Initialize(db)
  LocationRepo.Initialize(db)
	AnnounceRepo.Initialize(db)
	AchieveRepo.Initialize(db)
	SemesterRepo.Initialize(db)
	// UserRepo.Initialize(db)
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
