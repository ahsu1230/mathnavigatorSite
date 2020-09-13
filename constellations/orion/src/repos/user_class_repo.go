package repos

import (
	"context"
	"database/sql"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
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
	Initialize(context.Context, *sql.DB)
	SelectByClassId(context.Context, string) ([]domains.UserClass, error)
	SelectByUserId(context.Context, uint) ([]domains.UserClass, error)
	SelectByUserAndClass(context.Context, uint, string) (domains.UserClass, error)
	SelectByNew(context.Context) ([]domains.UserClass, error)
	Insert(context.Context, domains.UserClass) (uint, error)
	Update(context.Context, uint, domains.UserClass) error
	Delete(context.Context, uint) error
}

func (ur *userClassRepo) Initialize(ctx context.Context, db *sql.DB) {
	utils.LogWithContext(ctx, "userClassRepo.Initialize", logger.Fields{})
	ur.db = db
}

func (ur *userClassRepo) SelectByClassId(ctx context.Context, classId string) ([]domains.UserClass, error) {
	utils.LogWithContext(ctx, "userRepo.SelectByClassId", logger.Fields{"classId": classId})
	results := make([]domains.UserClass, 0)

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
		var userClasses domains.UserClass
		if errScan := rows.Scan(
			&userClasses.Id,
			&userClasses.CreatedAt,
			&userClasses.UpdatedAt,
			&userClasses.DeletedAt,
			&userClasses.ClassId,
			&userClasses.UserId,
			&userClasses.AccountId,
			&userClasses.State); errScan != nil {
			return results, errScan
		}
		results = append(results, userClasses)
	}
	return results, nil
}

func (ur *userClassRepo) SelectByUserId(ctx context.Context, userId uint) ([]domains.UserClass, error) {
	utils.LogWithContext(ctx, "userClassRepo.SelectByUserId", logger.Fields{"userId": userId})
	results := make([]domains.UserClass, 0)

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
		var userClasses domains.UserClass
		if errScan := rows.Scan(
			&userClasses.Id,
			&userClasses.CreatedAt,
			&userClasses.UpdatedAt,
			&userClasses.DeletedAt,
			&userClasses.ClassId,
			&userClasses.UserId,
			&userClasses.AccountId,
			&userClasses.State); errScan != nil {
			return results, errScan
		}
		results = append(results, userClasses)
	}
	return results, nil
}

func (ur *userClassRepo) SelectByUserAndClass(ctx context.Context, userId uint, classId string) (domains.UserClass, error) {
	utils.LogWithContext(ctx, "userClassRepo.SelectByUserAndClass", logger.Fields{
		"userId":  userId,
		"classId": classId})
	statement := "SELECT * FROM user_classes WHERE user_id=? AND class_id=?"
	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		err = appErrors.WrapDbPrepare(err, statement)
		return domains.UserClass{}, err
	}
	defer stmt.Close()

	var userClasses domains.UserClass
	row := stmt.QueryRow(userId, classId)
	if err = row.Scan(
		&userClasses.Id,
		&userClasses.CreatedAt,
		&userClasses.UpdatedAt,
		&userClasses.DeletedAt,
		&userClasses.ClassId,
		&userClasses.UserId,
		&userClasses.AccountId,
		&userClasses.State); err != nil {
		err = appErrors.WrapDbExec(err, statement, userId, classId)
		return domains.UserClass{}, err
	}
	return userClasses, nil
}

func (ur *userClassRepo) SelectByNew(ctx context.Context) ([]domains.UserClass, error) {
	utils.LogWithContext(ctx, "userClassRepo.SelectByNew", logger.Fields{})
	results := make([]domains.UserClass, 0)

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
		var userClasses domains.UserClass
		if errScan := rows.Scan(
			&userClasses.Id,
			&userClasses.CreatedAt,
			&userClasses.UpdatedAt,
			&userClasses.DeletedAt,
			&userClasses.ClassId,
			&userClasses.UserId,
			&userClasses.AccountId,
			&userClasses.State); errScan != nil {
			return results, errScan
		}
		results = append(results, userClasses)
	}
	return results, nil
}

func (ur *userClassRepo) Insert(ctx context.Context, userClasses domains.UserClass) (uint, error) {
	utils.LogWithContext(ctx, "userClassRepo.Insert", logger.Fields{"userClass": userClasses})
	statement := "INSERT INTO user_classes (" +
		"created_at, " +
		"updated_at, " +
		"class_id," +
		"user_id," +
		"account_id," +
		"state" +
		") VALUES (?, ?, ?, ?, ?, ?)"

	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		return 0, appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	now := time.Now().UTC()
	execResult, err := stmt.Exec(
		now,
		now,
		userClasses.ClassId,
		userClasses.UserId,
		userClasses.AccountId,
		userClasses.State,
	)
	if err != nil {
		return 0, appErrors.WrapDbExec(err, statement, userClasses)
	}

	rowId, err := execResult.LastInsertId()
	if err != nil {
		return 0, appErrors.WrapSQLBadInsertResult(err)
	}
	return uint(rowId), appErrors.ValidateDbResult(execResult, 1, "userClasses was not inserted")
}

func (ur *userClassRepo) Update(ctx context.Context, id uint, userClasses domains.UserClass) error {
	utils.LogWithContext(ctx, "userClassRepo.Update", logger.Fields{"userClass": userClasses})
	statement := "UPDATE user_classes SET " +
		"updated_at=?, " +
		"class_id=?, " +
		"user_id=?, " +
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
		userClasses.ClassId,
		userClasses.UserId,
		userClasses.AccountId,
		userClasses.State,
		id)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, userClasses, id)
	}
	return appErrors.ValidateDbResult(execResult, 1, "userClasses was not updated")
}

func (ur *userClassRepo) Delete(ctx context.Context, id uint) error {
	utils.LogWithContext(ctx, "userClassRepo.Delete", logger.Fields{"id": id})
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
func CreateTestUserClassRepo(ctx context.Context, db *sql.DB) UserClassRepoInterface {
	ur := &userClassRepo{}
	ur.Initialize(ctx, db)
	return ur
}
