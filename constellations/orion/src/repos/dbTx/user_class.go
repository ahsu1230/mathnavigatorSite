package dbTx

import (
	"database/sql"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"time"
)

func (dbTx *DbTx) CreateStmtSelectUserClassByUserId() string {
	return "SELECT * FROM user_classes WHERE user_id=?"
}

func (dbTx *DbTx) CreateStmtSelectUserClassByClasId() string {
	return "SELECT * FROM user_classes WHERE class_id=?"
}

func (dbTx *DbTx) CreateStmtSelectUserClassByBothIds() string {
	return "SELECT * FROM user_classes WHERE user_id=? AND class_id=?"
}

func (dbTx *DbTx) CreateStmtSelectNewUserClass() string {
	return "SELECT * FROM user_classes WHERE created_at>=?"
}

func (dbTx *DbTx) ScanUserClass(rows *sql.Rows) (domains.UserClass, error) {
	var userClass domains.UserClass
	if errScan := rows.Scan(
		&userClass.Id,
		&userClass.CreatedAt,
		&userClass.UpdatedAt,
		&userClass.DeletedAt,
		&userClass.ClassId,
		&userClass.UserId,
		&userClass.AccountId,
		&userClass.State); errScan != nil {
		return domains.UserClass{}, errScan
	}
	return userClass, nil
}

func (dbTx *DbTx) SelectOneUserClass(statement string, args ...interface{}) (domains.UserClass, error) {
	userClasses, err := dbTx.SelectManyUserClasses(statement, args...)
	if err != nil {
		return domains.UserClass{}, err
	}
	if len(userClasses) == 0 {
		return domains.UserClass{}, appErrors.ERR_SQL_NO_ROWS
	}
	return userClasses[0], nil
}

func (dbTx *DbTx) SelectManyUserClasses(statement string, args ...interface{}) ([]domains.UserClass, error) {
	results := make([]domains.UserClass, 0)
	rows, err := dbTx.Query(statement, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user, err := dbTx.ScanUserClass(rows)
		if err != nil {
			return results, err
		}
		results = append(results, user)
	}
	return results, nil
}

func (dbTx *DbTx) InsertUserClass(userClass domains.UserClass) (uint, error) {
	statement := "INSERT INTO user_classes (" +
		"created_at, " +
		"updated_at, " +
		"class_id," +
		"user_id," +
		"account_id," +
		"state" +
		") VALUES (?, ?, ?, ?, ?, ?)"

	now := time.Now().UTC()
	execResult, err := dbTx.Exec(
		statement,
		now,
		now,
		userClass.ClassId,
		userClass.UserId,
		userClass.AccountId,
		userClass.State,
	)
	if err != nil {
		return 0, appErrors.WrapDbExec(err, statement, userClass)
	}
	rowId, err := execResult.LastInsertId()
	if err != nil {
		return 0, appErrors.WrapSQLBadInsertResult(err)
	}
	return uint(rowId), appErrors.ValidateDbResult(execResult, 1, "userClass was not inserted")
}

func (dbTx *DbTx) UpdateUserClassById(userClassId uint, userClass domains.UserClass) error {
	statement := "UPDATE user_classes SET " +
		"updated_at=?, " +
		"class_id=?, " +
		"user_id=?, " +
		"account_id=?, " +
		"state=? " +
		"WHERE id=?"
	now := time.Now().UTC()
	execResult, err := dbTx.Exec(
		statement,
		now,
		userClass.ClassId,
		userClass.UserId,
		userClass.AccountId,
		userClass.State,
		userClassId,
	)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, userClass, userClassId)
	}
	return appErrors.ValidateDbResult(execResult, 1, "userClass was not updated")
}

func (dbTx *DbTx) DeleteUserClass(id uint) error {
	statement := "DELETE FROM user_classes WHERE id=?"
	result, err := dbTx.Exec(statement, id)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, id)
	}
	return appErrors.ValidateDbResult(result, 1, "user-class was not deleted")
}

func (dbTx *DbTx) DeleteUserClassByUserId(userId uint) error {
	statement := "DELETE FROM user_classes WHERE user_id=?"
	_, err := dbTx.Exec(statement, userId)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, userId)
	}
	return nil
}
