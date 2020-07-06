package repos

import (
	"database/sql"
	"time"
	//"fmt"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos/utils"
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
	//SearchAll(string) ([]domains.User,error)
	SelectAll(string, int, int) ([]domains.User, error)
	SelectById(uint) (domains.User, error)
	SelectByAccountId(uint) ([]domains.User, error)
	Insert(domains.User) error
	Update(uint, domains.User) error
	Delete(uint) error
}

func (ur *userRepo) Initialize(db *sql.DB) {
	ur.db = db
}

//TODO

// func (ur *userRepo) SearchAll(search string) ([]domains.User, error) {
// 	results := make([]domains.User, 0)

// 	//var query string

// 	query := fmt.Sprintf("SELECT * FROM users WHERE first_name CONTAINS %s",search)

// 	//query = "SELECT * FROM users WHERE first_name CONTAINS ?" //? OR middle_name contains ? OR last_name contains ? OR email contains ? "
// 	stmt, err := ur.db.Prepare(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer stmt.Close()

// 	return results, nil
// }

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
			&user.Notes); errScan != nil {
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
		&user.Notes)
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
			&user.Notes); errScan != nil {
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
		"notes" +
		") VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

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
		"notes=? " +
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
