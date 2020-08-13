package repos

import (
	"database/sql"
	"time"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/appErrors"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
)

// Global variable
var AchieveRepo AchieveRepoInterface = &achieveRepo{}

// Implements interface achieveRepoInterface
type achieveRepo struct {
	db *sql.DB // golang native db connection
}

// Interface to implement
type AchieveRepoInterface interface {
	Initialize(db *sql.DB)
	SelectAll() ([]domains.Achieve, error)
	SelectAllGroupedByYear() ([]domains.AchieveYearGroup, error)
	SelectById(uint) (domains.Achieve, error)
	Insert(domains.Achieve) error
	Update(uint, domains.Achieve) error
	Delete(uint) error
}

func (ar *achieveRepo) Initialize(db *sql.DB) {
	logger.Debug("Initialize AchieveRepo", logger.Fields{})
	ar.db = db
}

func (ar *achieveRepo) SelectAll() ([]domains.Achieve, error) {
	results := make([]domains.Achieve, 0)

	query := "SELECT * FROM achievements"

	stmt, err := ar.db.Prepare(query)
	if err != nil {
		return nil, appErrors.WrapDbPrepare(err, query)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, appErrors.WrapDbQuery(err, query, nil)
	}
	defer rows.Close()

	for rows.Next() {
		var achieve domains.Achieve
		if errScan := rows.Scan(
			&achieve.Id,
			&achieve.CreatedAt,
			&achieve.UpdatedAt,
			&achieve.DeletedAt,
			&achieve.Year,
			&achieve.Message,
			&achieve.Position); errScan != nil {
			return results, errScan
		}
		results = append(results, achieve)
	}
	return results, nil
}

func (ar *achieveRepo) SelectAllGroupedByYear() ([]domains.AchieveYearGroup, error) {
	results := make([]domains.AchieveYearGroup, 0)

	statement := "SELECT * FROM achievements ORDER BY year DESC, position ASC"
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

	var curYear uint = 0
	row := make([]domains.Achieve, 0)
	for rows.Next() {
		var achieve domains.Achieve
		if errScan := rows.Scan(
			&achieve.Id,
			&achieve.CreatedAt,
			&achieve.UpdatedAt,
			&achieve.DeletedAt,
			&achieve.Year,
			&achieve.Message,
			&achieve.Position); errScan != nil {
			return results, errScan
		}
		if achieve.Year != curYear {
			if len(row) > 0 {
				results = append(results, domains.AchieveYearGroup{Year: curYear, Achievements: row})
				row = nil
			}
			curYear = achieve.Year
		}
		row = append(row, achieve)
	}
	results = append(results, domains.AchieveYearGroup{Year: curYear, Achievements: row})

	return results, nil
}

func (ar *achieveRepo) SelectById(id uint) (domains.Achieve, error) {
	statement := "SELECT * FROM achievements WHERE id=?"
	stmt, err := ar.db.Prepare(statement)
	if err != nil {
		return domains.Achieve{}, appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	var achieve domains.Achieve
	row := stmt.QueryRow(id)
	if err = row.Scan(
		&achieve.Id,
		&achieve.CreatedAt,
		&achieve.UpdatedAt,
		&achieve.DeletedAt,
		&achieve.Year,
		&achieve.Message,
		&achieve.Position); err != nil {
		return domains.Achieve{}, appErrors.WrapDbExec(err, statement, id)
	}
	return achieve, nil
}

func (ar *achieveRepo) Insert(achieve domains.Achieve) error {
	statement := "INSERT INTO achievements (" +
		"created_at, " +
		"updated_at, " +
		"year, " +
		"message, " +
		"position " +
		") VALUES (?, ?, ?, ?, ?)"

	stmt, err := ar.db.Prepare(statement)
	if err != nil {
		return appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	now := time.Now().UTC()
	execResult, err := stmt.Exec(
		now,
		now,
		achieve.Year,
		achieve.Message,
		achieve.Position)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, achieve)
	}
	return appErrors.ValidateDbResult(execResult, 1, "achievement was not inserted")
}

func (ar *achieveRepo) Update(id uint, achieve domains.Achieve) error {
	statement := "UPDATE achievements SET " +
		"updated_at=?, " +
		"year=?, " +
		"message=?, " +
		"position=? " +
		"WHERE id=?"
	stmt, err := ar.db.Prepare(statement)
	if err != nil {
		return appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	now := time.Now().UTC()
	execResult, err := stmt.Exec(
		now,
		achieve.Year,
		achieve.Message,
		achieve.Position,
		id)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, id, achieve)
	}
	return appErrors.ValidateDbResult(execResult, 1, "achievement was not updated")
}

func (ar *achieveRepo) Delete(id uint) error {
	statement := "DELETE FROM achievements WHERE id=?"
	stmt, err := ar.db.Prepare(statement)
	if err != nil {
		return appErrors.WrapDbPrepare(err, statement)
	}
	defer stmt.Close()

	execResult, err := stmt.Exec(id)
	if err != nil {
		return appErrors.WrapDbExec(err, statement, id)
	}
	return appErrors.ValidateDbResult(execResult, 1, "achievement was not deleted")
}

// For Tests Only
func CreateTestAchieveRepo(db *sql.DB) AchieveRepoInterface {
	ar := &achieveRepo{}
	ar.Initialize(db)
	return ar
}
