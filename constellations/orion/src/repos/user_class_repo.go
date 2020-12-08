package repos

import (
	"context"
	"database/sql"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos/dbTx"
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

	tx := dbTx.New(ur.db)
	statement := tx.CreateStmtSelectUserClassByClasId()
	userClasses, err := tx.SelectManyUserClasses(statement, classId)
	if err != nil {
		return nil, err
	}
	return userClasses, nil
}

func (ur *userClassRepo) SelectByUserId(ctx context.Context, userId uint) ([]domains.UserClass, error) {
	utils.LogWithContext(ctx, "userClassRepo.SelectByUserId", logger.Fields{"userId": userId})

	tx := dbTx.New(ur.db)
	statement := tx.CreateStmtSelectUserClassByUserId()
	userClasses, err := tx.SelectManyUserClasses(statement, userId)
	if err != nil {
		return nil, err
	}
	return userClasses, nil
}

func (ur *userClassRepo) SelectByUserAndClass(ctx context.Context, userId uint, classId string) (domains.UserClass, error) {
	utils.LogWithContext(ctx, "userClassRepo.SelectByUserAndClass", logger.Fields{
		"userId":  userId,
		"classId": classId,
	})

	tx := dbTx.New(ur.db)
	statement := tx.CreateStmtSelectUserClassByBothIds()
	userClass, err := tx.SelectOneUserClass(statement, userId, classId)
	if err != nil {
		return domains.UserClass{}, err
	}
	return userClass, nil
}

func (ur *userClassRepo) SelectByNew(ctx context.Context) ([]domains.UserClass, error) {
	utils.LogWithContext(ctx, "userClassRepo.SelectByNew", logger.Fields{})
	now := time.Now().UTC()
	week := time.Hour * 24 * 7
	lastWeek := now.Add(-week)

	tx := dbTx.New(ur.db)
	statement := tx.CreateStmtSelectNewUserClass()
	userClasses, err := tx.SelectManyUserClasses(statement, lastWeek)
	if err != nil {
		return nil, err
	}
	return userClasses, nil
}

func (ur *userClassRepo) Insert(ctx context.Context, userClass domains.UserClass) (uint, error) {
	utils.LogWithContext(ctx, "userClassRepo.Insert", logger.Fields{"userClass": userClass})

	tx, err := dbTx.Begin(ur.db)
	if err != nil {
		return 0, err
	}
	userAfhId, err := tx.InsertUserClass(userClass)
	if err != nil {
		return 0, err
	}
	if err := tx.Commit(); err != nil {
		return 0, err
	}
	return userAfhId, nil
}

func (ur *userClassRepo) Update(ctx context.Context, id uint, userClass domains.UserClass) error {
	utils.LogWithContext(ctx, "userClassRepo.Update", logger.Fields{"userClass": userClass})
	tx, err := dbTx.Begin(ur.db)
	if err != nil {
		return err
	}
	if err := tx.UpdateUserClassById(id, userClass); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (ur *userClassRepo) Delete(ctx context.Context, id uint) error {
	utils.LogWithContext(ctx, "userClassRepo.Delete", logger.Fields{"id": id})

	tx, err := dbTx.Begin(ur.db)
	if err != nil {
		return err
	}
	if err := tx.DeleteUserClass(id); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

// For Tests Only
func CreateTestUserClassRepo(ctx context.Context, db *sql.DB) UserClassRepoInterface {
	ur := &userClassRepo{}
	ur.Initialize(ctx, db)
	return ur
}
