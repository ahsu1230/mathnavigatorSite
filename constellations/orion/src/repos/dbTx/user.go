package dbTx

import (
	"database/sql"
	"fmt"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"time"
)

func (dbTx *DbTx) CreateStmtSelectUserSearchOneTerm(lcSearch string) string {
	return fmt.Sprintf(
		"SELECT * FROM users WHERE "+
			"LOWER(`first_name`) LIKE '%%%s%%' OR "+
			"LOWER(`middle_name`) LIKE '%%%s%%' OR "+
			"LOWER(`last_name`) LIKE '%%%s%%' OR "+
			"LOWER(`email`) LIKE '%%%s%%'",
		lcSearch, lcSearch, lcSearch, lcSearch)
}

func (dbTx *DbTx) CreateStmtSelectUserSearchTwoTerms(searchTerms []string) string {
	return fmt.Sprintf(
		"SELECT * FROM users WHERE "+
			"(LOWER(`first_name`) LIKE '%%%s%%' AND "+
			"LOWER(`last_name`) LIKE '%%%s%%')",
		searchTerms[0], searchTerms[1])
}

func (dbTx *DbTx) CreateStmtSelectUserSearchThreeTerms(regexTerm string) string {
	return fmt.Sprintf(
		"SELECT * FROM users WHERE "+
			"LOWER(`first_name`) REGEXP '%s' OR "+
			"LOWER(`middle_name`) REGEXP '%s' OR "+
			"LOWER(`last_name`) REGEXP '%s' OR "+
			"LOWER(`email`) REGEXP '%s'",
		regexTerm, regexTerm, regexTerm, regexTerm)
}

func (dbTx *DbTx) CreateStmtSelectUsersAllWithLimitOffset() string {
	return "SELECT * FROM users LIMIT ? OFFSET ?"
}

func (dbTx *DbTx) CreateStmtSelectUserNamesWithLimitOffset() string {
	return "SELECT * FROM users WHERE ? IN (first_name,last_name,middle_name) LIMIT ? OFFSET ?"
}

func (dbTx *DbTx) CreateStmtSelectUserById() string {
	return "SELECT * FROM users WHERE id=?"
}

func (dbTx *DbTx) CreateStmtSelectUsersByAccountId() string {
	return "SELECT * FROM users WHERE account_id=?"
}

func (dbTx *DbTx) CreateStmtSelectUserByEmail() string {
	return "SELECT * FROM users WHERE email=?"
}

func (dbTx *DbTx) CreateStmtSelectUsersByNew() string {
	return "SELECT * FROM users WHERE created_at>=?"
}

func (dbTx *DbTx) ScanUser(rows *sql.Rows) (domains.User, error) {
	var user domains.User
	if err := rows.Scan(
		&user.Id,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
		&user.AccountId,
		&user.FirstName,
		&user.MiddleName,
		&user.LastName,
		&user.Email,
		&user.Phone,
		&user.IsAdminCreated,
		&user.IsGuardian,
		&user.School,
		&user.GraduationYear,
		&user.Notes); err != nil {
		return domains.User{}, err
	}
	return user, nil
}

func (dbTx *DbTx) SelectOneUser(statement string, args ...interface{}) (domains.User, error) {
	users, err := dbTx.SelectManyUsers(statement, args...)
	if err != nil {
		return domains.User{}, err
	}
	if len(users) == 0 {
		return domains.User{}, appErrors.ERR_SQL_NO_ROWS
	}
	return users[0], nil
}

func (dbTx *DbTx) SelectManyUsers(statement string, args ...interface{}) ([]domains.User, error) {
	results := make([]domains.User, 0)
	rows, err := dbTx.Query(statement, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user, err := dbTx.ScanUser(rows)
		if err != nil {
			return results, err
		}
		results = append(results, user)
	}
	return results, nil
}

func (dbTx *DbTx) InsertUser(user domains.User) (uint, error) {
	statement := "INSERT INTO users (" +
		"created_at, " +
		"updated_at, " +
		"account_id, " +
		"first_name, " +
		"middle_name, " +
		"last_name, " +
		"email, " +
		"phone, " +
		"is_admin_created, " +
		"is_guardian, " +
		"school, " +
		"graduation_year, " +
		"notes" +
		") VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	now := time.Now().UTC()
	execResult, err := dbTx.Exec(
		statement,
		now,
		now,
		user.AccountId,
		user.FirstName,
		user.MiddleName,
		user.LastName,
		user.Email,
		user.Phone,
		user.IsAdminCreated,
		user.IsGuardian,
		user.School,
		user.GraduationYear,
		user.Notes,
	)
	if err != nil {
		return 0, err
	}

	rowId, err := execResult.LastInsertId()
	if err != nil {
		return 0, appErrors.WrapSQLBadInsertResult(err)
	}
	return uint(rowId), appErrors.ValidateDbResult(execResult, 1, "user was not inserted")
}

func (dbTx *DbTx) UpdateUserById(userId uint, user domains.User) error {
	statement := "UPDATE users SET " +
		"updated_at=?, " +
		"account_id=?, " +
		"first_name=?, " +
		"middle_name=?, " +
		"last_name=?, " +
		"email=?, " +
		"phone=?, " +
		"is_admin_created=?, " +
		"is_guardian=?, " +
		"school=?, " +
		"graduation_year=?, " +
		"notes=? " +
		"WHERE id=?"

	now := time.Now().UTC()
	result, err := dbTx.Exec(
		statement,
		now,
		user.AccountId,
		user.FirstName,
		user.MiddleName,
		user.LastName,
		user.Email,
		user.Phone,
		user.IsAdminCreated,
		user.IsGuardian,
		user.School,
		user.GraduationYear,
		user.Notes,
		userId)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, userId, user)
	}
	return appErrors.ValidateDbResult(result, 1, "user was not updated")
}

func (dbTx *DbTx) DeleteUser(userId uint) error {
	statement := "DELETE FROM users WHERE id=?"
	result, err := dbTx.Exec(statement, userId)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, userId)
	}
	return appErrors.ValidateDbResult(result, 1, "user was not deleted")
}
