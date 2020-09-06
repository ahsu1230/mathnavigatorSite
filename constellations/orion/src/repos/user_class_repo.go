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
var UserClassesRepo UserClassesRepoInterface = &userClassesRepo{}

// Implements interface userRepoInterface
type userClassesRepo struct {
	db *sql.DB // golang native db connection
}

// Interface to implement
type UserClassesRepoInterface interface {
	Initialize(context.Context, *sql.DB)
	SelectByClassId(context.Context, string) ([]domains.UserClasses, error)
	SelectByUserId(context.Context, uint) ([]domains.UserClasses, error)
	SelectByUserAndClass(context.Context, uint, string) (domains.UserClasses, error)
	SelectByNew(context.Context) ([]domains.UserClasses, error)
	Insert(context.Context, domains.UserClasses) error
	Update(context.Context, uint, domains.UserClasses) error
	Delete(context.Context, uint) error
}

func (ur *userClassesRepo) Initialize(ctx context.Context, db *sql.DB) {
	utils.LogWithContext(ctx, "userClassRepo.Initialize", logger.Fields{})
	ur.db = db
}

func (ur *userClassesRepo) SelectByClassId(ctx context.Context, classId string) ([]domains.UserClasses, error) {
	utils.LogWithContext(ctx, "userRepo.SelectByClassId", logger.Fields{"classId": classId})
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

func (ur *userClassesRepo) SelectByUserId(ctx context.Context, userId uint) ([]domains.UserClasses, error) {
	utils.LogWithContext(ctx, "userClassRepo.SelectByUserId", logger.Fields{"userId": userId})
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

func (ur *userClassesRepo) SelectByUserAndClass(ctx context.Context, userId uint, classId string) (domains.UserClasses, error) {
	utils.LogWithContext(ctx, "userClassRepo.SelectByUserAndClass", logger.Fields{
		"userId":  userId,
		"classId": classId})
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

func (ur *userClassesRepo) SelectByNew(ctx context.Context) ([]domains.UserClasses, error) {
	utils.LogWithContext(ctx, "userClassRepo.SelectByNew", logger.Fields{})
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

func (ur *userClassesRepo) Insert(ctx context.Context, userClasses domains.UserClasses) error {
	utils.LogWithContext(ctx, "userClassRepo.Insert", logger.Fields{"userClass": userClasses})
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

func (ur *userClassesRepo) Update(ctx context.Context, id uint, userClasses domains.UserClasses) error {
	utils.LogWithContext(ctx, "userClassRepo.Update", logger.Fields{"userClass": userClasses})
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

func (ur *userClassesRepo) Delete(ctx context.Context, id uint) error {
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
func CreateTestUserClassesRepo(ctx context.Context, db *sql.DB) UserClassesRepoInterface {
	ur := &userClassesRepo{}
	ur.Initialize(ctx, db)
	return ur
}
