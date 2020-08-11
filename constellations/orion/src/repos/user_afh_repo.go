package repos

import (
	"database/sql"
	"time"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos/utils"
)

// Global variable
var UserAfhRepo UserAfhRepoInterface = &userAfhRepo{}

// Implements interface userAfhRepoInterface
type userAfhRepo struct {
	db *sql.DB
}

// Interface to implement
type UserAfhRepoInterface interface {
	Initialize(db *sql.DB)
	Insert(domains.UserAfh) error
	SelectByUserId(uint) ([]domains.UserAfh, error)
	SelectByAfhId(uint) ([]domains.UserAfh, error)
	SelectByBothIds(uint, uint) (domains.UserAfh, error)
	Update(uint, domains.UserAfh) error
	Delete(uint) error
}

func (ur *userAfhRepo) Initialize(db *sql.DB) {
	ur.db = db
}

func (ur *userAfhRepo) Insert(userAfh domains.UserAfh) error {
	statement := "INSERT INTO user_afh (" +
		"created_at, " +
		"updated_at, " +
		"user_id, " +
		"afh_id " +
		") VALUES (?, ?, ?, ?)"

	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()
	execResult, err := stmt.Exec(
		now,
		now,
		userAfh.UserId,
		userAfh.AfhId,
	)
	if err != nil {
		return err
	}
	return utils.HandleSqlExecResult(execResult, 1, "userAfh was not inserted")
}

func (ur *userAfhRepo) SelectByUserId(userId uint) ([]domains.UserAfh, error) {
	results := make([]domains.UserAfh, 0)

	statement := "SELECT * FROM user_afh WHERE user_id=?"
	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var userAfh domains.UserAfh
		if errScan := rows.Scan(
			&userAfh.Id,
			&userAfh.CreatedAt,
			&userAfh.UpdatedAt,
			&userAfh.DeletedAt,
			&userAfh.UserId,
			&userAfh.AfhId); errScan != nil {
			return results, errScan
		}
		results = append(results, userAfh)
	}
	return results, nil
}

func (ur *userAfhRepo) SelectByAfhId(afhId uint) ([]domains.UserAfh, error) {
	results := make([]domains.UserAfh, 0)

	statement := "SELECT * FROM user_afh WHERE afh_id=?"
	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(afhId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var userAfh domains.UserAfh
		if errScan := rows.Scan(
			&userAfh.Id,
			&userAfh.CreatedAt,
			&userAfh.UpdatedAt,
			&userAfh.DeletedAt,
			&userAfh.UserId,
			&userAfh.AfhId); errScan != nil {
			return results, errScan
		}
		results = append(results, userAfh)
	}
	return results, nil
}

func (ur *userAfhRepo) SelectByBothIds(userId, afhId uint) (domains.UserAfh, error) {
	statement := "SELECT * FROM user_afh WHERE user_id=? AND afh_id=?"
	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		return domains.UserAfh{}, err
	}
	defer stmt.Close()

	var userAfh domains.UserAfh
	row := stmt.QueryRow(userId, afhId)
	errScan := row.Scan(
		&userAfh.Id,
		&userAfh.CreatedAt,
		&userAfh.UpdatedAt,
		&userAfh.DeletedAt,
		&userAfh.UserId,
		&userAfh.AfhId)
	return userAfh, errScan
}

func (ur *userAfhRepo) Update(id uint, userAfh domains.UserAfh) error {
	statement := "UPDATE user_afh SET " +
		"user_id=?, " +
		"afh_id=?, " +
		"updated_at=? " +
		"WHERE id=?"
	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()
	execResult, err := stmt.Exec(
		userAfh.UserId,
		userAfh.AfhId,
		now,
		id,
	)
	if err != nil {
		return err
	}
	return utils.HandleSqlExecResult(execResult, 1, "userAfh was not updated")
}

func (ur *userAfhRepo) Delete(id uint) error {
	statement := "DELETE FROM user_afh WHERE id=?"
	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	execResult, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	return utils.HandleSqlExecResult(execResult, 1, "userAfh was not deleted")
}

// For Tests Only
func CreateTestUserAfhRepo(db *sql.DB) UserAfhRepoInterface {
	ur := &userAfhRepo{}
	ur.Initialize(db)
	return ur
}
