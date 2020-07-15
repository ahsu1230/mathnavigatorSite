package repos

import (
	"database/sql"
	"time"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos/utils"
)

// Global Variable
var TransactionRepo TransactionRepoInterface = &transactionRepo{}

// Implements interface transactionRepoInterface
type transactionRepo struct {
	db *sql.DB
}

//Interface to implement
type TransactionRepoInterface interface {
	Initialize(db *sql.DB)
	SelectAll() ([]domains.Transaction, error)
	SelectById(uint) (domains.Transaction, error)
	Insert(domains.Transaction) error
	Update(uint, domains.Transaction) error
	Delete(uint) error
}

func (tr *transactionRepo) Initialize(db *sql.DB) {
	tr.db = db
}

func (tr *transactionRepo) SelectAll() ([]domains.Transaction, error) {
	results := make([]domains.Transaction, 0)

	stmt, err := tr.db.Prepare("SELECT * FROM transactions")
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
		var transaction domains.Transaction
		if errScan := rows.Scan(
			&transaction.Id,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
			&transaction.DeletedAt,
			&transaction.Amount,
			&transaction.PaymentType,
			&transaction.PaymentNotes,
			&transaction.AccountId); errScan != nil {
			return results, errScan
		}
		results = append(results, transaction)
	}

	return results, nil
}

func (tr *transactionRepo) SelectById(id uint) (domains.Transaction, error) {
	statement := "SELECT * FROM transactions WHERE id=?"
	stmt, err := tr.db.Prepare(statement)
	if err != nil {
		return domains.Transaction{}, err
	}
	defer stmt.Close()

	var transaction domains.Transaction
	row := stmt.QueryRow(id)
	errScan := row.Scan(
		&transaction.Id,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
		&transaction.DeletedAt,
		&transaction.Amount,
		&transaction.PaymentType,
		&transaction.PaymentNotes,
		&transaction.AccountId)
	return transaction, errScan
}

func (tr *transactionRepo) Insert(transaction domains.Transaction) error {
	stmt, err := tr.db.Prepare("INSERT INTO transactions (" +
		"created_at, " +
		"updated_at, " +
		"amount, " +
		"payment_type, " +
		"payment_notes, " +
		"account_id " +
		") VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()
	result, err := stmt.Exec(
		now,
		now,
		transaction.Amount,
		transaction.PaymentType,
		transaction.PaymentNotes,
		transaction.AccountId)
	if err != nil {
		return err
	}
	return utils.HandleSqlExecResult(result, 1, "transaction was not inserted")
}

func (tr *transactionRepo) Update(id uint, transaction domains.Transaction) error {
	stmt, err := tr.db.Prepare("UPDATE transactions SET " +
		"updated_at=?, " +
		"amount=?, " +
		"payment_type=?, " +
		"payment_notes=?, " +
		"account_id=? " +
		"WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()
	result, err := stmt.Exec(
		now,
		transaction.Amount,
		transaction.PaymentType,
		transaction.PaymentNotes,
		transaction.AccountId,
		id)
	if err != nil {
		return err
	}
	return utils.HandleSqlExecResult(result, 1, "transaction was not updated")
}

func (tr *transactionRepo) Delete(id uint) error {
	statement := "DELETE FROM transactions WHERE id=?"
	stmt, err := tr.db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	execResult, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	return utils.HandleSqlExecResult(execResult, 1, "transaction was not deleted")
}

func CreateTestTransactionRepo(db *sql.DB) TransactionRepoInterface {
	tr := &transactionRepo{}
	tr.Initialize(db)
	return tr
}
