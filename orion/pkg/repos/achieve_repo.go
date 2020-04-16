package repos

import (
	"database/sql"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"time"
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
	SelectAll(bool) ([]domains.Achieve, error)
	SelectAllUnpublished() ([]domains.Achieve, error)
	SelectAllGroupedByYear() ([]domains.AchieveYearGroup, error)
	SelectById(uint) (domains.Achieve, error)
	Insert(domains.Achieve) error
	Update(uint, domains.Achieve) error
	Publish([]uint) error
	Delete(uint) error
}

func (ar *achieveRepo) Initialize(db *sql.DB) {
	ar.db = db
}

func (ar *achieveRepo) SelectAll(publishedOnly bool) ([]domains.Achieve, error) {
	results := make([]domains.Achieve, 0)

	var query string
	if publishedOnly {
		query = "SELECT * FROM achievements WHERE published_at IS NOT NULL"
	} else {
		query = "SELECT * FROM achievements"
	}
	stmt, err := ar.db.Prepare(query)
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
		var achieve domains.Achieve
		if errScan := rows.Scan(
			&achieve.Id,
			&achieve.CreatedAt,
			&achieve.UpdatedAt,
			&achieve.DeletedAt,
			&achieve.PublishedAt,
			&achieve.Year,
			&achieve.Message); errScan != nil {
			return results, errScan
		}
		results = append(results, achieve)
	}
	return results, nil
}

func (ar *achieveRepo) SelectAllUnpublished() ([]domains.Achieve, error) {
	results := make([]domains.Achieve, 0)

	stmt, err := ar.db.Prepare("SELECT * FROM achievements WHERE published_at IS NULL")
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
		var achieve domains.Achieve
		if errScan := rows.Scan(
			&achieve.Id,
			&achieve.CreatedAt,
			&achieve.UpdatedAt,
			&achieve.DeletedAt,
			&achieve.PublishedAt,
			&achieve.Year,
			&achieve.Message); errScan != nil {
			return results, errScan
		}
		results = append(results, achieve)
	}
	return results, nil
}

func (ar *achieveRepo) SelectAllGroupedByYear() ([]domains.AchieveYearGroup, error) {
	results := make([]domains.AchieveYearGroup, 0)

	stmt, err := ar.db.Prepare("SELECT * FROM achievements ORDER BY year DESC")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
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
			&achieve.PublishedAt,
			&achieve.Year,
			&achieve.Message); errScan != nil {
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
		return domains.Achieve{}, err
	}
	defer stmt.Close()

	var achieve domains.Achieve
	row := stmt.QueryRow(id)
	errScan := row.Scan(
		&achieve.Id,
		&achieve.CreatedAt,
		&achieve.UpdatedAt,
		&achieve.DeletedAt,
		&achieve.PublishedAt,
		&achieve.Year,
		&achieve.Message)
	return achieve, errScan
}

func (ar *achieveRepo) Insert(achieve domains.Achieve) error {
	statement := "INSERT INTO achievements (" +
		"created_at, " +
		"updated_at, " +
		"year, " +
		"message" +
		") VALUES (?, ?, ?, ?)"

	stmt, err := ar.db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()
	execResult, err := stmt.Exec(
		now,
		now,
		achieve.Year,
		achieve.Message)
	if err != nil {
		return err
	}
	return handleSqlExecResult(execResult, 1, "achievement was not inserted")
}

func (ar *achieveRepo) Update(id uint, achieve domains.Achieve) error {
	statement := "UPDATE achievements SET " +
		"updated_at=?, " +
		"year=?, " +
		"message=? " +
		"WHERE id=?"
	stmt, err := ar.db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()
	execResult, err := stmt.Exec(
		now,
		achieve.Year,
		achieve.Message,
		id)
	if err != nil {
		return err
	}
	return handleSqlExecResult(execResult, 1, "achievement was not updated")
}

func (ar *achieveRepo) Publish(ids []uint) error {
	errorList := make([]domains.PublishErrorBody, 0)

	// Begin Transaction
	tx, err := ar.db.Begin()
	str := "UPDATE achievements SET published_at=? WHERE id=? AND published_at IS NULL"
	stmt, err := tx.Prepare(str)
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()
	for _, id := range ids {
		_, err := stmt.Exec(now, id)
		if err != nil {
			errorList = append(errorList, domains.PublishErrorBody{RowId: id, Error: err})
		}
	}

	// End Transaction
	tx.Commit()

	return domains.ConcatErrors(errorList)
}

func (ar *achieveRepo) Delete(id uint) error {
	statement := "DELETE FROM achievements WHERE id=?"
	stmt, err := ar.db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	execResult, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	return handleSqlExecResult(execResult, 1, "achievement was not deleted")
}

// For Tests Only
func CreateTestAchieveRepo(db *sql.DB) AchieveRepoInterface {
	ar := &achieveRepo{}
	ar.Initialize(db)
	return ar
}
