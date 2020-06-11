package integration_tests

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/domains"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/middlewares"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/repos"
	"github.com/ahsu1230/mathnavigatorSite/constellations/orion/src/router"
	"github.com/gin-gonic/gin"
)

var db *sql.DB
var handler router.Handler

func TestMain(m *testing.M) {
	fmt.Println("Setting up Test Environment...")

	var configPath string
	if os.Getenv("TEST_ENV") == "test_ci" {
		configPath = "./configs/ci.yml"
	} else {
		configPath = "./configs/local.yml"
	}
	config := middlewares.RetrieveConfigurations(configPath)

	fmt.Println("Connecting to database...")
	configDb := config.Database
	db = setupTestDatabase(configDb.Host, configDb.Port, configDb.Username, configDb.Password, configDb.DbName)
	defer db.Close()

	fmt.Println("Setting up router...")
	handler = setupTestRouter()
	os.Exit(m.Run())
}

// Helper methods for Database
func setupTestDatabase(host string, port int, username string, password string, dbName string) *sql.DB {
	// Open first connection to create database
	dbConn := repos.Open(host, port, username, password, dbName)

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
	dbConn = repos.Open(host, port, username, password, dbName)
	if err := dbConn.Ping(); err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}

	// Can now start operations with newly created test database
	fmt.Println("Starting migrations...")
	repos.Migrate(dbConn, "file://../repos/migrations")

	fmt.Println("Initializing repos...")
	repos.SetupRepos(dbConn)

	return dbConn
}

func resetTable(t *testing.T, tableName string) error {
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

func resetAllTables(t *testing.T) {
	resetTable(t, domains.TABLE_SESSIONS)
	resetTable(t, domains.TABLE_CLASSES)
	resetTable(t, domains.TABLE_PROGRAMS)
	resetTable(t, domains.TABLE_SEMESTERS)
	resetTable(t, domains.TABLE_LOCATIONS)
	resetTable(t, domains.TABLE_ACHIEVEMENTS)
	resetTable(t, domains.TABLE_USERS)
}

// Helper methods for Router
func setupTestRouter() router.Handler {
	fmt.Println("Initializing Router...")
	gin.SetMode(gin.TestMode)
	engine := gin.Default()
	newHandler := router.Handler{Engine: engine}
	newHandler.SetupApiEndpoints()
	return newHandler
}

func sendHttpRequest(t *testing.T, method, url string, body io.Reader) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Errorf("http request error: %v\n", err)
	}
	w := httptest.NewRecorder()
	handler.Engine.ServeHTTP(w, req)
	return w
}

func createJsonBody(v interface{}) io.Reader {
	marshal, _ := json.Marshal(v)
	return bytes.NewBuffer(marshal)
}
