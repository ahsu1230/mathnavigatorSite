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
var AnnounceRepo AnnounceRepoInterface = &announceRepo{}

type announceRepo struct {
	db *sql.DB
}

type AnnounceRepoInterface interface {
	Initialize(context.Context, *sql.DB)
	SelectAll(context.Context) ([]domains.Announce, error)
	SelectByAnnounceId(context.Context, uint) (domains.Announce, error)
	Insert(context.Context, domains.Announce) (uint, error)
	Update(context.Context, uint, domains.Announce) error
	Delete(context.Context, uint) error
}

func (ar *announceRepo) Initialize(ctx context.Context, db *sql.DB) {
	utils.LogWithContext(ctx, "announceRepo.Initialize", logger.Fields{})
	ar.db = db
}

func (ar *announceRepo) SelectAll(ctx context.Context) ([]domains.Announce, error) {
	utils.LogWithContext(ctx, "announceRepo.SelectAll", logger.Fields{})
	results := make([]domains.Announce, 0)

	statement := "SELECT * FROM announcements ORDER BY posted_at DESC"
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
		var announce domains.Announce
		if errScan := rows.Scan(
			&announce.Id,
			&announce.CreatedAt,
			&announce.UpdatedAt,
			&announce.DeletedAt,
			&announce.PostedAt,
			&announce.Author,
			&announce.Message,
			&announce.OnHomePage); errScan != nil {
			return results, errScan
		}
		results = append(results, announce)
	}

	return results, nil
}

func (ar *announceRepo) SelectByAnnounceId(ctx context.Context, id uint) (domains.Announce, error) {
	utils.LogWithContext(ctx, "announceRepo.SelectByAnnounceId", logger.Fields{"id": id})
	statement := "SELECT * FROM announcements WHERE id=?"
	stmt, err := ar.db.Prepare(statement)
	if err != nil {
		return domains.Announce{}, appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	var announce domains.Announce
	row := stmt.QueryRow(id)
	if err := row.Scan(
		&announce.Id,
		&announce.CreatedAt,
		&announce.UpdatedAt,
		&announce.DeletedAt,
		&announce.PostedAt,
		&announce.Author,
		&announce.Message,
		&announce.OnHomePage); err != nil {
		return domains.Announce{}, appErrors.WrapDbExec(err, statement, id)
	}

	return announce, nil
}

func (ar *announceRepo) Insert(ctx context.Context, announce domains.Announce) (uint, error) {
	utils.LogWithContext(ctx, "announceRepo.Insert", logger.Fields{"announce": announce})
	statement := "INSERT INTO announcements (" +
		"created_at, " +
		"updated_at, " +
		"posted_at, " +
		"author, " +
		"message," +
		"on_home_page" +
		") VALUES (?, ?, ?, ?, ?, ?)"
	stmt, err := ar.db.Prepare(statement)
	if err != nil {
		return 0, appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	now := time.Now().UTC()
	result, err := stmt.Exec(
		now,
		now,
		announce.PostedAt,
		announce.Author,
		announce.Message,
		announce.OnHomePage)
	if err != nil {
		return 0, appErrors.WrapDbExec(err, statement, announce)
	}

	rowId, err := result.LastInsertId()
	if err != nil {
		return 0, appErrors.WrapSQLBadInsertResult(err)
	}
	return uint(rowId), appErrors.ValidateDbResult(result, 1, "announcement was not inserted")
}

func (ar *announceRepo) Update(ctx context.Context, id uint, announce domains.Announce) error {
	utils.LogWithContext(ctx, "announceRepo.Update", logger.Fields{"announce": announce})
	statement := "UPDATE announcements SET " +
		"updated_at=?, " +
		"posted_at=?, " +
		"author=?, " +
		"message=?, " +
		"on_home_page=? " +
		"WHERE id=?"
	stmt, err := ar.db.Prepare(statement)
	if err != nil {
		return appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	now := time.Now().UTC()
	result, err := stmt.Exec(
		now,
		announce.PostedAt,
		announce.Author,
		announce.Message,
		announce.OnHomePage,
		id)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, announce, id)
	}
	return appErrors.ValidateDbResult(result, 1, "announcement was not updated")
}

func (ar *announceRepo) Delete(ctx context.Context, id uint) error {
	utils.LogWithContext(ctx, "Repo.Delete", logger.Fields{"id": id})
	statement := "DELETE FROM announcements WHERE id=?"
	stmt, err := ar.db.Prepare(statement)
	if err != nil {
		return appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, id)
	}

	return appErrors.ValidateDbResult(result, 1, "announcement was not deleted")
}

func CreateTestAnnounceRepo(ctx context.Context, db *sql.DB) AnnounceRepoInterface {
	ar := &announceRepo{}
	ar.Initialize(ctx, db)
	return ar
}
