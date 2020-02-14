package database

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	"github.com/golang-migrate/migrate/source/file"
)

var dbSql *sql.DB   // go/sql db connection
var DbSqlx *sqlx.DB // sqlx wrapper over db connection

func createConnectionInfo(host string, port int, user string, pass string) string {
	dbSchema := "mathnavdb"
	info := fmt.Sprintf("%s:%s@(%s)/%s", user, pass, host, dbSchema)
	info += "?charset=utf8&parseTime=True&loc=Local"
	fmt.Println(info)
	return info
}

func OpenDb(host string, port int, user string, pass string) {
	connection := createConnectionInfo(host, port, user, pass)

	// Connect to database using sql
	var err error
	dbSql, err = sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}

	// Wrap database connection with sqlx
	DbSqlx = sqlx.NewDb(dbSql, "mysql")
}

func Migrate() {
	// Create driver using sql db connection
	driver, err1 := mysql.WithInstance(dbSql, &mysql.Config{})
	if err1 != nil {
		panic(err1)
	}

	// Setup migrations
	migrationsPath := "file://database/migrations"
	m, err2 := migrate.NewWithDatabaseInstance(migrationsPath, "mysql", driver)
	if err2 != nil {
		panic(err2)
	}

	// Execute migrations
	version, _, _ := m.Version()
	fmt.Println("Previous migration version: ", version)
	err3 := m.Up()
	if err3 != nil && err3 != migrate.ErrNoChange {
		panic(err3)
	}
	version, _, _ = m.Version()
	fmt.Println("Current migration version: ", version)
}

func CloseDb() {
	DbSqlx.Close()
}
