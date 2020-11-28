package dbTx

import (
	"database/sql"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"time"
)

func (dbTx *DbTx) CreateStmtSelectAccountById() string {
	return "SELECT * FROM accounts WHERE id=?"
}

func (dbTx *DbTx) CreateStmtSelectAccountByEmail() string {
	return "SELECT * FROM accounts WHERE primary_email=?"
}

func (dbTx *DbTx) CreateStmtSelectAccountByNegativeBalance() string {
	return "SELECT accounts.*, SUM(amount) FROM accounts " +
		"JOIN transactions ON accounts.id=transactions.account_id " +
		"GROUP BY account_id " +
		"HAVING SUM(amount) < 0"
}

func (dbTx *DbTx) ScanAccount(rows *sql.Rows) (domains.Account, error) {
	var account domains.Account
	if err := rows.Scan(
		&account.Id,
		&account.CreatedAt,
		&account.UpdatedAt,
		&account.DeletedAt,
		&account.PrimaryEmail,
		&account.Password,
	); err != nil {
		return domains.Account{}, err
	}
	return account, nil
}

func (dbTx *DbTx) SelectOneAccount(statement string, args ...interface{}) (domains.Account, error) {
	accounts, err := dbTx.SelectManyAccounts(statement, args...)
	if err != nil {
		return domains.Account{}, err
	}
	if len(accounts) == 0 {
		return domains.Account{}, appErrors.ERR_SQL_NO_ROWS
	}
	return accounts[0], nil
}

func (dbTx *DbTx) SelectManyAccounts(statement string, args ...interface{}) ([]domains.Account, error) {
	results := make([]domains.Account, 0)
	rows, err := dbTx.Query(statement, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user, err := dbTx.ScanAccount(rows)
		if err != nil {
			return results, err
		}
		results = append(results, user)
	}
	return results, nil
}

func (dbTx *DbTx) SelectManyAccountBalances(statement string, args ...interface{}) ([]domains.AccountBalance, error) {
	results := make([]domains.AccountBalance, 0)
	rows, err := dbTx.Query(statement, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var accountSum domains.AccountBalance
		if err := rows.Scan(
			&accountSum.Account.Id,
			&accountSum.Account.CreatedAt,
			&accountSum.Account.UpdatedAt,
			&accountSum.Account.DeletedAt,
			&accountSum.Account.PrimaryEmail,
			&accountSum.Account.Password,
			&accountSum.Balance,
		); err != nil {
			return results, err
		}
		results = append(results, accountSum)
	}
	return results, nil
}

func (dbTx *DbTx) InsertAccount(account domains.Account) (uint, error) {
	statement := "INSERT INTO accounts (" +
		"created_at, " +
		"updated_at, " +
		"primary_email, " +
		"password" +
		") VALUES (?, ?, ?, ?)"
	now := time.Now().UTC()
	result, err := dbTx.Exec(
		statement,
		now,
		now,
		account.PrimaryEmail,
		account.Password,
	)
	if err != nil {
		return 0, err
	}
	rowId, err := result.LastInsertId()
	if err != nil {
		return 0, appErrors.WrapSQLBadInsertResult(err)
	}
	return uint(rowId), nil
}

func (dbTx *DbTx) UpdateAccountById(accountId uint, account domains.Account) error {
	statement := "UPDATE accounts SET " +
		"updated_at=?, " +
		"primary_email=?, " +
		"password=? " +
		"WHERE id=?"
	now := time.Now().UTC()
	result, err := dbTx.Exec(
		statement,
		now,
		account.PrimaryEmail,
		account.Password,
		accountId)
	if err != nil {
		err = appErrors.WrapDbExec(err, statement, accountId, account)
		return err
	}
	return appErrors.ValidateDbResult(result, 1, "account was not updated")
}

func (dbTx *DbTx) DeleteAccount(id uint) error {
	statement := "DELETE FROM accounts WHERE id=?"
	result, err := dbTx.Exec(statement, id)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, id)
	}
	return appErrors.ValidateDbResult(result, 1, "account was not deleted")
}
