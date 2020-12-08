package repos

import (
	"context"
	"database/sql"
	"time"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos/cache"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos/utils"
)

// Global variable
var ClassRepo ClassRepoInterface = &classRepo{}

// Implements interface classRepoInterface
type classRepo struct {
	db *sql.DB // golang native db connection
}

// Interface to implement
type ClassRepoInterface interface {
	Initialize(context.Context, *sql.DB)
	SelectAll(context.Context, bool) ([]domains.Class, error)
	SelectAllUnpublished(context.Context) ([]domains.Class, error)
	SelectByClassId(context.Context, string) (domains.Class, error)
	SelectByProgramId(context.Context, string) ([]domains.Class, error)
	SelectBySemesterId(context.Context, string) ([]domains.Class, error)
	SelectByProgramAndSemesterId(context.Context, string, string) ([]domains.Class, error)
	Insert(context.Context, domains.Class) (uint, error)
	Update(context.Context, string, domains.Class) error
	Publish(context.Context, []string) []error
	Archive(context.Context, string) error
	Delete(context.Context, string) error
}

func (cr *classRepo) Initialize(ctx context.Context, db *sql.DB) {
	utils.LogWithContext(ctx, "classRepo.Initialize", logger.Fields{})
	cr.db = db
}

func (cr *classRepo) SelectAll(ctx context.Context, publishedOnly bool) ([]domains.Class, error) {
	utils.LogWithContext(ctx, "classRepo.SelectAll", logger.Fields{
		"publishedOnly": publishedOnly})
	results := make([]domains.Class, 0)

	var query string
	if publishedOnly {
		query = "SELECT * FROM classes WHERE deleted_at IS NULL AND published_at IS NOT NULL"
	} else {
		query = "SELECT * FROM classes WHERE deleted_at IS NULL"
	}
	stmt, err := cr.db.Prepare(query)
	if err != nil {
		return nil, appErrors.WrapDbPrepare(err, query)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, appErrors.WrapDbQuery(err, query)
	}
	defer rows.Close()

	for rows.Next() {
		var class domains.Class
		if errScan := rows.Scan(
			&class.Id,
			&class.CreatedAt,
			&class.UpdatedAt,
			&class.DeletedAt,
			&class.PublishedAt,
			&class.ProgramId,
			&class.SemesterId,
			&class.ClassKey,
			&class.ClassId,
			&class.LocationId,
			&class.TimesStr,
			&class.GoogleClassCode,
			&class.FullState,
			&class.PricePerSession,
			&class.PriceLumpSum,
			&class.PaymentNotes); errScan != nil {
			return results, errScan
		}
		results = append(results, class)
	}
	return results, nil
}

func (cr *classRepo) SelectAllUnpublished(ctx context.Context) ([]domains.Class, error) {
	utils.LogWithContext(ctx, "classRepo.SelectAllUnpublished", logger.Fields{})
	results := make([]domains.Class, 0)

	statement := "SELECT * FROM classes WHERE deleted_at IS NULL AND published_at IS NULL"
	stmt, err := cr.db.Prepare(statement)
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
		var class domains.Class
		if errScan := rows.Scan(
			&class.Id,
			&class.CreatedAt,
			&class.UpdatedAt,
			&class.DeletedAt,
			&class.PublishedAt,
			&class.ProgramId,
			&class.SemesterId,
			&class.ClassKey,
			&class.ClassId,
			&class.LocationId,
			&class.TimesStr,
			&class.GoogleClassCode,
			&class.FullState,
			&class.PricePerSession,
			&class.PriceLumpSum,
			&class.PaymentNotes); errScan != nil {
			return results, errScan
		}
		results = append(results, class)
	}
	return results, nil
}

