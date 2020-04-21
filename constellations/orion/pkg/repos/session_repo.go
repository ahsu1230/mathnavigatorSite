package repos

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/pkg/domains"
	"time"
)

// Global variable
var SessionRepo SessionRepoInterface = &sessionRepo{}

type sessionRepo struct {
	db *sql.DB
}

type SessionRepoInterface interface {
	Initialize(db *sql.DB)
	SelectAllByClassId(string, bool) ([]domains.Session, error)
	SelectAllUnpublished() ([]domains.Session, error)
	SelectBySessionId(uint) (domains.Session, error)
	Insert([]domains.Session) error
	Update(uint, domains.Session) error
	Publish([]uint) error
	Delete([]uint) error
}

func (sr *sessionRepo) Initialize(db *sql.DB) {
	sr.db = db
}

func (sr *sessionRepo) SelectAllByClassId(classId string, publishedOnly bool) ([]domains.Session, error) {
	results := make([]domains.Session, 0)

	var statement string
	if publishedOnly {
		statement = "SELECT * FROM sessions WHERE class_id=? AND published_at IS NOT NULL ORDER BY starts_at ASC"
	} else {
		statement = "SELECT * FROM sessions WHERE class_id=? ORDER BY starts_at ASC"
	}
	stmt, err := sr.db.Prepare(statement)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(classId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var session domains.Session
		if errScan := rows.Scan(
			&session.Id,
			&session.CreatedAt,
			&session.UpdatedAt,
			&session.DeletedAt,
			&session.PublishedAt,
			&session.ClassId,
			&session.StartsAt,
			&session.EndsAt,
			&session.Canceled,
			&session.Notes); errScan != nil {
			return results, errScan
		}
		results = append(results, session)
	}

	return results, nil
}

func (sr *sessionRepo) SelectAllUnpublished() ([]domains.Session, error) {
	results := make([]domains.Session, 0)

	stmt, err := sr.db.Prepare("SELECT * FROM sessions WHERE published_at IS NULL")
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
		var session domains.Session
		if errScan := rows.Scan(
			&session.Id,
			&session.CreatedAt,
			&session.UpdatedAt,
			&session.DeletedAt,
			&session.PublishedAt,
			&session.ClassId,
			&session.StartsAt,
			&session.EndsAt,
			&session.Canceled,
			&session.Notes); errScan != nil {
			return results, errScan
		}
		results = append(results, session)
	}
	return results, nil
}

func (sr *sessionRepo) SelectBySessionId(id uint) (domains.Session, error) {
	stmt, err := sr.db.Prepare("SELECT * FROM sessions WHERE id=?")
	if err != nil {
		return domains.Session{}, err
	}
	defer stmt.Close()

	var session domains.Session
	row := stmt.QueryRow(id)
	errScan := row.Scan(
		&session.Id,
		&session.CreatedAt,
		&session.UpdatedAt,
		&session.DeletedAt,
		&session.PublishedAt,
		&session.ClassId,
		&session.StartsAt,
		&session.EndsAt,
		&session.Canceled,
		&session.Notes)

	return session, errScan
}

func (sr *sessionRepo) Insert(sessions []domains.Session) error {
	var errorString string

	tx, err := sr.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("INSERT INTO sessions (" +
		"created_at, " +
		"updated_at, " +
		"class_id, " +
		"starts_at, " +
		"ends_at, " +
		"canceled, " +
		"notes" +
		") VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()
	for _, session := range sessions {
		if err := session.Validate(); err != nil {
			errorString = appendError(errorString, fmt.Sprint(session.Id), err)
			continue
		}
		result, err := stmt.Exec(
			now,
			now,
			session.ClassId,
			session.StartsAt,
			session.EndsAt,
			session.Canceled,
			session.Notes)
		if err != nil {
			errorString = appendError(errorString, fmt.Sprint(session.Id), err)
			continue
		}
		if err = handleSqlExecResult(result, 1, "session was not inserted"); err != nil {
			errorString = appendError(errorString, fmt.Sprint(session.Id), err)
		}
	}
	errorString = appendError(errorString, "", tx.Commit())

	if len(errorString) == 0 {
		return nil
	}
	return errors.New(errorString)
}

func (sr *sessionRepo) Update(id uint, session domains.Session) error {
	stmt, err := sr.db.Prepare("UPDATE sessions SET " +
		"updated_at=?, " +
		"published_at=?, " +
		"class_id=?, " +
		"starts_at=?, " +
		"ends_at=?, " +
		"canceled=?, " +
		"notes=? " +
		"WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()
	result, err := stmt.Exec(
		now,
		session.PublishedAt,
		session.ClassId,
		session.StartsAt,
		session.EndsAt,
		session.Canceled,
		session.Notes,
		id)
	if err != nil {
		return err
	}

	return handleSqlExecResult(result, 1, "session was not updated")
}

func (sr *sessionRepo) Publish(ids []uint) error {
	var errorString string

	tx, err := sr.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("UPDATE sessions SET published_at=? WHERE id=? AND published_at IS NULL")
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()
	for _, id := range ids {
		_, err := stmt.Exec(now, id)
		if err != nil {
			errorString = appendError(errorString, fmt.Sprint(id), err)
		}
	}
	errorString = appendError(errorString, "", tx.Commit())

	if len(errorString) == 0 {
		return nil
	}
	return errors.New(errorString)
}

func (sr *sessionRepo) Delete(ids []uint) error {
	var errorString string

	tx, err := sr.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("DELETE FROM sessions WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, id := range ids {
		result, err := stmt.Exec(id)
		if err != nil {
			errorString = appendError(errorString, fmt.Sprint(id), err)
			continue
		}
		if err = handleSqlExecResult(result, 1, "session was not deleted"); err != nil {
			errorString = appendError(errorString, fmt.Sprint(id), err)
		}
	}
	errorString = appendError(errorString, "", tx.Commit())

	if len(errorString) == 0 {
		return nil
	}
	return errors.New(errorString)
}

func CreateTestSessionRepo(db *sql.DB) SessionRepoInterface {
	sr := &sessionRepo{}
	sr.Initialize(db)
	return sr
}
