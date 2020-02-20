package programs_test

import (
    "fmt"
    "testing"
    sqlmock "github.com/DATA-DOG/go-sqlmock"
    "github.com/ahsu1230/mathnavigatorSite/orion/domains"
    "github.com/ahsu1230/mathnavigatorSite/orion/stores"
)

func TestGetAllPrograms(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
    dbx := stores.CreateDbSqlx(db)

    // Mock expected outcome
    rows := sqlmock.NewRows(domains.ProgramColumns).
        AddRow(1, 1000, 1000, nil, "prog1", "Program1", 2, 3, "description1").
        AddRow(2, 2000, 2000, nil, "prog2", "Program2", 5, 6, "description2")
    mock.ExpectQuery("^SELECT (.+) FROM programs").WillReturnRows(rows)

    // Execute method
    _, err = stores.GetAllPrograms(dbx)
	if err != nil {
		t.Errorf("Error was not expected: %s", err)
	}

    // Check results
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestGetProgramById(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
    dbx := stores.CreateDbSqlx(db)

    // Mock expected outcome with valid programId
    rows := sqlmock.NewRows(domains.ProgramColumns).
        AddRow(1, 1000, 1000, nil, "prog1", "Program1", 2, 3, "description1")
    mock.ExpectQuery("^SELECT (.+) FROM programs WHERE program_id=?").
            WithArgs("prog1").
            WillReturnRows(rows)
    // Execute command and check expectations
    stores.GetProgramById(dbx, "prog1")
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}

    // Mock expected outcome with invalid programId
    mock.ExpectQuery("^SELECT (.+) FROM programs WHERE program_id=?").
            WithArgs("prog2").
            WillReturnError(fmt.Errorf("not found"))
    // Execute command and check expectations
    stores.GetProgramById(dbx, "prog2")
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
