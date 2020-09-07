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
var SessionRepo SessionRepoInterface = &sessionRepo{}

type sessionRepo struct {
	db *sql.DB
}

type SessionRepoInterface interface {
	Initialize(context.Context, *sql.DB)
	SelectAllByClassId(context.Context, string) ([]domains.Session, error)
	SelectBySessionId(context.Context, uint) (domains.Session, error)
	Insert(context.Context, []domains.Session) ([]uint, []error)
	Update(context.Context, uint, domains.Session) error
	Delete(context.Context, []uint) []error
}

func (sr *sessionRepo) Initialize(ctx context.Context, db *sql.DB) {
	utils.LogWithContext(ctx, "sessionRepo.Initialize", logger.Fields{})
	sr.db = db
}

func (sr *sessionRepo) SelectAllByClassId(ctx context.Context, classId string) ([]domains.Session, error) {
	utils.LogWithContext(ctx, "sessionRepo.SelectAllByClassId", logger.Fields{"classId": classId})
	results := make([]domains.Session, 0)

	statement := "SELECT * FROM sessions WHERE class_id=? ORDER BY starts_at ASC"

	stmt, err := sr.db.Prepare(statement)
	if err != nil {
		return nil, appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()
	rows, err := stmt.Query(classId)
	if err != nil {
		return nil, appErrors.WrapDbQuery(err, statement, classId)
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
			&session.Notes); errScan != nil {
			return results, errScan
		}
		results = append(results, session)
	}

	return results, nil
}

func (sr *sessionRepo) SelectBySessionId(ctx context.Context, id uint) (domains.Session, error) {
	utils.LogWithContext(ctx, "sessionRepo.SelectBySessionId", logger.Fields{"id": id})
	statement := "SELECT * FROM sessions WHERE id=?"
	stmt, err := sr.db.Prepare(statement)
	if err != nil {
		return domains.Session{}, appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	var session domains.Session
	row := stmt.QueryRow(id)
	if err = row.Scan(
		&session.Id,
		&session.CreatedAt,
		&session.UpdatedAt,
		&session.DeletedAt,
		&session.ClassId,
		&session.StartsAt,
		&session.EndsAt,
		&session.Canceled,
		&session.Notes); err != nil {
		return domains.Session{}, appErrors.WrapDbExec(err, statement, id)
	}

	return session, nil
}

func (sr *sessionRepo) Insert(ctx context.Context, sessions []domains.Session) ([]uint, []error) {
	utils.LogWithContext(ctx, "sessionRepo.Insert", logger.Fields{"sessions": sessions})
	tx, err := sr.db.Begin()
	if err != nil {
		return []uint{}, []error{appErrors.WrapDbTxBegin(err)}
	}
	statement := "INSERT INTO sessions (" +
		"created_at, " +
		"updated_at, " +
		"class_id, " +
		"starts_at, " +
		"ends_at, " +
		"canceled, " +
		"notes" +
		") VALUES (?, ?, ?, ?, ?, ?, ?)"
	stmt, err := tx.Prepare(statement)
	if err != nil {
		return []uint{}, []error{appErrors.WrapDbPrepare(err, statement)}
	}
	defer stmt.Close()

	var ids []uint
	var errorList []error
	now := time.Now().UTC()
	for _, session := range sessions {
		result, err := stmt.Exec(
			now,
			now,
			session.ClassId,
			session.StartsAt,
			session.EndsAt,
			session.Canceled,
			session.Notes)
		if err != nil {
			errorList = append(errorList, appErrors.WrapDbExec(err, statement, session))
			continue
		} else if err = appErrors.ValidateDbResult(result, 1, "session was not inserted"); err != nil {
			errorList = append(errorList, err)
		}
		rowId, err := result.LastInsertId()
		if err != nil {
			errorList = append(errorList, appErrors.WrapSQLBadInsertResult(err))
		} else {
			ids = append(ids, uint(rowId))
		}
	}
	if err = tx.Commit(); err != nil {
		// TODO: Commit failed, need to rollback?
		return []uint{}, append(errorList, appErrors.WrapDbTxCommit(err))
	}
	return ids, errorList
}

func (sr *sessionRepo) Update(ctx context.Context, id uint, session domains.Session) error {
	utils.LogWithContext(ctx, "sessionRepo.Update", logger.Fields{"session": session})
	statement := "UPDATE sessions SET " +
		"updated_at=?, " +
		"class_id=?, " +
		"starts_at=?, " +
		"ends_at=?, " +
		"canceled=?, " +
		"notes=? " +
		"WHERE id=?"
	stmt, err := sr.db.Prepare(statement)
	if err != nil {
		return appErrors.WrapDbPrepare(err, statement)
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
		id)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, session, id)
	}

	return appErrors.ValidateDbResult(result, 1, "session was not updated")
}

func (sr *sessionRepo) Delete(ctx context.Context, ids []uint) []error {
	utils.LogWithContext(ctx, "sessionRepo.Delete", logger.Fields{"ids": ids})
	tx, err := sr.db.Begin()
	if err != nil {
		return []error{appErrors.WrapDbTxBegin(err)}
	}
	statement := "DELETE FROM sessions WHERE id=?"
	stmt, err := tx.Prepare(statement)
	if err != nil {
		return []error{appErrors.WrapDbPrepare(err, statement)}
	}
	defer stmt.Close()

	var errorList []error
	for _, id := range ids {
		result, err := stmt.Exec(id)
		if err != nil {
			err = appErrors.WrapDbExec(err, statement, id)
			errorList = append(errorList, err)
			continue
		}
		if err = appErrors.ValidateDbResult(result, 1, "session was not deleted"); err != nil {
			errorList = append(errorList, err)
		}
	}

	if err = tx.Commit(); err != nil {
		// TODO: Commit failed, need to rollback?
		return append(errorList, appErrors.WrapDbTxCommit(err))
	}
	return errorList
}

func CreateTestSessionRepo(ctx context.Context, db *sql.DB) SessionRepoInterface {
	sr := &sessionRepo{}
	sr.Initialize(ctx, db)
	return sr
}
