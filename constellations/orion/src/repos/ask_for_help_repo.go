package repos

import (
	"database/sql"
	"time"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos/utils"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
)

// Global variable
var AskForHelpRepo AskForHelpRepoInterface = &askForHelpRepo{}

// Implements interface askForHelpRepoInterface
type askForHelpRepo struct {
	db *sql.DB
}

// Interface to implement
type AskForHelpRepoInterface interface {
	Initialize(db *sql.DB)
	SelectAll() ([]domains.AskForHelp, error)
	SelectById(uint) (domains.AskForHelp, error)
	Insert(domains.AskForHelp) error
	Update(uint, domains.AskForHelp) error
	Delete(uint) error
}

func (ar *askForHelpRepo) Initialize(db *sql.DB) {
	ar.db = db
}

func (ar *askForHelpRepo) SelectAll() ([]domains.AskForHelp, error) {
	results := make([]domains.AskForHelp, 0)

	stmt, err := ar.db.Prepare("SELECT * FROM ask_for_help")
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
		var askForHelp domains.AskForHelp
		if errScan := rows.Scan(
			&askForHelp.Id,
			&askForHelp.CreatedAt,
			&askForHelp.UpdatedAt,
			&askForHelp.DeletedAt,
			&askForHelp.Title,
			&askForHelp.Date,
			&askForHelp.TimeString,
			&askForHelp.Subject,
			&askForHelp.LocationId); errScan != nil {
			return results, errScan
		}
		results = append(results, askForHelp)
	}

	return results, nil
}

func (ar *askForHelpRepo) SelectById(id uint) (domains.AskForHelp, error) {
	statement := "SELECT * FROM ask_for_help WHERE id=?"
	stmt, err := ar.db.Prepare(statement)
	if err != nil {
		return domains.AskForHelp{}, err
	}
	defer stmt.Close()

	var askForHelp domains.AskForHelp
	row := stmt.QueryRow(id)
	errScan := row.Scan(
		&askForHelp.Id,
		&askForHelp.CreatedAt,
		&askForHelp.UpdatedAt,
		&askForHelp.DeletedAt,
		&askForHelp.Title,
		&askForHelp.Date,
		&askForHelp.TimeString,
		&askForHelp.Subject,
		&askForHelp.LocationId)
	return askForHelp, errScan
}

func (ar *askForHelpRepo) Insert(askForHelp domains.AskForHelp) error {
	stmt, err := ar.db.Prepare("INSERT INTO ask_for_help (" +
		"created_at, " +
		"updated_at, " +
		"title, " +
		"date, " +
		"time_string, " +
		"subject, " +
		"location_id" +
		") VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()
	result, err := stmt.Exec(
		now,
		now,
		askForHelp.Title,
		askForHelp.Date,
		askForHelp.TimeString,
		askForHelp.Subject,
		askForHelp.LocationId)
	if err != nil {
		return err
	}
	return utils.HandleSqlExecResult(result, 1, "ask for help was not inserted")
}

func (ar *askForHelpRepo) Update(id uint, askForHelp domains.AskForHelp) error {
	stmt, err := ar.db.Prepare("UPDATE ask_for_help SET " +
		"updated_at=?, " +
		"id=?, " +
		"title=?, " +
		"date=?, " +
		"time_string=?, " +
		"subject=?, " +
		"location_id=? " +
		"WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()
	result, err := stmt.Exec(
		now,
		askForHelp.Id,
		askForHelp.Title,
		askForHelp.Date,
		askForHelp.TimeString,
		askForHelp.Subject,
		askForHelp.LocationId,
		id)
	if err != nil {
		return err
	}
	return utils.HandleSqlExecResult(result, 1, "ask for help was not updated")
}

func (ar *askForHelpRepo) Delete(id uint) error {
	statement := "DELETE FROM ask_for_help WHERE id=?"
	stmt, err := ar.db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	execResult, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	return utils.HandleSqlExecResult(execResult, 1, "ask for help was not deleted")
}
func CreateTestAFHRepo(db *sql.DB) AskForHelpRepoInterface {
	ar := &askForHelpRepo{}
	ar.Initialize(db)
	return ar
}
