package repos

import (
	"database/sql"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"time"
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
	SelectAll() ([]domains.Class, error)
	SelectByClassId(string) (domains.Class, error)
	SelectByProgramId(string) ([]domains.Class, error)
	SelectBySemesterId(string) ([]domains.Class, error)
	SelectByProgramAndSemesterId(string, string) ([]domains.Class, error)
	Insert(domains.Class) error
	Update(string, domains.Class) error
	Delete(string) error
}

func (cr *classRepo) Initialize(db *sql.DB) {
	cr.db = db
}

func (cr *classRepo) SelectAll() ([]domains.Class, error) {
	results := make([]domains.Class, 0)

	stmt, err := cr.db.Prepare("SELECT * FROM classes")
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
			&class.ProgramId,
			&class.SemesterId,
			&class.ClassKey,
			&class.ClassId,
			&class.LocationId,
			&class.Times,
			&class.StartDate,
			&class.EndDate); errScan != nil {
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
		&class.ProgramId,
		&class.SemesterId,
		&class.ClassKey,
		&class.ClassId,
		&class.LocationId,
		&class.Times,
		&class.StartDate,
		&class.EndDate)
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
			&class.ProgramId,
			&class.SemesterId,
			&class.ClassKey,
			&class.ClassId,
			&class.LocationId,
			&class.Times,
			&class.StartDate,
			&class.EndDate); errScan != nil {
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
			&class.ProgramId,
			&class.SemesterId,
			&class.ClassKey,
			&class.ClassId,
			&class.LocationId,
			&class.Times,
			&class.StartDate,
			&class.EndDate); errScan != nil {
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
			&class.ProgramId,
			&class.SemesterId,
			&class.ClassKey,
			&class.ClassId,
			&class.LocationId,
			&class.Times,
			&class.StartDate,
			&class.EndDate); errScan != nil {
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
		"end_date " +
		") VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

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
		class.EndDate)
	if err != nil {
		return err
	}
	return handleSqlExecResult(execResult, 1, "class was not inserted")
}

func (cr *classRepo) Update(classId string, class domains.Class) error {
	statement := "UPDATE classes SET " +
		"updated_at, " +
		"program_id, " +
		"semester_id, " +
		"class_key, " +
		"class_id, " +
		"location_id, " +
		"times, " +
		"start_date, " +
		"end_date, " +
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
		classId)
	if err != nil {
		return err
	}
	return handleSqlExecResult(execResult, 1, "class was not updated")
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
	return handleSqlExecResult(execResult, 1, "class was not deleted")
}

// For Tests Only
func CreateTestClassRepo(db *sql.DB) ClassRepoInterface {
	cr := &classRepo{}
	cr.Initialize(db)
	return cr
}

func generateClassId(class domains.Class) string {
	classId := class.ProgramId + "_" + class.SemesterId
	if len(class.ClassKey) != 0 {
		return classId + "_" + class.ClassKey
	}
	return classId
}
