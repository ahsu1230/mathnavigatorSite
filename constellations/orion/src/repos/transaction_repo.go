package repos

import (
	"database/sql"
	"time"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
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
	SelectByAccountId(uint) ([]domains.Transaction, error)
	SelectById(uint) (domains.Transaction, error)
	Insert(domains.Transaction) error
	Update(uint, domains.Transaction) error
	Delete(uint) error
}

func (tr *transactionRepo) Initialize(db *sql.DB) {
	utils.LogWithContext("transactionRepo.Initialize", logger.Fields{})
	tr.db = db
}

func (tr *transactionRepo) SelectByAccountId(accountId uint) ([]domains.Transaction, error) {
	utils.LogWithContext("transactionRepo.SelectByAccountId", logger.Fields{"accountId": accountId})
	results := make([]domains.Transaction, 0)

	statement := "SELECT * FROM transactions WHERE account_id=?"
	stmt, err := tr.db.Prepare(statement)
	if err != nil {
		return nil, appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()
	rows, err := stmt.Query(domains.NewNullUint(accountId))
	if err != nil {
		return nil, appErrors.WrapDbQuery(err, statement, accountId)
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
	utils.LogWithContext("transactionRepo.SelectById", logger.Fields{"id": id})
	statement := "SELECT * FROM transactions WHERE id=?"
	stmt, err := tr.db.Prepare(statement)
	if err != nil {
		return domains.Transaction{}, appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	var transaction domains.Transaction
	row := stmt.QueryRow(id)
	if err = row.Scan(
		&transaction.Id,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
		&transaction.DeletedAt,
		&transaction.Amount,
		&transaction.PaymentType,
		&transaction.PaymentNotes,
		&transaction.AccountId); err != nil {
		return domains.Transaction{}, appErrors.WrapDbExec(err, statement, id)
	}
	return transaction, nil
}

func (tr *transactionRepo) Insert(transaction domains.Transaction) error {
	utils.LogWithContext("transactionRepo.Insert", logger.Fields{"transaction": transaction})
	statement := "INSERT INTO transactions (" +
		"created_at, " +
		"updated_at, " +
		"amount, " +
		"payment_type, " +
		"payment_notes, " +
		"account_id " +
		") VALUES (?, ?, ?, ?, ?, ?)"
	stmt, err := tr.db.Prepare(statement)
	if err != nil {
		return appErrors.WrapDbPrepare(err, statement)
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
		return appErrors.WrapDbExec(err, statement, transaction)
	}
	return appErrors.ValidateDbResult(result, 1, "transaction was not inserted")
}

func (tr *transactionRepo) Update(id uint, transaction domains.Transaction) error {
	utils.LogWithContext("transactionRepo.Update", logger.Fields{"transaction": transaction})
	statement := "UPDATE transactions SET " +
		"updated_at=?, " +
		"amount=?, " +
		"payment_type=?, " +
		"payment_notes=?, " +
		"account_id=? " +
		"WHERE id=?"
	stmt, err := tr.db.Prepare(statement)
	if err != nil {
		return appErrors.WrapDbPrepare(err, statement)
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
		return appErrors.WrapDbExec(err, statement, transaction, id)
	}
	return appErrors.ValidateDbResult(result, 1, "transaction was not updated")
}

func (tr *transactionRepo) Delete(id uint) error {
	utils.LogWithContext("transactionRepo.Delete", logger.Fields{"id": id})
	statement := "DELETE FROM transactions WHERE id=?"
	stmt, err := tr.db.Prepare(statement)
	if err != nil {
		return appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	execResult, err := stmt.Exec(id)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, id)
	}
	return appErrors.ValidateDbResult(execResult, 1, "transaction was not deleted")
}

func CreateTestTransactionRepo(db *sql.DB) TransactionRepoInterface {
	tr := &transactionRepo{}
	tr.Initialize(db)
	return tr
}
