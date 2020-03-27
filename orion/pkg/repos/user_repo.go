package repos

import (
	"database/sql"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/sql_helper"
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
	SelectAll() ([]domains.User, error)
	SelectById(uint) (domains.User, error)
	Insert(domains.User) error
	Update(uint, domains.User) error
	Delete(uint) error
}

func (ur *userRepo) Initialize(db *sql.DB) {
	ur.db = db
}

func (ur *userRepo) SelectAll() ([]domains.User, error) {
	results := make([]domains.User, 0)

	stmt, err := ur.db.Prepare("SELECT * FROM users")
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
			&user.GuardianId); errScan != nil {
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
		&user.GuardianId)
	return user, errScan
}

func (ur *userRepo) Insert(user domains.User) error {
	statement := "INSERT INTO users (" +
		"created_at, " +
		"updated_at, " +
		"first_name, " +
		"last_name" +
		"middle_name, " +
		"email" +
		"phone, " +
		"is_guardian" +
		"guardian_id" +
		") VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"

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
		sql.NullString{String: user.MiddleName, Valid: user.MiddleName != ""},
		user.Email,
		user.Phone,
		user.IsGuardian,
		sql_helper.NullUint{Uint: user.GuardianId, Valid: user.GuardianId != 0})
	if err != nil {
		return err
	}
	return handleSqlExecResult(execResult, 1, "user was not inserted")
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
		"guardian_id=? " +
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
		sql.NullString{String: user.MiddleName, Valid: user.MiddleName != ""},
		user.Email,
		user.Phone,
		user.IsGuardian,
		sql_helper.NullUint{Uint: user.GuardianId, Valid: user.GuardianId != 0},
		id)
	if err != nil {
		return err
	}
	return handleSqlExecResult(execResult, 1, "user was not updated")
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
	return handleSqlExecResult(execResult, 1, "user was not deleted")
}

// For Tests Only
func CreateTestUserRepo(db *sql.DB) UserRepoInterface {
	ur := &userRepo{}
	ur.Initialize(db)
	return ur
}
