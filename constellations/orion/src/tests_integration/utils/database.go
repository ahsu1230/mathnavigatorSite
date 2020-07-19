package utils

import (
	"database/sql"
	"fmt"
	"strconv"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	repoUtils "github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos/utils"
)

var db *sql.DB

func SetupTestDatabase(host string, port int, username string, password string, dbName string) {
	// Open first connection to create database
	dbConn := repoUtils.Open(host, port, username, password, dbName)

	fmt.Println("Creating test database...")
	fmt.Println("host:"+host, "port:"+strconv.Itoa(port), "username:"+username, "password:"+password, "dbName:"+dbName)
	tx, err := dbConn.Begin()
	if err != nil {
		panic(err.Error())
	}
	tx.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s;", dbName))
	tx.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbName))
	tx.Commit()

	// Must close db connection and open a new one with dbName
	dbConn.Close()
	fmt.Println("Reopening test database...")
	dbConn = repoUtils.Open(host, port, username, password, dbName)
	if err := dbConn.Ping(); err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}

	// Can now start operations with newly created test database
	fmt.Println("Starting migrations...")
	repoUtils.Migrate(dbConn, "file://../repos/migrations")

	fmt.Println("Initializing repoUtils...")
	repos.SetupRepos(dbConn)

	db = dbConn
}

func ResetTable(t *testing.T, tableName string) error {
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
	ResetTable(t, domains.TABLE_SESSIONS)
	ResetTable(t, domains.TABLE_CLASSES)
	ResetTable(t, domains.TABLE_PROGRAMS)
	ResetTable(t, domains.TABLE_SEMESTERS)
	ResetTable(t, domains.TABLE_LOCATIONS)
	ResetTable(t, domains.TABLE_ACHIEVEMENTS)
	ResetTable(t, domains.TABLE_USERS)
	ResetTable(t, domains.TABLE_ACCOUNTS)
	ResetTable(t, domains.TABLE_ASKFORHELP)
	ResetTable(t, domains.TABLE_USERAFH)
	ResetTable(t, domains.TABLE_TRANSACTIONS)
}

func CloseDb() {
	db.Close()
}
