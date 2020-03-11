package repos

import (
	"database/sql"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"time"
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
	ar.db = db
}

func (ar *announceRepo) SelectAll() ([]domains.Announce, error) {
	results := make([]domains.Announce, 0)

	stmt, err := ar.db.Prepare("SELECT * FROM announcements")
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
		var announce domains.Announce
		if errScan := rows.Scan(
			&announce.Id,
			&announce.CreatedAt,
			&announce.UpdatedAt,
			&announce.DeletedAt,
			&announce.PostedAt,
			&announce.Author,
			&announce.Message); errScan != nil {
			return results, errScan
		}
		results = append(results, announce)
	}

	return results, nil
}

func (ar *announceRepo) SelectByAnnounceId(id uint) (domains.Announce, error) {
	stmt, err := ar.db.Prepare("SELECT * FROM announcements WHERE id=?")
	if err != nil {
		return domains.Announce{}, err
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
		&announce.Message)

	return announce, errScan
}

func (ar *announceRepo) Insert(announce domains.Announce) error {
	stmt, err := ar.db.Prepare("INSERT INTO announcements (" +
		"created_at, " +
		"updated_at, " +
		"posted_at, " +
		"author, " +
		"message" +
		") VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()
	result, err := stmt.Exec(
		now,
		now,
		announce.PostedAt,
		announce.Author,
		announce.Message)
	if err != nil {
		return err
	}

	return handleSqlExecResult(result, 1, "announcement was not inserted")
}

func (ar *announceRepo) Update(id uint, announce domains.Announce) error {
	stmt, err := ar.db.Prepare("UPDATE announcements SET " +
		"updated_at=?, " +
		"posted_at=?, " +
		"author=?, " +
		"message=? " +
		"WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()
	result, err := stmt.Exec(
		now,
		announce.PostedAt,
		announce.Author,
		announce.Message,
		id)
	if err != nil {
		return err
	}

	return handleSqlExecResult(result, 1, "announcement was not updated")
}

func (ar *announceRepo) Delete(id uint) error {
	stmt, err := ar.db.Prepare("DELETE FROM announcements WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	return handleSqlExecResult(result, 1, "announcement was not deleted")
}

func CreateTestAnnounceRepo(db *sql.DB) AnnounceRepoInterface {
	ar := &announceRepo{}
	ar.Initialize(db)
	return ar
}
