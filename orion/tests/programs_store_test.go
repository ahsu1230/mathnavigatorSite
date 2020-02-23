package programs_test

import (
    "database/sql"
    "fmt"
    "testing"
    sqlmock "github.com/DATA-DOG/go-sqlmock"
    "github.com/ahsu1230/mathnavigatorSite/orion/domains"
    "github.com/ahsu1230/mathnavigatorSite/orion/stores"
)

func TestStoreGetAllPrograms(t *testing.T) {
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

func TestStoreGetProgramById(t *testing.T) {
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

func TestStoreCreateProgram(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
    dbx := stores.CreateDbSqlx(db)

    // Mock expected SQL statement
    program := domains.Program {
        1, 1000, 1000, sql.NullInt64{Int64:0, Valid:false}, "prog1", "Program1", 2, 3, "description1",
    }
    mock.ExpectExec("^INSERT INTO programs").
        // WithArgs("prog1").
        WillReturnResult(sqlmock.NewResult(1, 1))
    // Execute command and check expectations
    stores.InsertProgram(dbx, program)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}

    // Mock expected outcome when inserting same program again
    // Test unique program_id
    mock.ExpectExec("^INSERT INTO programs").
        WillReturnError(fmt.Errorf("asdfd"))
    // Execute command and check expectations
    stores.InsertProgram(dbx, program)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

// func TestStoreUpdateProgram(t *testing.T) {
//     db, mock, err := sqlmock.New()
//     if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()
//     dbx := stores.CreateDbSqlx(db)
//
//     // Mock expected outcome with valid programId
//     rows := sqlmock.NewRows(domains.ProgramColumns).
//         AddRow(1, 1000, 1000, nil, "prog1", "Program1", 2, 3, "description1")
//     mock.ExpectQuery("^SELECT (.+) FROM programs WHERE program_id=?").
//             WithArgs("prog1").
//             WillReturnRows(rows)
//     // Execute command and check expectations
//     stores.GetProgramById(dbx, "prog1")
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("There were unfulfilled expectations: %s", err)
// 	}
//
//     // Mock expected outcome with invalid programId
//     mock.ExpectQuery("^SELECT (.+) FROM programs WHERE program_id=?").
//             WithArgs("prog2").
//             WillReturnError(fmt.Errorf("not found"))
//     // Execute command and check expectations
//     stores.GetProgramById(dbx, "prog2")
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("There were unfulfilled expectations: %s", err)
// 	}
// }

func TestStoreDeleteProgram(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
    dbx := stores.CreateDbSqlx(db)

    // Mock expected outcome with valid programId
    sqlmock.NewRows(domains.ProgramColumns).
        AddRow(1, 1000, 1000, nil, "prog1", "Program1", 2, 3, "description1")
    mock.ExpectExec("^DELETE FROM programs WHERE program_id=?").
            WithArgs("prog1")
    // Execute command and check expectations
    stores.DeleteProgram(dbx, "prog1")
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}

    // Mock expected outcome with invalid programId
    mock.ExpectExec("^DELETE FROM programs WHERE program_id=?").
            WithArgs("prog1").
            WillReturnError(fmt.Errorf("not found"))
    // Execute command and check expectations
    stores.DeleteProgram(dbx, "prog1")
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
