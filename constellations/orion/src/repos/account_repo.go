package repos

import (
	"database/sql"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos/utils"
	"time"
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
	statement := "SELECT * FROM accounts WHERE id=? AND deleted_at IS NULL"
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
	statement := "SELECT * FROM accounts WHERE primary_email=? AND deleted_at IS NULL"
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

// Sets the "deleted_at" column in the following tables (accounts, transactions,
// users, user_afh, user_classes) to time.Now().UTC()

func (acc *accountRepo) Delete(id uint) error {
	utils.LogWithContext("accountRepo.Delete", logger.Fields{"id": id})
	now := time.Now().UTC()

	trans, err := acc.db.Begin()
	if err != nil {
		return appErrors.WrapDbTxBegin(err)
	}

	// Accounts
	query := "UPDATE accounts SET deleted_at=? WHERE id=?"
	stmt, err := acc.db.Prepare(query)
	if err != nil {
		trans.Rollback()
		err = appErrors.WrapDbPrepare(err, query)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(domains.NewNullTime(now), id)
	if err != nil {
		trans.Rollback()
		err = appErrors.WrapDbExec(err, query, id)
		return err
	}

	// Transactions
	query2 := "UPDATE transactions SET deleted_at=? WHERE account_id=?"
	stmt2, err := acc.db.Prepare(query2)
	if err != nil {
		trans.Rollback()
		err = appErrors.WrapDbPrepare(err, query2)
		return err
	}
	defer stmt2.Close()

	_, err = stmt2.Exec(now, id)
	if err != nil {
		trans.Rollback()
		err = appErrors.WrapDbExec(err, query2, id)
		return err
	}

	// Get User Ids
	userIds := make([]uint, 0)
	query3 := "SELECT id FROM users WHERE account_id=?"
	stmt3, err := acc.db.Prepare(query3)
	if err != nil {
		return appErrors.WrapDbPrepare(err, query3)
	}
	defer stmt3.Close()
	rows, err := stmt3.Query(domains.NewNullUint(id))
	if err != nil {
		return appErrors.WrapDbQuery(err, query3, id)
	}
	defer rows.Close()
	for rows.Next() {
		var user domains.User
		if errScan := rows.Scan(
			&user.Id); errScan != nil {
			return errScan
		}
		userIds = append(userIds, user.Id)
	}

	// Users
	query4 := "UPDATE users SET deleted_at=? WHERE account_id=?"
	stmt4, err := acc.db.Prepare(query4)
	if err != nil {
		trans.Rollback()
		err = appErrors.WrapDbPrepare(err, query4)
		return err
	}
	defer stmt4.Close()

	_, err = stmt4.Exec(now, id)
	if err != nil {
		trans.Rollback()
		err = appErrors.WrapDbExec(err, query4, id)
		return err
	}

	// User Classes
	query5 := "UPDATE user_classes SET deleted_at=? WHERE account_id=?"
	stmt5, err := acc.db.Prepare(query5)
	if err != nil {
		trans.Rollback()
		err = appErrors.WrapDbPrepare(err, query5)
		return err
	}
	defer stmt5.Close()

	_, err = stmt5.Exec(now, id)
	if err != nil {
		trans.Rollback()
		err = appErrors.WrapDbExec(err, query5, id)
		return err
	}

	// User AFH
	for _, uid := range userIds {
		query6 := "UPDATE user_afh SET deleted_at=? WHERE user_id=?"
		stmt6, err := acc.db.Prepare(query6)
		if err != nil {
			trans.Rollback()
			err = appErrors.WrapDbPrepare(err, query6)
			return err
		}
		defer stmt6.Close()

		_, err = stmt6.Exec(now, uid)
		if err != nil {
			trans.Rollback()
			err = appErrors.WrapDbExec(err, query6, uid)
			return err
		}
	}
	trans.Commit()
	return nil
}

// For Tests Only
func CreateTestAccountRepo(db *sql.DB) AccountRepoInterface {
	acc := &accountRepo{}
	acc.Initialize(db)
	return acc
}
