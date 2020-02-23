package store

import (
    "database/sql"
    "fmt"

    "github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

func CreateDbSqlx(db *sql.DB) *sqlx.DB {
    return sqlx.NewDb(db, "mysql")
}

func createConnectionInfo(host string, port int, user string, pass string) string {
	dbSchema := "mathnavdb"
	info := fmt.Sprintf("%s:%s@(%s)/%s", user, pass, host, dbSchema)
	info += "?charset=utf8&parseTime=True&loc=Local"
	return info
}

func Open(host string, port int, user string, pass string) (*sql.DB, *sqlx.DB) {
    connection := createConnectionInfo(host, port, user, pass)
	dbSql, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}
    return dbSql, CreateDbSqlx(dbSql)
}

func Migrate(dbSql *sql.DB) {
    // Create driver using sql db connection
    fmt.Println("Performing DB Migrations...")
    driver, err1 := mysql.WithInstance(dbSql, &mysql.Config{})
    if err1 != nil {
        panic(err1)
    }

    // Setup migrations
    migrationsPath := "file://store/migrations"
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

func Close(dbSql *sql.DB, dbSqlx *sqlx.DB) error {
	dbSqlx.Close()
	dbSqlx = nil
	dbSql = nil
    return nil
}
