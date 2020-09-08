package utils

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/logger"
)

func createConnectionInfo(host string, port int, user string, pass string, dbName string) string {
	info := fmt.Sprintf("%s:%s@(%s:%d)/%s", user, pass, host, port, dbName)
	info += "?charset=utf8&parseTime=True&loc=UTC&multiStatements=true"
	return info
}

func Open(host string, port int, user string, pass string, dbName string) *sql.DB {
	connection := createConnectionInfo(host, port, user, pass, dbName)
	db, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}
	return db
}

func Migrate(db *sql.DB, migrationsPath string) {
	// Create driver using sql db connection
	logger.Info("Performing DB Migrations...", logger.Fields{
		"migrationsPath": migrationsPath,
	})
	driver, err1 := mysql.WithInstance(db, &mysql.Config{})
	if err1 != nil {
		panic(err1)
	}

	// Setup migrations
	m, err2 := migrate.NewWithDatabaseInstance(migrationsPath, "mysql", driver)
	if err2 != nil {
		panic(err2)
	}

	// Execute migrations
	version, _, _ := m.Version()
	logger.Info("Previous migration version: ", logger.Fields{"version": version})
	err3 := m.Up()
	if err3 != nil && err3 != migrate.ErrNoChange {
		panic(err3)
	}
	version, _, _ = m.Version()
	logger.Info("Current migration version: ", logger.Fields{"version": version})
}

func Close(db *sql.DB) error {
	err := db.Close()
	db = nil
	return err
}
