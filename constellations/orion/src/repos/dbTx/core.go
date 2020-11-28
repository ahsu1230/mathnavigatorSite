package dbTx

import (
	"database/sql"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
)

type DbTx struct {
	db *sql.DB
	tx *sql.Tx
}

// A constructor for Read-only transactions
// Does not create a transaction
func New(db *sql.DB) *DbTx {
	return &DbTx{db, nil}
}

// A constructor for write transactions
// Creates a transaction (calls db.Tx.Begin())
func Begin(db *sql.DB) (*DbTx, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, appErrors.WrapDbTxBegin(err)
	}
	newDbTx := DbTx{db, tx}
	return &newDbTx, nil
}

func (dbTx *DbTx) Commit() error {
	if err := dbTx.tx.Commit(); err != nil {
		return appErrors.WrapDbTxCommit(err)
	}
	return nil
}

func (dbTx *DbTx) Rollback() error {
	if err := dbTx.tx.Rollback(); err != nil {
		return appErrors.WrapDbTxRollback(err)
	}
	return nil
}

func (dbTx *DbTx) Exec(statement string, args ...interface{}) (sql.Result, error) {
	stmt, err := dbTx.db.Prepare(statement)
	if err != nil {
		return nil, appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	execResult, err := stmt.Exec(args...)
	if err != nil {
		return nil, appErrors.WrapDbExec(err, statement, args...)
	}
	return execResult, nil
}

func (dbTx *DbTx) Query(statement string, args ...interface{}) (*sql.Rows, error) {
	stmt, err := dbTx.db.Prepare(statement)
	if err != nil {
		return nil, appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()
	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, appErrors.WrapDbQuery(err, statement, args...)
	}
	return rows, nil
}
