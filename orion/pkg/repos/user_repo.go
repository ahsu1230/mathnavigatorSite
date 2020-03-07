package repos

import (
	"database/sql"
	"time"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
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
	SelectById(string) (domains.User, error)
	Insert(domains.User) error
	Update(string, domains.User) error
	Delete(string) error
}

func (pr *userRepo) Initialize(db *sql.DB) {
	pr.db = db
}

func (pr *userRepo) SelectAll() ([]domains.User, error) {
	results := make([]domains.User, 0)

	stmt, err := pr.db.Prepare("SELECT * FROM users")
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
			&user.GuardianId);
			errScan != nil {
			return results, errScan
		}
		results = append(results, user)
	}
	return results, nil
}

func (pr *userRepo) SelectById(id string) (domains.User, error) {
	statement := "SELECT * FROM users WHERE id=?"
	stmt, err := pr.db.Prepare(statement)
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

func (pr *userRepo) Insert(user domains.User) error {
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

	stmt, err := pr.db.Prepare(statement)
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
		user.GuardianId)
	if err != nil {
		return err
	}
	return handleSqlExecResult(execResult, 1, "user was not inserted")
}

func (pr *userRepo) Update(id string, user domains.User) error {
	statement := "UPDATE users SET " +
		"updated_at=?, " +
		"id=?, " +
		"first_name=?, " +
		"last_name=?, " +
		"middle_name=?, " +
		"email=?, " +
		"phone=?, " +
		"is_guardian=?, " +
		"guardian_id=? " +
		"WHERE id=?"
	stmt, err := pr.db.Prepare(statement)
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
		user.GuardianId,
		id)
	if err != nil {
		return err
	}
	return handleSqlExecResult(execResult, 1, "user was not updated")
}

func (pr *userRepo) Delete(id string) error {
	statement := "DELETE FROM users WHERE id=?"
	stmt, err := pr.db.Prepare(statement)
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
	pr := &userRepo{}
	pr.Initialize(db)
	return pr
}
