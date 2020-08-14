package repos

import (
	"database/sql"
	"fmt"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos/utils"
	"strings"
	"time"
)

// Global variable
var UserRepo UserRepoInterface = &userRepo{}

// Implements interface userRepoInterface
type userRepo struct {
	db *sql.DB // golang native db connection
}

// Interface to implement
type UserRepoInterface interface {
	Initialize(db *sql.DB)
	SearchUsers(string) ([]domains.User, error)
	SelectAll(string, int, int) ([]domains.User, error)
	SelectById(uint) (domains.User, error)
	SelectByAccountId(uint) ([]domains.User, error)
	SelectByNew() ([]domains.User, error)
	Insert(domains.User) error
	Update(uint, domains.User) error
	Delete(uint) error
}

func (ur *userRepo) Initialize(db *sql.DB) {
	ur.db = db
}

//TODO

func (ur *userRepo) SearchUsers(search string) ([]domains.User, error) {
	results := make([]domains.User, 0)

	lcSearch := strings.ToLower(search)
	query := fmt.Sprintf("SELECT * FROM users WHERE LOWER(`first_name`) LIKE '%%%s%%' OR LOWER(`middle_name`) LIKE '%%%s%%' OR LOWER(`last_name`) LIKE '%%%s%%' OR LOWER(`email`) LIKE '%%%s%%'", lcSearch, lcSearch, lcSearch, lcSearch)

	stmt, err := ur.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user domains.User
		if errScan := rows.Scan(
			&user.Id,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt,
			&user.FirstName,
			&user.LastName,
			&user.MiddleName,
			&user.Email,
			&user.Phone,
			&user.IsGuardian,
			&user.AccountId,
			&user.Notes,
			&user.School,
			&user.GraduationYear); errScan != nil {
			return results, errScan
		}
		results = append(results, user)
	}

	return results, nil
}

func (ur *userRepo) SelectAll(search string, pageSize, offset int) ([]domains.User, error) {
	results := make([]domains.User, 0)

	getAll := len(search) == 0
	var query string
	if getAll {
		query = "SELECT * FROM users LIMIT ? OFFSET ?"
	} else {
		query = "SELECT * FROM users WHERE ? IN (first_name,last_name,middle_name) LIMIT ? OFFSET ?"
	}
	stmt, err := ur.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var rows *sql.Rows
	if getAll {
		rows, err = stmt.Query(pageSize, offset)
	} else {
		rows, err = stmt.Query(search, pageSize, offset)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user domains.User
		if errScan := rows.Scan(
			&user.Id,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt,
			&user.FirstName,
			&user.LastName,
			&user.MiddleName,
			&user.Email,
			&user.Phone,
			&user.IsGuardian,
			&user.AccountId,
			&user.Notes,
			&user.School,
			&user.GraduationYear); errScan != nil {
			return results, errScan
		}
		results = append(results, user)
	}
	return results, nil
}

func (ur *userRepo) SelectById(id uint) (domains.User, error) {
	statement := "SELECT * FROM users WHERE id=?"
	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		return domains.User{}, err
	}
	defer stmt.Close()

	var user domains.User
	row := stmt.QueryRow(id)
	errScan := row.Scan(
		&user.Id,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
		&user.FirstName,
		&user.LastName,
		&user.MiddleName,
		&user.Email,
		&user.Phone,
		&user.IsGuardian,
		&user.AccountId,
		&user.Notes,
		&user.School,
		&user.GraduationYear)
	return user, errScan
}

func (ur *userRepo) SelectByAccountId(accountId uint) ([]domains.User, error) {
	results := make([]domains.User, 0)

	stmt, err := ur.db.Prepare("SELECT * FROM users WHERE account_id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(domains.NewNullUint(accountId))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user domains.User
		if errScan := rows.Scan(
			&user.Id,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt,
			&user.FirstName,
			&user.LastName,
			&user.MiddleName,
			&user.Email,
			&user.Phone,
			&user.IsGuardian,
			&user.AccountId,
			&user.Notes,
			&user.School,
			&user.GraduationYear); errScan != nil {
			return results, errScan
		}
		results = append(results, user)
	}
	return results, nil
}

func (ur *userRepo) SelectByNew() ([]domains.User, error) {
	results := make([]domains.User, 0)

	now := time.Now().UTC()
	week := time.Hour * 24 * 7
	lastWeek := now.Add(-week)
	stmt, err := ur.db.Prepare("SELECT * FROM users WHERE created_at>=?")
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
		var user domains.User
		if errScan := rows.Scan(
			&user.Id,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt,
			&user.FirstName,
			&user.LastName,
			&user.MiddleName,
			&user.Email,
			&user.Phone,
			&user.IsGuardian,
			&user.AccountId,
			&user.Notes,
			&user.School,
			&user.GraduationYear); errScan != nil {
			return results, errScan
		}
		results = append(results, user)
	}
	return results, nil
}

func (ur *userRepo) Insert(user domains.User) error {
	statement := "INSERT INTO users (" +
		"created_at, " +
		"updated_at, " +
		"first_name, " +
		"last_name," +
		"middle_name, " +
		"email," +
		"phone, " +
		"is_guardian," +
		"account_id," +
		"notes," +
		"school," +
		"graduation_year" +
		") VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()
	execResult, err := stmt.Exec(
		now,
		now,
		user.FirstName,
		user.LastName,
		user.MiddleName,
		user.Email,
		user.Phone,
		user.IsGuardian,
		user.AccountId,
		user.Notes,
		user.School,
		user.GraduationYear,
	)
	if err != nil {
		return err
	}
	return utils.HandleSqlExecResult(execResult, 1, "user was not inserted")
}

func (ur *userRepo) Update(id uint, user domains.User) error {
	statement := "UPDATE users SET " +
		"updated_at=?, " +
		"first_name=?, " +
		"last_name=?, " +
		"middle_name=?, " +
		"email=?, " +
		"phone=?, " +
		"is_guardian=?, " +
		"account_id=?, " +
		"notes=?, " +
		"school=?, " +
		"graduation_year=? " +
		"WHERE id=?"
	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()
	execResult, err := stmt.Exec(
		now,
		user.FirstName,
		user.LastName,
		user.MiddleName,
		user.Email,
		user.Phone,
		user.IsGuardian,
		user.AccountId,
		user.Notes,
		user.School,
		user.GraduationYear,
		id)
	if err != nil {
		return err
	}
	return utils.HandleSqlExecResult(execResult, 1, "user was not updated")
}

func (ur *userRepo) Delete(id uint) error {
	statement := "DELETE FROM users WHERE id=?"
	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	execResult, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	return utils.HandleSqlExecResult(execResult, 1, "user was not deleted")
}

// For Tests Only
func CreateTestUserRepo(db *sql.DB) UserRepoInterface {
	ur := &userRepo{}
	ur.Initialize(db)
	return ur
}
