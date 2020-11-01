package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func openDb(host string, port int, user string, pass string, dbName string) *sql.DB {
	connection := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=UTC", user, pass, host, port, dbName)
	db, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	fmt.Println("Aquila service starting...")

	// Connect to database
	dbHost := os.Getenv("DB_HOST")
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbDefault := os.Getenv("DB_DEFAULT")
	fmt.Println("Connecting to database...")
	db := openDb(dbHost, dbPort, dbUser, dbPassword, dbDefault)
	defer db.Close()
	fmt.Println("Database connected!")

	// Run web server
	engine := gin.Default()
	engine.Run(":8002")
}
