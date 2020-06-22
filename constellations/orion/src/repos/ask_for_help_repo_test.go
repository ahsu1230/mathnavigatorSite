package repos_test

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
)

func initAFHTest(t *testing.T) (*sql.DB, sqlmock.Sqlmock, repos.AskForHelpRepoInterface) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	repo := repos.CreateTestAFHRepo(db)
	return db, mock, repo
}
