package repos

import (
	"database/sql"
	"time"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
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
	SelectByNew() ([]domains.UserAfh, error)
	Update(uint, domains.UserAfh) error
	Delete(uint) error
}

func (ur *userAfhRepo) Initialize(db *sql.DB) {
	utils.LogWithContext("userAfhRepo.Initialize", logger.Fields{})
	ur.db = db
}

func (ur *userAfhRepo) Insert(userAfh domains.UserAfh) error {
	utils.LogWithContext("userAfhRepo.Insert", logger.Fields{"userAfh": userAfh})
	statement := "INSERT INTO user_afh (" +
		"created_at, " +
		"updated_at, " +
		"user_id, " +
		"afh_id, " +
		"account_id" + 
		") VALUES (?, ?, ?, ?, ?)"

	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		return appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	now := time.Now().UTC()
	execResult, err := stmt.Exec(
		now,
		now,
		userAfh.UserId,
		userAfh.AfhId,
		userAfh.AccountId,
	)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, userAfh.UserId, userAfh.AfhId)
	}
	return appErrors.ValidateDbResult(execResult, 1, "userAfh was not inserted")
}

func (ur *userAfhRepo) SelectByUserId(userId uint) ([]domains.UserAfh, error) {
	utils.LogWithContext("userAfhRepo.SelectByUserId", logger.Fields{"userId": userId})
	results := make([]domains.UserAfh, 0)

	statement := "SELECT * FROM user_afh WHERE user_id=?"
	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		return nil, appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()
	rows, err := stmt.Query(userId)
	if err != nil {
		return nil, appErrors.WrapDbQuery(err, statement, userId)
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
			&userAfh.AfhId,
			&userAfh.AccountId); errScan != nil {
			return results, errScan
		}
		results = append(results, userAfh)
	}
	return results, nil
}

func (ur *userAfhRepo) SelectByAfhId(afhId uint) ([]domains.UserAfh, error) {
	utils.LogWithContext("userAfhRepo.SelectByAfhId", logger.Fields{"afhId": afhId})
	results := make([]domains.UserAfh, 0)

	statement := "SELECT * FROM user_afh WHERE afh_id=?"
	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		return nil, appErrors.WrapDbPrepare(err, statement)
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
			&userAfh.AfhId,
			&userAfh.AccountId); errScan != nil {
			return results, errScan
		}
		results = append(results, userAfh)
	}
	return results, nil
}

func (ur *userAfhRepo) SelectByBothIds(userId, afhId uint) (domains.UserAfh, error) {
	utils.LogWithContext("userAfhRepo.SelectByBothIds", logger.Fields{"userId": userId, "afhId": afhId})
	statement := "SELECT * FROM user_afh WHERE user_id=? AND afh_id=?"
	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		err = appErrors.WrapDbPrepare(err, statement)
		return domains.UserAfh{}, err
	}
	defer stmt.Close()

	var userAfh domains.UserAfh
	row := stmt.QueryRow(userId, afhId)
	if err = row.Scan(
		&userAfh.Id,
		&userAfh.CreatedAt,
		&userAfh.UpdatedAt,
		&userAfh.DeletedAt,
		&userAfh.UserId,
		&userAfh.AfhId,
		&userAfh.AccountId); err != nil {
		err = appErrors.WrapDbExec(err, statement, userId, afhId)
		return domains.UserAfh{}, err
	}
	return userAfh, nil
}
func (ur *userAfhRepo) SelectByNew() ([]domains.UserAfh, error) {
	utils.LogWithContext("userAfhRepo.SelectByNew", logger.Fields{})
	results := make([]domains.UserAfh, 0)

	now := time.Now().UTC()
	week := time.Hour * 24 * 7
	lastWeek := now.Add(-week)
	statement := "SELECT * FROM user_afh WHERE created_at>=?"

	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(lastWeek)
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
			&userAfh.AfhId,
			&userAfh.AccountId); errScan != nil {
			return results, errScan
		}
		results = append(results, userAfh)
	}
	return results, nil
}

func (ur *userAfhRepo) Update(id uint, userAfh domains.UserAfh) error {
	utils.LogWithContext("userAfhRepo.Update", logger.Fields{"userAfh": userAfh})
	statement := "UPDATE user_afh SET " +
		"user_id=?, " +
		"afh_id=?, " +
		"account_id=?, " +
		"updated_at=? " +
		"WHERE id=?"
	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		return appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	now := time.Now().UTC()
	execResult, err := stmt.Exec(
		userAfh.UserId,
		userAfh.AfhId,
		userAfh.AccountId,
		now,
		id,
	)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, userAfh.UserId, userAfh.AfhId, id)
	}
	return appErrors.ValidateDbResult(execResult, 1, "userAfh was not updated")
}

func (ur *userAfhRepo) Delete(id uint) error {
	utils.LogWithContext("userAfhRepo.Delete", logger.Fields{"id": id})
	statement := "DELETE FROM user_afh WHERE id=?"
	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		return appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	execResult, err := stmt.Exec(id)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, id)
	}
	return appErrors.ValidateDbResult(execResult, 1, "userAfh was not deleted")
}

// For Tests Only
func CreateTestUserAfhRepo(db *sql.DB) UserAfhRepoInterface {
	ur := &userAfhRepo{}
	ur.Initialize(db)
	return ur
}
