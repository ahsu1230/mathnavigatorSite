package repos

import (
	"database/sql"
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
	SelectAllByClassId(string) ([]domains.Session, error)
	SelectBySessionId(uint) (domains.Session, error)
	Insert(domains.Session) error
	Update(uint, domains.Session) error
	Delete(uint) error
	SelectAllUnpublished() ([]uint, error)
	Publish([]uint) error
}

func (sr *sessionRepo) Initialize(db *sql.DB) {
	sr.db = db
}

func (sr *sessionRepo) SelectAllByClassId(classId string) ([]domains.Session, error) {
	results := make([]domains.Session, 0)

	stmt, err := sr.db.Prepare("SELECT * FROM sessions WHERE class_id=? ORDER BY starts_at ASC")
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
			&session.ClassId,
			&session.StartsAt,
			&session.EndsAt,
			&session.Canceled,
			&session.Notes,
			&session.PublishedAt); errScan != nil {
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
		&session.ClassId,
		&session.StartsAt,
		&session.EndsAt,
		&session.Canceled,
		&session.Notes,
		&session.PublishedAt)

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

func (sr *sessionRepo) Update(id uint, session domains.Session) error {
	stmt, err := sr.db.Prepare("UPDATE sessions SET " +
		"updated_at=?, " +
		"class_id=?, " +
		"starts_at=?, " +
		"ends_at=?, " +
		"canceled=?, " +
		"notes=?, " +
		"published_at=? " +
		"WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()
	result, err := stmt.Exec(
		now,
		session.ClassId,
		session.StartsAt,
		session.EndsAt,
		session.Canceled,
		session.Notes,
		session.PublishedAt,
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

func (sr *sessionRepo) SelectAllUnpublished() ([]uint, error) {
	results := make([]uint, 0)

	stmt, err := sr.db.Prepare("SELECT id FROM sessions WHERE published_at IS NULL")
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
		var id uint
		if errScan := rows.Scan(&id); errScan != nil {
			return results, errScan
		}
		results = append(results, id)
	}
	return results, nil
}

func (sr *sessionRepo) Publish(ids []uint) error {
	for _, id := range ids {
		session, err := sr.SelectBySessionId(id)
		if err != nil {
			return err
		}
		if !session.PublishedAt.Valid {
			now := time.Now().UTC()
			session.PublishedAt.Scan(now)
			sr.Update(id, session)
		}
	}
	return nil
}

func CreateTestSessionRepo(db *sql.DB) SessionRepoInterface {
	sr := &sessionRepo{}
	sr.Initialize(db)
	return sr
}
