package repos

import (
	"database/sql"
	"errors"
	"time"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
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
	Initialize(db *sql.DB)
	SelectAll(bool) ([]domains.Class, error)
	SelectAllUnpublished() ([]domains.Class, error)
	SelectByClassId(string) (domains.Class, error)
	SelectByProgramId(string) ([]domains.Class, error)
	SelectBySemesterId(string) ([]domains.Class, error)
	SelectByProgramAndSemesterId(string, string) ([]domains.Class, error)
	Insert(domains.Class) error
	Update(string, domains.Class) error
	Publish([]string) error
	Delete(string) error
}

func (cr *classRepo) Initialize(db *sql.DB) {
	cr.db = db
}

func (cr *classRepo) SelectAll(publishedOnly bool) ([]domains.Class, error) {
	results := make([]domains.Class, 0)

	var query string
	if publishedOnly {
		query = "SELECT * FROM classes WHERE published_at IS NOT NULL"
	} else {
		query = "SELECT * FROM classes"
	}
	stmt, err := cr.db.Prepare(query)
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
			&class.Times,
			&class.StartDate,
			&class.EndDate,
			&class.GoogleClassCode,
			&class.FullState,
			&class.PricePerSession,
			&class.PriceLump); errScan != nil {
			return results, errScan
		}
		results = append(results, class)
	}
	return results, nil
}

func (cr *classRepo) SelectAllUnpublished() ([]domains.Class, error) {
	results := make([]domains.Class, 0)

	stmt, err := cr.db.Prepare("SELECT * FROM classes WHERE published_at IS NULL")
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
			&class.Times,
			&class.StartDate,
			&class.EndDate,
			&class.GoogleClassCode,
			&class.FullState,
			&class.PricePerSession,
			&class.PriceLump); errScan != nil {
			return results, errScan
		}
		results = append(results, class)
	}
	return results, nil
}

func (cr *classRepo) SelectByClassId(classId string) (domains.Class, error) {
	statement := "SELECT * FROM classes WHERE class_id=?"
	stmt, err := cr.db.Prepare(statement)
	if err != nil {
		return domains.Class{}, err
	}
	defer stmt.Close()

	var class domains.Class
	row := stmt.QueryRow(classId)
	errScan := row.Scan(
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
		&class.Times,
		&class.StartDate,
		&class.EndDate,
		&class.GoogleClassCode,
		&class.FullState,
		&class.PricePerSession,
		&class.PriceLump)
	return class, errScan
}

func (cr *classRepo) SelectByProgramId(programId string) ([]domains.Class, error) {
	results := make([]domains.Class, 0)

	stmt, err := cr.db.Prepare("SELECT * FROM classes WHERE program_id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(programId)
	if err != nil {
		return nil, err
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
			&class.Times,
			&class.StartDate,
			&class.EndDate,
			&class.GoogleClassCode,
			&class.FullState,
			&class.PricePerSession,
			&class.PriceLump); errScan != nil {
			return results, errScan
		}
		results = append(results, class)
	}
	return results, nil
}

func (cr *classRepo) SelectBySemesterId(semesterId string) ([]domains.Class, error) {
	results := make([]domains.Class, 0)

	stmt, err := cr.db.Prepare("SELECT * FROM classes WHERE semester_id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(semesterId)
	if err != nil {
		return nil, err
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
			&class.Times,
			&class.StartDate,
			&class.EndDate,
			&class.GoogleClassCode,
			&class.FullState,
			&class.PricePerSession,
			&class.PriceLump); errScan != nil {
			return results, errScan
		}
		results = append(results, class)
	}
	return results, nil
}

func (cr *classRepo) SelectByProgramAndSemesterId(programId, semesterId string) ([]domains.Class, error) {
	results := make([]domains.Class, 0)

	stmt, err := cr.db.Prepare("SELECT * FROM classes WHERE program_id=? AND semester_id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(programId, semesterId)
	if err != nil {
		return nil, err
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
			&class.Times,
			&class.StartDate,
			&class.EndDate,
			&class.GoogleClassCode,
			&class.FullState,
			&class.PricePerSession,
			&class.PriceLump); errScan != nil {
			return results, errScan
		}
		results = append(results, class)
	}
	return results, nil
}

func (cr *classRepo) Insert(class domains.Class) error {
	statement := "INSERT INTO classes (" +
		"created_at, " +
		"updated_at, " +
		"program_id, " +
		"semester_id, " +
		"class_key, " +
		"class_id, " +
		"location_id, " +
		"times, " +
		"start_date, " +
		"end_date, " +
		"google_class_code, " +
		"full_state, " +
		"price_per_session, " +
		"price_lump " +
		") VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	stmt, err := cr.db.Prepare(statement)
	if err != nil {
		return err
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
		class.Times,
		class.StartDate,
		class.EndDate,
		class.GoogleClassCode,
		class.FullState,
		class.PricePerSession,
		class.PriceLump,
	)
	if err != nil {
		return err
	}
	return utils.HandleSqlExecResult(execResult, 1, "class was not inserted")
}

func (cr *classRepo) Update(classId string, class domains.Class) error {
	statement := "UPDATE classes SET " +
		"updated_at=?, " +
		"program_id=?, " +
		"semester_id=?, " +
		"class_key=?, " +
		"class_id=?, " +
		"location_id=?, " +
		"times=?, " +
		"start_date=?, " +
		"end_date=?, " +
		"google_class_code=?, " +
		"full_state=?, " +
		"price_per_session=?, " +
		"price_lump=? " +
		"WHERE class_id=?"
	stmt, err := cr.db.Prepare(statement)
	if err != nil {
		return err
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
		class.Times,
		class.StartDate,
		class.EndDate,
		class.GoogleClassCode,
		class.FullState,
		class.PricePerSession,
		class.PriceLump,
		classId)
	if err != nil {
		return err
	}
	return utils.HandleSqlExecResult(execResult, 1, "class was not updated")
}

func (cr *classRepo) Publish(classIds []string) error {
	var errorString string

	tx, err := cr.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("UPDATE classes SET published_at=? WHERE class_id=? AND published_at IS NULL")
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()
	for _, classId := range classIds {
		_, err := stmt.Exec(now, classId)
		if err != nil {
			errorString = utils.AppendError(errorString, classId, err)
		}
	}
	errorString = utils.AppendError(errorString, "", tx.Commit())

	if len(errorString) == 0 {
		return nil
	}
	return errors.New(errorString)
}

func (cr *classRepo) Delete(classId string) error {
	statement := "DELETE FROM classes WHERE class_id=?"
	stmt, err := cr.db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	execResult, err := stmt.Exec(classId)
	if err != nil {
		return err
	}
	return utils.HandleSqlExecResult(execResult, 1, "class was not deleted")
}

// For Tests Only
func CreateTestClassRepo(db *sql.DB) ClassRepoInterface {
	cr := &classRepo{}
	cr.Initialize(db)
	return cr
}

func generateClassId(class domains.Class) string {
	classId := class.ProgramId + "_" + class.SemesterId
	if class.ClassKey.Valid {
		return classId + "_" + class.ClassKey.String
	}
	return classId
}
