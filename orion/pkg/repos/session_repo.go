package repos

import (
	"database/sql"
	"errors"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
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
	Insert(domains.Session) error
	Publish([]uint) error
	Update(uint, domains.Session) error
	Delete(uint) error
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

func (sr *sessionRepo) Insert(session domains.Session) error {
	stmt, err := sr.db.Prepare("INSERT INTO sessions (" +
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
	result, err := stmt.Exec(
		now,
		now,
		session.ClassId,
		session.StartsAt,
		session.EndsAt,
		session.Canceled,
		session.Notes)
	if err != nil {
		return err
	}

	return handleSqlExecResult(result, 1, "session was not inserted")
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
		execResult, err := stmt.Exec(now, id)
		if err != nil {
			errorString = appendError(errorString, "", id, err)
			continue
		}
		err1 := handleSqlExecResult(execResult, 0, "session was not published") // session is already published, 0 rows affected
		err2 := handleSqlExecResult(execResult, 1, "session was not published") // session was not published, 1 row affected
		if err1 != nil && err2 != nil {
			errorString = appendError(errorString, "", id, err1)
		}
	}
	errorString = appendError(errorString, "", 0, tx.Commit())

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

func (sr *sessionRepo) Delete(id uint) error {
	stmt, err := sr.db.Prepare("DELETE FROM sessions WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	return handleSqlExecResult(result, 1, "session was not deleted")
}

func CreateTestSessionRepo(db *sql.DB) SessionRepoInterface {
	sr := &sessionRepo{}
	sr.Initialize(db)
	return sr
}
