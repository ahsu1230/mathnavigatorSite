package utils

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	repoUtils "github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos/utils"
)

var db *sql.DB

func SetupTestDatabase(host string, port int, username string, password string, dbName string) {
	// Open first connection to create database
	dbConn := repoUtils.Open(host, port, username, password, dbName)

	logger.Message("Creating test database connection...")
	logFields := logger.Fields{
		"host":      host,
		"port":      strconv.Itoa(port),
		"username":  username,
		"password":  password,
		"defaultDb": dbName,
	}
	logger.Debug("Database properties", logFields)

	tx, err := dbConn.Begin()
	if err != nil {
		logger.Error("Error connecting to db", err, logFields)
		panic(err.Error())
	}
	tx.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s;", dbName))
	tx.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbName))
	tx.Commit()

	// Must close db connection and open a new one with dbName
	dbConn.Close()
	logger.Message("Reopening test database...")
	dbConn = repoUtils.Open(host, port, username, password, dbName)
	if err := dbConn.Ping(); err != nil {
		logger.Error("Error reconnecting to db", err, logFields)
		panic(err.Error())
	}

	// Can now start operations with newly created test database
	logger.Message("Starting migrations...")
	repoUtils.Migrate(dbConn, "file://../repos/migrations")

	logger.Message("Initializing repoUtils...")
	ctx := context.Background()
	repos.SetupRepos(ctx, dbConn)

	db = dbConn
}

func ResetTable(t *testing.T, tableName string) error {
	logger.Debug("Resetting table", logger.Fields{"table": tableName})
	_, err := db.Exec(fmt.Sprintf("DELETE FROM %s; ", tableName))
	if err != nil {
		t.Fatalf("Error deleting table rows: %s", err)
	}
	_, err = db.Exec(fmt.Sprintf("ALTER TABLE %s AUTO_INCREMENT=1;", tableName))
	if err != nil {
		t.Fatalf("Error altering table auto-increment: %s", err)
	}
	return nil
}

func ResetAllTables(t *testing.T) {
	ResetTable(t, domains.TABLE_USER_CLASSES)
	ResetTable(t, domains.TABLE_USER_AFHS)
	ResetTable(t, domains.TABLE_USERS)
	ResetTable(t, domains.TABLE_TRANSACTIONS)
	ResetTable(t, domains.TABLE_ACCOUNTS)

	ResetTable(t, domains.TABLE_ACHIEVEMENTS)
	ResetTable(t, domains.TABLE_ANNOUNCEMENTS)

	ResetTable(t, domains.TABLE_SESSIONS)
	ResetTable(t, domains.TABLE_CLASSES)
	ResetTable(t, domains.TABLE_ASKFORHELP)
	ResetTable(t, domains.TABLE_PROGRAMS)
	ResetTable(t, domains.TABLE_SEMESTERS)
	ResetTable(t, domains.TABLE_LOCATIONS)
}

func CloseDb() {
	logger.Message("Closing DB connection")
	db.Close()
}