func (cr *classRepo) SelectByClassId(ctx context.Context, classId string) (domains.Class, error) {
	utils.LogWithContext(ctx, "classRepo.SelectByClassId", logger.Fields{"classId": classId})
	statement := "SELECT * FROM classes WHERE class_id=? AND deleted_at IS NULL"
	stmt, err := cr.db.Prepare(statement)
	if err != nil {
		return domains.Class{}, appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	var class domains.Class
	row := stmt.QueryRow(classId)
	if err = row.Scan(
		&class.Id,
		&class.CreatedAt,
		&class.UpdatedAt,
		&class.DeletedAt,
		&class.PublishedAt,
		&class.ProgramId,
		&class.SemesterId,
		&class.ClassKey,
		&class.ClassId,
		&class.LocationId,
		&class.TimesStr,
		&class.GoogleClassCode,
		&class.FullState,
		&class.PricePerSession,
		&class.PriceLumpSum,
		&class.PaymentNotes); err != nil {
		return domains.Class{}, appErrors.WrapDbExec(err, statement, classId)
	}
	return class, nil
}

func (cr *classRepo) SelectByProgramId(ctx context.Context, programId string) ([]domains.Class, error) {
	utils.LogWithContext(ctx, "classRepo.SelectByProgramId", logger.Fields{"programId": programId})
	results := make([]domains.Class, 0)

	statement := "SELECT * FROM classes WHERE program_id=? AND deleted_at IS NULL"
	stmt, err := cr.db.Prepare(statement)
	if err != nil {
		return nil, appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()
	rows, err := stmt.Query(programId)
	if err != nil {
		return nil, appErrors.WrapDbQuery(err, statement, programId)
	}
	defer rows.Close()

	for rows.Next() {
		var class domains.Class
		if errScan := rows.Scan(
			&class.Id,
			&class.CreatedAt,
			&class.UpdatedAt,
			&class.DeletedAt,
			&class.PublishedAt,
			&class.ProgramId,
			&class.SemesterId,
			&class.ClassKey,
			&class.ClassId,
			&class.LocationId,
			&class.TimesStr,
			&class.GoogleClassCode,
			&class.FullState,
			&class.PricePerSession,
			&class.PriceLumpSum,
			&class.PaymentNotes); errScan != nil {
			return results, errScan
		}
		results = append(results, class)
	}
	return results, nil
}

func (cr *classRepo) SelectBySemesterId(ctx context.Context, semesterId string) ([]domains.Class, error) {
	utils.LogWithContext(ctx, "classRepo.SelectBySemesterId", logger.Fields{"semesterId": semesterId})
	results := make([]domains.Class, 0)

	statement := "SELECT * FROM classes WHERE semester_id=? AND deleted_at IS NULL"
	stmt, err := cr.db.Prepare(statement)
	if err != nil {
		return nil, appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()
	rows, err := stmt.Query(semesterId)
	if err != nil {
		return nil, appErrors.WrapDbQuery(err, statement, semesterId)
	}
	defer rows.Close()

	for rows.Next() {
		var class domains.Class
		if errScan := rows.Scan(
			&class.Id,
			&class.CreatedAt,
			&class.UpdatedAt,
			&class.DeletedAt,
			&class.PublishedAt,
			&class.ProgramId,
			&class.SemesterId,
			&class.ClassKey,
			&class.ClassId,
			&class.LocationId,
			&class.TimesStr,
			&class.GoogleClassCode,
			&class.FullState,
			&class.PricePerSession,
			&class.PriceLumpSum,
			&class.PaymentNotes); errScan != nil {
			return results, errScan
		}
		results = append(results, class)
	}
	return results, nil
}

func (cr *classRepo) SelectByProgramAndSemesterId(ctx context.Context, programId, semesterId string) ([]domains.Class, error) {
	utils.LogWithContext(ctx, "classRepo.SelectByProgramAndSemesterId", logger.Fields{
		"programId":  programId,
		"semesterId": semesterId})
	results := make([]domains.Class, 0)

	statement := "SELECT * FROM classes WHERE program_id=? AND semester_id=? AND deleted_at IS NULL"
	stmt, err := cr.db.Prepare(statement)
	if err != nil {
		return nil, appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()
	rows, err := stmt.Query(programId, semesterId)
	if err != nil {
		return nil, appErrors.WrapDbQuery(err, statement, semesterId)
	}
	defer rows.Close()

	for rows.Next() {
		var class domains.Class
		if errScan := rows.Scan(
			&class.Id,
			&class.CreatedAt,
			&class.UpdatedAt,
			&class.DeletedAt,
			&class.PublishedAt,
			&class.ProgramId,
			&class.SemesterId,
			&class.ClassKey,
			&class.ClassId,
			&class.LocationId,
			&class.TimesStr,
			&class.GoogleClassCode,
			&class.FullState,
			&class.PricePerSession,
			&class.PriceLumpSum,
			&class.PaymentNotes); errScan != nil {
			return results, errScan
		}
		results = append(results, class)
	}
	return results, nil
}

func (cr *classRepo) Insert(ctx context.Context, class domains.Class) (uint, error) {
	utils.LogWithContext(ctx, "classRepo.Insert", logger.Fields{"class": class})
	statement := "INSERT INTO classes (" +
		"created_at, " +
		"updated_at, " +
		"program_id, " +
		"semester_id, " +
		"class_key, " +
		"class_id, " +
		"location_id, " +
		"times_str, " +
		"google_class_code, " +
		"full_state, " +
		"price_per_session, " +
		"price_lump_sum, " +
		"payment_notes " +
		") VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := cr.db.Prepare(statement)
	if err != nil {
		return 0, appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	now := time.Now().UTC()
	execResult, err := stmt.Exec(
		now,
		now,
		class.ProgramId,
		class.SemesterId,
		class.ClassKey,
		generateClassId(class),
		class.LocationId,
		class.TimesStr,
		class.GoogleClassCode,
		class.FullState,
		class.PricePerSession,
		class.PriceLumpSum,
		class.PaymentNotes,
	)
	if err != nil {
		return 0, appErrors.WrapDbExec(err, statement, class)
	}

	invalidateClassesCache(ctx)

	rowId, err := execResult.LastInsertId()
	if err != nil {
		return 0, appErrors.WrapSQLBadInsertResult(err)
	}
	return uint(rowId), appErrors.ValidateDbResult(execResult, 1, "class was not inserted")
}

func (cr *classRepo) Update(ctx context.Context, classId string, class domains.Class) error {
	utils.LogWithContext(ctx, "classRepo.Update", logger.Fields{"classId": classId})
	statement := "UPDATE classes SET " +
		"updated_at=?, " +
		"program_id=?, " +
		"semester_id=?, " +
		"class_key=?, " +
		"class_id=?, " +
		"location_id=?, " +
		"times_str=?, " +
		"google_class_code=?, " +
		"full_state=?, " +
		"price_per_session=?, " +
		"price_lump_sum=?, " +
		"payment_notes=? " +
		"WHERE class_id=?"
	stmt, err := cr.db.Prepare(statement)
	if err != nil {
		return appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	now := time.Now().UTC()
	execResult, err := stmt.Exec(
		now,
		class.ProgramId,
		class.SemesterId,
		class.ClassKey,
		generateClassId(class),
		class.LocationId,
		class.TimesStr,
		class.GoogleClassCode,
		class.FullState,
		class.PricePerSession,
		class.PriceLumpSum,
		class.PaymentNotes,
		classId)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, class, classId)
	}

	invalidateClassesCache(ctx)
	return appErrors.ValidateDbResult(execResult, 1, "class was not updated")
}

func (cr *classRepo) Publish(ctx context.Context, classIds []string) []error {
	utils.LogWithContext(ctx, "classRepo.Publish", logger.Fields{"classIds": classIds})
	tx, err := cr.db.Begin()
	if err != nil {
		return []error{appErrors.WrapDbTxBegin(err)}
	}
	statement := "UPDATE classes SET published_at=? WHERE class_id=? AND published_at IS NULL"
	stmt, err := tx.Prepare(statement)
	if err != nil {
		return []error{appErrors.WrapDbPrepare(err, statement)}
	}
	defer stmt.Close()

	var errorList []error
	now := time.Now().UTC()
	for _, classId := range classIds {
		result, err := stmt.Exec(now, classId)
		if err != nil {
			err = appErrors.WrapDbExec(err, statement, classId)
			errorList = append(errorList, err)
			continue
		}
		if err = appErrors.ValidateDbResult(result, 1, "class was not inserted"); err != nil {
			errorList = append(errorList, err)
		}
	}

	if err = tx.Commit(); err != nil {
		// TODO: Commit failed, need to rollback?
		return append(errorList, appErrors.WrapDbTxCommit(err))
	}

	invalidateClassesCache(ctx)
	return errorList
}

func (cr *classRepo) Archive(ctx context.Context, classId string) error {
	utils.LogWithContext(ctx, "classRepo.Archive", logger.Fields{"classId": classId})
	statement := "UPDATE classes SET deleted_at=? WHERE class_id=?"
	stmt, err := cr.db.Prepare(statement)
	if err != nil {
		return appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	now := time.Now().UTC()
	execResult, err := stmt.Exec(now, classId)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, classId)
	}

	invalidateClassesCache(ctx)
	return appErrors.ValidateDbResult(execResult, 1, "class was not archived")
}

func (cr *classRepo) Delete(ctx context.Context, classId string) error {
	utils.LogWithContext(ctx, "classRepo.Delete", logger.Fields{"classId": classId})
	statement := "DELETE FROM classes WHERE class_id=?"
	stmt, err := cr.db.Prepare(statement)
	if err != nil {
		return appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	execResult, err := stmt.Exec(classId)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, classId)
	}

	invalidateClassesCache(ctx)
	return appErrors.ValidateDbResult(execResult, 1, "class was not deleted")
}

// Helper functions
func generateClassId(class domains.Class) string {
	classId := class.ProgramId + "_" + class.SemesterId
	if class.ClassKey.Valid {
		return classId + "_" + class.ClassKey.String
	}
	return classId
}

// Call this function whenever classes have changed, cache must be invalidated!
func invalidateClassesCache(ctx context.Context) {
	cache.Delete(ctx, cache.KEY_PROGRAM_CLASSES_BY_SEMESTER)
}

// For Tests Only
func CreateTestClassRepo(ctx context.Context, db *sql.DB) ClassRepoInterface {
	cr := &classRepo{}
	cr.Initialize(ctx, db)
	return cr
}
