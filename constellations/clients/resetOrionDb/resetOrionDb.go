package main

import (
	"database/sql"
	"fmt"
	// "os"
	// "strconv"
	_ "github.com/go-sql-driver/mysql"
)

// This script exists to easily reset contents of a MySQL server
// However, if you are using Docker, it's better to use these commands:
//
// docker-compose kill db-mysql
// docker-compose rm db-mysql
// docker-compose up -d db-mysql
//
// If you are using MySQL to support your orion webserver,
// this CLI connects directly to the MySQL server.
// This script is for testing purposes only.
// It should NEVER be used for a non-local orion webserver!
//
// You can run this CLI using:
// go run resetOrionDb.go
//
// or via a binary:
// go build resetOrionDb.go
// ./resetOrionDb

func main() {
	fmt.Println("Fetching database information from environment variables...")
	dbHost := os.Getenv("DB_HOST")
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbDefault := os.Getenv("DB_DEFAULT")

	// Or feel free to hard code the values
	// dbHost := "127.0.0.1"
	// dbPort := 3308
	// dbUser := "user"
	// dbPassword := "password"
	// dbDefault := "mathnavdb"

	connection := createConnectionInfo(dbHost, dbPort, dbUser, dbPassword, dbDefault)
	fmt.Println("Connecting to database... ", connection)
	db, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	resetAllTables(db)
	fmt.Println("Done!")
}

func createConnectionInfo(host string, port int, user string, pass string, dbName string) string {
	info := fmt.Sprintf("%s:%s@(%s:%d)/%s", user, pass, host, port, dbName)
	info += "?charset=utf8&parseTime=True&loc=UTC"
	return info
}

func resetTable(db *sql.DB, tableName string) error {
	_, err := db.Exec(fmt.Sprintf("DELETE FROM %s; ", tableName))
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(fmt.Sprintf("ALTER TABLE %s AUTO_INCREMENT=1;", tableName))
	if err != nil {
		panic(err)
	}
	return nil
}

func resetAllTables(db *sql.DB) {
	resetTable(db, "sessions")
	resetTable(db, "classes")
	resetTable(db, "programs")
	resetTable(db, "semesters")
	resetTable(db, "locations")
	resetTable(db, "achievements")
	resetTable(db, "announcements")
	resetTable(db, "transactions")
	resetTable(db, "users")
	resetTable(db, "accounts")
	resetTable(db, "ask_for_help")
	resetTable(db, "user_classes")
	resetTable(db, "user_afh")
}
