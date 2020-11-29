package repos

import (
	"context"
	"database/sql"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos/dbTx"
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
	SelectByEmail(context.Context, string) (domains.User, error)
	SelectByNew(context.Context) ([]domains.User, error)
	Insert(context.Context, domains.User) (uint, error)
	Update(context.Context, uint, domains.User) error
	Delete(context.Context, uint) error
	FullDelete(context.Context, uint) error
}

func (ur *userRepo) Initialize(ctx context.Context, db *sql.DB) {
	utils.LogWithContext(ctx, "userRepo.Initialize", logger.Fields{})
	ur.db = db
}

func (ur *userRepo) SearchUsers(ctx context.Context, search string) ([]domains.User, error) {
	utils.LogWithContext(ctx, "userRepo.SelectUsers", logger.Fields{"search": search})

	tx := dbTx.New(ur.db)

	var query string
	lcSearch := strings.ToLower(search)
	searchTerms := strings.Split(lcSearch, " ")
	if len(searchTerms) == 1 {
		// Generic one term search
		query = tx.CreateStmtSelectUserSearchOneTerm(searchTerms[0])
	} else if len(searchTerms) == 2 {
		// Two term search most likely means (firstName, lastName) search
		query = tx.CreateStmtSelectUserSearchTwoTerms(searchTerms)
	} else {
		// Generic multi-term search
		regexTerm := strings.Join(searchTerms, "|")
		query = tx.CreateStmtSelectUserSearchThreeTerms(regexTerm)

	}
	utils.LogWithContext(ctx, "userRepo.SelectUsers.Query", logger.Fields{"query": query})

	users, err := tx.SelectManyUsers(query)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *userRepo) SelectAll(ctx context.Context, search string, pageSize, offset int) ([]domains.User, error) {
	utils.LogWithContext(ctx, "userRepo.SelectAll", logger.Fields{
		"search":   search,
		"pageSize": pageSize,
		"offset":   offset,
	})

	tx := dbTx.New(ur.db)
	getAll := len(search) == 0
	var query string
	var args []interface{}
	if getAll {
		query = tx.CreateStmtSelectUsersAllWithLimitOffset()
		args = []interface{}{pageSize, offset}
	} else {
		query = tx.CreateStmtSelectUserNamesWithLimitOffset()
		args = []interface{}{search, pageSize, offset}
	}

	users, err := tx.SelectManyUsers(query, args...)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *userRepo) SelectById(ctx context.Context, id uint) (domains.User, error) {
	utils.LogWithContext(ctx, "userRepo.SelectById", logger.Fields{"id": id})
	tx := dbTx.New(ur.db)
	user, err := tx.SelectOneUser(tx.CreateStmtSelectUserById(), id)
	if err != nil {
		return domains.User{}, err
	}
	return user, nil
}

func (ur *userRepo) SelectByAccountId(ctx context.Context, accountId uint) ([]domains.User, error) {
	utils.LogWithContext(ctx, "userRepo.SelectByAccountId", logger.Fields{"accountId": accountId})
	tx := dbTx.New(ur.db)
	users, err := tx.SelectManyUsers(tx.CreateStmtSelectUsersByAccountId(), domains.NewNullUint(accountId))
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *userRepo) SelectByEmail(ctx context.Context, email string) (domains.User, error) {
	utils.LogWithContext(ctx, "userRepo.SelectById", logger.Fields{"email": email})
	tx := dbTx.New(ur.db)
	user, err := tx.SelectOneUser(tx.CreateStmtSelectUserByEmail(), email)
	if err != nil {
		return domains.User{}, err
	}
	return user, nil
}

func (ur *userRepo) SelectByNew(ctx context.Context) ([]domains.User, error) {
	utils.LogWithContext(ctx, "userRepo.SelectByNew", logger.Fields{})
	now := time.Now().UTC()
	week := time.Hour * 24 * 7
	lastWeek := now.Add(-week)

	tx := dbTx.New(ur.db)
	users, err := tx.SelectManyUsers(tx.CreateStmtSelectUsersByNew(), lastWeek)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *userRepo) Insert(ctx context.Context, user domains.User) (uint, error) {
	utils.LogWithContext(ctx, "userRepo.Insert", logger.Fields{"user": user})
	tx, err := dbTx.Begin(ur.db)
	if err != nil {
		return 0, err
	}
	userId, err := tx.InsertUser(user)
	if err != nil {
		return 0, err
	}
	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return userId, nil
}

func (ur *userRepo) Update(ctx context.Context, id uint, user domains.User) error {
	utils.LogWithContext(ctx, "userRepo.Update", logger.Fields{"id": id, "user": user})

	tx, err := dbTx.Begin(ur.db)
	if err != nil {
		return err
	}
	if err := tx.UpdateUserById(id, user); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (ur *userRepo) Delete(ctx context.Context, id uint) error {
	utils.LogWithContext(ctx, "userRepo.Delete", logger.Fields{"id": id})

	tx, err := dbTx.Begin(ur.db)
	if err != nil {
		return err
	}
	if err := tx.DeleteUser(id); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

// A Delete method which cascadingly deletes
// - Deletes all userClasses
// - Deletes all userAfhs
// - Deletes the user
func (ur *userRepo) FullDelete(ctx context.Context, userId uint) error {
	utils.LogWithContext(ctx, "userRepo.Delete", logger.Fields{"id": userId})

	tx, err := dbTx.Begin(ur.db)
	if err != nil {
		return err
	}
	if errDel := tx.DeleteUserClassByUserId(userId); errDel != nil {
		tx.Rollback()
		return errDel
	}
	if errDel := tx.DeleteUserAfhByUserId(userId); errDel != nil {
		tx.Rollback()
		return errDel
	}
	if errDel := tx.DeleteUser(userId); errDel != nil {
		tx.Rollback()
		return errDel
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

// For Tests Only
func CreateTestUserRepo(ctx context.Context, db *sql.DB) UserRepoInterface {
	ur := &userRepo{}
	ur.Initialize(ctx, db)
	return ur
}
