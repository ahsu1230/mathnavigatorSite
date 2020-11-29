package repos

import (
	"context"
	"database/sql"
	"time"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos/dbTx"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos/utils"
)

// Global variable
var UserAfhRepo UserAfhRepoInterface = &userAfhRepo{}

// Implements interface userAfhRepoInterface
type userAfhRepo struct {
	db *sql.DB
}

// Interface to implement
type UserAfhRepoInterface interface {
	Initialize(context.Context, *sql.DB)
	SelectByUserId(context.Context, uint) ([]domains.UserAfh, error)
	SelectByAfhId(context.Context, uint) ([]domains.UserAfh, error)
	SelectByBothIds(context.Context, uint, uint) (domains.UserAfh, error)
	SelectByNew(context.Context) ([]domains.UserAfh, error)
	Insert(context.Context, domains.UserAfh) (uint, error)
	Update(context.Context, uint, domains.UserAfh) error
	Delete(context.Context, uint) error
}

func (ur *userAfhRepo) Initialize(ctx context.Context, db *sql.DB) {
	utils.LogWithContext(ctx, "userAfhRepo.Initialize", logger.Fields{})
	ur.db = db
}

func (ur *userAfhRepo) SelectByUserId(ctx context.Context, userId uint) ([]domains.UserAfh, error) {
	utils.LogWithContext(ctx, "userAfhRepo.SelectByUserId", logger.Fields{"userId": userId})

	tx := dbTx.New(ur.db)
	statement := tx.CreateStmtSelectUserAfhByUserId()
	userAfhs, err := tx.SelectManyUserAfhs(statement, userId)
	if err != nil {
		return nil, err
	}
	return userAfhs, nil
}

func (ur *userAfhRepo) SelectByAfhId(ctx context.Context, afhId uint) ([]domains.UserAfh, error) {
	utils.LogWithContext(ctx, "userAfhRepo.SelectByAfhId", logger.Fields{"afhId": afhId})

	tx := dbTx.New(ur.db)
	statement := tx.CreateStmtSelectUserAfhByAfhId()
	userAfhs, err := tx.SelectManyUserAfhs(statement, afhId)
	if err != nil {
		return nil, err
	}
	return userAfhs, nil
}

func (ur *userAfhRepo) SelectByBothIds(ctx context.Context, userId, afhId uint) (domains.UserAfh, error) {
	utils.LogWithContext(ctx, "userAfhRepo.SelectByBothIds", logger.Fields{"userId": userId, "afhId": afhId})

	tx := dbTx.New(ur.db)
	statement := tx.CreateStmtSelectUserAfhByBothIds()
	userAfh, err := tx.SelectOneUserAfh(statement, userId, afhId)
	if err != nil {
		return domains.UserAfh{}, err
	}
	return userAfh, nil
}
func (ur *userAfhRepo) SelectByNew(ctx context.Context) ([]domains.UserAfh, error) {
	utils.LogWithContext(ctx, "userAfhRepo.SelectByNew", logger.Fields{})

	now := time.Now().UTC()
	week := time.Hour * 24 * 7
	lastWeek := now.Add(-week)

	tx := dbTx.New(ur.db)
	statement := tx.CreateStmtSelectNewUserAfh()
	userAfhs, err := tx.SelectManyUserAfhs(statement, lastWeek)
	if err != nil {
		return nil, err
	}
	return userAfhs, nil
}

func (ur *userAfhRepo) Insert(ctx context.Context, userAfh domains.UserAfh) (uint, error) {
	utils.LogWithContext(ctx, "userAfhRepo.Insert", logger.Fields{"userAfh": userAfh})
	tx, err := dbTx.Begin(ur.db)
	if err != nil {
		return 0, err
	}
	userAfhId, err := tx.InsertUserAfh(userAfh)
	if err != nil {
		return 0, err
	}
	if err := tx.Commit(); err != nil {
		return 0, err
	}
	return userAfhId, nil
}

func (ur *userAfhRepo) Update(ctx context.Context, id uint, userAfh domains.UserAfh) error {
	utils.LogWithContext(ctx, "userAfhRepo.Update", logger.Fields{"userAfh": userAfh})
	tx, err := dbTx.Begin(ur.db)
	if err != nil {
		return err
	}
	if err := tx.UpdateUserAfhById(id, userAfh); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (ur *userAfhRepo) Delete(ctx context.Context, id uint) error {
	utils.LogWithContext(ctx, "userAfhRepo.Delete", logger.Fields{"id": id})

	tx, err := dbTx.Begin(ur.db)
	if err != nil {
		return err
	}
	if err := tx.DeleteUserAfh(id); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

// For Tests Only
func CreateTestUserAfhRepo(ctx context.Context, db *sql.DB) UserAfhRepoInterface {
	ur := &userAfhRepo{}
	ur.Initialize(ctx, db)
	return ur
}
