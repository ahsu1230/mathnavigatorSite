package repos

import "database/sql"

// Global variable
var AskForHelpRepo AskForHelpRepoInterface = &askForHelpRepo{}

// Implements interface askForHelpRepoInterface
type askForHelpRepo struct {
	db *sql.DB
}

// Interface to implement
type AskForHelpRepoInterface interface {
	Initialize(db *sql.DB)
}

func (ar *askForHelpRepo) Initialize(db *sql.DB) {
	ar.db = db
}

func CreateTestAFHRepo(db *sql.DB) AskForHelpRepoInterface {
	ar := &askForHelpRepo{}
	ar.Initialize(db)
	return ar
}
