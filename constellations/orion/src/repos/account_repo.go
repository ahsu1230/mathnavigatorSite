package repos

import (
	"database/sql"
	"time"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
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
	Initialize(db *sql.DB)
	SelectById(uint) (domains.Account, error)
	SelectByPrimaryEmail(string) (domains.Account, error)
	SelectAllNegativeBalances() ([]domains.AccountSum, error)
	Insert(domains.Account) error
	Update(uint, domains.Account) error
	Delete(uint) error
}

func (acc *accountRepo) Initialize(db *sql.DB) {
	utils.LogWithContext("accountRepo.Initialize", logger.Fields{})
	acc.db = db
}

func (acc *accountRepo) SelectById(id uint) (domains.Account, error) {
	utils.LogWithContext("accountRepo.SelectById", logger.Fields{"id": id})
	statement := "SELECT * FROM accounts WHERE id=?"
	stmt, err := acc.db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		return domains.Account{}, appErrors.WrapDbPrepare(err, statement)
	}

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
		return domains.Account{}, appErrors.WrapDbQuery(err, statement, id)
	}
	return account, nil
}

func (acc *accountRepo) SelectByPrimaryEmail(primaryEmail string) (domains.Account, error) {
	utils.LogWithContext("accountRepo.SelectByPrimaryEmail",
		logger.Fields{"primaryEmail": primaryEmail},
	)
	statement := "SELECT * FROM accounts WHERE primary_email=?"
	stmt, err := acc.db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		return domains.Account{}, appErrors.WrapDbPrepare(err, statement)
	}

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

func (acc *accountRepo) SelectAllNegativeBalances() ([]domains.AccountSum, error) {
	results := make([]domains.AccountSum, 0)

	statement := "SELECT accounts.*, SUM(amount) FROM accounts JOIN transactions ON accounts.id=transactions.account_id GROUP BY account_id HAVING SUM(amount) < 0"
	stmt, err := acc.db.Prepare(statement)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var accountSum domains.AccountSum
		if errScan := rows.Scan(
			&accountSum.Account.Id,
			&accountSum.Account.CreatedAt,
			&accountSum.Account.UpdatedAt,
			&accountSum.Account.DeletedAt,
			&accountSum.Account.PrimaryEmail,
			&accountSum.Account.Password,
			&accountSum.Balance); errScan != nil {
			return results, errScan
		}
		results = append(results, accountSum)
	}
	return results, nil
}

func (acc *accountRepo) Insert(account domains.Account) error {
	utils.LogWithContext("accountRepo.Insert", logger.Fields{"account": account})
	statement := "INSERT INTO accounts (" +
		"created_at, " +
		"updated_at, " +
		"primary_email, " +
		"password" +
		") VALUES (?, ?, ?, ?)"

	stmt, err := acc.db.Prepare(statement)
	if err != nil {
		return appErrors.WrapDbPrepare(err, statement)
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
		return appErrors.WrapDbExec(err, statement, account)
	}
	return appErrors.ValidateDbResult(execResult, 1, "account was not inserted")
}

func (acc *accountRepo) Update(id uint, account domains.Account) error {
	utils.LogWithContext("accountRepo.Update", logger.Fields{
		"id":      id,
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
	return appErrors.ValidateDbResult(execResult, 1, "account was not updated")
}

func (acc *accountRepo) Delete(id uint) error {
	utils.LogWithContext("accountRepo.Delete", logger.Fields{"id": id})
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
	return appErrors.ValidateDbResult(execResult, 1, "account was not deleted")
}

// For Tests Only
func CreateTestAccountRepo(db *sql.DB) AccountRepoInterface {
	acc := &accountRepo{}
	acc.Initialize(db)
	return acc
}
