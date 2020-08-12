package repos

import (
	"database/sql"
	"time"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
)

// Global variable
var AccountRepo AccountRepoInterface = &accountRepo{}

// Implements interface userRepoInterface
type accountRepo struct {
	db *sql.DB // golang native db connection
}

// Interface to implement
type AccountRepoInterface interface {
	Initialize(db *sql.DB)
	SelectById(uint) (domains.Account, error)
	SelectByPrimaryEmail(string) (domains.Account, error)
	Insert(domains.Account) error
	Update(uint, domains.Account) error
	Delete(uint) error
}

func (acc *accountRepo) Initialize(db *sql.DB) {
	logger.Debug("Initialize AccountRepo", logger.Fields{})
	acc.db = db
}

func (acc *accountRepo) SelectById(id uint) (domains.Account, error) {
	logger.Info("accountRepo.SelectById", logger.Fields{"id": id})
	statement := "SELECT * FROM accounts WHERE id=?"
	stmt, err := acc.db.Prepare(statement)
	if err != nil {
		err = appErrors.WrapDbPrepare(err, statement)
		return domains.Account{}, err
	}
	defer stmt.Close()

	var account domains.Account
	row := stmt.QueryRow(id)
	if err = row.Scan(
		&account.Id,
		&account.CreatedAt,
		&account.UpdatedAt,
		&account.DeletedAt,
		&account.PrimaryEmail,
		&account.Password,
	); err != nil {
		err = appErrors.WrapDbQuery(err, statement, id)
		return domains.Account{}, err
	}
	return account, err
}

func (acc *accountRepo) SelectByPrimaryEmail(primaryEmail string) (domains.Account, error) {
	logger.Info("accountRepo.SelectByPrimaryEmail", logger.Fields{"primaryEmail": primaryEmail})
	statement := "SELECT * FROM accounts WHERE primary_email=?"
	stmt, err := acc.db.Prepare(statement)
	if err != nil {
		err = appErrors.WrapDbPrepare(err, statement)
		return domains.Account{}, err
	}
	defer stmt.Close()

	var account domains.Account
	row := stmt.QueryRow(primaryEmail)
	if err = row.Scan(
		&account.Id,
		&account.CreatedAt,
		&account.UpdatedAt,
		&account.DeletedAt,
		&account.PrimaryEmail,
		&account.Password,
	); err != nil {
		err = appErrors.WrapDbQuery(err, statement, primaryEmail)
		return domains.Account{}, err
	}
	return account, nil
}

func (acc *accountRepo) Insert(account domains.Account) error {
	logger.Info("accountRepo.Insert", logger.Fields{"account": account})
	statement := "INSERT INTO accounts (" +
		"created_at, " +
		"updated_at, " +
		"primary_email, " +
		"password" +
		") VALUES (?, ?, ?, ?)"

	stmt, err := acc.db.Prepare(statement)
	if err != nil {
		err = appErrors.WrapDbPrepare(err, statement)
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()
	execResult, err := stmt.Exec(
		now,
		now,
		account.PrimaryEmail,
		account.Password,
	)
	if err != nil {
		err = appErrors.WrapDbExec(err, statement, account)
		return err
	}
	return appErrors.HandleDbResult(execResult, 1, "account was not inserted")
}

func (acc *accountRepo) Update(id uint, account domains.Account) error {
	logger.Info("accountRepo.Update", logger.Fields{
		"id": id, 
		"account": account,
	})
	statement := "UPDATE accounts SET " +
		"updated_at=?, " +
		"primary_email=?, " +
		"password=? " +
		"WHERE id=?"
	stmt, err := acc.db.Prepare(statement)
	if err != nil {
		err = appErrors.WrapDbPrepare(err, statement)
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()
	execResult, err := stmt.Exec(
		now,
		account.PrimaryEmail,
		account.Password,
		id)
	if err != nil {
		err = appErrors.WrapDbExec(err, statement, id, account)
		return err
	}
	return appErrors.HandleDbResult(execResult, 1, "account was not updated")
}

func (acc *accountRepo) Delete(id uint) error {
	logger.Info("accountRepo.Delete", logger.Fields{"id": id})
	statement := "DELETE FROM accounts WHERE id=?"
	stmt, err := acc.db.Prepare(statement)
	if err != nil {
		err = appErrors.WrapDbPrepare(err, statement)
		return err
	}
	defer stmt.Close()

	execResult, err := stmt.Exec(id)
	if err != nil {
		err = appErrors.WrapDbExec(err, statement, id)
		return err
	}
	return appErrors.HandleDbResult(execResult, 1, "account was not deleted")
}

// For Tests Only
func CreateTestAccountRepo(db *sql.DB) AccountRepoInterface {
	acc := &accountRepo{}
	acc.Initialize(db)
	return acc
}
