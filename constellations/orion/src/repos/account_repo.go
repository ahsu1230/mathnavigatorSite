package repos

import (
	"context"
	"database/sql"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos/dbTx"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos/utils"
)

// Global variable
var AccountRepo AccountRepoInterface = &accountRepo{}

// Implements interface userRepoInterface
type accountRepo struct {
	db *sql.DB // golang native db connection
}

// Interface to implement
type AccountRepoInterface interface {
	Initialize(context.Context, *sql.DB)
	SelectById(context.Context, uint) (domains.Account, error)
	SelectByPrimaryEmail(context.Context, string) (domains.Account, error)
	SelectAllNegativeBalances(context.Context) ([]domains.AccountBalance, error)
	InsertWithUser(context.Context, domains.Account, domains.User) (uint, error)
	Update(context.Context, uint, domains.Account) error
	Delete(context.Context, uint) error
	FullDelete(context.Context, uint) error
}

func (acc *accountRepo) Initialize(ctx context.Context, db *sql.DB) {
	utils.LogWithContext(ctx, "accountRepo.Initialize", logger.Fields{})
	acc.db = db
}

func (acc *accountRepo) SelectById(ctx context.Context, id uint) (domains.Account, error) {
	utils.LogWithContext(ctx, "accountRepo.SelectById", logger.Fields{"id": id})
	tx := dbTx.New(acc.db)
	account, err := tx.SelectOneAccount(tx.CreateStmtSelectAccountById(), id)
	if err != nil {
		return domains.Account{}, err
	}
	return account, nil
}

func (acc *accountRepo) SelectByPrimaryEmail(ctx context.Context, primaryEmail string) (domains.Account, error) {
	utils.LogWithContext(ctx, "accountRepo.SelectByPrimaryEmail",
		logger.Fields{"primaryEmail": primaryEmail},
	)
	tx := dbTx.New(acc.db)
	account, err := tx.SelectOneAccount(tx.CreateStmtSelectAccountByEmail(), primaryEmail)
	if err != nil {
		return domains.Account{}, err
	}
	return account, nil
}

func (acc *accountRepo) SelectAllNegativeBalances(ctx context.Context) ([]domains.AccountBalance, error) {
	utils.LogWithContext(ctx, "accountRepo.SelectAllNegativeBalances", logger.Fields{})
	tx := dbTx.New(acc.db)
	accounts, err := tx.SelectManyAccountBalances(tx.CreateStmtSelectAccountByNegativeBalance())
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (acc *accountRepo) InsertWithUser(ctx context.Context, account domains.Account, user domains.User) (uint, error) {
	utils.LogWithContext(ctx, "accountRepo.Insert", logger.Fields{
		"account": account,
		"user":    user,
	})

	tx, err := dbTx.Begin(acc.db)
	if err != nil {
		return 0, err
	}

	accountId, err := tx.InsertAccount(account)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	user.AccountId = accountId
	_, err = tx.InsertUser(user)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if err := tx.Commit(); err != nil {
		return 0, err
	}
	return uint(accountId), nil
}

func (acc *accountRepo) Update(ctx context.Context, id uint, account domains.Account) error {
	utils.LogWithContext(ctx, "accountRepo.Update", logger.Fields{
		"id":      id,
		"account": account,
	})

	tx, err := dbTx.Begin(acc.db)
	if err != nil {
		return err
	}
	if err := tx.UpdateAccountById(id, account); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (acc *accountRepo) Delete(ctx context.Context, id uint) error {
	utils.LogWithContext(ctx, "accountRepo.Delete", logger.Fields{"id": id})
	tx, err := dbTx.Begin(acc.db)
	if err != nil {
		return err
	}
	if err := tx.DeleteAccount(id); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

// A Delete method which cascadingly deletes
// - Search for all users associated with account. For each user,
// - Deletes all userClasses
// - Deletes all userAfhs
// - Deletes the user
// - Finally, delete the account
func (acc *accountRepo) FullDelete(ctx context.Context, accountId uint) error {
	utils.LogWithContext(ctx, "accountRepo.FullDelete", logger.Fields{"id": accountId})
	tx, err := dbTx.Begin(acc.db)
	if err != nil {
		return err
	}

	statement := tx.CreateStmtSelectUsersByAccountId()
	users, err := tx.SelectManyUsers(statement, accountId)
	if err != nil {
		return err
	}

	for i := 0; i < len(users); i++ {
		user := users[i]
		userId := user.Id
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
	}
	if err := tx.DeleteAccount(accountId); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

// For Tests Only
func CreateTestAccountRepo(ctx context.Context, db *sql.DB) AccountRepoInterface {
	acc := &accountRepo{}
	acc.Initialize(ctx, db)
	return acc
}
