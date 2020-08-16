package repos

import (
	"database/sql"
	"log"
	"time"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
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
	InsertWithUser(domains.Account, domains.User) error
	Update(uint, domains.Account) error
	Delete(uint) error
}

func (acc *accountRepo) Initialize(db *sql.DB) {
	acc.db = db
}

func (acc *accountRepo) SelectById(id uint) (domains.Account, error) {
	statement := "SELECT * FROM accounts WHERE id=?"
	stmt, err := acc.db.Prepare(statement)
	if err != nil {
		return domains.Account{}, err
	}
	defer stmt.Close()

	var account domains.Account
	row := stmt.QueryRow(id)
	errScan := row.Scan(
		&account.Id,
		&account.CreatedAt,
		&account.UpdatedAt,
		&account.DeletedAt,
		&account.PrimaryEmail,
		&account.Password)
	return account, errScan
}

func (acc *accountRepo) SelectByPrimaryEmail(primary_email string) (domains.Account, error) {
	statement := "SELECT * FROM accounts WHERE primary_email=?"
	stmt, err := acc.db.Prepare(statement)
	if err != nil {
		return domains.Account{}, err
	}
	defer stmt.Close()

	var account domains.Account
	row := stmt.QueryRow(primary_email)

	errScan := row.Scan(
		&account.Id,
		&account.CreatedAt,
		&account.UpdatedAt,
		&account.DeletedAt,
		&account.PrimaryEmail,
		&account.Password)

	return account, errScan
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
	statement := "INSERT INTO accounts (" +
		"created_at, " +
		"updated_at, " +
		"primary_email, " +
		"password" +
		") VALUES (?, ?, ?, ?)"

	stmt, err := acc.db.Prepare(statement)
	if err != nil {
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
		return err
	}
	return utils.HandleSqlExecResult(execResult, 1, "account was not inserted")
}

func (acc *accountRepo) InsertWithUser(account domains.Account, user domains.User) error {
	tx, err := acc.db.Begin()
	if err != nil {
		return err
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
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()
	_, err = stmt.Exec(
		now,
		now,
		account.PrimaryEmail,
		account.Password,
	)
	if err != nil {
		tx.Rollback()
		return err
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
		return err
	}
	defer stmt2.Close()

	_, err = stmt2.Exec(
		now,
		now,
		user.FirstName,
		user.LastName,
		user.MiddleName,
		user.Email,
		user.Phone,
		user.IsGuardian,
		user.AccountId,
		user.Notes,
		user.School,
		user.GraduationYear,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (acc *accountRepo) Update(id uint, account domains.Account) error {
	statement := "UPDATE accounts SET " +
		"updated_at=?, " +
		"primary_email=?, " +
		"password=? " +
		"WHERE id=?"
	stmt, err := acc.db.Prepare(statement)
	if err != nil {
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
		return err
	}
	return utils.HandleSqlExecResult(execResult, 1, "account was not updated")
}

func (acc *accountRepo) Delete(id uint) error {
	statement := "DELETE FROM accounts WHERE id=?"
	stmt, err := acc.db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	execResult, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	return utils.HandleSqlExecResult(execResult, 1, "account was not deleted")
}

// For Tests Only
func CreateTestAccountRepo(db *sql.DB) AccountRepoInterface {
	acc := &accountRepo{}
	acc.Initialize(db)
	return acc
}
