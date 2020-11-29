package dbTx

import (
	"database/sql"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"time"
)

func (dbTx *DbTx) CreateStmtSelectUserAfhByUserId() string {
	return "SELECT * FROM user_afhs WHERE user_id=?"
}

func (dbTx *DbTx) CreateStmtSelectUserAfhByAfhId() string {
	return "SELECT * FROM user_afhs WHERE afh_id=?"
}

func (dbTx *DbTx) CreateStmtSelectUserAfhByBothIds() string {
	return "SELECT * FROM user_afhs WHERE user_id=? AND afh_id=?"
}

func (dbTx *DbTx) CreateStmtSelectNewUserAfh() string {
	return "SELECT * FROM user_afhs WHERE created_at>=?"
}

func (dbTx *DbTx) ScanUserAfh(rows *sql.Rows) (domains.UserAfh, error) {
	var userAfh domains.UserAfh
	if errScan := rows.Scan(
		&userAfh.Id,
		&userAfh.CreatedAt,
		&userAfh.UpdatedAt,
		&userAfh.DeletedAt,
		&userAfh.AfhId,
		&userAfh.UserId,
		&userAfh.AccountId); errScan != nil {
		return domains.UserAfh{}, errScan
	}
	return userAfh, nil
}

func (dbTx *DbTx) SelectOneUserAfh(statement string, args ...interface{}) (domains.UserAfh, error) {
	userAfhs, err := dbTx.SelectManyUserAfhs(statement, args...)
	if err != nil {
		return domains.UserAfh{}, err
	}
	if len(userAfhs) == 0 {
		return domains.UserAfh{}, appErrors.ERR_SQL_NO_ROWS
	}
	return userAfhs[0], nil
}

func (dbTx *DbTx) SelectManyUserAfhs(statement string, args ...interface{}) ([]domains.UserAfh, error) {
	results := make([]domains.UserAfh, 0)
	rows, err := dbTx.Query(statement, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user, err := dbTx.ScanUserAfh(rows)
		if err != nil {
			return results, err
		}
		results = append(results, user)
	}
	return results, nil
}

func (dbTx *DbTx) InsertUserAfh(userAfh domains.UserAfh) (uint, error) {
	statement := "INSERT INTO user_afhs (" +
		"created_at, " +
		"updated_at, " +
		"afh_id, " +
		"user_id, " +
		"account_id" +
		") VALUES (?, ?, ?, ?, ?)"
	now := time.Now().UTC()
	execResult, err := dbTx.Exec(
		statement,
		now,
		now,
		userAfh.AfhId,
		userAfh.UserId,
		userAfh.AccountId,
	)
	if err != nil {
		return 0, appErrors.WrapDbExec(err, statement, userAfh.UserId, userAfh.AfhId)
	}
	rowId, err := execResult.LastInsertId()
	if err != nil {
		return 0, appErrors.WrapSQLBadInsertResult(err)
	}
	return uint(rowId), appErrors.ValidateDbResult(execResult, 1, "userAfh was not inserted")
}

func (dbTx *DbTx) UpdateUserAfhById(userAfhId uint, userAfh domains.UserAfh) error {
	statement := "UPDATE user_afhs SET " +
		"afh_id=?, " +
		"user_id=?, " +
		"account_id=?, " +
		"updated_at=? " +
		"WHERE id=?"
	now := time.Now().UTC()
	execResult, err := dbTx.Exec(
		statement,
		userAfh.AfhId,
		userAfh.UserId,
		userAfh.AccountId,
		now,
		userAfhId,
	)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, userAfh.UserId, userAfh.AfhId, userAfhId)
	}
	return appErrors.ValidateDbResult(execResult, 1, "userAfh was not updated")
}

func (dbTx *DbTx) DeleteUserAfh(id uint) error {
	statement := "DELETE FROM user_afhs WHERE id=?"
	result, err := dbTx.Exec(statement, id)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, id)
	}
	return appErrors.ValidateDbResult(result, 1, "user-afh was not deleted")
}

func (dbTx *DbTx) DeleteUserAfhByUserId(userId uint) error {
	statement := "DELETE FROM user_afhs WHERE user_id=?"
	_, err := dbTx.Exec(statement, userId)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, userId)
	}
	return nil
}
