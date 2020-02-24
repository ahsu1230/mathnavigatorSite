package repos

import (
	"database/sql"
)

func SetupRepos(db *sql.DB) {
	ProgramRepo.Initialize(db)
}

func QueryStatement(db *sql.DB, statement string) (*sql.Rows, error) {
	stmt, err := db.Prepare(statement)
	defer stmt.Close()
	if err != nil {
		return nil, err
    }
    
    rows, err := stmt.Query()
	if err != nil {
		return nil,  err
	}
	return rows, err
}

func ExecStatement(db *sql.DB) {

}