package repos

import (
	"database/sql"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos/utils"
	"time"
)

// Global variable
var UserClassRepo UserClassRepoInterface = &userClassRepo{}

// Implements interface userRepoInterface
type userClassRepo struct {
	db *sql.DB // golang native db connection
}

// Interface to implement
type UserClassRepoInterface interface {
	Initialize(db *sql.DB)
	SelectByClassId(string) ([]domains.UserClass, error)
	SelectByUserId(uint) ([]domains.UserClass, error)
	SelectByUserAndClass(uint, string) (domains.UserClass, error)
	Insert(domains.UserClass) error
	Update(uint, domains.UserClass) error
	Delete(uint) error
}

func (ur *userClassRepo) Initialize(db *sql.DB) {
	ur.db = db
}

func (ur *userClassRepo) SelectByClassId(classId string) ([]domains.UserClass, error) {
	results := make([]domains.UserClass, 0)

	stmt, err := ur.db.Prepare("SELECT * FROM userclass WHERE class_id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(domains.NewNullString(classId))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var userClass domains.UserClass
		if errScan := rows.Scan(
			&userClass.Id,
			&userClass.CreatedAt,
			&userClass.UpdatedAt,
			&userClass.DeletedAt,
			&userClass.UserId,
			&userClass.ClassId,
			&userClass.AccountId,
			&userClass.State); errScan != nil {
			return results, errScan
		}
		results = append(results, userClass)
	}
	return results, nil
}

func (ur *userClassRepo) SelectByUserId(userId uint) ([]domains.UserClass, error) {
	results := make([]domains.UserClass, 0)

	stmt, err := ur.db.Prepare("SELECT * FROM userclass WHERE user_id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(domains.NewNullUint(userId))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var userClass domains.UserClass
		if errScan := rows.Scan(
			&userClass.Id,
			&userClass.CreatedAt,
			&userClass.UpdatedAt,
			&userClass.DeletedAt,
			&userClass.UserId,
			&userClass.ClassId,
			&userClass.AccountId,
			&userClass.State); errScan != nil {
			return results, errScan
		}
		results = append(results, userClass)
	}
	return results, nil
}

func (ur *userClassRepo) SelectByUserAndClass(userId uint, classId string) (domains.UserClass, error) {
	statement := "SELECT * FROM userclass WHERE user_id=? AND class_id=?"
	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		return domains.UserClass{}, err
	}
	defer stmt.Close()

	var userClass domains.UserClass
	row := stmt.QueryRow(userId, classId)
	errScan := row.Scan(
		&userClass.Id,
		&userClass.CreatedAt,
		&userClass.UpdatedAt,
		&userClass.DeletedAt,
		&userClass.UserId,
		&userClass.ClassId,
		&userClass.AccountId,
		&userClass.State)
	return userClass, errScan
}

func (ur *userClassRepo) Insert(userClass domains.UserClass) error {
	statement := "INSERT INTO userclass (" +
		"created_at, " +
		"updated_at, " +
		"user_id," +
		"class_id," +
		"account_id," +
		"state" +
		") VALUES (?, ?, ?, ?, ?, ?)"

	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()
	execResult, err := stmt.Exec(
		now,
		now,
		userClass.UserId,
		userClass.ClassId,
		userClass.AccountId,
		userClass.State,
	)
	if err != nil {
		return err
	}
	return utils.HandleSqlExecResult(execResult, 1, "userClass was not inserted")
}

func (ur *userClassRepo) Update(id uint, userClass domains.UserClass) error {
	statement := "UPDATE userclass SET " +
		"updated_at=?, " +
		"user_id=?, " +
		"class_id=?, " +
		"account_id=?, " +
		"state=? " +
		"WHERE id=?"
	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()
	execResult, err := stmt.Exec(
		now,
		userClass.UserId,
		userClass.ClassId,
		userClass.AccountId,
		userClass.State,
		id)
	if err != nil {
		return err
	}
	return utils.HandleSqlExecResult(execResult, 1, "userClass was not updated")
}

func (ur *userClassRepo) Delete(id uint) error {
	statement := "DELETE FROM userclass WHERE id=?"
	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	execResult, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	return utils.HandleSqlExecResult(execResult, 1, "userClass was not deleted")
}

// For Tests Only
func CreateTestUserClassRepo(db *sql.DB) UserClassRepoInterface {
	ur := &userClassRepo{}
	ur.Initialize(db)
	return ur
}
