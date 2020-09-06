package repos

import (
	"context"
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
	Initialize(context.Context, *sql.DB)
	SelectById(context.Context, uint) (domains.Account, error)
	SelectByPrimaryEmail(context.Context, string) (domains.Account, error)
	SelectAllNegativeBalances(context.Context) ([]domains.AccountSum, error)
	InsertWithUser(context.Context, domains.Account, domains.User) error
	Update(context.Context, uint, domains.Account) error
	Delete(context.Context, uint) error
}

func (acc *accountRepo) Initialize(ctx context.Context, db *sql.DB) {
	utils.LogWithContext(ctx, "accountRepo.Initialize", logger.Fields{})
	acc.db = db
}

func (acc *accountRepo) SelectById(ctx context.Context, id uint) (domains.Account, error) {
	utils.LogWithContext(ctx, "accountRepo.SelectById", logger.Fields{"id": id})
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

func (acc *accountRepo) SelectByPrimaryEmail(ctx context.Context, primaryEmail string) (domains.Account, error) {
	utils.LogWithContext(ctx, "accountRepo.SelectByPrimaryEmail",
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

func (acc *accountRepo) SelectAllNegativeBalances(ctx context.Context) ([]domains.AccountSum, error) {
	utils.LogWithContext(ctx, "accountRepo.SelectAllNegativeBalances", logger.Fields{})
	results := make([]domains.AccountSum, 0)

	statement := "SELECT accounts.*, SUM(amount) FROM accounts JOIN transactions ON accounts.id=transactions.account_id GROUP BY account_id HAVING SUM(amount) < 0"
	stmt, err := acc.db.Prepare(statement)
	if err != nil {
		return nil, appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, appErrors.WrapDbQuery(err, statement)
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

func (acc *accountRepo) InsertWithUser(ctx context.Context, account domains.Account, user domains.User) error {
	utils.LogWithContext(ctx, "accountRepo.Insert", logger.Fields{"account": account})
	tx, err := acc.db.Begin()
	if err != nil {
		return appErrors.WrapDbTxBegin(err)
	}

	statement := "INSERT INTO accounts (" +
		"created_at, " +
		"updated_at, " +
		"primary_email, " +
		"password" +
		") VALUES (?, ?, ?, ?)"
	stmt, err := acc.db.Prepare(statement)
	if err != nil {
		tx.Rollback()
		return appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	now := time.Now().UTC()
	result1, err := stmt.Exec(
		now,
		now,
		account.PrimaryEmail,
		account.Password,
	)
	if err != nil {
		tx.Rollback()
		return appErrors.WrapDbExec(err, statement, account)
	}
	if err = appErrors.ValidateDbResult(result1, 1, "account was not inserted"); err != nil {
		tx.Rollback()
		return appErrors.WrapDbExec(err, statement, account)
	}

	lastAccountId, err := result1.LastInsertId()
	if err != nil {
		tx.Rollback()
		return appErrors.WrapDbExec(err, statement, account)
	}
	statement2 := "INSERT INTO users (" +
		"created_at, " +
		"updated_at, " +
		"first_name, " +
		"last_name," +
		"middle_name, " +
		"email," +
		"phone, " +
		"is_guardian," +
		"account_id," +
		"notes," +
		"school," +
		"graduation_year" +
		") VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt2, err := acc.db.Prepare(statement2)
	if err != nil {
		tx.Rollback()
		return appErrors.WrapDbPrepare(err, statement2)
	}
	defer stmt2.Close()

	result2, err := stmt2.Exec(
		now,
		now,
		user.FirstName,
		user.LastName,
		user.MiddleName,
		user.Email,
		user.Phone,
		user.IsGuardian,
		lastAccountId,
		user.Notes,
		user.School,
		user.GraduationYear,
	)
	if err != nil {
		tx.Rollback()
		return appErrors.WrapDbExec(err, statement, user)
	}
	if err = appErrors.ValidateDbResult(result2, 1, "user was not inserted"); err != nil {
		tx.Rollback()
		return appErrors.WrapDbExec(err, statement, user)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return appErrors.WrapDbTxCommit(err)
	}
	return nil
}

func (acc *accountRepo) Update(ctx context.Context, id uint, account domains.Account) error {
	utils.LogWithContext(ctx, "accountRepo.Update", logger.Fields{
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

func (acc *accountRepo) Delete(ctx context.Context, id uint) error {
	utils.LogWithContext(ctx, "accountRepo.Delete", logger.Fields{"id": id})
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
func CreateTestAccountRepo(ctx context.Context, db *sql.DB) AccountRepoInterface {
	acc := &accountRepo{}
	acc.Initialize(ctx, db)
	return acc
}
