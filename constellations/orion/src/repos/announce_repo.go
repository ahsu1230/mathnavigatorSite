package repos

import (
	"database/sql"
	"time"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
)

// Global variable
var AnnounceRepo AnnounceRepoInterface = &announceRepo{}

type announceRepo struct {
	db *sql.DB
}

type AnnounceRepoInterface interface {
	Initialize(db *sql.DB)
	SelectAll() ([]domains.Announce, error)
	SelectByAnnounceId(uint) (domains.Announce, error)
	Insert(domains.Announce) error
	Update(uint, domains.Announce) error
	Delete(uint) error
}

func (ar *announceRepo) Initialize(db *sql.DB) {
	logger.Debug("Initialize AnnounceRepo", logger.Fields{})
	ar.db = db
}

func (ar *announceRepo) SelectAll() ([]domains.Announce, error) {
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

func (ar *announceRepo) SelectByAnnounceId(id uint) (domains.Announce, error) {
	statement := "SELECT * FROM announcements WHERE id=?"
	stmt, err := ar.db.Prepare(statement)
	if err != nil {
		return domains.Announce{}, appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	var announce domains.Announce
	row := stmt.QueryRow(id)
	errScan := row.Scan(
		&announce.Id,
		&announce.CreatedAt,
		&announce.UpdatedAt,
		&announce.DeletedAt,
		&announce.PostedAt,
		&announce.Author,
		&announce.Message,
		&announce.OnHomePage)

	return announce, errScan
}

func (ar *announceRepo) Insert(announce domains.Announce) error {
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
		return appErrors.WrapDbPrepare(err, statement)
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
		return appErrors.WrapDbExec(err, statement, announce)
	}
	return appErrors.ValidateDbResult(result, 1, "announcement was not inserted")
}

func (ar *announceRepo) Update(id uint, announce domains.Announce) error {
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

func (ar *announceRepo) Delete(id uint) error {
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

func CreateTestAnnounceRepo(db *sql.DB) AnnounceRepoInterface {
	ar := &announceRepo{}
	ar.Initialize(db)
	return ar
}
