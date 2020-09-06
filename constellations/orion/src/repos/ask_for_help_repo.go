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
var AskForHelpRepo AskForHelpRepoInterface = &askForHelpRepo{}

// Implements interface askForHelpRepoInterface
type askForHelpRepo struct {
	db *sql.DB
}

// Interface to implement
type AskForHelpRepoInterface interface {
	Initialize(context.Context, *sql.DB)
	SelectAll(context.Context) ([]domains.AskForHelp, error)
	SelectById(context.Context, uint) (domains.AskForHelp, error)
	Insert(context.Context, domains.AskForHelp) error
	Update(context.Context, uint, domains.AskForHelp) error
	Delete(context.Context, uint) error
}

func (ar *askForHelpRepo) Initialize(ctx context.Context, db *sql.DB) {
	utils.LogWithContext(ctx, "afhRepo.Initialize", logger.Fields{})
	ar.db = db
}

func (ar *askForHelpRepo) SelectAll(ctx context.Context) ([]domains.AskForHelp, error) {
	utils.LogWithContext(ctx, "afhRepo.SelectAll", logger.Fields{})
	results := make([]domains.AskForHelp, 0)

	statement := "SELECT * FROM ask_for_help"
	stmt, err := ar.db.Prepare(statement)
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
			&askForHelp.LocationId,
			&askForHelp.Notes); errScan != nil {
			return results, errScan
		}
		results = append(results, askForHelp)
	}

	return results, nil
}

func (ar *askForHelpRepo) SelectById(ctx context.Context, id uint) (domains.AskForHelp, error) {
	utils.LogWithContext(ctx, "afhRepo.SelectById", logger.Fields{"id": id})
	statement := "SELECT * FROM ask_for_help WHERE id=?"
	stmt, err := ar.db.Prepare(statement)
	if err != nil {
		return domains.AskForHelp{}, appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	var askForHelp domains.AskForHelp
	row := stmt.QueryRow(id)

	if err = row.Scan(
		&askForHelp.Id,
		&askForHelp.CreatedAt,
		&askForHelp.UpdatedAt,
		&askForHelp.DeletedAt,
		&askForHelp.Title,
		&askForHelp.Date,
		&askForHelp.TimeString,
		&askForHelp.Subject,
		&askForHelp.LocationId,
		&askForHelp.Notes); err != nil {
		return domains.AskForHelp{}, appErrors.WrapDbExec(err, statement, id)
	}
	return askForHelp, nil
}

func (ar *askForHelpRepo) Insert(ctx context.Context, askForHelp domains.AskForHelp) error {
	utils.LogWithContext(ctx, "afhRepo.Insert", logger.Fields{"afh": askForHelp})
	statement := "INSERT INTO ask_for_help (" +
		"created_at, " +
		"updated_at, " +
		"title, " +
		"date, " +
		"time_string, " +
		"subject, " +
		"location_id, " +
		"notes" +
		") VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := ar.db.Prepare(statement)
	if err != nil {
		return appErrors.WrapDbPrepare(err, statement)
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
		askForHelp.LocationId,
		askForHelp.Notes)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, askForHelp)
	}
	return appErrors.ValidateDbResult(result, 1, "ask for help was not inserted")
}

func (ar *askForHelpRepo) Update(ctx context.Context, id uint, askForHelp domains.AskForHelp) error {
	utils.LogWithContext(ctx, "afhRepo.Update", logger.Fields{"afh": askForHelp})
	statement := "UPDATE ask_for_help SET " +
		"updated_at=?, " +
		"id=?, " +
		"title=?, " +
		"date=?, " +
		"time_string=?, " +
		"subject=?, " +
		"location_id=?, " +
		"notes=? " +
		"WHERE id=?"
	stmt, err := ar.db.Prepare(statement)
	if err != nil {
		return appErrors.WrapDbPrepare(err, statement)
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
		askForHelp.Notes,
		id)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, askForHelp)
	}
	return appErrors.ValidateDbResult(result, 1, "ask for help was not updated")
}

func (ar *askForHelpRepo) Delete(ctx context.Context, id uint) error {
	utils.LogWithContext(ctx, "afhRepo.Delete", logger.Fields{"id": id})
	statement := "DELETE FROM ask_for_help WHERE id=?"
	stmt, err := ar.db.Prepare(statement)
	if err != nil {
		return appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	execResult, err := stmt.Exec(id)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, id)
	}
	return appErrors.ValidateDbResult(execResult, 1, "ask for help was not deleted")
}

func CreateTestAFHRepo(ctx context.Context, db *sql.DB) AskForHelpRepoInterface {
	ar := &askForHelpRepo{}
	ar.Initialize(ctx, db)
	return ar
}
