package repos_test

import (
    "database/sql"
    "reflect"
    "testing"
    sqlmock "github.com/DATA-DOG/go-sqlmock"
    "github.com/ahsu1230/mathnavigatorSite/orion/domains"
    "github.com/ahsu1230/mathnavigatorSite/orion/repos"
)

func initTest(t *testing.T) (*sql.DB, sqlmock.Sqlmock, repos.ProgramRepoInterface) {
    db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
    repo := repos.CreateTestProgramRepo(db)
    return db, mock, repo
}

func TestGetAllPrograms(t *testing.T) {
    db, mock, repo := initTest(t)
    defer db.Close()

    // Mock DB queries and execute
    rows := sqlmock.NewRows([]string{"Id", "Name", "ProgramId", "Grade1", "Grade2", "Description"}).
                            AddRow(1, "Program1", "prog1", 2, 3, "descript1")
				mock.ExpectPrepare("SELECT (.+) FROM programs").ExpectQuery().WillReturnRows(rows)
    got, err := repo.SelectAll()
    if err != nil {
        t.Errorf("Unexpected error %v", err)
    }

    // Validate results
    want := []domains.Program{
        {
            Id:         1,
            Name:       "Program1",
            ProgramId:  "prog1",
            Grade1: 2,
            Grade2: 3,
            Description: "descript1",
        },
    }
    if !reflect.DeepEqual(got, want) {
        t.Errorf("SelectAll() = %v, want %v", got, want)
    }
}




















// func TestStoreGetAllPrograms(t *testing.T) {
//     db, dbx, mock := initTest(t)
// 	defer db.Close()

//     // Mock expected outcome
//     rows := sqlmock.NewRows(domains.ProgramColumns).
//         AddRow(1, 1000, 1000, nil, "prog1", "Program1", 2, 3, "description1").
//         AddRow(2, 2000, 2000, nil, "prog2", "Program2", 5, 6, "description2")
//     mock.ExpectQuery("^SELECT (.+) FROM programs").WillReturnRows(rows)

//     // Execute method
// 	if _, err := store.GetAllPrograms(dbx); err != nil {
// 		t.Errorf("Error was not expected: %s", err)
// 	}

//     // Check results
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("There were unfulfilled expectations: %s", err)
// 	}
// }

// func TestStoreGetProgramById(t *testing.T) {
//     db, dbx, mock := initTest(t)
//     defer db.Close()

//     // Mock expected outcome with valid programId
//     rows := sqlmock.NewRows(domains.ProgramColumns).
//         AddRow(1, 1000, 1000, nil, "prog1", "Program1", 2, 3, "description1")
//     mock.ExpectQuery("^SELECT (.+) FROM programs WHERE program_id=?").
//             WithArgs("prog1").
//             WillReturnRows(rows)
//     // Execute command and check expectations
//     store.GetProgramById(dbx, "prog1")
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("There were unfulfilled expectations: %s", err)
// 	}

//     // Mock expected outcome with invalid programId
//     mock.ExpectQuery("^SELECT (.+) FROM programs WHERE program_id=?").
//             WithArgs("prog2").
//             WillReturnError(fmt.Errorf("not found"))
//     // Execute command and check expectations
//     store.GetProgramById(dbx, "prog2")
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("There were unfulfilled expectations: %s", err)
// 	}
// }

// func TestStoreCreateProgram(t *testing.T) {
//     db, dbx, mock := initTest(t)
// 	defer db.Close()

//     // Mock expected SQL statement
//     program := domains.Program {
//         1, 1000, 1000, sql.NullInt64{Int64:0, Valid:false}, "prog1", "Program1", 2, 3, "description1",
//     }
//     mock.ExpectExec("^INSERT INTO programs").
//         WillReturnResult(sqlmock.NewResult(1, 1)) // last rowId is 1, 1 row affected
//     // Execute command and check expectations
//     store.InsertProgram(dbx, program)
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("There were unfulfilled expectations: %s", err)
// 	}

//     // Mock expected outcome when inserting same program again
//     // Test unique program_id
//     mock.ExpectExec("^INSERT INTO programs").
//         WillReturnError(fmt.Errorf("non-unique program_id"))
//     // Execute command and check expectations
//     if err := store.InsertProgram(dbx, program); err == nil {
//         t.Errorf("Expected an error but was none")
//     }
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("There were unfulfilled expectations: %s", err)
// 	}
// }

// // func TestStoreUpdateProgram(t *testing.T) {
// //     db, dbx, mock := initTest()
// // 	    defer db.Close()
// //
// //     // Mock expected outcome with valid programId
// //     rows := sqlmock.NewRows(domains.ProgramColumns).
// //         AddRow(1, 1000, 1000, nil, "prog1", "Program1", 2, 3, "description1")
// //     mock.ExpectQuery("^SELECT (.+) FROM programs WHERE program_id=?").
// //             WithArgs("prog1").
// //             WillReturnRows(rows)
// //     // Execute command and check expectations
// //     store.GetProgramById(dbx, "prog1")
// // 	if err := mock.ExpectationsWereMet(); err != nil {
// // 		t.Errorf("There were unfulfilled expectations: %s", err)
// // 	}
// //
// //     // Mock expected outcome with invalid programId
// //     mock.ExpectQuery("^SELECT (.+) FROM programs WHERE program_id=?").
// //             WithArgs("prog2").
// //             WillReturnError(fmt.Errorf("not found"))
// //     // Execute command and check expectations
// //     store.GetProgramById(dbx, "prog2")
// // 	if err := mock.ExpectationsWereMet(); err != nil {
// // 		t.Errorf("There were unfulfilled expectations: %s", err)
// // 	}
// // }

// func TestStoreDeleteProgram(t *testing.T) {
//     db, dbx, mock := initTest(t)
// 	defer db.Close()

//     // Mock expected outcome with valid programId
//     sqlmock.NewRows(domains.ProgramColumns).
//         AddRow(1, 1000, 1000, nil, "prog1", "Program1", 2, 3, "description1")
//     mock.ExpectExec("^DELETE FROM programs WHERE program_id=?").
//             WithArgs("prog1")
//     // Execute command and check expectations
//     store.DeleteProgram(dbx, "prog1")
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("There were unfulfilled expectations: %s", err)
// 	}

//     // Mock expected outcome with invalid programId
//     mock.ExpectExec("^DELETE FROM programs WHERE program_id=?").
//             WithArgs("prog1").
//             WillReturnError(fmt.Errorf("not found"))
//     // Execute command and check expectations
//     store.DeleteProgram(dbx, "prog1")
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("There were unfulfilled expectations: %s", err)
// 	}
// }
