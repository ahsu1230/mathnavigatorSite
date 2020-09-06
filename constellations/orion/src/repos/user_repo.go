package repos

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
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
	Initialize(context.Context, *sql.DB)
	SearchUsers(context.Context, string) ([]domains.User, error)
	SelectAll(context.Context, string, int, int) ([]domains.User, error)
	SelectById(context.Context, uint) (domains.User, error)
	SelectByAccountId(context.Context, uint) ([]domains.User, error)
	SelectByNew(context.Context) ([]domains.User, error)
	Insert(context.Context, domains.User) error
	Update(context.Context, uint, domains.User) error
	Delete(context.Context, uint) error
}

func (ur *userRepo) Initialize(ctx context.Context, db *sql.DB) {
	utils.LogWithContext(ctx, "userRepo.Initialize", logger.Fields{})
	ur.db = db
}

func (ur *userRepo) SearchUsers(ctx context.Context, search string) ([]domains.User, error) {
	utils.LogWithContext(ctx, "userRepo.SelectUsers", logger.Fields{"search": search})
	results := make([]domains.User, 0)

	lcSearch := strings.ToLower(search)
	query := fmt.Sprintf("SELECT * FROM users WHERE LOWER(`first_name`) "+
		"LIKE '%%%s%%' OR "+
		"LOWER(`middle_name`) LIKE '%%%s%%' OR "+
		"LOWER(`last_name`) LIKE '%%%s%%' OR "+
		"LOWER(`email`) LIKE '%%%s%%'",
		lcSearch, lcSearch, lcSearch, lcSearch)

	stmt, err := ur.db.Prepare(query)
	if err != nil {
		return nil, appErrors.WrapDbPrepare(err, query)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, appErrors.WrapDbQuery(err, query, search)
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

func (ur *userRepo) SelectAll(ctx context.Context, search string, pageSize, offset int) ([]domains.User, error) {
	utils.LogWithContext(ctx, "userRepo.SelectAll", logger.Fields{
		"search":   search,
		"pageSize": pageSize,
		"offset":   offset,
	})

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
		return nil, appErrors.WrapDbPrepare(err, query)
	}
	defer stmt.Close()

	var rows *sql.Rows
	if getAll {
		rows, err = stmt.Query(pageSize, offset)
	} else {
		rows, err = stmt.Query(search, pageSize, offset)
	}
	if err != nil {
		return nil, appErrors.WrapDbQuery(err, query, search, pageSize, offset)
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

func (ur *userRepo) SelectById(ctx context.Context, id uint) (domains.User, error) {
	utils.LogWithContext(ctx, "userRepo.SelectById", logger.Fields{"id": id})

	statement := "SELECT * FROM users WHERE id=?"
	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		return domains.User{}, appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	var user domains.User
	row := stmt.QueryRow(id)
	if errScan := row.Scan(
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
		return domains.User{}, appErrors.WrapDbQuery(errScan, statement, id)
	}
	return user, nil
}

func (ur *userRepo) SelectByAccountId(ctx context.Context, accountId uint) ([]domains.User, error) {
	utils.LogWithContext(ctx, "userRepo.SelectByAccountId", logger.Fields{"accountId": accountId})
	results := make([]domains.User, 0)

	query := "SELECT * FROM users WHERE account_id=?"
	stmt, err := ur.db.Prepare(query)
	if err != nil {
		return nil, appErrors.WrapDbPrepare(err, query)
	}
	defer stmt.Close()
	rows, err := stmt.Query(domains.NewNullUint(accountId))
	if err != nil {
		return nil, appErrors.WrapDbQuery(err, query, accountId)
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

func (ur *userRepo) SelectByNew(ctx context.Context) ([]domains.User, error) {
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

func (ur *userRepo) Insert(ctx context.Context, user domains.User) error {
	utils.LogWithContext(ctx, "userRepo.Insert", logger.Fields{"user": user})
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
		return appErrors.WrapDbPrepare(err, statement)
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
		return appErrors.WrapDbExec(err, statement, user)
	}
	return appErrors.ValidateDbResult(execResult, 1, "user was not inserted")
}

func (ur *userRepo) Update(ctx context.Context, id uint, user domains.User) error {
	utils.LogWithContext(ctx, "userRepo.Update", logger.Fields{"id": id, "user": user})
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
		return appErrors.WrapDbPrepare(err, statement)
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
		return appErrors.WrapDbExec(err, statement, id, user)
	}
	return appErrors.ValidateDbResult(execResult, 1, "user was not updated")
}

func (ur *userRepo) Delete(ctx context.Context, id uint) error {
	utils.LogWithContext(ctx, "userRepo.Delete", logger.Fields{"id": id})
	statement := "DELETE FROM users WHERE id=?"
	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		return appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	execResult, err := stmt.Exec(id)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, id)
	}
	return appErrors.ValidateDbResult(execResult, 1, "user was not deleted")
}

// For Tests Only
func CreateTestUserRepo(ctx context.Context, db *sql.DB) UserRepoInterface {
	ur := &userRepo{}
	ur.Initialize(ctx, db)
	return ur
}
