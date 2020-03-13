package integration_tests

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/middlewares"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/repos"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/router"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var db *sql.DB
var handler router.Handler

func TestMain(m *testing.M) {
	fmt.Println("Setting up Test Environment...")
	config := middlewares.RetrieveConfigurations("./config_integrations.yaml")
	configDb := config.Database
	db = setupTestDatabase(configDb.Host, configDb.Port, configDb.Username, configDb.Password, configDb.DbName)
	defer db.Close()
	handler = setupTestRouter()
	os.Exit(m.Run())
}

// Helper methods for Database
func setupTestDatabase(host string, port int, username string, password string, dbName string) *sql.DB {
	// Open first connection to create database
	dbConn := repos.Open(host, port, username, password, "")
	tx, err := dbConn.Begin()
	if err != nil {
		panic(err.Error())
	}
	tx.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s;", dbName))
	tx.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbName))
	tx.Exec(fmt.Sprintf("USE %s;", dbName))
	tx.Commit()

	// Close connection and start a new connection using test database
	dbConn.Close()
	dbConn = repos.Open(host, port, username, password, dbName)
	// Can now start operations with newly created test database

	fmt.Println("Starting migrations...")
	repos.Migrate(dbConn, "file://../repos/migrations")

	fmt.Println("Initializing repos...")
	repos.ProgramRepo.Initialize(dbConn)
	repos.AnnounceRepo.Initialize(dbConn)
	// Initialize other tables here...

	if err := dbConn.Ping(); err != nil {
		panic(err.Error())
	}
	return dbConn
}

func refreshTable(t *testing.T, tableName string) error {
	stmt, err := db.Prepare(fmt.Sprintf("TRUNCATE TABLE %s;", tableName))
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		t.Fatalf("Error truncating tables: %s", err)
	}
	return nil
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
