package repos

import (
	"database/sql"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"

	"time"
)

// Global variable
var UserClassesRepo UserClassesRepoInterface = &userClassesRepo{}

// Implements interface userRepoInterface
type userClassesRepo struct {
	db *sql.DB // golang native db connection
}

// Interface to implement
type UserClassesRepoInterface interface {
	Initialize(db *sql.DB)
	SelectByClassId(string) ([]domains.UserClasses, error)
	SelectByUserId(uint) ([]domains.UserClasses, error)
	SelectByUserAndClass(uint, string) (domains.UserClasses, error)
	SelectByNew() ([]domains.UserClasses, error)
	Insert(domains.UserClasses) error
	Update(uint, domains.UserClasses) error
	Delete(uint) error
}

func (ur *userClassesRepo) Initialize(db *sql.DB) {
	logger.Debug("Initialize UserClassRepo", logger.Fields{})
	ur.db = db
}

func (ur *userClassesRepo) SelectByClassId(classId string) ([]domains.UserClasses, error) {
	results := make([]domains.UserClasses, 0)

	statement := "SELECT * FROM user_classes WHERE class_id=?"
	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		return nil, appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()
	rows, err := stmt.Query(domains.NewNullString(classId))
	if err != nil {
		return nil, appErrors.WrapDbQuery(err, statement, classId)
	}
	defer rows.Close()

	for rows.Next() {
		var userClasses domains.UserClasses
		if errScan := rows.Scan(
			&userClasses.Id,
			&userClasses.CreatedAt,
			&userClasses.UpdatedAt,
			&userClasses.DeletedAt,
			&userClasses.UserId,
			&userClasses.ClassId,
			&userClasses.AccountId,
			&userClasses.State); errScan != nil {
			return results, errScan
		}
		results = append(results, userClasses)
	}
	return results, nil
}

func (ur *userClassesRepo) SelectByUserId(userId uint) ([]domains.UserClasses, error) {
	results := make([]domains.UserClasses, 0)

	statement := "SELECT * FROM user_classes WHERE user_id=?"
	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		return nil, appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()
	rows, err := stmt.Query(domains.NewNullUint(userId))
	if err != nil {
		return nil, appErrors.WrapDbQuery(err, statement, userId)
	}
	defer rows.Close()

	for rows.Next() {
		var userClasses domains.UserClasses
		if errScan := rows.Scan(
			&userClasses.Id,
			&userClasses.CreatedAt,
			&userClasses.UpdatedAt,
			&userClasses.DeletedAt,
			&userClasses.UserId,
			&userClasses.ClassId,
			&userClasses.AccountId,
			&userClasses.State); errScan != nil {
			return results, errScan
		}
		results = append(results, userClasses)
	}
	return results, nil
}

func (ur *userClassesRepo) SelectByUserAndClass(userId uint, classId string) (domains.UserClasses, error) {
	statement := "SELECT * FROM user_classes WHERE user_id=? AND class_id=?"
	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		err = appErrors.WrapDbPrepare(err, statement)
		return domains.UserClasses{}, err
	}
	defer stmt.Close()

	var userClasses domains.UserClasses
	row := stmt.QueryRow(userId, classId)
	if err = row.Scan(
		&userClasses.Id,
		&userClasses.CreatedAt,
		&userClasses.UpdatedAt,
		&userClasses.DeletedAt,
		&userClasses.UserId,
		&userClasses.ClassId,
		&userClasses.AccountId,
		&userClasses.State); err != nil {
		err = appErrors.WrapDbExec(err, statement, userId, classId)
		return domains.UserClasses{}, err
	}
	return userClasses, nil
}

func (ur *userClassesRepo) SelectByNew() ([]domains.UserClasses, error) {
	results := make([]domains.UserClasses, 0)

	now := time.Now().UTC()
	week := time.Hour * 24 * 7
	lastWeek := now.Add(-week)
	stmt, err := ur.db.Prepare("SELECT * FROM user_classes WHERE created_at>=?")

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
		var userClasses domains.UserClasses
		if errScan := rows.Scan(
			&userClasses.Id,
			&userClasses.CreatedAt,
			&userClasses.UpdatedAt,
			&userClasses.DeletedAt,
			&userClasses.UserId,
			&userClasses.ClassId,
			&userClasses.AccountId,
			&userClasses.State); errScan != nil {
			return results, errScan
		}
		results = append(results, userClasses)
	}
	return results, nil
}

func (ur *userClassesRepo) Insert(userClasses domains.UserClasses) error {
	statement := "INSERT INTO user_classes (" +
		"created_at, " +
		"updated_at, " +
		"user_id," +
		"class_id," +
		"account_id," +
		"state" +
		") VALUES (?, ?, ?, ?, ?, ?)"

	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		return appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	now := time.Now().UTC()
	execResult, err := stmt.Exec(
		now,
		now,
		userClasses.UserId,
		userClasses.ClassId,
		userClasses.AccountId,
		userClasses.State,
	)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, userClasses)
	}
	return appErrors.ValidateDbResult(execResult, 1, "userClasses was not inserted")
}

func (ur *userClassesRepo) Update(id uint, userClasses domains.UserClasses) error {
	statement := "UPDATE user_classes SET " +
		"updated_at=?, " +
		"user_id=?, " +
		"class_id=?, " +
		"account_id=?, " +
		"state=? " +
		"WHERE id=?"
	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		return appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	now := time.Now().UTC()
	execResult, err := stmt.Exec(
		now,
		userClasses.UserId,
		userClasses.ClassId,
		userClasses.AccountId,
		userClasses.State,
		id)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, userClasses, id)
	}
	return appErrors.ValidateDbResult(execResult, 1, "userClasses was not updated")
}

func (ur *userClassesRepo) Delete(id uint) error {
	statement := "DELETE FROM user_classes WHERE id=?"
	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		return appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	execResult, err := stmt.Exec(id)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, id)
	}
	return appErrors.ValidateDbResult(execResult, 1, "userClasses was not deleted")
}

// For Tests Only
func CreateTestUserClassesRepo(db *sql.DB) UserClassesRepoInterface {
	ur := &userClassesRepo{}
	ur.Initialize(db)
	return ur
}
